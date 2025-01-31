package controllers

import (
	"gateway/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

// GenerateOrderIDs generates a comma-separated string of order IDs based on the page and limit.
func GenerateOrderIDs(page int, limit int) string {
	start := (page - 1) * limit
	end := start + limit

	// Generate order IDs string (e.g., for page 1 and limit 10: "1,2,3,4,...,10")
	orderIDs := make([]string, 0, limit)
	for i := start + 1; i <= end; i++ {
		orderIDs = append(orderIDs, strconv.Itoa(i))
	}

	// Convert the order IDs slice to a string, separated by commas
	return strings.Join(orderIDs, ",")
}

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

	// Retrieve pagination from query parameters
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	// Generate order IDs string based on pagination
	orderIDsStr := GenerateOrderIDs(page, limit)

	// Fetch orders data from orders service
	orders, err := services.FetchOrders(strings.Split(orderIDsStr, ","), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}

	if len(orders) == 0 {
		c.JSON(http.StatusOK, gin.H{"orders": []interface{}{}, "total": 0})
		return
	}

	// Collect unique client IDs and product IDs from the orders
	clientIDs := make(map[string]bool)
	productIDs := make(map[string]bool)
	for _, order := range orders {
		if clientID, ok := order["client_id"].(float64); ok { // Handle float64 case
			clientIDs[strconv.Itoa(int(clientID))] = true
		} else if clientIDStr, ok := order["client_id"].(string); ok {
			clientIDs[clientIDStr] = true
		}

		if productID, ok := order["product_id"].(float64); ok { // Handle float64 case
			productIDs[strconv.Itoa(int(productID))] = true
		} else if productIDStr, ok := order["product_id"].(string); ok {
			productIDs[productIDStr] = true
		}
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

	ordersTotal, err := services.CalculateTotalAmount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to calculate total amount"})
		return
	}

	// Transform orders to match frontend expectations
	var formattedOrders []map[string]interface{}
	for _, order := range orders {
		// Ensure proper type assertion for client_id and product_id
		var clientID, productID float64
		if id, ok := order["client_id"].(float64); ok {
			clientID = id // Directly use float64
		}

		if id, ok := order["product_id"].(float64); ok {
			productID = id // Directly use float64
		}

		var clientName, productName string
		var price float64

		// Find the client and product by ID
		for _, client := range clients {
			if client["id"].(float64) == clientID {
				clientName = client["name"].(string)
				break
			}
		}

		for _, product := range products {
			if product["id"].(float64) == productID {
				productName = product["name"].(string)
				price = product["price"].(float64)
				break
			}
		}

		// Append the formatted order
		formattedOrders = append(formattedOrders, map[string]interface{}{
			"id":          order["id"],
			"productName": productName,
			"clientName":  clientName,
			"price":       price,
		})
	}

	// Return the aggregated data
	c.JSON(http.StatusOK, gin.H{
		"orders": formattedOrders,
		"ordersTotal":  ordersTotal,
	})
}