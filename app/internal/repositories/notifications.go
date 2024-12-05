package repositories

import (
	"errors"
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/models"

	"gorm.io/gorm"
)

var (
	ErrNotificationNotFound = errors.New("уведомление не найдено")
)

type NotificationRepository interface {
	CreateNotification(notification *models.Notification) error
	GetNotificationByID(id int) (*models.Notification, error)
	GetAllNotifications() ([]models.Notification, error)
	UpdateNotification(notification *models.Notification) error
	DeleteNotification(id int) error
}

type notificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &notificationRepository{
		db: db,
	}
}

func (r *notificationRepository) CreateNotification(notification *models.Notification) error {
	return r.db.Create(notification).Error
}

func (r *notificationRepository) GetNotificationByID(id int) (*models.Notification, error) {
	var notification models.Notification
	if err := r.db.First(&notification, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotificationNotFound
		}
		return nil, err
	}
	return &notification, nil
}

func (r *notificationRepository) GetAllNotifications() ([]models.Notification, error) {
	var notifications []models.Notification
	if err := r.db.Find(&notifications).Error; err != nil {
		return nil, err
	}
	return notifications, nil
}

func (r *notificationRepository) UpdateNotification(notification *models.Notification) error {
	return r.db.Save(notification).Error
}

func (r *notificationRepository) DeleteNotification(id int) error {
	if err := r.db.Delete(&models.Notification{}, id).Error; err != nil {
		return err
	}
	return nil
}
