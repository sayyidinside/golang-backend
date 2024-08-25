package utils

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadENV() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	envPath := filepath.Join(pwd, "./.env")

	err = godotenv.Load(filepath.Join(envPath))
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func LoadENVTest() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	envPath := filepath.Join(pwd, "../.env")

	err = godotenv.Load(filepath.Join(envPath))
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

// Function to abstarct the process of getting env data that also set default value
// if the env attribute empty or doesn't exist
func GetENVWithDefault(key string, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}
