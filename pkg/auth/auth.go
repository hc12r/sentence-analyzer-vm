package auth

// This package is kept as a placeholder for future authentication utilities
// Basic authentication has been removed in favor of Kong with OpenID Connect

// AuthInfo represents information about the authenticated user
// This can be expanded in the future if needed
type AuthInfo struct {
	UserID string
	Roles  []string
}

// Note: Authentication is now handled by Kong with OpenID Connect
// This package can be used for auth-related utilities in the future
