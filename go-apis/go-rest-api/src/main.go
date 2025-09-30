package main

import (
	"go-rest-api/src/config"
	"go-rest-api/src/routes"
	"log"
	"net/http"
)

func main() {

	// Load configuration at startup
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Print configuration details
	log.Printf("Configuration loaded:")
	log.Printf("Environment: %s", cfg.Environment)
	log.Printf("Server Port: %s", cfg.ServerPort)
	log.Printf("External API URL: %s", cfg.ExternalAPIURL)

	router := routes.SetupRoutes(cfg)
	log.Printf("Starting server on port %s", cfg.ServerPort)
	if err := http.ListenAndServe(":"+cfg.ServerPort, router); err != nil {
		log.Fatal(err)
	}
}
