package routes

import (
	"clients/client/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/clients", handlers.GetClients)
}
