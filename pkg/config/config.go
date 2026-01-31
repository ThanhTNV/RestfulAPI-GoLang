package config

import "os"

// Config holds application configuration
type Config struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	ServerPort string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	return &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres"),
		DBName:     getEnv("DB_NAME", "restapi"),
		DBPort:     getEnv("DB_PORT", "5432"),
		ServerPort: getEnv("PORT", "8080"),
	}
}

// getEnv gets an environment variable with a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
