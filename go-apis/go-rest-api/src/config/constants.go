package config

// Environment variable keys
const (
	EnvExternalAPIURL = "EXTERNAL_API_URL"
	// Add other environment variables here
	EnvServerPort  = "SERVER_PORT"
	EnvEnvironment = "APP_ENV"
)

// Default values for environment variables
const (
	DefaultExternalAPIURL = "http://153.58.140.40:8080"
	DefaultServerPort     = "8080"
	DefaultEnvironment    = "development"
)
