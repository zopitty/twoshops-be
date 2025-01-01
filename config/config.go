package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found, proceeding with environment variables")
	}
}

func GetGoogleAPIKey() string {
	key := os.Getenv("GOOGLE_API_KEY")
	if key == "" {
		log.Fatal("GOOGLE_API_KEY not set")
	}
	fmt.Println("api key loaded:", key)
	return key
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT not set")
	}
	fmt.Println("port loaded:", port)
	return port
}
