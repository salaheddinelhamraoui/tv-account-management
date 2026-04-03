// Package config loads application configuration from environment variables.
package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// Config groups all application configuration sections.
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	AWS      AWSConfig
	Upload   UploadConfig
}

// ServerConfig contains HTTP server settings.
type ServerConfig struct {
	Port    string
	GinMode string
}

// DatabaseConfig contains database connection settings.
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

// JWTConfig contains token signing and expiration settings.
type JWTConfig struct {
	Secret              string
	ExpiresIn           time.Duration
	RefreshTokenExpires time.Duration
}

// AWSConfig contains S3-compatible storage settings.
type AWSConfig struct {
	Region          string
	AccessKeyID     string
	SecretAccessKey string
	S3Bucket        string
	S3Endpoint      string
}

// UploadConfig contains local upload constraints.
type UploadConfig struct {
	Path        string
	MaxFileSize int64
}

// Load reads application configuration from the environment.
func Load() (*Config, error) {
	_ = godotenv.Load()

	jwtExpiresIn, _ := time.ParseDuration(getEnv("JWT_EXPIRATION_HOURS", "24h"))
	refreshExpiresIn, _ := time.ParseDuration(getEnv("REFRESH_TOKEN_EXPIRES", "72h"))
	maxUploadSize, _ := strconv.ParseInt(getEnv("MAX_UPLOAD_SIZE", "10485760"), 10, 64)

	return &Config{
		Server: ServerConfig{
			Port:    getEnv("PORT", "8080"),
			GinMode: getEnv("GIN_MODE", "debug"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "password"),
			Name:     getEnv("DB_NAME", "tv_accounts_management"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		}, JWT: JWTConfig{
			Secret:              getEnv("JWT_SECRET", "secret"),
			ExpiresIn:           jwtExpiresIn,
			RefreshTokenExpires: refreshExpiresIn,
		},
		AWS: AWSConfig{
			Region:          getEnv("AWS_REGION", "eu-central-1"),
			AccessKeyID:     getEnv("AWS_ACCESS_KEY_ID", "test"),
			SecretAccessKey: getEnv("AWS_SECRET_ACCESS_KEY", "test"),
			S3Bucket:        getEnv("AWS_S3_BUCKET_NAME", "tv-management-uploads"),
			S3Endpoint:      getEnv("AWS_S3_ENDPOINT", "http://localhost:9000"),
		},
		Upload: UploadConfig{
			Path:        getEnv("UPLOAD_PATH", "./uploads"),
			MaxFileSize: maxUploadSize,
		},
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
