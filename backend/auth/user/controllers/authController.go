package controllers

import (
	"auth/user/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SignUpRequest defines the expected request payload.
type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// SignUp handles user registration.
func SignUp(c *gin.Context) {
	var req SignUpRequest

	// Parse the JSON request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Register user
	err := services.RegisterUser(req.Username, req.Password)
	if err != nil {
		// Return appropriate error response
		if err.Error() == "user already exists" {
			c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Registration failed"})
		}
		return
	}

	// Return 200 OK on success (no token)
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
