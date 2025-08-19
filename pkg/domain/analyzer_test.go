package domain

import (
	"testing"
)

func TestAnalyzeSentence(t *testing.T) {
	tests := []struct {
		name     string
		sentence string
		want     SentenceAnalysisResponse
	}{
		{
			name:     "empty string",
			sentence: "",
			want: SentenceAnalysisResponse{
				WordCount:      0,
				VowelCount:     0,
				ConsonantCount: 0,
			},
		},
		{
			name:     "simple sentence",
			sentence: "Hello World",
			want: SentenceAnalysisResponse{
				WordCount:      2,
				VowelCount:     3,
				ConsonantCount: 7,
			},
		},
		{
			name:     "complex sentence with punctuation",
			sentence: "The quick brown fox jumps over the lazy dog!",
			want: SentenceAnalysisResponse{
				WordCount:      9,
				VowelCount:     11,
				ConsonantCount: 24,
			},
		},
		{
			name:     "sentence with numbers",
			sentence: "There are 42 apples and 7 oranges.",
			want: SentenceAnalysisResponse{
				WordCount:      7,
				VowelCount:     10,
				ConsonantCount: 14,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnalyzeSentence(tt.sentence)
			if got.WordCount != tt.want.WordCount {
				t.Errorf("WordCount = %v, want %v", got.WordCount, tt.want.WordCount)
			}
			if got.VowelCount != tt.want.VowelCount {
				t.Errorf("VowelCount = %v, want %v", got.VowelCount, tt.want.VowelCount)
			}
			if got.ConsonantCount != tt.want.ConsonantCount {
				t.Errorf("ConsonantCount = %v, want %v", got.ConsonantCount, tt.want.ConsonantCount)
			}
		})
	}
}
