package routes

import (
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/handlers"
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupScheduleRoutes(router *gin.RouterGroup, scheduleHandler *handlers.ScheduleHandler) {
	scheduleRoutes := router.Group("/schedules", middleware.JWTMiddleware())
	{
		scheduleRoutes.POST("/", scheduleHandler.CreateScheduleHandler)
		scheduleRoutes.GET("/", scheduleHandler.GetAllSchedulesHandler)
		scheduleRoutes.GET("/:id", scheduleHandler.GetScheduleHandler)
		scheduleRoutes.PUT("/:id", scheduleHandler.UpdateScheduleHandler)
		scheduleRoutes.DELETE("/:id", scheduleHandler.DeleteScheduleHandler)
		scheduleRoutes.GET("/filter", scheduleHandler.FilterSchedulesByUserHandler)
	}
}
