package models

import "time"

type Schedule struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	UserID      int       `gorm:"not null" json:"user_id"`
	ScheduleDay string    `gorm:"size:10;not null" json:"schedule_day"`
	StartTime   string    `gorm:"not null" json:"start_time"`
	EndTime     string    `gorm:"not null" json:"end_time"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
