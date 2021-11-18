package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

// Config function to get env value
func Config(key string) string {
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error : Error load .env file")
	}
	// Return the value of variable from .env file
	return os.Getenv(key)
}
