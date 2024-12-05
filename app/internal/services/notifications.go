package services

import (
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/models"
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/repositories"
)

type NotificationService interface {
	CreateNotification(notification *models.Notification) error
	GetNotificationByID(id int) (*models.Notification, error)
	GetAllNotifications() ([]models.Notification, error)
	UpdateNotification(id int, input *models.Notification) error
	DeleteNotification(id int) error
}

type notificationService struct {
	repo repositories.NotificationRepository
}

func NewNotificationService(repo repositories.NotificationRepository) NotificationService {
	return &notificationService{
		repo: repo,
	}
}

func (s *notificationService) CreateNotification(notification *models.Notification) error {
	return s.repo.CreateNotification(notification)
}

func (s *notificationService) GetNotificationByID(id int) (*models.Notification, error) {
	return s.repo.GetNotificationByID(id)
}

func (s *notificationService) GetAllNotifications() ([]models.Notification, error) {
	return s.repo.GetAllNotifications()
}

func (s *notificationService) UpdateNotification(id int, input *models.Notification) error {
	notification, err := s.repo.GetNotificationByID(id)
	if err != nil {
		return err
	}

	// Обновляем поля
	notification.Message = input.Message
	notification.NotificationType = input.NotificationType
	notification.Status = input.Status

	return s.repo.UpdateNotification(notification)
}

func (s *notificationService) DeleteNotification(id int) error {
	return s.repo.DeleteNotification(id)
}
