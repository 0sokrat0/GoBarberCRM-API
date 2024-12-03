package models

type Notifications struct {
	ID               uint   `gorm:"primaryKey" json:"id"`                    // Уникальный идентификатор уведомления
	ClientID         uint   `gorm:"not null" json:"client_id"`               // Ссылка на клиента
	Message          string `gorm:"type:text;not null" json:"message"`       // Сообщение уведомления
	NotificationType string `gorm:"size:50" json:"notification_type"`        // Тип уведомления (например, Telegram)
	SentAt           string `gorm:"autoCreateTime" json:"sent_at"`           // Время отправки уведомления
	Status           string `gorm:"size:50;default:'pending'" json:"status"` // Статус уведомления
}
