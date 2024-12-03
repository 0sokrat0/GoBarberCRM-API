package schedules

import "github.com/gin-gonic/gin"

func SetupScheduleRoutes(router *gin.Engine) {
	scheduleRoutes := router.Group("/schedules")
	{
		scheduleRoutes.POST("/", CreateScheduleHandler)
		scheduleRoutes.GET("/", GetAllSchedulesHandler)
		scheduleRoutes.GET("/:id", GetScheduleHandler)
		scheduleRoutes.PUT("/:id", UpdateScheduleHandler)
		scheduleRoutes.DELETE("/:id", DeleteScheduleHandler)
		scheduleRoutes.GET("/filter/:id", FilterSchedulesByUserHandler)

	}
}
