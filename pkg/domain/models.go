package domain

// SentenceAnalysisRequest represents the request body
type SentenceAnalysisRequest struct {
	Sentence string `json:"sentence"`
}

// SentenceAnalysisResponse represents the response body
type SentenceAnalysisResponse struct {
	WordCount      int `json:"word_count"`
	VowelCount     int `json:"vowel_count"`
	ConsonantCount int `json:"consonant_count"`
}
