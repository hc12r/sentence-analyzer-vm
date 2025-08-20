package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHandleLogin(t *testing.T) {
	// Set test credentials
	os.Setenv("LOGIN_USERNAME", "testuser")
	os.Setenv("LOGIN_PASSWORD", "testpass")
	os.Setenv("JWT_SECRET_KEY", "test-secret-key")
	defer func() {
		os.Unsetenv("LOGIN_USERNAME")
		os.Unsetenv("LOGIN_PASSWORD")
		os.Unsetenv("JWT_SECRET_KEY")
	}()

	tests := []struct {
		name           string
		method         string
		requestBody    interface{}
		wantStatusCode int
		wantToken      bool
	}{
		{
			name:   "valid credentials",
			method: http.MethodPost,
			requestBody: LoginRequest{
				Username: "testuser",
				Password: "testpass",
			},
			wantStatusCode: http.StatusOK,
			wantToken:      true,
		},
		{
			name:   "invalid credentials",
			method: http.MethodPost,
			requestBody: LoginRequest{
				Username: "testuser",
				Password: "wrongpass",
			},
			wantStatusCode: http.StatusUnauthorized,
			wantToken:      false,
		},
		{
			name:           "invalid method",
			method:         http.MethodGet,
			requestBody:    nil,
			wantStatusCode: http.StatusMethodNotAllowed,
			wantToken:      false,
		},
		{
			name:           "invalid request body",
			method:         http.MethodPost,
			requestBody:    "not a valid json",
			wantStatusCode: http.StatusBadRequest,
			wantToken:      false,
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

			req, err := http.NewRequest(tt.method, "/login", bytes.NewBuffer(reqBody))
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(HandleLogin)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.wantStatusCode {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.wantStatusCode)
			}

			if tt.wantToken {
				var response LoginResponse
				err = json.Unmarshal(rr.Body.Bytes(), &response)
				if err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Token == "" {
					t.Errorf("Expected a token in the response, got empty string")
				}
			}
		})
	}
}

func TestHandleLoginWithDefaultCredentials(t *testing.T) {
	// Ensure environment variables are not set
	os.Unsetenv("LOGIN_USERNAME")
	os.Unsetenv("LOGIN_PASSWORD")
	os.Setenv("JWT_SECRET_KEY", "test-secret-key")
	defer os.Unsetenv("JWT_SECRET_KEY")

	// Create a request with the default credentials
	reqBody, _ := json.Marshal(LoginRequest{
		Username: "admin",
		Password: "password",
	})

	req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleLogin)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response LoginResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if response.Token == "" {
		t.Errorf("Expected a token in the response, got empty string")
	}
}
