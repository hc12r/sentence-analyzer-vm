package config

import (
	"os"
	"strconv"
)

// Config holds all configuration for the application
type Config struct {
	Port int
}

// LoadConfig loads configuration from environment variables with defaults
func LoadConfig() Config {
	config := Config{
		Port: 8080, // Default port
	}

	// Override with environment variables if set
	if portStr := os.Getenv("PORT"); portStr != "" {
		if port, err := strconv.Atoi(portStr); err == nil {
			config.Port = port
		}
	}

	return config
}
