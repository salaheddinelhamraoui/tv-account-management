package models

import (
	"time"

	"gorm.io/gorm"
)

// ProviderPackage represents a package exposed by an IPTV provider.
type ProviderPackage struct {
	ID                uint   `json:"id" gorm:"primaryKey"`
	ProviderID        uint   `json:"provider_id" gorm:"not null;index"`
	DisplayName       string `json:"display_name" gorm:"not null"`
	ExternalPackageID string `json:"external_package_id" gorm:"not null"`
	DurationMonths    int    `json:"duration_months" gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Provider     IPTVProvider  `json:"-" gorm:"constraint:OnDelete:CASCADE"`
	IPTVAccounts []IPTVAccount `json:"-"`
}
