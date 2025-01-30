package handlers

import (
	"clients/client/models"
	"clients/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetClients(c *gin.Context) {
	clientIds := c.DefaultQuery("client_ids", "")
	if clientIds == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Client IDs are required"})
		return
	}

	var clients []models.Client
	if err := config.GetDB().Where("id IN (?)", clientIds).Find(&clients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch clients"})
		return
	}

	c.JSON(http.StatusOK, clients)
}
