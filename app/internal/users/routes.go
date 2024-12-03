package users

import "github.com/gin-gonic/gin"

func SetupUserRoutes(router *gin.Engine) {
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", CreateUserHandler)
		userRoutes.GET("/", GetAllUsersHandler)
		userRoutes.GET("/:id", GetUserHandler)
		userRoutes.PUT("/:id", UpdateUserHandler)
		userRoutes.DELETE("/:id", DeleteUserHandler)
	}
}
