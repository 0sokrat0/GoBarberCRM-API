package models

type Clients struct {
	ID          uint   `gorm:"primaryKey" json:"id"`             // Уникальный идентификатор клиента
	FirstName   string `gorm:"not null" json:"first_name"`       // Имя клиента
	LastName    string `gorm:"not null" json:"last_name"`        // Фамилия клиента
	Email       string `gorm:"unique;size:255" json:"email"`     // Email клиента
	PhoneNumber string `gorm:"size:20" json:"phone_number"`      // Номер телефона клиента
	TgID        int64  `gorm:"unique" json:"tg_id"`              // Telegram ID клиента (для связи через Telegram)
	TgNickname  string `gorm:"size:100" json:"tg_nickname"`      // Никнейм в Telegram (необязательно)
	CreatedAt   string `gorm:"autoCreateTime" json:"created_at"` // Дата создания записи
	UpdatedAt   string `gorm:"autoUpdateTime" json:"updated_at"` // Дата последнего обновления записи
}
