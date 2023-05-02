package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Define configuration variables
var (
	EIDAS_PATH                 = ""
	VERIFIABLE_CREDENTIAL_PATH = ""
	SESSIONS_PATH              = ""
	TRUSSIHEALTH_ASSERTION     = ""
	ENCRYPTION_KEY             = ""
	VIDCHAIN_API               = ""
	IPFS_URL                   = ""
	ISSUER_DID                 = ""
	CERTIFICATE_PASSWORD       = ""
	PASSWORD                   = ""
)

func init() {

	// Load environment variables from .env file
	env := os.Getenv("APP_ENV")

	switch env {
	case "local":
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading local .env file")
		}
	case "deployment":
		break
	default: // testing cases
		err := godotenv.Load("../../../.env")
		if err != nil {
			log.Fatal("Error loading test .env file")
		}
	}
	VIDCHAIN_API = GetEnvVariable("VIDCHAIN_API")
	EIDAS_PATH = VIDCHAIN_API + "/api/v1/eidas/signatures"
	VERIFIABLE_CREDENTIAL_PATH = VIDCHAIN_API + "/api/v1/verifiable-credentials"
	SESSIONS_PATH = VIDCHAIN_API + "/api/v1/sessions"
	TRUSSIHEALTH_ASSERTION = GetEnvVariable("TRUSSIHEALTH_ASSERTION")
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
