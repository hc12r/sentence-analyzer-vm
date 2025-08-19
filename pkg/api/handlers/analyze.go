package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/hc12r/sentence-analyzer-vm/pkg/domain"
)

// HandleAnalyzeSentence handles the sentence analysis endpoint
func HandleAnalyzeSentence(w http.ResponseWriter, r *http.Request) {
	// Only allow POST method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse request body
	var req domain.SentenceAnalysisRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Analyze the sentence
	result := domain.AnalyzeSentence(req.Sentence)

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
