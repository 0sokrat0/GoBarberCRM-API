package notifications

import (
	"net/http"
	"strconv"

	"github.com/0sokrat0/GoGRAFFApi.git/app/pkg/db"
	"github.com/0sokrat0/GoGRAFFApi.git/app/pkg/db/models"
	"github.com/gin-gonic/gin"
)

func CreateNotificationHandler(c *gin.Context) {
	var notification models.Notifications
	if err := c.ShouldBindJSON(&notification); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&notification).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create notification"})
		return
	}

	c.JSON(http.StatusCreated, notification)
}

func GetAllNotificationsHandler(c *gin.Context) {
	var notifications []models.Notifications
	if err := db.DB.Find(&notifications).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get notifications"})
		return
	}

	c.JSON(http.StatusOK, notifications)
}

func GetNotificationHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	var notification models.Notifications
	if err := db.DB.First(&notification, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get notification"})
		return
	}

	c.JSON(http.StatusOK, notification)
}

func UpdateNotificationHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	var notification models.Notifications
	if err := db.DB.First(&notification, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get notification"})
		return
	}

	var input models.Notifications
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Model(&notification).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update notification"})
		return
	}
}

func DeleteNotificationHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	if err := db.DB.Delete(&models.Notifications{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete notification"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "notification deleted successfully"})
}
