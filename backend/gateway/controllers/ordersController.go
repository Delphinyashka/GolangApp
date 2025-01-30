package controllers

import (
	"gateway/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"strconv"
	"strings"
)

func GetOrders(c *gin.Context) {
	// Retrieve JWT from Authorization header
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
		return
	}

	// Verify JWT token through the auth service
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
	orders, err := services.FetchOrders(orderIDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}

	// Collect unique client IDs from the orders
	clientIDs := make(map[string]bool)
	for _, order := range orders {
		clientIDs[order["client_id"].(string)] = true
	}

	// Fetch clients in a batch from the clients service
	clients, err := services.FetchClientsBatch(clientIDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch clients"})
		return
	}

	// Fetch products based on orders
	productIDs := make(map[string]bool)
	for _, order := range orders {
		productIDs[order["product_id"].(string)] = true
	}
	products, err := services.FetchProductsBatch(productIDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	// Aggregate the data
	aggregatedData := map[string]interface{}{
		"orders":  orders,
		"clients": clients,
		"products": products,
	}

	// Return the aggregated data as a response
	c.JSON(http.StatusOK, aggregatedData)
}
