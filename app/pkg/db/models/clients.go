package models

import "time"

type Clients struct {
	ID          uint      `gorm:"primaryKey" json:"id"`               // ID
	FirstName   string    `json:"first_name"`                         // First Name
	LastName    string    `json:"last_name"`                          // Last Name
	Email       string    `gorm:"unique;size:255;index" json:"email"` // Email
	PhoneNumber string    `gorm:"size:20" json:"phone_number"`        //Phone Number
	TgID        int64     `gorm:"unique" json:"tg_id"`                // Telegram ID
	TgNickname  string    `gorm:"size:100" json:"tg_nickname"`        // Telegram Nickname
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`   // Created At
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`   // Updated At
}
