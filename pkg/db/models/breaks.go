package models

import "gorm.io/gorm"

type Breaks struct {
	gorm.Model
	ID         uint   `gorm:"primaryKey" json:"id"`             // Уникальный идентификатор перерыва
	UserID     uint   `gorm:"not null" json:"user_id"`          // Ссылка на сотрудника
	BreakStart string `gorm:"not null" json:"break_start"`      // Время начала перерыва
	BreakEnd   string `gorm:"not null" json:"break_end"`        // Время окончания перерыва
	CreatedAt  string `gorm:"autoCreateTime" json:"created_at"` // Дата создания
	UpdatedAt  string `gorm:"autoUpdateTime" json:"updated_at"` // Дата последнего обновления
}
