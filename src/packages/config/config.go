package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	// Load environment variables from .env file
	err := godotenv.Load("../../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// Get the value of an environment variable
func GetEnvVariable(key string) string {
	return os.Getenv(key)
}

// Define configuration variables
var (
	TRUSSIHEALTH_ASSERTION = GetEnvVariable("TRUSSIHEALTH_ASSERTION")
)
