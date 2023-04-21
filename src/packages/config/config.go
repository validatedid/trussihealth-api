package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
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
	TRUSSIHEALTH_ASSERTION     = GetEnvVariable("TRUSSIHEALTH_ASSERTION")
	ENCRYPTION_KEY             = GetEnvVariable("ENCRYPTION_KEY")
	VIDCHAIN_API               = GetEnvVariable("VIDCHAIN_API")
	EIDAS_PATH                 = VIDCHAIN_API + "/api/v1/eidas/signatures"
	VERIFIABLE_CREDENTIAL_PATH = VIDCHAIN_API + "/api/v1/verifiable-credentials"
	SESSIONS_PATH              = VIDCHAIN_API + "/api/v1/sessions"
	IPFS_URL                   = GetEnvVariable("IPFS_URL")
)
