package domain

import (
	"github.com/hc12r/sentence-analyzer-vm/internal/analyzer"
)

// AnalyzeSentence counts words, vowels, and consonants in a sentence
// This function acts as an adapter between the internal analyzer and the public API
func AnalyzeSentence(sentence string) SentenceAnalysisResponse {
	// Use the internal analyzer to perform the analysis
	result := analyzer.AnalyzeSentence(sentence)

	// Convert the internal result to the public response format
	return SentenceAnalysisResponse{
		WordCount:      result.WordCount,
		VowelCount:     result.VowelCount,
		ConsonantCount: result.ConsonantCount,
	}
}
