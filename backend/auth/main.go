package main

import (
	"auth/config"
	"auth/routes"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// Initialize the Gin router
	r := gin.Default()

	// Setup CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"}, // Frontend URL
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
	}))

	// Initialize the database connection
	config.InitDB()

	// Setup API routes
	routes.SetupRoutes(r)

	// Running the Gin server
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("could not start the server: %v", err)
	}

	// Printing server start confirmation
	fmt.Println("Server started at http://localhost:8081")
}
