package middleware

import (
	"net/http"

	"github.com/hc12r/sentence-analyzer-vm/pkg/auth"
)

// BasicAuth middleware for authentication
func BasicAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		// Validate credentials using the auth package
		if !auth.ValidateCredentials(username, password) {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// If authentication is successful, call the next handler
		next(w, r)
	}
}
