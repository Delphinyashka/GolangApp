package routes

import (
	"products/product/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/products", handlers.GetProducts)
}
