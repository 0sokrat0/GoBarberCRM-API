package models

import "time"

type Bookings struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	ClientID    uint      `gorm:"not null;index" json:"client_id"`
	ServiceID   uint      `gorm:"not null;index" json:"service_id"`
	UserID      uint      `gorm:"not null;index" json:"user_id"`
	BookingTime time.Time `gorm:"not null" json:"booking_time"`
	Status      string    `gorm:"size:50;default:'pending'" json:"status"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Client  Clients  `gorm:"foreignKey:ClientID" json:"client"`
	Service Services `gorm:"foreignKey:ServiceID" json:"service"`
	User    Users    `gorm:"foreignKey:UserID" json:"user"`
}
