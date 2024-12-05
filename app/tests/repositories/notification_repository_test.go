package repositories

import (
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/models"
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/repositories"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNotificationRepository_CreateNotification(t *testing.T) {
	db := setupTestDB(t, &models.Notification{})
	repo := repositories.NewNotificationRepository(db)

	notification := &models.Notification{
		ClientID:         1,
		Message:          "Test notification",
		NotificationType: "Email",
		Status:           "pending",
	}

	err := repo.CreateNotification(notification)
	require.NoError(t, err)
	require.NotZero(t, notification.ID)
}

func TestNotificationRepository_GetNotificationByID(t *testing.T) {
	db := setupTestDB(t, &models.Notification{})
	repo := repositories.NewNotificationRepository(db)

	notification := &models.Notification{
		ClientID:         1,
		Message:          "Test notification",
		NotificationType: "Email",
		Status:           "pending",
	}
	err := repo.CreateNotification(notification)
	require.NoError(t, err)

	fetchedNotification, err := repo.GetNotificationByID(notification.ID)
	require.NoError(t, err)
	assert.Equal(t, notification.ID, fetchedNotification.ID)
	assert.Equal(t, notification.Message, fetchedNotification.Message)
}

func TestNotificationRepository_GetAllNotifications(t *testing.T) {
	db := setupTestDB(t, &models.Notification{})
	repo := repositories.NewNotificationRepository(db)

	notifications := []models.Notification{
		{ClientID: 1, Message: "Notification 1", NotificationType: "Email", Status: "sent"},
		{ClientID: 2, Message: "Notification 2", NotificationType: "SMS", Status: "pending"},
	}

	for i := range notifications {
		err := repo.CreateNotification(&notifications[i])
		require.NoError(t, err)
	}

	fetchedNotifications, err := repo.GetAllNotifications()
	require.NoError(t, err)
	assert.Len(t, fetchedNotifications, 2)
}

func TestNotificationRepository_UpdateNotification(t *testing.T) {
	db := setupTestDB(t, &models.Notification{})
	repo := repositories.NewNotificationRepository(db)

	notification := &models.Notification{
		ClientID:         1,
		Message:          "Old message",
		NotificationType: "Email",
		Status:           "pending",
	}
	err := repo.CreateNotification(notification)
	require.NoError(t, err)

	notification.Message = "Updated message"
	err = repo.UpdateNotification(notification)
	require.NoError(t, err)

	updatedNotification, err := repo.GetNotificationByID(notification.ID)
	require.NoError(t, err)
	assert.Equal(t, "Updated message", updatedNotification.Message)
}

func TestNotificationRepository_DeleteNotification(t *testing.T) {
	db := setupTestDB(t, &models.Notification{})
	repo := repositories.NewNotificationRepository(db)

	notification := &models.Notification{
		ClientID:         1,
		Message:          "Test notification",
		NotificationType: "Email",
		Status:           "pending",
	}
	err := repo.CreateNotification(notification)
	require.NoError(t, err)

	err = repo.DeleteNotification(notification.ID)
	require.NoError(t, err)

	_, err = repo.GetNotificationByID(notification.ID)
	assert.Error(t, err)
	assert.Equal(t, repositories.ErrNotificationNotFound, err)
}
