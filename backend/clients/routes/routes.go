package routes

import (
	"clients/client/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	clientGroup := router.Group("/clients")
	{
		clientGroup.GET("/", handlers.GetClients)
	}
}
