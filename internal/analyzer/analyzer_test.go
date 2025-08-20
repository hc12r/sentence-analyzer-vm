package analyzer

import (
	"testing"
)

func TestAnalyzeSentence(t *testing.T) {
	tests := []struct {
		name     string
		sentence string
		want     SentenceAnalysisResult
	}{
		{
			name:     "empty string",
			sentence: "",
			want: SentenceAnalysisResult{
				WordCount:      0,
				VowelCount:     0,
				ConsonantCount: 0,
			},
		},
		{
			name:     "simple sentence",
			sentence: "Hello World",
			want: SentenceAnalysisResult{
				WordCount:      2,
				VowelCount:     3,
				ConsonantCount: 7,
			},
		},
		{
			name:     "complex sentence with punctuation",
			sentence: "The quick brown fox jumps over the lazy dog!",
			want: SentenceAnalysisResult{
				WordCount:      9,
				VowelCount:     11,
				ConsonantCount: 24,
			},
		},
		{
			name:     "sentence with numbers",
			sentence: "There are 42 apples and 7 oranges.",
			want: SentenceAnalysisResult{
				WordCount:      7,
				VowelCount:     10,
				ConsonantCount: 14,
			},
		},
		{
			name:     "sentence with special characters",
			sentence: "Hello, world! How are you today? I'm fine, thank you.",
			want: SentenceAnalysisResult{
				WordCount:      10,
				VowelCount:     16,
				ConsonantCount: 22,
			},
		},
		{
			name:     "sentence with uppercase letters",
			sentence: "HELLO WORLD",
			want: SentenceAnalysisResult{
				WordCount:      2,
				VowelCount:     3,
				ConsonantCount: 7,
			},
		},
		{
			name:     "sentence with mixed case",
			sentence: "HeLLo WoRLd",
			want: SentenceAnalysisResult{
				WordCount:      2,
				VowelCount:     3,
				ConsonantCount: 7,
			},
		},
		{
			name:     "sentence with only vowels",
			sentence: "aeiou",
			want: SentenceAnalysisResult{
				WordCount:      1,
				VowelCount:     5,
				ConsonantCount: 0,
			},
		},
		{
			name:     "sentence with only consonants",
			sentence: "bcdfg",
			want: SentenceAnalysisResult{
				WordCount:      1,
				VowelCount:     0,
				ConsonantCount: 5,
			},
		},
		{
			name:     "sentence with multiple spaces",
			sentence: "Hello    World",
			want: SentenceAnalysisResult{
				WordCount:      2,
				VowelCount:     3,
				ConsonantCount: 7,
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
