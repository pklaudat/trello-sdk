package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pklaudat/trello-sdk/internal/server"
)

func LoadSettings() (apiKey, apiToken, hostname string) {
	apiKey, exists := os.LookupEnv("API_KEY")
	if !exists {
		panic("API_KEY environment variable not set")
	}

	apiToken, exists = os.LookupEnv("API_TOKEN")

	if !exists {
		panic("API_TOKEN environment variable not set")
	}

	host, exists := os.LookupEnv("TRELLO_API_ENDPOINT")

	if !exists {
		panic("TRELLO_API_ENDPOINT environment variable not set")
	}

	return apiKey, apiToken, host
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	key, token, host := LoadSettings()

	server.InitializeRoutes(key, token, host)
}
