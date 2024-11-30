package models

import "gorm.io/gorm"

type Bookings struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey" json:"id"`                    // Уникальный идентификатор бронирования
	ClientID    uint   `gorm:"not null" json:"client_id"`               // Ссылка на клиента
	ServiceID   uint   `gorm:"not null" json:"service_id"`              // Ссылка на услугу
	UserID      uint   `gorm:"not null" json:"user_id"`                 // Ссылка на сотрудника
	BookingTime string `gorm:"not null" json:"booking_time"`            // Время бронирования
	Status      string `gorm:"size:50;default:'pending'" json:"status"` // Статус бронирования
	CreatedAt   string `gorm:"autoCreateTime" json:"created_at"`        // Дата создания
	UpdatedAt   string `gorm:"autoUpdateTime" json:"updated_at"`        // Дата последнего обновления
}
