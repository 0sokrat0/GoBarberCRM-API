package routes

import (
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(router *gin.RouterGroup, authHandler *handlers.AuthHandler) {
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/login", authHandler.LoginHandler)
		authRoutes.POST("/register", authHandler.RegisterHandler)
	}
}
