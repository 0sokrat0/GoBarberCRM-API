package models

import "time"

type Break struct {
	ID         int       `gorm:"primaryKey" json:"id"`
	UserID     int       `gorm:"not null;index" json:"user_id"`
	BreakStart time.Time `gorm:"not null" json:"break_start"`
	BreakEnd   time.Time `gorm:"not null" json:"break_end"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	User User `gorm:"foreignKey:UserID" json:"user"`
}
