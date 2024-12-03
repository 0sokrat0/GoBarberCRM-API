package clients

import "github.com/gin-gonic/gin"

func SetupClientRoutes(router *gin.Engine) {
	clientRoutes := router.Group("/clients")
	{
		clientRoutes.POST("/", CreateClientHandler)
		clientRoutes.GET("/", GetAllClientsHandler)
		clientRoutes.GET("/telegram/:tg_id", GetClientByTelegramIDHandler)
		clientRoutes.GET("/filter", FilterClientsByNameHandler)
		clientRoutes.GET("/:id", GetClientHandler)
		clientRoutes.PUT("/:id", UpdateClientHandler)
		clientRoutes.DELETE("/:id", DeleteClientHandler)
		clientRoutes.POST("/quick_add", QuickAddClientHandler)
		clientRoutes.GET("/search", SearchClientHandler)
		clientRoutes.GET("/check", CheckClientExistenceHandler)
	}
}
