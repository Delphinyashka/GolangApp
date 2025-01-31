package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"products/config"
	"products/product/models"
	"strconv"
	"strings"
)

func GetProducts(c *gin.Context) {
	productIds := c.DefaultQuery("ids", "")
	if productIds == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product IDs are required"})
		return
	}

	productIdsStr := strings.Split(productIds, ",")
	var ids []int
	for _, idStr := range productIdsStr {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
			return
		}
		ids = append(ids, id)
	}

	var Products []models.Product
	if err := config.GetDB().Where("id IN (?)", ids).Find(&Products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	c.JSON(http.StatusOK, Products)
}
