package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hc12r/sentence-analyzer-vm/pkg/domain"
)

func TestHandleAnalyzeSentence(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		requestBody    interface{}
		wantStatusCode int
		wantResponse   *domain.SentenceAnalysisResponse
	}{
		{
			name:           "valid request",
			method:         http.MethodPost,
			requestBody:    domain.SentenceAnalysisRequest{Sentence: "Hello World"},
			wantStatusCode: http.StatusOK,
			wantResponse: &domain.SentenceAnalysisResponse{
				WordCount:      2,
				VowelCount:     3,
				ConsonantCount: 7,
			},
		},
		{
			name:           "invalid method",
			method:         http.MethodGet,
			requestBody:    nil,
			wantStatusCode: http.StatusMethodNotAllowed,
			wantResponse:   nil,
		},
		{
			name:           "invalid request body",
			method:         http.MethodPost,
			requestBody:    "not a valid json",
			wantStatusCode: http.StatusBadRequest,
			wantResponse:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var reqBody []byte
			var err error

			if tt.requestBody != nil {
				switch v := tt.requestBody.(type) {
				case string:
					reqBody = []byte(v)
				default:
					reqBody, err = json.Marshal(tt.requestBody)
					if err != nil {
						t.Fatalf("Failed to marshal request body: %v", err)
					}
				}
			}

			req, err := http.NewRequest(tt.method, "/analyze", bytes.NewBuffer(reqBody))
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(HandleAnalyzeSentence)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.wantStatusCode {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.wantStatusCode)
			}

			if tt.wantResponse != nil {
				var got domain.SentenceAnalysisResponse
				err = json.Unmarshal(rr.Body.Bytes(), &got)
				if err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if got.WordCount != tt.wantResponse.WordCount {
					t.Errorf("WordCount = %v, want %v", got.WordCount, tt.wantResponse.WordCount)
				}
				if got.VowelCount != tt.wantResponse.VowelCount {
					t.Errorf("VowelCount = %v, want %v", got.VowelCount, tt.wantResponse.VowelCount)
				}
				if got.ConsonantCount != tt.wantResponse.ConsonantCount {
					t.Errorf("ConsonantCount = %v, want %v", got.ConsonantCount, tt.wantResponse.ConsonantCount)
				}
			}
		})
	}
}
