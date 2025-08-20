package middleware

import (
	"net/http"
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
