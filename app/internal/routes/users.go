package routes

import (
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/handlers"
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.RouterGroup, userHandler *handlers.UserHandler) {
	userRoutes := router.Group("/users", middleware.JWTMiddleware())
	{
		userRoutes.POST("/", userHandler.CreateUserHandler)
		userRoutes.GET("/", userHandler.GetAllUsersHandler)
		userRoutes.GET("/:id", userHandler.GetUserHandler)
		userRoutes.PUT("/:id", userHandler.UpdateUserHandler)
		userRoutes.DELETE("/:id", userHandler.DeleteUserHandler)
		//userRoutes.POST("/authenticate", userHandler.AuthenticateUserHandler)
	}
}
