package server

import (
	"fmt"
	"net/http"

	"github.com/hc12r/sentence-analyzer-vm/internal/middleware"
	"github.com/hc12r/sentence-analyzer-vm/pkg/api/handlers"
	"github.com/hc12r/sentence-analyzer-vm/pkg/config"
	"github.com/hc12r/sentence-analyzer-vm/pkg/docs"
)

// SetupAndRun configures and starts the HTTP server
func SetupAndRun() error {
	// Load configuration
	cfg := config.LoadConfig()

	// Register handlers without authentication (Kong handles auth)
	http.HandleFunc("/analyze", middleware.NoAuth(handlers.HandleAnalyzeSentence))

	// Register health endpoint without authentication
	http.HandleFunc("/health", handlers.HandleHealth)

	// Register Swagger documentation endpoints
	http.HandleFunc("/swagger", docs.HandleSwaggerUI)
	http.HandleFunc("/swagger/openapi.yaml", docs.HandleSwaggerYAML)

	// Start server
	port := fmt.Sprintf(":%d", cfg.Port)
	fmt.Printf("Server starting on port %s...\n", port)
	return http.ListenAndServe(port, nil)
}
