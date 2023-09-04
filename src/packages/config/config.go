package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Define configuration variables
var (
	VC_API               = ""
	ESEAL_API            = ""
	ENCRYPTION_KEY       = ""
	IPFS_URL             = ""
	ISSUER_DID           = ""
	CERTIFICATE_PASSWORD = ""
	PASSWORD             = ""
)

func init() {

	// Load environment variables from .env file
	env := os.Getenv("APP_ENV")

	switch env {
	case "deployment":
		break
	default: // testing cases
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading test .env file")
		}
	}
	VC_API = GetEnvVariable("VC_API")
	ESEAL_API = GetEnvVariable("ESEAL_API")
	ENCRYPTION_KEY = GetEnvVariable("ENCRYPTION_KEY")
	IPFS_URL = GetEnvVariable("IPFS_URL")
	ISSUER_DID = GetEnvVariable("ISSUER_DID")
	CERTIFICATE_PASSWORD = GetEnvVariable("CERTIFICATE_PASSWORD")
	PASSWORD = GetEnvVariable("PASSWORD")
}

// Get the value of an environment variable
func GetEnvVariable(key string) string {
	return os.Getenv(key)
}
