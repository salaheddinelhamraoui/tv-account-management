package models

import "time"

// EmailLog records the delivery status of an email.
type EmailLog struct {
	ID     uint `json:"id" gorm:"primaryKey"`
	UserID uint `json:"user_id" gorm:"not null;index"`

	Type    string `json:"type"` // credentials, reminder, broadcast
	Subject string `json:"subject"`
	Status  string `json:"status"` // sent, failed

	ErrorMessage string `json:"error_message"`

	SentAt time.Time

	User User `json:"-" gorm:"constraint:OnDelete:CASCADE"`
}
