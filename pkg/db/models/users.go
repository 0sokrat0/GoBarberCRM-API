package models

import "time"

// Users представляет сущность пользователя
type Users struct {
	ID           uint      `gorm:"primaryKey" json:"id"`                    // Уникальный идентификатор пользователя
	Username     string    `gorm:"unique;size:50;not null" json:"username"` // Логин пользователя
	PasswordHash string    `gorm:"not null" json:"-"`                       // Хеш пароля
	Role         string    `gorm:"size:50" json:"role"`                     // Роль пользователя
	Email        string    `gorm:"unique;size:100" json:"email"`            // Email пользователя
	PhoneNumber  string    `gorm:"size:20" json:"phone_number"`             // Номер телефона
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`        // Дата создания
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`        // Дата обновления
}
