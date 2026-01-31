package routes

import (
	"github.com/ThanhTNV/RestfulAPI-GoLang/internal/api/controllers"
	"github.com/ThanhTNV/RestfulAPI-GoLang/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all application routes
func SetupRoutes(router *gin.Engine, userController *controllers.UserController) {
	// Apply global middleware
	router.Use(middleware.ErrorHandler())

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// User routes
		users := v1.Group("/users")
		{
			users.GET("", userController.GetUsers)
			users.GET("/:id", userController.GetUser)
			users.POST("", userController.CreateUser)
			users.PUT("/:id", userController.UpdateUser)
			users.DELETE("/:id", userController.DeleteUser)
		}
	}
}
