package database

import (
	"fmt"
	"log"

	"github.com/ThanhTNV/RestfulAPI-GoLang/internal/models"
	"github.com/ThanhTNV/RestfulAPI-GoLang/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB initializes the database connection
func InitDB(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate the schema
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database connection established")
	return db
}
