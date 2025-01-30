package routes

import (
	"auth/user/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures the API routes.
func SetupRoutes(router *gin.Engine) {
	userGroup := router.Group("/user")
	{
		userGroup.POST("/signUp", controllers.SignUp)
	}
}
