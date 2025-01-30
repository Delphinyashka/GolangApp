package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"products/config"
	"products/product/models"
)

func GetProducts(c *gin.Context) {
	productIds := c.DefaultQuery("order_ids", "")
	if productIds == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product IDs are required"})
		return
	}

	var clients []models.Product
	if err := config.GetDB().Where("id IN (?)", productIds).Find(&clients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	c.JSON(http.StatusOK, clients)
}
