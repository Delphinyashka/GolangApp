package main

import (
	"clients/config"
	"log"
	"clients/routes"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func main() {
	r := gin.Default()

	routes.SetupRoutes(r)

	config.InitDB()

	if err := r.Run(":8083"); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
