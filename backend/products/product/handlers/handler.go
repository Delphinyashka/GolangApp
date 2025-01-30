package handlers

import (
	"products/product/models"
	"products/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProducts(c *gin.Context) {
	orderIds := c.DefaultQuery("order_ids", "")
	if orderIds == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Order IDs are required"})
		return
	}

	var clients []models.Product
	if err := config.GetDB().Where("id IN (?)", orderIds).Find(&clients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}

	c.JSON(http.StatusOK, clients)
}
