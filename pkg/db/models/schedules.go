package models

type Schedules struct {
	ID          uint   `gorm:"primaryKey" json:"id"`                 // Уникальный идентификатор расписания
	UserID      uint   `gorm:"not null" json:"user_id"`              // Ссылка на сотрудника
	ScheduleDay string `gorm:"size:10;not null" json:"schedule_day"` // День недели (например, понедельник)
	StartTime   string `gorm:"not null" json:"start_time"`           // Время начала рабочего дня
	EndTime     string `gorm:"not null" json:"end_time"`             // Время окончания рабочего дня
	CreatedAt   string `gorm:"autoCreateTime" json:"created_at"`     // Дата создания
	UpdatedAt   string `gorm:"autoUpdateTime" json:"updated_at"`     // Дата последнего обновления
}
