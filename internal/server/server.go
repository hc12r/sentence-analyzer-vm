package server

import (
	"fmt"
	"net/http"

	"github.com/hc12r/sentence-analyzer-vm/internal/middleware"
	"github.com/hc12r/sentence-analyzer-vm/pkg/api/handlers"
	"github.com/hc12r/sentence-analyzer-vm/pkg/config"
	"github.com/hc12r/sentence-analyzer-vm/pkg/docs"
)

// SetupRoutes configures all the routes for the HTTP server
func SetupRoutes() {
	// Register login endpoint without authentication
	http.HandleFunc("/login", handlers.HandleLogin)

	// Register handlers with JWT authentication
	http.HandleFunc("/analyze", middleware.JWTAuth(handlers.HandleAnalyzeSentence))

	// Register health endpoint without authentication
	http.HandleFunc("/health", handlers.HandleHealth)

	// Register Swagger documentation endpoints
	http.HandleFunc("/swagger", docs.HandleSwaggerUI)
	http.HandleFunc("/swagger/openapi.yaml", docs.HandleSwaggerYAML)
}

// SetupAndRun configures and starts the HTTP server
func SetupAndRun() error {
	// Load configuration
	cfg := config.LoadConfig()

	// Setup routes
	SetupRoutes()

	// Start server
	port := fmt.Sprintf(":%d", cfg.Port)
	fmt.Printf("Server starting on port %s...\n", port)
	return http.ListenAndServe(port, nil)
}
