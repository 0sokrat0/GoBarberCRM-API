package services

import "github.com/gin-gonic/gin"

func SetupServiceRoutes(router *gin.Engine) {
	serviceRoutes := router.Group("/services")
	{
		serviceRoutes.POST("/", CreateServiceHandler)
		serviceRoutes.GET("/", GetAllServicesHandler)
		serviceRoutes.GET("/:id", GetServiceHandler)
		serviceRoutes.PUT("/:id", UpdateServiceHandler)
		serviceRoutes.DELETE("/:id", DeleteServiceHandler)
		serviceRoutes.PUT("/:id/deactivate", DeactivateServiceHandler)

	}
}
