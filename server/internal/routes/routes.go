package routes

import (
	"task-manager/internal/controllers"
	"task-manager/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	auth := router.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
		auth.POST("/refresh", controllers.RefreshToken)
		auth.POST("/logout", controllers.Logout)
	}

	task := router.Group("/tasks")
	task.Use(middleware.AuthMiddleware())
	{
		task.POST("/", controllers.CreateTask)
		task.GET("/", controllers.GetTasks)
		task.PUT("/:id", controllers.UpdateTask)
		task.DELETE("/:id", controllers.DeleteTask)
	}
}