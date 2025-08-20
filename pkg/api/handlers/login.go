package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/hc12r/sentence-analyzer-vm/pkg/auth"
)

// LoginRequest represents the login request body
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse represents the login response body
type LoginResponse struct {
	Token string `json:"token"`
}

// HandleLogin handles the login endpoint
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	// Only allow POST method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse request body
	var req LoginRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Get credentials from environment variables
	expectedUsername := os.Getenv("LOGIN_USERNAME")
	if expectedUsername == "" {
		expectedUsername = "admin" // Default username
	}

	expectedPassword := os.Getenv("LOGIN_PASSWORD")
	if expectedPassword == "" {
		expectedPassword = "password" // Default password
	}

	// Validate credentials
	if req.Username != expectedUsername || req.Password != expectedPassword {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token, err := auth.GenerateToken(req.Username, []string{"user"})
	if err != nil {
		log.Printf("Error generating token: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Write response
	response := LoginResponse{
		Token: token,
	}
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
