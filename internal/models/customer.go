package models

import (
	"time"

	"gorm.io/gorm"
)

// Customer represents a customer tied to a user account.
type Customer struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	UserID         uint   `json:"user_id" gorm:"not null;index"`
	Email          string `json:"email" gorm:"index"`
	WhatsappNumber string `json:"whatsapp_number"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	User         User          `json:"-" gorm:"constraint:OnDelete:CASCADE"`
	IPTVAccounts []IPTVAccount `json:"iptv_accounts"`
}
