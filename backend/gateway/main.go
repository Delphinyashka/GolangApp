package main

import (
	"gateway/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()

	routes.SetupRoutes(router)

	if err := router.Run(":8082"); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
