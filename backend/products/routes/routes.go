package routes

import (
	"products/product/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	clientGroup := router.Group("/clients")
	{
		clientGroup.GET("/", handlers.GetProducts)
	}
}
