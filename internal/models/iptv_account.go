package models

import (
	"time"

	"gorm.io/gorm"
)

// IPTVAccount represents a provisioned IPTV account for a customer.
type IPTVAccount struct {
	ID uint `json:"id" gorm:"primaryKey"`

	CustomerID uint `json:"customer_id" gorm:"not null;index"`
	ProviderID uint `json:"provider_id" gorm:"not null;index"`
	PackageID  uint `json:"package_id" gorm:"not null;index"`

	ProviderUserID string `json:"provider_user_id"` // ID from IPTV panel

	Username  string `json:"username" gorm:"not null"`
	Password  string `json:"password" gorm:"not null"`
	ServerURL string `json:"server_url" gorm:"not null"`
	M3UURL    string `json:"m3u_url"`

	MaxConnections int    `json:"max_connections"`
	Status         string `json:"status"` // optional cache

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Customer Customer        `json:"-" gorm:"constraint:OnDelete:CASCADE"`
	Provider IPTVProvider    `json:"-" gorm:"constraint:OnDelete:CASCADE"`
	Package  ProviderPackage `json:"-" gorm:"constraint:OnDelete:CASCADE"`

	Subscriptions []Subscription `json:"subscriptions"`
}
