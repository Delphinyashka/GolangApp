package routes

import (
	"gateway/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	apiGroup := router.Group("/api")
	{
		// Endpoint to fetch orders with pagination and order IDs
		apiGroup.GET("/orders", controllers.GetOrders)
	}
}
