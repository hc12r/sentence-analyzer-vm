package config

import (
	"os"
	"testing"
)

func TestLoadConfigDefault(t *testing.T) {
	// Save current environment variable
	oldPort := os.Getenv("PORT")

	// Clean up after the test
	defer func() {
		os.Setenv("PORT", oldPort)
	}()

	// Unset PORT environment variable to test default value
	os.Unsetenv("PORT")

	// Load config
	config := LoadConfig()

	// Check default port value
	if config.Port != 8080 {
		t.Errorf("Expected default port to be 8080, got %d", config.Port)
	}
}

func TestLoadConfigFromEnv(t *testing.T) {
	// Save current environment variable
	oldPort := os.Getenv("PORT")

	// Clean up after the test
	defer func() {
		os.Setenv("PORT", oldPort)
	}()

	// Set PORT environment variable
	os.Setenv("PORT", "9090")

	// Load config
	config := LoadConfig()

	// Check port value from environment
	if config.Port != 9090 {
		t.Errorf("Expected port to be 9090, got %d", config.Port)
	}
}

func TestLoadConfigInvalidPort(t *testing.T) {
	// Save current environment variable
	oldPort := os.Getenv("PORT")

	// Clean up after the test
	defer func() {
		os.Setenv("PORT", oldPort)
	}()

	// Set invalid PORT environment variable
	os.Setenv("PORT", "invalid")

	// Load config
	config := LoadConfig()

	// Check that default port is used when environment variable is invalid
	if config.Port != 8080 {
		t.Errorf("Expected port to be default 8080 when invalid, got %d", config.Port)
	}
}

func TestLoadConfigZeroPort(t *testing.T) {
	// Save current environment variable
	oldPort := os.Getenv("PORT")

	// Clean up after the test
	defer func() {
		os.Setenv("PORT", oldPort)
	}()

	// Set PORT environment variable to zero
	os.Setenv("PORT", "0")

	// Load config
	config := LoadConfig()

	// Check that zero port is accepted
	if config.Port != 0 {
		t.Errorf("Expected port to be 0, got %d", config.Port)
	}
}

func TestLoadConfigNegativePort(t *testing.T) {
	// Save current environment variable
	oldPort := os.Getenv("PORT")

	// Clean up after the test
	defer func() {
		os.Setenv("PORT", oldPort)
	}()

	// Set PORT environment variable to negative value
	os.Setenv("PORT", "-1")

	// Load config
	config := LoadConfig()

	// Check that negative port is accepted (though it's not valid for a real server)
	if config.Port != -1 {
		t.Errorf("Expected port to be -1, got %d", config.Port)
	}
}

func TestConfigStruct(t *testing.T) {
	// Test creating a Config struct directly
	config := Config{
		Port: 7070,
	}

	// Check the value
	if config.Port != 7070 {
		t.Errorf("Expected port to be 7070, got %d", config.Port)
	}
}
