package routes

import (
	"net/http"

	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/bookings"
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/breaks"
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/clients"
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/notifications"
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/schedules"
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/services"
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/users"
	"github.com/0sokrat0/GoGRAFFApi.git/app/pkg/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Middleware
	router.Use(middleware.RequestLogger())
	router.Use(middleware.CORSMiddleware())

	// Base Routes
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Routes
	users.SetupUserRoutes(router)
	clients.SetupClientRoutes(router)
	bookings.SetupBookingRoutes(router)
	services.SetupServiceRoutes(router)
	schedules.SetupScheduleRoutes(router)
	breaks.SetupBreakRoutes(router)
	notifications.SetupNotificationRoutes(router)

	return router
}
