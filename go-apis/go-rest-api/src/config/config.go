package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	ExternalAPIURL string
	ServerPort     string
	Environment    string
}

func LoadConfig() (*Config, error) {

	projectRoot := findProjectRoot()

	// First load the base .env file
	baseEnvPath := filepath.Join(projectRoot, ".env")
	if err := godotenv.Load(baseEnvPath); err != nil {
		// Only log if file exists but couldn't be loaded
		if !os.IsNotExist(err) {
			log.Printf("Warning: Error loading .env file - %v", err)
		}
	}
	// Get environment from .env file or default to dev
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	// Then load the environment specific file
	// Then load the environment specific file using project root
	envFile := filepath.Join(projectRoot, fmt.Sprintf(".env.%s", env))
	if err := godotenv.Load(envFile); err != nil {
		// Only log if file exists but couldn't be loaded
		if !os.IsNotExist(err) {
			log.Printf("Warning: Error loading %s file - %v", envFile, err)
		}
	}

	// Now get values from environment after both files are loaded
	cfg := &Config{
		ExternalAPIURL: os.Getenv("EXTERNAL_API_URL"),
		ServerPort:     os.Getenv("SERVER_PORT"),
		Environment:    env,
	}

	// Apply defaults only if values are empty
	if cfg.ExternalAPIURL == "" {
		log.Println("EXTERNAL_API_URL not set, using default")
		cfg.ExternalAPIURL = DefaultExternalAPIURL
	}
	if cfg.ServerPort == "" {
		log.Println("SERVER_PORT not set, using default")
		cfg.ServerPort = DefaultServerPort
	}

	// Validate configuration
	if err := cfg.validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	return cfg, nil
}

// Add validation method to Config struct
func (c *Config) validate() error {
	if c.ExternalAPIURL == "" {
		return fmt.Errorf("EXTERNAL_API_URL is required")
	}
	if c.ServerPort == "" {
		return fmt.Errorf("SERVER_PORT is required")
	}
	if c.Environment == "" {
		return fmt.Errorf("Environment is required")
	}
	return nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// Add this helper function to find the project root
func findProjectRoot() string {
	// When running normally, use working directory
	if dir, err := os.Getwd(); err == nil {
		// If we're in the src directory, move up one level
		if filepath.Base(dir) == "src" {
			return filepath.Dir(dir)
		}
		return dir
	}

	// Fallback to executable directory
	if exe, err := os.Executable(); err == nil {
		return filepath.Dir(exe)
	}

	// Last resort: use current directory
	return "."
}
