package main

import (
	"auth/config"
	"auth/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"}, // Frontend URL
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	config.InitDB()

	routes.SetupRoutes(r)

	if err := r.Run(":8081"); err != nil {
		log.Fatalf("could not start the server: %v", err)
	}
}
