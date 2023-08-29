package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVars() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
}
