package bookings

import "github.com/gin-gonic/gin"

func SetupBookingRoutes(router *gin.Engine) {
	bookingRoutes := router.Group("/bookings")
	{
		bookingRoutes.POST("/", CreateBookingHandler)
		bookingRoutes.GET("/", GetAllBookingsHandler)
		bookingRoutes.GET("/:id", GetBookingHandler)
		bookingRoutes.PUT("/:id", UpdateBookingHandler)
		bookingRoutes.DELETE("/:id", DeleteBookingHandler)
		bookingRoutes.GET("/client/:client_id", GetBookingsByClientHandler)
		bookingRoutes.GET("/user/:user_id", GetBookingsByUserHandler)
		bookingRoutes.GET("/service/:service_id", GetBookingsByServiceHandler)
		bookingRoutes.GET("/check", CheckBookingAvailabilityHandler)
	}
}
