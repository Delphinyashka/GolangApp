package controllers

import (
	"gateway/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func GetOrders(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
		return
	}

	isValid, err := services.VerifyJWTToken(token)

	if err != nil || !isValid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return
	}

	// Retrieve pagination and order IDs from query parameters
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	orderIDsStr := c.DefaultQuery("orderIds", "")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)
	orderIDs := strings.Split(orderIDsStr, ",")

	// Fetch orders data from orders service
	orders, err := services.FetchOrders(orderIDs, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}

	if len(orders) == 0 {
		c.JSON(http.StatusOK, gin.H{"orders": []interface{}{}, "total": 0})
		return
	}

	// Collect unique client IDs from the orders
	clientIDs := make(map[string]bool)
	productIDs := make(map[string]bool)
	for _, order := range orders {
		clientIDs[order["client_id"].(string)] = true
		productIDs[order["product_id"].(string)] = true
	}

	// Fetch clients and products
	clients, err := services.FetchClientsBatch(clientIDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch clients"})
		return
	}
	products, err := services.FetchProductsBatch(productIDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	// Create lookup maps for easy reference
	clientMap := make(map[string]interface{})
	productMap := make(map[string]interface{})
	for _, client := range clients {
		clientMap[client["id"].(string)] = client
	}
	for _, product := range products {
		productMap[product["id"].(string)] = product
	}

	// Transform orders to match frontend expectations
	var formattedOrders []map[string]interface{}
	for _, order := range orders {
		client := clientMap[order["client_id"].(string)]
		product := productMap[order["product_id"].(string)]

		formattedOrders = append(formattedOrders, map[string]interface{}{
			"id":          order["id"],
			"productName": product.(map[string]interface{})["name"],
			"clientName":  client.(map[string]interface{})["name"],
			"price":       product.(map[string]interface{})["price"],
		})
	}

	// Return the aggregated data
	c.JSON(http.StatusOK, gin.H{
		"orders": formattedOrders,
		"total":  len(orders),
	})
}
