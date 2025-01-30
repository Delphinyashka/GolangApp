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
var jwtSecret = []byte("your_secret_key")

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

	// Generate JWT token
	tokenString, err := generateJWT(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return "", fmt.Errorf("failed to generate token")
	}

	// Set JWT in HttpOnly cookie
	expiration := time.Now().Add(time.Minute * 5)
	c.SetCookie("jwt", tokenString, int(expiration.Sub(time.Now()).Seconds()), "/", "", true, true)

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})

	return tokenString, nil
}

// generateJWT creates a new JWT token for authenticated users.
func generateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Minute * 5).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// LogoutUser clears the JWT cookie.
func LogoutUser(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "/", "", true, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
