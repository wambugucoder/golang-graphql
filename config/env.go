package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// ExtractKey -> Extract Required Key form .env file if present
func ExtractKey(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("No key found for " + key)
	}
	return os.Getenv(key)
}
