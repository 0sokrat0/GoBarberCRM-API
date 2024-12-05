package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/handlers"
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/middleware"
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/repositories"
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/routes"
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/services"

	"gorm.io/gorm"
)

func SetupRouter(database *gorm.DB) *gin.Engine {
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

	// Initialize repositories
	authRepo := repositories.NewAuthUserRepository(database)
	userRepo := repositories.NewUserRepository(database)
	clientRepo := repositories.NewClientRepository(database)
	bookingRepo := repositories.NewBookingRepository(database)
	serviceRepo := repositories.NewServiceRepository(database)
	scheduleRepo := repositories.NewScheduleRepository(database)
	breakRepo := repositories.NewBreakRepository(database)
	notificationRepo := repositories.NewNotificationRepository(database)

	// Initialize services
	authHandler := handlers.NewAuthHandler(authRepo)
	userService := services.NewUserService(userRepo)
	clientService := services.NewClientService(clientRepo)
	bookingService := services.NewBookingService(bookingRepo)
	serviceService := services.NewServiceService(serviceRepo)
	scheduleService := services.NewScheduleService(scheduleRepo)
	breakService := services.NewBreakService(breakRepo)
	notificationService := services.NewNotificationService(notificationRepo)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userService)
	clientHandler := handlers.NewClientHandler(clientService)
	bookingHandler := handlers.NewBookingHandler(bookingService)
	serviceHandler := handlers.NewServiceHandler(serviceService)
	scheduleHandler := handlers.NewScheduleHandler(scheduleService)
	breakHandler := handlers.NewBreakHandler(breakService)
	notificationHandler := handlers.NewNotificationHandler(notificationService)

	// Public routes (без JWT)
	api := router.Group("/api")
	{
		routes.SetupAuthRoutes(api, authHandler) // Routes for authentication (public)
	}

	// Protected routes (с JWT)
	protected := router.Group("/api")
	protected.Use(middleware.JWTMiddleware()) // Применяем JWT middleware ко всем маршрутам этой группы
	{
		routes.SetupUserRoutes(protected, userHandler)                 // Routes for user management
		routes.SetupClientRoutes(protected, clientHandler)             // Routes for client management
		routes.SetupBookingRoutes(protected, bookingHandler)           // Routes for bookings
		routes.SetupServiceRoutes(protected, serviceHandler)           // Routes for services
		routes.SetupScheduleRoutes(protected, scheduleHandler)         // Routes for schedules
		routes.SetupBreakRoutes(protected, breakHandler)               // Routes for breaks
		routes.SetupNotificationRoutes(protected, notificationHandler) // Routes for notifications
	}

	return router
}
