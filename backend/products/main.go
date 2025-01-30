package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"products/config"
	"products/routes"
)

var db *gorm.DB

func main() {
	r := gin.Default()

	routes.SetupRoutes(r)

	config.InitDB()

	if err := r.Run(":8085"); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
