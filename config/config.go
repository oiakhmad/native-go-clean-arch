package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadConfig loads environment variables from a .env file
func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}
}

// Get retrieves an environment variable or returns a default value
func Get(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
