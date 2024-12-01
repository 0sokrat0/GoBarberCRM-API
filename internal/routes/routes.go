package routes

import (
	"net/http"

	"github.com/0sokrat0/GoGRAFFApi.git/internal/bookings"
	"github.com/0sokrat0/GoGRAFFApi.git/internal/breaks"
	"github.com/0sokrat0/GoGRAFFApi.git/internal/clients"
	"github.com/0sokrat0/GoGRAFFApi.git/internal/notifications"
	"github.com/0sokrat0/GoGRAFFApi.git/internal/schedules"
	"github.com/0sokrat0/GoGRAFFApi.git/internal/services"
	"github.com/0sokrat0/GoGRAFFApi.git/internal/users"
	"github.com/0sokrat0/GoGRAFFApi.git/pkg/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter initializes the Gin router and API routes
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
	setupUserRoutes(router)
	setupClientRoutes(router)
	setupBookingRoutes(router)
	setupServiceRoutes(router)
	setupScheduleRoutes(router)
	setupBreakRoutes(router)
	setupNotificationRoutes(router)

	return router
}

func setupUserRoutes(router *gin.Engine) {
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", users.CreateUserHandler)
		userRoutes.GET("/", users.GetAllUsersHandler)
		userRoutes.GET("/:id", users.GetUserHandler)
		userRoutes.PUT("/:id", users.UpdateUserHandler)
		userRoutes.DELETE("/:id", users.DeleteUserHandler)
	}
}

func setupClientRoutes(router *gin.Engine) {
	clientRoutes := router.Group("/clients")
	{
		clientRoutes.POST("/", clients.CreateClientHandler)
		clientRoutes.GET("/", clients.GetAllClientsHandler)
		clientRoutes.GET("/telegram/:tg_id", clients.GetClientByTelegramIDHandler)
		clientRoutes.GET("/filter", clients.FilterClientsByNameHandler)
		clientRoutes.GET("/:id", clients.GetClientHandler)
		clientRoutes.PUT("/:id", clients.UpdateClientHandler)
		clientRoutes.DELETE("/:id", clients.DeleteClientHandler)
		clientRoutes.POST("/quick_add", clients.QuickAddClientHandler)
		clientRoutes.GET("/search", clients.SearchClientHandler)
		clientRoutes.GET("/check", clients.CheckClientExistenceHandler)
	}
}

func setupBookingRoutes(router *gin.Engine) {
	bookingRoutes := router.Group("/bookings")
	{
		bookingRoutes.POST("/", bookings.CreateBookingHandler)
		bookingRoutes.GET("/", bookings.GetAllBookingsHandler)
		bookingRoutes.GET("/:id", bookings.GetBookingHandler)
		bookingRoutes.PUT("/:id", bookings.UpdateBookingHandler)
		bookingRoutes.DELETE("/:id", bookings.DeleteBookingHandler)
		bookingRoutes.GET("/client/:client_id", bookings.GetBookingsByClientHandler)
		bookingRoutes.GET("/user/:user_id", bookings.GetBookingsByUserHandler)
		bookingRoutes.GET("/service/:service_id", bookings.GetBookingsByServiceHandler)
		bookingRoutes.GET("/check", bookings.CheckBookingAvailabilityHandler)
	}
}

func setupServiceRoutes(router *gin.Engine) {
	serviceRoutes := router.Group("/services")
	{
		serviceRoutes.POST("/", services.CreateServiceHandler)
		serviceRoutes.GET("/", services.GetAllServicesHandler)
		serviceRoutes.GET("/:id", services.GetServiceHandler)
		serviceRoutes.PUT("/:id", services.UpdateServiceHandler)
		serviceRoutes.DELETE("/:id", services.DeleteServiceHandler)
		serviceRoutes.PUT("/:id/deactivate", services.DeactivateServiceHandler)

	}
}

func setupScheduleRoutes(router *gin.Engine) {
	scheduleRoutes := router.Group("/schedules")
	{
		scheduleRoutes.POST("/", schedules.CreateScheduleHandler)
		scheduleRoutes.GET("/", schedules.GetAllSchedulesHandler)
		scheduleRoutes.GET("/:id", schedules.GetScheduleHandler)
		scheduleRoutes.PUT("/:id", schedules.UpdateScheduleHandler)
		scheduleRoutes.DELETE("/:id", schedules.DeleteScheduleHandler)
		scheduleRoutes.GET("/filter/:id", schedules.FilterSchedulesByUserHandler)

	}
}

func setupBreakRoutes(router *gin.Engine) {
	breakRoutes := router.Group("/breaks")
	{
		breakRoutes.POST("/", breaks.CreateBreakHandler)
		breakRoutes.GET("/", breaks.GetAllBreaksHandler)
		breakRoutes.GET("/:id", breaks.GetBreakHandler)
		breakRoutes.PUT("/:id", breaks.UpdateBreakHandler)
		breakRoutes.DELETE("/:id", breaks.DeleteBreakHandler)
	}
}

func setupNotificationRoutes(router *gin.Engine) {
	notificationRoutes := router.Group("/notifications")
	{
		notificationRoutes.POST("/", notifications.CreateNotificationHandler)
		notificationRoutes.GET("/", notifications.GetAllNotificationsHandler)
		notificationRoutes.GET("/:id", notifications.GetNotificationHandler)
		notificationRoutes.PUT("/:id", notifications.UpdateNotificationHandler)
		notificationRoutes.DELETE("/:id", notifications.DeleteNotificationHandler)
	}
}
