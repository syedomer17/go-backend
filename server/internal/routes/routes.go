package routes

import (
	"task-manager/internal/controllers"
	
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	auth := router.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}
}