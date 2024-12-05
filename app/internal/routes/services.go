package routes

import (
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/handlers"
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupServiceRoutes(router *gin.RouterGroup, serviceHandler *handlers.ServiceHandler) {
	serviceRoutes := router.Group("/services", middleware.JWTMiddleware())
	{
		serviceRoutes.POST("/", serviceHandler.CreateServiceHandler)
		serviceRoutes.GET("/", serviceHandler.GetAllServicesHandler)
		serviceRoutes.GET("/:id", serviceHandler.GetServiceHandler)
		serviceRoutes.PUT("/:id", serviceHandler.UpdateServiceHandler)
		serviceRoutes.DELETE("/:id", serviceHandler.DeleteServiceHandler)
		serviceRoutes.PUT("/:id/deactivate", serviceHandler.DeactivateServiceHandler)
	}
}
