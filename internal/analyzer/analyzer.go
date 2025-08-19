package analyzer

import (
	"strings"
)

// AnalyzeSentence counts words, vowels, and consonants in a sentence
func AnalyzeSentence(sentence string) SentenceAnalysisResult {
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

	return SentenceAnalysisResult{
		WordCount:      wordCount,
		VowelCount:     vowelCount,
		ConsonantCount: consonantCount,
	}
}
