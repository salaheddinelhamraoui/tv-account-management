package models

import (
	"time"

	"gorm.io/gorm"
)

// User model represents a user in the system with authentication and role-based access control.
type User struct {
	ID        uint     `json:"id" gorm:"primaryKey"`
	Email     string   `gorm:"uniqueIndex:idx_user_email;not null"`
	Password  string   `json:"-" gorm:"not null;size:255"`
	FirstName string   `json:"first_name" gorm:"not null"`
	LastName  string   `json:"last_name" gorm:"not null"`
	IsActive  bool     `json:"is_active" gorm:"default:true"`
	Role      UserRole `json:"role" gorm:"type:varchar(20);default:'user'"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	RefreshTokens []RefreshToken `json:"-"`
	Customers     []Customer     `json:"-"`
}

// UserRole (user or admin)
type UserRole string

const (
	// UserRoleAdmin is the role for administrators with elevated privileges.
	UserRoleAdmin UserRole = "admin"
	// UserRoleUser is the default role for regular users.
	UserRoleUser UserRole = "user"
)

// RefreshToken represents a user's refresh token used for authentication.
type RefreshToken struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	UserID    uint           `json:"user_id" gorm:"not null;index"`
	Token     string         `json:"token" gorm:"not null;uniqueIndex"`
	ExpiresAt time.Time      `gorm:"not null;index"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	User User `json:"-" gorm:"constraint:OnDelete:CASCADE"`
}
