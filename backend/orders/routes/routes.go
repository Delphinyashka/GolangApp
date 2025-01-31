package routes

import (
	"orders/order/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/orders", handlers.GetOrders)
}
