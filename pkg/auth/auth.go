package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWT errors
var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("token has expired")
	ErrNoToken      = errors.New("no token provided")
)

// AuthInfo represents information about the authenticated user
type AuthInfo struct {
	UserID string
	Roles  []string
}

// JWTClaims represents the claims in the JWT token
type JWTClaims struct {
	UserID string   `json:"user_id"`
	Roles  []string `json:"roles"`
	jwt.RegisteredClaims
}

// Config holds the JWT configuration
type Config struct {
	SecretKey     string
	TokenDuration time.Duration
}

// LoadConfig loads the JWT configuration from environment variables
func LoadConfig() Config {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		secretKey = "default-secret-key-for-development-only"
	}

	// Default token duration is 24 hours
	tokenDuration := 24 * time.Hour

	return Config{
		SecretKey:     secretKey,
		TokenDuration: tokenDuration,
	}
}

// GenerateToken generates a JWT token for the given user
func GenerateToken(userID string, roles []string) (string, error) {
	config := LoadConfig()

	claims := JWTClaims{
		UserID: userID,
		Roles:  roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.TokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "sentence-analyzer-api",
			Subject:   userID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.SecretKey))
}

// ValidateToken validates the JWT token and returns the claims
func ValidateToken(tokenString string) (*JWTClaims, error) {
	config := LoadConfig()

	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.SecretKey), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidToken
}

// ExtractTokenFromRequest extracts the JWT token from the Authorization header
func ExtractTokenFromRequest(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoToken
	}

	// Check if it's a Bearer token
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", ErrInvalidToken
	}

	return parts[1], nil
}

// GetAuthInfoFromRequest extracts the AuthInfo from the request
func GetAuthInfoFromRequest(r *http.Request) (*AuthInfo, error) {
	tokenString, err := ExtractTokenFromRequest(r)
	if err != nil {
		return nil, err
	}

	claims, err := ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}

	return &AuthInfo{
		UserID: claims.UserID,
		Roles:  claims.Roles,
	}, nil
}

// Context key type to avoid collisions
type contextKey string

// AuthInfoKey is the key used to store AuthInfo in the context
const AuthInfoKey contextKey = "auth_info"

// WithAuthInfo adds AuthInfo to the context
func WithAuthInfo(ctx context.Context, authInfo *AuthInfo) context.Context {
	return context.WithValue(ctx, AuthInfoKey, authInfo)
}

// GetAuthInfo retrieves AuthInfo from the context
func GetAuthInfo(ctx context.Context) (*AuthInfo, bool) {
	authInfo, ok := ctx.Value(AuthInfoKey).(*AuthInfo)
	return authInfo, ok
}
