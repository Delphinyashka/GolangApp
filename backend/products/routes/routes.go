package routes

import (
	"products/product/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	productGroup := router.Group("/products")
	{
		productGroup.GET("/", handlers.GetProducts)
	}
}
