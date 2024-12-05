package models

import (
	"time"
)

type AuthUser struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"unique;not null" json:"username"`         // Уникальное имя пользователя
	Password  string    `gorm:"not null" json:"password"`                // Храните хэшированный пароль
	UserID    *uint     `json:"user_id"`                                 // Опциональная связь с сотрудником
	User      *User     `gorm:"foreignKey:UserID" json:"user,omitempty"` // Связь с моделью User
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`        // Время создания записи
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`        // Время последнего обновления
}
