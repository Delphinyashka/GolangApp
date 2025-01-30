package services

import (
	"auth/config"
	"auth/user/models"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func AuthenticateUser(username, password string, c *gin.Context) (string, error) {
	var user models.User

	err := config.GetDB().Where("username = ?", username).First(&user).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return "", fmt.Errorf("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return "", fmt.Errorf("invalid credentials")
	}

	// Generate tokens
	expirationTimeJwt := time.Now().Add(time.Minute * 5)
	expirationTimeRefresh := time.Now().Add(time.Hour * 1)

	jwtTokenString, err := generateToken(username, expirationTimeJwt, []byte(os.Getenv("JWT_TOKEN_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return "", fmt.Errorf("failed to generate token")
	}

	refreshTokenString, err := generateToken(username, expirationTimeRefresh, []byte(os.Getenv("REFRESH_TOKEN_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate refresh token"})
		return "", fmt.Errorf("failed to generate refresh token")
	}

	// Set tokens as cookies
	secure := false
	if gin.Mode() == gin.ReleaseMode {
		secure = true
	}

	c.SetCookie("jwt", jwtTokenString, int(expirationTimeJwt.Sub(time.Now()).Seconds()), "/", "", secure, false)
	c.SetCookie("refresh", refreshTokenString, int(expirationTimeRefresh.Sub(time.Now()).Seconds()), "/", "", secure, true)

	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
	c.Writer.Header().Set("Vary", "Origin")

	c.JSON(http.StatusOK, gin.H{"refresh": expirationTimeRefresh.Unix()})
	return jwtTokenString, nil
}
