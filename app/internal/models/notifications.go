package models

import "time"

type Notification struct {
	ID               int       `gorm:"primaryKey" json:"id"`
	ClientID         int       `gorm:"not null" json:"client_id"`
	Message          string    `gorm:"type:text;not null" json:"message"`
	NotificationType string    `gorm:"size:50" json:"notification_type"`
	SentAt           time.Time `gorm:"autoCreateTime" json:"sent_at"`
	Status           string    `gorm:"size:50;default:'pending'" json:"status"`
}
