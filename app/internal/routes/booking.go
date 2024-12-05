package routes

import (
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/handlers"
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupBookingRoutes(router *gin.RouterGroup, bookingHandler *handlers.BookingHandler) {
	bookingRoutes := router.Group("/bookings", middleware.JWTMiddleware())
	{
		bookingRoutes.POST("/", bookingHandler.CreateBookingHandler)
		bookingRoutes.GET("/", bookingHandler.GetAllBookingsHandler)
		bookingRoutes.GET("/:id", bookingHandler.GetBookingHandler)
		bookingRoutes.PUT("/:id", bookingHandler.UpdateBookingHandler)
		bookingRoutes.DELETE("/:id", bookingHandler.DeleteBookingHandler)
		bookingRoutes.GET("/client/:client_id", bookingHandler.GetBookingsByClientHandler)
		bookingRoutes.GET("/user/:user_id", bookingHandler.GetBookingsByUserHandler)
		bookingRoutes.GET("/service/:service_id", bookingHandler.GetBookingsByServiceHandler)
		bookingRoutes.GET("/availability", bookingHandler.CheckBookingAvailabilityHandler)
	}
}
