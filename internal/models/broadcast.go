// Package models defines the application's persistence models.
package models

import "time"

// Broadcast stores a broadcast message sent by the application.
type Broadcast struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Subject string `json:"subject" gorm:"not null"`
	Message string `json:"message" gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
