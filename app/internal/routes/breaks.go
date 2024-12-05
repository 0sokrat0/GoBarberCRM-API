package routes

import (
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/handlers"
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupBreakRoutes(router *gin.RouterGroup, breakHandler *handlers.BreakHandler) {
	breakRoutes := router.Group("/breaks", middleware.JWTMiddleware())
	{
		breakRoutes.POST("/", breakHandler.CreateBreakHandler)
		breakRoutes.GET("/", breakHandler.GetAllBreaksHandler)
		breakRoutes.GET("/:id", breakHandler.GetBreakHandler)
		breakRoutes.PUT("/:id", breakHandler.UpdateBreakHandler)
		breakRoutes.DELETE("/:id", breakHandler.DeleteBreakHandler)
	}
}
