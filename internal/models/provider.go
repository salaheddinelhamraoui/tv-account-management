package models

import (
	"time"

	"gorm.io/gorm"
)

// IPTVProvider represents an upstream IPTV service provider.
type IPTVProvider struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"not null"`
	APIURL   string `json:"api_url" gorm:"not null"`
	APIToken string `json:"-" gorm:"not null"`
	Type     string `json:"type" gorm:"default:'xtream'"`
	IsActive bool   `json:"is_active" gorm:"default:true"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Packages     []ProviderPackage `json:"packages"`
	IPTVAccounts []IPTVAccount     `json:"-"`
}
