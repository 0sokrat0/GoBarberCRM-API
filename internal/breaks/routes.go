package breaks

import "github.com/gin-gonic/gin"

func SetupBreakRoutes(router *gin.Engine) {
	breakRoutes := router.Group("/breaks")
	{
		breakRoutes.POST("/", CreateBreakHandler)
		breakRoutes.GET("/", GetAllBreaksHandler)
		breakRoutes.GET("/:id", GetBreakHandler)
		breakRoutes.PUT("/:id", UpdateBreakHandler)
		breakRoutes.DELETE("/:id", DeleteBreakHandler)
	}
}
