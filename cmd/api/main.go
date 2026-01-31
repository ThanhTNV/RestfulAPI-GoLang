package main

import (
	"log"

	"github.com/ThanhTNV/RestfulAPI-GoLang/internal/api/controllers"
	"github.com/ThanhTNV/RestfulAPI-GoLang/internal/api/routes"
	"github.com/ThanhTNV/RestfulAPI-GoLang/internal/repository"
	"github.com/ThanhTNV/RestfulAPI-GoLang/internal/service"
	"github.com/ThanhTNV/RestfulAPI-GoLang/pkg/config"
	"github.com/ThanhTNV/RestfulAPI-GoLang/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	db := database.InitDB(cfg)

	// Dependency Injection: Wire up the layers
	// Repository layer
	userRepo := repository.NewUserRepository(db)

	// Service layer
	userService := service.NewUserService(userRepo)

	// Controller layer
	userController := controllers.NewUserController(userService)

	// Initialize Gin router
	router := gin.Default()

	// Setup routes
	routes.SetupRoutes(router, userController)

	// Start server
	log.Printf("Server starting on port %s", cfg.ServerPort)
	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
