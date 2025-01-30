package routes

import (
	"orders/order/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	ordersGroup := router.Group("/orders")
	{
		ordersGroup.GET("/", handlers.GetOrders)
	}
}
