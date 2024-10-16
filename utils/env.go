package utils

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

// GetEnv retrieves an environment variable value by key.
// It loads the environment variables from the .env file if not already loaded.
func GetEnv(key string) string {
    err := godotenv.Load()
    if err != nil {
        log.Printf("Warning: Could not load .env file (using system env variables if available)")
    }

    return os.Getenv(key)
}
