package routes

import (
	"auth/user/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	userGroup := router.Group("/user")
	{
		userGroup.POST("/signUp", controllers.SignUp)
		userGroup.POST("/signIn", controllers.SignIn)
		userGroup.POST("/refresh", controllers.Refresh)
		router.GET("/verify", controllers.VerifyToken)
	}
}
