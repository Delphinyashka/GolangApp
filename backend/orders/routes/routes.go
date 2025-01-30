package routes

import (
	"orders/order/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	clientGroup := router.Group("/clients")
	{
		clientGroup.GET("/", handlers.GetOrders)
	}
}
