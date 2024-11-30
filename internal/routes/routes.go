package routes

import (
	"github.com/0sokrat0/GoGRAFFApi.git/internal/bookings"
	"github.com/0sokrat0/GoGRAFFApi.git/internal/breaks"
	"github.com/0sokrat0/GoGRAFFApi.git/internal/clients"
	"github.com/0sokrat0/GoGRAFFApi.git/internal/notifications"
	"github.com/0sokrat0/GoGRAFFApi.git/internal/schedules"
	"github.com/0sokrat0/GoGRAFFApi.git/internal/services"
	"github.com/0sokrat0/GoGRAFFApi.git/internal/users"
	"github.com/gin-gonic/gin"
)

// SetupRouter создает маршруты для приложения
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Маршрут для проверки работоспособности API
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Маршруты для пользователей
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", users.CreateUserHandler)      // Создание пользователя
		userRoutes.GET("/", users.GetAllUsersHandler)      // Получение всех пользователей
		userRoutes.GET("/:id", users.GetUserHandler)       // Получение пользователя по ID
		userRoutes.PUT("/:id", users.UpdateUserHandler)    // Обновление пользователя
		userRoutes.DELETE("/:id", users.DeleteUserHandler) // Удаление пользователя
	}

	clientsRoutes := router.Group("/clients") // Группа маршрутов для клиентов
	{
		clientsRoutes.POST("/", clients.CreateClientHandler)      // Создание клиента
		clientsRoutes.GET("/", clients.GetAllClientsHandler)      // Получение всех клиентов
		clientsRoutes.GET("/:id", clients.GetClientHandler)       // Получение клиента по ID
		clientsRoutes.PUT("/:id", clients.UpdateClientHandler)    // Обновление клиента
		clientsRoutes.DELETE("/:id", clients.DeleteClientHandler) // Удаление клиента
	}

	bookingsRoutes := router.Group("/bookings") // Группа маршрутов для бронирований
	{
		bookingsRoutes.POST("/", bookings.CreateBookingHandler)      // Создание бронирования
		bookingsRoutes.GET("/", bookings.GetAllBookingsHandler)      // Получение всех бронирований
		bookingsRoutes.GET("/:id", bookings.GetBookingHandler)       // Получение бронирования по ID
		bookingsRoutes.PUT("/:id", bookings.UpdateBookingHandler)    // Обновление бронирования
		bookingsRoutes.DELETE("/:id", bookings.DeleteBookingHandler) // Удаление бронирования
	}
	servicesRoutes := router.Group("/services") // Группа маршрутов для сервисов
	{
		servicesRoutes.POST("/", services.CreateServiceHandler)      // Создание сервиса
		servicesRoutes.GET("/", services.GetAllServicesHandler)      // Получение всех сервисов
		servicesRoutes.GET("/:id", services.GetServiceHandler)       // Получение сервиса по ID
		servicesRoutes.PUT("/:id", services.UpdateServiceHandler)    // Обновление сервиса
		servicesRoutes.DELETE("/:id", services.DeleteServiceHandler) // Удаление сервиса
	}

	schedulesRoutes := router.Group("/schedules") // Группа маршрутов для расписания
	{
		schedulesRoutes.POST("/", schedules.CreateScheduleHandler)      // Создание расписания
		schedulesRoutes.GET("/", schedules.GetAllSchedulesHandler)      // Получение всех расписаний
		schedulesRoutes.GET("/:id", schedules.GetScheduleHandler)       // Получение расписания по ID
		schedulesRoutes.PUT("/:id", schedules.UpdateScheduleHandler)    // Обновление расписания
		schedulesRoutes.DELETE("/:id", schedules.DeleteScheduleHandler) // Удаление расписания
	}
	breaksRoutes := router.Group("/breaks") // Группа маршрутов для перерывов
	{
		breaksRoutes.POST("/", breaks.CreateBreakHandler)      // Создание перерыва
		breaksRoutes.GET("/", breaks.GetAllBreaksHandler)      // Получение всех перерывов
		breaksRoutes.GET("/:id", breaks.GetBreakHandler)       // Получение перерыва по ID
		breaksRoutes.PUT("/:id", breaks.UpdateBreakHandler)    // Обновление перерыва
		breaksRoutes.DELETE("/:id", breaks.DeleteBreakHandler) // Удаление перерыва
	}
	notificationsRoutes := router.Group("/notifications") // Группа маршрутов для уведомлений клиентов
	{
		notificationsRoutes.POST("/", notifications.CreateNotificationHandler)      // Создание уведомления
		notificationsRoutes.GET("/", notifications.GetAllNotificationsHandler)      // Получение всех уведомлений
		notificationsRoutes.GET("/:id", notifications.GetNotificationHandler)       // Получение уведомления по ID
		notificationsRoutes.PUT("/:id", notifications.UpdateNotificationHandler)    // Обновление уведомления
		notificationsRoutes.DELETE("/:id", notifications.DeleteNotificationHandler) // Удаление уведомления
	}

	// Добавляй другие группы маршрутов (clients, bookings и т.д.)
	return router
}
