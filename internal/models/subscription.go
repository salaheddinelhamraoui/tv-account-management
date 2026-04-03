package models

import (
	"time"

	"gorm.io/gorm"
)

// Subscription is linked to IPTV accounts and represent the active subscription details, including duration and validity period.
type Subscription struct {
	ID            uint `json:"id" gorm:"primaryKey"`
	IPTVAccountID uint `json:"iptv_account_id" gorm:"not null;index"`

	DurationMonths int       `json:"duration_months" gorm:"not null"`
	StartDate      time.Time `json:"start_date" gorm:"not null"`
	EndDate        time.Time `json:"end_date" gorm:"not null;index"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	IPTVAccount IPTVAccount `json:"-" gorm:"constraint:OnDelete:CASCADE"`
}
