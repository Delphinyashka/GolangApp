package handlers

import (
	"clients/client/models"
	"clients/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func GetClients(c *gin.Context) {
	clientIds := c.DefaultQuery("ids", "")
	if clientIds == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Client IDs are required"})
		return
	}

	clientIdsStr := strings.Split(clientIds, ",")
	var ids []int
	for _, idStr := range clientIdsStr {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid client ID"})
			return
		}
		ids = append(ids, id)
	}

	var clients []models.Client
	if err := config.GetDB().Where("id IN (?)", ids).Find(&clients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch clients"})
		return
	}

	c.JSON(http.StatusOK, clients)
}
