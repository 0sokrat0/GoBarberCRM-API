package routes

import (
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/handlers"
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupNotificationRoutes(router *gin.RouterGroup, notificationHandler *handlers.NotificationHandler) {
	notificationRoutes := router.Group("/notifications", middleware.JWTMiddleware())
	{
		notificationRoutes.POST("/", notificationHandler.CreateNotificationHandler)
		notificationRoutes.GET("/", notificationHandler.GetAllNotificationsHandler)
		notificationRoutes.GET("/:id", notificationHandler.GetNotificationHandler)
		notificationRoutes.PUT("/:id", notificationHandler.UpdateNotificationHandler)
		notificationRoutes.DELETE("/:id", notificationHandler.DeleteNotificationHandler)
	}
}
