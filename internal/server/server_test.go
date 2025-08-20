package server

import (
	"net/http"
	"reflect"
	"testing"
	"time"
)

// TestSetupRoutes tests that all routes are registered correctly
func TestSetupRoutes(t *testing.T) {
	// Save the original routes
	originalServeMux := http.DefaultServeMux

	// Create a new ServeMux for testing
	http.DefaultServeMux = http.NewServeMux()

	// Restore the original ServeMux when the test completes
	defer func() {
		http.DefaultServeMux = originalServeMux
	}()

	// Call the function under test
	SetupRoutes()

	// Get all registered routes
	registeredRoutes := make(map[string]bool)

	// Use reflection to access the internal map of routes in the DefaultServeMux
	// This is a bit hacky but necessary since there's no public API to get registered routes
	muxValue := reflect.ValueOf(http.DefaultServeMux).Elem()
	muxMap := muxValue.FieldByName("m")

	if muxMap.IsValid() {
		// If we can access the map, check each route
		for _, k := range muxMap.MapKeys() {
			route := k.String()
			registeredRoutes[route] = true
		}
	} else {
		// If we can't access the map (Go version changes, etc.),
		// at least check that the routes respond to requests
		expectedRoutes := []string{
			"/analyze",
			"/health",
			"/swagger",
			"/swagger/openapi.yaml",
		}

		for _, route := range expectedRoutes {
			// Create a test request
			req, err := http.NewRequest("GET", route, nil)
			if err != nil {
				t.Fatalf("Failed to create request for %s: %v", route, err)
			}

			// Check if a handler is registered for this route
			_, pattern := http.DefaultServeMux.Handler(req)
			if pattern == "" {
				t.Errorf("No handler registered for route %s", route)
			} else {
				registeredRoutes[pattern] = true
			}
		}
	}

	// Check that all expected routes were registered
	expectedRoutes := []string{
		"/analyze",
		"/health",
		"/swagger",
		"/swagger/openapi.yaml",
	}

	for _, route := range expectedRoutes {
		if !registeredRoutes[route] {
			t.Errorf("Route %s was not registered", route)
		}
	}
}

// TestSetupAndRunPortConfiguration tests that the port is configured correctly
func TestSetupAndRunPortConfiguration(t *testing.T) {
	// This is a limited test since we can't easily mock http.ListenAndServe
	// We're just testing that the function doesn't panic

	// Save the original DefaultServeMux
	originalServeMux := http.DefaultServeMux

	// Create a new ServeMux for testing
	http.DefaultServeMux = http.NewServeMux()

	// Restore the original ServeMux when the test completes
	defer func() {
		http.DefaultServeMux = originalServeMux
	}()

	// Create a channel to catch panics
	done := make(chan bool)

	// Run the function in a goroutine
	go func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("SetupAndRun panicked: %v", r)
			}
			done <- true
		}()

		// We expect this to block forever, so we'll time out
		// We're just checking that it doesn't panic before starting the server
		SetupAndRun()
	}()

	// Wait a short time to allow any panics to occur
	select {
	case <-done:
		// If we get here, the function returned, which is unexpected
		t.Error("SetupAndRun returned unexpectedly")
	case <-time.After(100 * time.Millisecond):
		// This is expected - the function should block on http.ListenAndServe
	}
}
