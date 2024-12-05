package models

import "time"

type Client struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `gorm:"unique;size:255;index" json:"email"`
	PhoneNumber string    `gorm:"size:20" json:"phone_number"`
	TgID        int64     `gorm:"unique" json:"tg_id"`
	TgNickname  string    `gorm:"size:100" json:"tg_nickname"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
