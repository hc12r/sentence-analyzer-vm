package auth

import (
	"net/http"
)

// Credentials for basic authentication
type Credentials struct {
	Username string
	Password string
}

// GetDefaultCredentials returns the default credentials for the application
// In a real application, this would be retrieved from a secure storage
func GetDefaultCredentials() Credentials {
	return Credentials{
		Username: "admin",
		Password: "password",
	}
}

// ValidateCredentials checks if the provided username and password are valid
func ValidateCredentials(username, password string) bool {
	validCreds := GetDefaultCredentials()
	return username == validCreds.Username && password == validCreds.Password
}

// BasicAuth middleware for authentication
// This is kept for backward compatibility but applications should use
// the middleware in internal/middleware package instead
func BasicAuth(next http.HandlerFunc) http.HandlerFunc {
	// Import the middleware from internal package
	// This is a simple wrapper to maintain backward compatibility
	return func(w http.ResponseWriter, r *http.Request) {
		// This is a simplified implementation
		// In a real application, we would use the internal middleware

		// Get the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Extract username and password from the Authorization header
		username, password, ok := r.BasicAuth()
		if !ok {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Invalid credentials format", http.StatusUnauthorized)
			return
		}

		// Validate credentials
		if !ValidateCredentials(username, password) {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// If authentication is successful, call the next handler
		next(w, r)
	}
}
