package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv memuat variabel dari file .env
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// GetEnv mengembalikan nilai environment variable
func GetEnv(key string) string {
	return os.Getenv(key)
}
