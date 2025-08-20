package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNoAuth(t *testing.T) {
	// Create a mock handler that we'll wrap with NoAuth
	mockHandlerCalled := false
	mockHandler := func(w http.ResponseWriter, r *http.Request) {
		mockHandlerCalled = true
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("handler called"))
	}

	// Create a test request
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Apply the NoAuth middleware to our mock handler
	handlerWithMiddleware := NoAuth(mockHandler)

	// Call the handler with our test request and response recorder
	handlerWithMiddleware(rr, req)

	// Check if the mock handler was called
	if !mockHandlerCalled {
		t.Error("NoAuth middleware did not call the next handler")
	}

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("NoAuth middleware returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := "handler called"
	if rr.Body.String() != expected {
		t.Errorf("NoAuth middleware returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

// Test with different HTTP methods
func TestNoAuthWithDifferentMethods(t *testing.T) {
	methods := []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS", "HEAD"}

	for _, method := range methods {
		t.Run(method, func(t *testing.T) {
			// Create a mock handler
			mockHandlerCalled := false
			mockHandler := func(w http.ResponseWriter, r *http.Request) {
				mockHandlerCalled = true
				if r.Method != method {
					t.Errorf("Expected method %s, got %s", method, r.Method)
				}
				w.WriteHeader(http.StatusOK)
			}

			// Create a test request with the current method
			req, err := http.NewRequest(method, "/test", nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			// Create a response recorder
			rr := httptest.NewRecorder()

			// Apply the NoAuth middleware to our mock handler
			handlerWithMiddleware := NoAuth(mockHandler)

			// Call the handler with our test request and response recorder
			handlerWithMiddleware(rr, req)

			// Check if the mock handler was called
			if !mockHandlerCalled {
				t.Error("NoAuth middleware did not call the next handler")
			}

			// Check the status code
			if status := rr.Code; status != http.StatusOK {
				t.Errorf("NoAuth middleware returned wrong status code: got %v want %v", status, http.StatusOK)
			}
		})
	}
}

// Test with custom headers
func TestNoAuthWithCustomHeaders(t *testing.T) {
	// Create a mock handler that checks headers
	mockHandler := func(w http.ResponseWriter, r *http.Request) {
		// Check if the custom header was passed through
		if r.Header.Get("X-Custom-Header") != "test-value" {
			t.Error("Custom header was not passed through middleware")
		}
		w.WriteHeader(http.StatusOK)
	}

	// Create a test request with a custom header
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("X-Custom-Header", "test-value")

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Apply the NoAuth middleware to our mock handler
	handlerWithMiddleware := NoAuth(mockHandler)

	// Call the handler with our test request and response recorder
	handlerWithMiddleware(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("NoAuth middleware returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
