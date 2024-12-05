package models

import "time"

type Bookings struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	ClientID    int       `gorm:"not null;index" json:"client_id"`
	ServiceID   int       `gorm:"not null;index" json:"service_id"`
	UserID      int       `gorm:"not null;index" json:"user_id"`
	BookingTime time.Time `gorm:"not null" json:"booking_time"`
	Status      string    `gorm:"size:50;default:'pending'" json:"status"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Client  Client  `gorm:"foreignKey:ClientID" json:"client"`
	Service Service `gorm:"foreignKey:ServiceID" json:"service"`
	User    User    `gorm:"foreignKey:UserID" json:"user"`
}
