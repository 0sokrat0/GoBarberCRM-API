package models

import "time"

type Users struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Username     string    `gorm:"unique;size:50;not null" json:"username"`
	PasswordHash string    `gorm:"not null" json:"password_hash"`
	Role         string    `gorm:"size:50" json:"role"`
	Email        string    `gorm:"unique;size:100;index" json:"email"`
	PhoneNumber  string    `gorm:"size:20" json:"phone_number"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	LastLoginAt  time.Time `gorm:"index" json:"last_login_at,omitempty"`
}
