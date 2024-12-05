package routes

import (
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/handlers"
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupClientRoutes(router *gin.RouterGroup, clientHandler *handlers.ClientHandler) {
	clientRoutes := router.Group("/clients", middleware.JWTMiddleware())
	{
		clientRoutes.POST("/", clientHandler.CreateClientHandler)
		clientRoutes.GET("/", clientHandler.GetAllClientsHandler)
		clientRoutes.GET("/telegram/:tg_id", clientHandler.GetClientByTelegramIDHandler)
		clientRoutes.GET("/filter", clientHandler.FilterClientsByNameHandler)
		clientRoutes.GET("/:id", clientHandler.GetClientHandler)
		clientRoutes.PUT("/:id", clientHandler.UpdateClientHandler)
		clientRoutes.DELETE("/:id", clientHandler.DeleteClientHandler)
		clientRoutes.POST("/quick_add", clientHandler.QuickAddClientHandler)
		clientRoutes.GET("/search", clientHandler.SearchClientHandler)
		clientRoutes.GET("/check", clientHandler.CheckClientExistenceHandler)
	}
}
