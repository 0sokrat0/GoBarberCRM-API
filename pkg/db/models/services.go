package models

type Services struct {
	ID          uint    `gorm:"primaryKey" json:"id"`             // Уникальный идентификатор услуги
	Name        string  `gorm:"size:255;not null" json:"name"`    // Название услуги
	Description string  `gorm:"type:text" json:"description"`     // Описание услуги
	Price       float64 `gorm:"not null" json:"price"`            // Цена услуги
	Duration    int     `gorm:"not null" json:"duration"`         // Продолжительность услуги (в минутах)
	CreatedAt   string  `gorm:"autoCreateTime" json:"created_at"` // Дата создания
	UpdatedAt   string  `gorm:"autoUpdateTime" json:"updated_at"` // Дата последнего обновления
}
