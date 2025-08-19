package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// SentenceAnalysisRequest represents the request body
type SentenceAnalysisRequest struct {
	Sentence string `json:"sentence"`
}

// SentenceAnalysisResponse represents the response body
type SentenceAnalysisResponse struct {
	WordCount     int `json:"word_count"`
	VowelCount    int `json:"vowel_count"`
	ConsonantCount int `json:"consonant_count"`
}

// Credentials for basic authentication
type Credentials struct {
	Username string
	Password string
}

// analyzeSentence counts words, vowels, and consonants in a sentence
func analyzeSentence(sentence string) SentenceAnalysisResponse {
	// Count words
	words := strings.Fields(sentence)
	wordCount := len(words)
	
	// Count vowels and consonants
	vowelCount := 0
	consonantCount := 0
	
	for _, char := range strings.ToLower(sentence) {
		if char >= 'a' && char <= 'z' {
			if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
				vowelCount++
			} else {
				consonantCount++
			}
		}
	}
	
	return SentenceAnalysisResponse{
		WordCount:     wordCount,
		VowelCount:    vowelCount,
		ConsonantCount: consonantCount,
	}
}

// basicAuth middleware for authentication
func basicAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Check if it's Basic auth
		if !strings.HasPrefix(authHeader, "Basic ") {
			http.Error(w, "Invalid authentication method", http.StatusUnauthorized)
			return
		}

		// Decode the base64 encoded credentials
		payload, err := base64.StdEncoding.DecodeString(authHeader[6:])
		if err != nil {
			http.Error(w, "Invalid credentials format", http.StatusUnauthorized)
			return
		}

		// Split username and password
		pair := strings.SplitN(string(payload), ":", 2)
		if len(pair) != 2 {
			http.Error(w, "Invalid credentials format", http.StatusUnauthorized)
			return
		}

		// In a real application, you would check against a database or secure storage
		// For this example, we use hardcoded credentials
		validCredentials := Credentials{
			Username: "admin",
			Password: "password",
		}

		// Validate credentials
		if pair[0] != validCredentials.Username || pair[1] != validCredentials.Password {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// If authentication is successful, call the next handler
		next(w, r)
	}
}

// handleAnalyzeSentence handles the sentence analysis endpoint
func handleAnalyzeSentence(w http.ResponseWriter, r *http.Request) {
	// Only allow POST method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	// Parse request body
	var req SentenceAnalysisRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	
	// Analyze the sentence
	result := analyzeSentence(req.Sentence)
	
	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	// Write response
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(result); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// handleHealth handles the health check endpoint
func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}

func main() {
	// Register handlers with basic authentication
	http.HandleFunc("/analyze", basicAuth(handleAnalyzeSentence))
	
	// Register health endpoint without authentication
	http.HandleFunc("/health", handleHealth)
	
	// Start server
	port := ":8080"
	fmt.Printf("Server starting on port %s...\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}