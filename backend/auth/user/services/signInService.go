package services

import (
	"auth/config"
	"auth/user/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Secret key for signing JWTs (move this to an environment variable in production)
var jwtSecret = []byte("SQX1234567")

// AuthenticateUser checks credentials and returns a JWT token if valid.
func AuthenticateUser(username, password string, c *gin.Context) (string, error) {
	var user models.User

	// Retrieve user by username
	err := config.GetDB().Where("username = ?", username).First(&user).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return "", fmt.Errorf("user not found")
	}

	// Compare hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return "", fmt.Errorf("invalid credentials")
	}

	// Set expiration time for 7 days
	expiration := time.Now().Add(time.Minute * 5) // 7 days expiration

	// Generate JWT token
	tokenString, err := generateJWT(username, expiration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return "", fmt.Errorf("failed to generate token")
	}

	// Set JWT in HttpOnly cookie
	secure := false
	if gin.Mode() == gin.ReleaseMode {
		// If in production (secure environment), set Secure flag to true
		secure = true
	}
	c.SetCookie("jwt", tokenString, int(expiration.Sub(time.Now()).Seconds()), "/", "", secure, true)
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true") // Allow credentials
	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // Your frontend domain here

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})

	return tokenString, nil
}

// generateJWT creates a new JWT token for authenticated users.
func generateJWT(username string, expiration time.Time) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      expiration.Unix(), // Use the same expiration time
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// LogoutUser clears the JWT cookie.
func LogoutUser(c *gin.Context) {
	// Clear the JWT cookie by setting its expiration to a past date
	c.SetCookie("jwt", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
