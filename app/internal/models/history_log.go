package models

type HistoryLogs struct {
	ID        uint   `gorm:"primaryKey" json:"id"`             // Уникальный идентификатор записи
	UserID    uint   `gorm:"not null" json:"user_id"`          // Ссылка на пользователя
	Action    string `gorm:"not null" json:"action"`           // Описание действия
	CreatedAt string `gorm:"autoCreateTime" json:"created_at"` // Дата действия
}
