package middleware

import (
	"log"
	"net/http"

	"github.com/hc12r/sentence-analyzer-vm/pkg/auth"
)

// NoAuth middleware that simply passes through all requests
// This is a placeholder for when Kong handles authentication
func NoAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// With Kong and OpenID Connect, authentication is handled at the API gateway level
		// This middleware simply passes through all requests
		next(w, r)
	}
}

// JWTAuth middleware that validates JWT tokens
func JWTAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Skip authentication for OPTIONS requests (for CORS)
		if r.Method == http.MethodOptions {
			next(w, r)
			return
		}

		// Extract and validate the token
		authInfo, err := auth.GetAuthInfoFromRequest(r)
		if err != nil {
			log.Printf("Authentication error: %v", err)

			switch err {
			case auth.ErrNoToken:
				http.Error(w, "Authentication required", http.StatusUnauthorized)
			case auth.ErrExpiredToken:
				http.Error(w, "Token has expired", http.StatusUnauthorized)
			case auth.ErrInvalidToken:
				http.Error(w, "Invalid token", http.StatusUnauthorized)
			default:
				http.Error(w, "Authentication error", http.StatusInternalServerError)
			}
			return
		}

		// Store auth info in the request context for later use
		ctx := r.Context()
		ctx = auth.WithAuthInfo(ctx, authInfo)
		r = r.WithContext(ctx)

		// Call the next handler
		next(w, r)
	}
}
