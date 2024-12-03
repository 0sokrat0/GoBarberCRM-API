package notifications

import "github.com/gin-gonic/gin"

func SetupNotificationRoutes(router *gin.Engine) {
	notificationRoutes := router.Group("/notifications")
	{
		notificationRoutes.POST("/", CreateNotificationHandler)
		notificationRoutes.GET("/", GetAllNotificationsHandler)
		notificationRoutes.GET("/:id", GetNotificationHandler)
		notificationRoutes.PUT("/:id", UpdateNotificationHandler)
		notificationRoutes.DELETE("/:id", DeleteNotificationHandler)
	}
}
