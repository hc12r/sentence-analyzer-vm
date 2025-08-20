package auth

import (
	"testing"
)

func TestAuthInfo(t *testing.T) {
	// Test creating an AuthInfo struct
	authInfo := AuthInfo{
		UserID: "user123",
		Roles:  []string{"admin", "user"},
	}

	// Verify the fields are set correctly
	if authInfo.UserID != "user123" {
		t.Errorf("Expected UserID to be 'user123', got '%s'", authInfo.UserID)
	}

	if len(authInfo.Roles) != 2 {
		t.Errorf("Expected 2 roles, got %d", len(authInfo.Roles))
	}

	if authInfo.Roles[0] != "admin" {
		t.Errorf("Expected first role to be 'admin', got '%s'", authInfo.Roles[0])
	}

	if authInfo.Roles[1] != "user" {
		t.Errorf("Expected second role to be 'user', got '%s'", authInfo.Roles[1])
	}
}

func TestEmptyAuthInfo(t *testing.T) {
	// Test creating an empty AuthInfo struct
	authInfo := AuthInfo{}

	// Verify the fields are empty
	if authInfo.UserID != "" {
		t.Errorf("Expected empty UserID, got '%s'", authInfo.UserID)
	}

	if authInfo.Roles != nil && len(authInfo.Roles) != 0 {
		t.Errorf("Expected empty Roles, got %v", authInfo.Roles)
	}
}

func TestAuthInfoWithEmptyRoles(t *testing.T) {
	// Test creating an AuthInfo struct with empty roles
	authInfo := AuthInfo{
		UserID: "user123",
		Roles:  []string{},
	}

	// Verify the fields are set correctly
	if authInfo.UserID != "user123" {
		t.Errorf("Expected UserID to be 'user123', got '%s'", authInfo.UserID)
	}

	if len(authInfo.Roles) != 0 {
		t.Errorf("Expected 0 roles, got %d", len(authInfo.Roles))
	}
}
