package controllers

import (
	"auth/user/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthRequest defines the expected request payload.
type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// SignUp handles user registration.
func SignUp(c *gin.Context) {
	var req AuthRequest

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

func SignIn(c *gin.Context) {
	var req AuthRequest

	// Parse JSON request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Authenticate user and generate token (errors already handled inside AuthenticateUser)
	_, err := services.AuthenticateUser(req.Username, req.Password, c)
	if err != nil {
		return // Do not set another JSON response
	}

	// Return success response (token is already set in cookie by service)
	//c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

// VerifyToken checks if the JWT is valid
/*func VerifyToken(c *gin.Context) {
	tokenString, err := c.Cookie("token") // Assuming token is stored in "token" cookie
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Parse and validate the JWT
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Authenticated"})
}*/
