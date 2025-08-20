package auth

import (
	"context"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func TestGenerateAndValidateToken(t *testing.T) {
	// Set a test secret key
	os.Setenv("JWT_SECRET_KEY", "test-secret-key")
	defer os.Unsetenv("JWT_SECRET_KEY")

	// Generate a token
	userID := "test-user"
	roles := []string{"admin", "user"}
	token, err := GenerateToken(userID, roles)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	// Validate the token
	claims, err := ValidateToken(token)
	if err != nil {
		t.Fatalf("Failed to validate token: %v", err)
	}

	// Check the claims
	if claims.UserID != userID {
		t.Errorf("Expected UserID to be %s, got %s", userID, claims.UserID)
	}

	if len(claims.Roles) != len(roles) {
		t.Errorf("Expected %d roles, got %d", len(roles), len(claims.Roles))
	}

	for i, role := range roles {
		if claims.Roles[i] != role {
			t.Errorf("Expected role %s, got %s", role, claims.Roles[i])
		}
	}

	if claims.Issuer != "sentence-analyzer-api" {
		t.Errorf("Expected Issuer to be sentence-analyzer-api, got %s", claims.Issuer)
	}

	if claims.Subject != userID {
		t.Errorf("Expected Subject to be %s, got %s", userID, claims.Subject)
	}
}

func TestExpiredToken(t *testing.T) {
	// Set a test secret key
	os.Setenv("JWT_SECRET_KEY", "test-secret-key")
	defer os.Unsetenv("JWT_SECRET_KEY")

	// Create a token that expires immediately
	claims := JWTClaims{
		UserID: "test-user",
		Roles:  []string{"user"},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(-time.Hour)), // Expired 1 hour ago
			IssuedAt:  jwt.NewNumericDate(time.Now().Add(-2 * time.Hour)),
			NotBefore: jwt.NewNumericDate(time.Now().Add(-2 * time.Hour)),
			Issuer:    "sentence-analyzer-api",
			Subject:   "test-user",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("test-secret-key"))
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	// Validate the token
	_, err = ValidateToken(tokenString)
	if err != ErrExpiredToken {
		t.Errorf("Expected ErrExpiredToken, got %v", err)
	}
}

func TestInvalidToken(t *testing.T) {
	// Set a test secret key
	os.Setenv("JWT_SECRET_KEY", "test-secret-key")
	defer os.Unsetenv("JWT_SECRET_KEY")

	// Test with an invalid token
	_, err := ValidateToken("invalid-token")
	if err != ErrInvalidToken {
		t.Errorf("Expected ErrInvalidToken, got %v", err)
	}
}

func TestExtractTokenFromRequest(t *testing.T) {
	// Create a request with an Authorization header
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer test-token")

	// Extract the token
	token, err := ExtractTokenFromRequest(req)
	if err != nil {
		t.Fatalf("Failed to extract token: %v", err)
	}

	if token != "test-token" {
		t.Errorf("Expected token to be test-token, got %s", token)
	}
}

func TestExtractTokenFromRequestNoToken(t *testing.T) {
	// Create a request without an Authorization header
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Extract the token
	_, err = ExtractTokenFromRequest(req)
	if err != ErrNoToken {
		t.Errorf("Expected ErrNoToken, got %v", err)
	}
}

func TestExtractTokenFromRequestInvalidFormat(t *testing.T) {
	// Create a request with an invalid Authorization header
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Authorization", "InvalidFormat")

	// Extract the token
	_, err = ExtractTokenFromRequest(req)
	if err != ErrInvalidToken {
		t.Errorf("Expected ErrInvalidToken, got %v", err)
	}
}

func TestContextFunctions(t *testing.T) {
	// Create an AuthInfo
	authInfo := &AuthInfo{
		UserID: "test-user",
		Roles:  []string{"admin", "user"},
	}

	// Add it to the context
	ctx := context.Background()
	ctx = WithAuthInfo(ctx, authInfo)

	// Get it from the context
	retrievedAuthInfo, ok := GetAuthInfo(ctx)
	if !ok {
		t.Fatalf("Failed to get AuthInfo from context")
	}

	// Check the values
	if retrievedAuthInfo.UserID != authInfo.UserID {
		t.Errorf("Expected UserID to be %s, got %s", authInfo.UserID, retrievedAuthInfo.UserID)
	}

	if len(retrievedAuthInfo.Roles) != len(authInfo.Roles) {
		t.Errorf("Expected %d roles, got %d", len(authInfo.Roles), len(retrievedAuthInfo.Roles))
	}

	for i, role := range authInfo.Roles {
		if retrievedAuthInfo.Roles[i] != role {
			t.Errorf("Expected role %s, got %s", role, retrievedAuthInfo.Roles[i])
		}
	}
}

func TestGetAuthInfoFromEmptyContext(t *testing.T) {
	// Try to get AuthInfo from an empty context
	ctx := context.Background()
	_, ok := GetAuthInfo(ctx)
	if ok {
		t.Errorf("Expected GetAuthInfo to return false for empty context")
	}
}
