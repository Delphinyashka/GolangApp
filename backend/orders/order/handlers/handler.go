package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"orders/config"
	"orders/order/models"
	"strconv"
	"strings"
)

func GetOrders(c *gin.Context) {
	orderIds := c.DefaultQuery("ids", "")
	if orderIds == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Order IDs are required"})
		return
	}

	orderIdsStr := strings.Split(orderIds, ",")
	var ids []int
	for _, idStr := range orderIdsStr {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
			return
		}
		ids = append(ids, id)
	}

	var orders []models.Order
	if err := config.GetDB().Where("id IN (?)", ids).Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func GetTotalAmount(c *gin.Context) {
	var result struct {
		TotalAmount int64 `json:"totalAmount"` // Changed the type to int64
	}

	if err := config.GetDB().Model(&models.Order{}).
		Select("COUNT(*) as totalAmount").
		Scan(&result).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to calculate total amount"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"totalAmount": result.TotalAmount,
	})
}
