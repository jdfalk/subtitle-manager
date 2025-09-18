// file: pkg/gcommonauth/auth_test.go
// version: 1.1.0
// guid: fb1a0786-29d8-4305-8c9c-8762fb8845c7
package gcommonauth

import (
	"testing"
	"time"

	"github.com/jdfalk/subtitle-manager/pkg/testutil"
)

func TestSetUserRole(t *testing.T) {
	db := testutil.GetTestDB(t)
	defer db.Close()
	if err := CreateUser(db, "u", "p", "", "user"); err != nil {
		t.Fatalf("create: %v", err)
	}
	if err := SetUserRole(db, "u", "admin"); err != nil {
		t.Fatalf("set role: %v", err)
	}
	ok, err := CheckPermission(db, 1, "all")
	if err != nil {
		t.Fatalf("check: %v", err)
	}
	if !ok {
		t.Fatal("permission not granted")
	}
}

func TestOneTimeToken(t *testing.T) {
	db := testutil.GetTestDB(t)
	defer db.Close()
	if err := CreateUser(db, "u", "p", "e@example.com", "user"); err != nil {
		t.Fatalf("create: %v", err)
	}
	token, err := GenerateOneTimeToken(db, 1, time.Minute)
	if err != nil {
		t.Fatalf("gen token: %v", err)
	}
	id, err := ConsumeOneTimeToken(db, token)
	if err != nil {
		t.Fatalf("consume: %v", err)
	}
	if id != 1 {
		t.Fatalf("unexpected user id %d", id)
	}
	if _, err := ConsumeOneTimeToken(db, token); err == nil {
		t.Fatal("expected second consume to fail")
	}
}

func TestListUsers(t *testing.T) {
	db := testutil.GetTestDB(t)
	defer db.Close()
	if err := CreateUser(db, "u1", "p1", "e1@example.com", "user"); err != nil {
		t.Fatalf("create: %v", err)
	}
	if err := CreateUser(db, "u2", "p2", "e2@example.com", "admin"); err != nil {
		t.Fatalf("create2: %v", err)
	}
	users, err := ListUsers(db)
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if len(users) != 2 {
		t.Fatalf("expected 2 users, got %d", len(users))
	}
	if users[0].GetUsername() != "u1" || users[1].GetUsername() != "u2" {
		t.Fatalf("unexpected users %+v", users)
	}
}

// TestValidateAPIKey tests API key validation.
func TestValidateAPIKey(t *testing.T) {
	db := testutil.GetTestDB(t)
	defer db.Close()

	// Create a user
	if err := CreateUser(db, "testuser", "password", "test@example.com", "user"); err != nil {
		t.Fatalf("create user: %v", err)
	}

	// Generate an API key
	apiKeyObj, err := GenerateAPIKey(db, 1)
	if err != nil {
		t.Fatalf("generate API key: %v", err)
	}
	// Extract the key string from the gcommon APIKey
	apiKeyStr := apiKeyObj.GetId()

	// Test valid API key
	validatedAPIKey, err := ValidateAPIKey(db, apiKeyStr)
	if err != nil {
		t.Fatalf("validate API key: %v", err)
	}
	// Extract user ID from validated gcommon APIKey
	if userIdStr := validatedAPIKey.GetUserId(); userIdStr != "1" {
		t.Errorf("expected user ID '1', got '%s'", userIdStr)
	}

	// Test invalid API key
	_, err = ValidateAPIKey(db, "invalid-key")
	if err == nil {
		t.Error("expected error for invalid API key, got nil")
	}

	// Test empty API key
	_, err = ValidateAPIKey(db, "")
	if err == nil {
		t.Error("expected error for empty API key, got nil")
	}
}

// TestResetPassword tests password reset functionality.
func TestResetPassword(t *testing.T) {
	db := testutil.GetTestDB(t)
	defer db.Close()

	// Create a user
	if err := CreateUser(db, "testuser", "oldpassword", "test@example.com", "user"); err != nil {
		t.Fatalf("create user: %v", err)
	}

	// Reset password
	newPassword, newAPIKey, err := ResetPassword(db, 1)
	if err != nil {
		t.Fatalf("reset password: %v", err)
	}

	// Verify new password is generated
	if newPassword == "" {
		t.Error("expected new password to be generated, got empty string")
	}
	if len(newPassword) != 12 {
		t.Errorf("expected password length 12, got %d", len(newPassword))
	}

	// Verify new API key is generated
	if newAPIKey == "" {
		t.Error("expected new API key to be generated, got empty string")
	}

	// Test that the new API key is valid
	validatedNewAPIKey, err := ValidateAPIKey(db, newAPIKey)
	if err != nil {
		t.Fatalf("validate new API key: %v", err)
	}
	if userIdStr := validatedNewAPIKey.GetUserId(); userIdStr != "1" {
		t.Errorf("expected user ID '1', got '%s'", userIdStr)
	}

	// Test that old password no longer works
	validUser, err := AuthenticateUser(db, "testuser", "oldpassword")
	if err == nil {
		t.Error("expected error when authenticating with old password after reset")
	}
	if validUser != 0 {
		t.Error("expected user ID 0 when using old password after reset")
	}

	// Test that new password works
	validUser, err = AuthenticateUser(db, "testuser", newPassword)
	if err != nil {
		t.Fatalf("authenticate with new password: %v", err)
	}
	if validUser != 1 {
		t.Errorf("expected user ID 1 with new password, got %d", validUser)
	}
}

// TestGetOrCreateUser tests user creation and retrieval.
func TestGetOrCreateUser(t *testing.T) {
	db := testutil.GetTestDB(t)
	defer db.Close()

	// Test creating a new user
	userID1, err := GetOrCreateUser(db, "newuser", "new@example.com", "user")
	if err != nil {
		t.Fatalf("get or create user: %v", err)
	}
	if userID1 == 0 {
		t.Error("expected non-zero user ID for new user")
	}

	// Test getting existing user (same email)
	userID2, err := GetOrCreateUser(db, "anotheruser", "new@example.com", "admin")
	if err != nil {
		t.Fatalf("get existing user: %v", err)
	}
	if userID2 != userID1 {
		t.Errorf("expected same user ID %d, got %d", userID1, userID2)
	}

	// Test creating user with different email
	userID3, err := GetOrCreateUser(db, "thirduser", "third@example.com", "user")
	if err != nil {
		t.Fatalf("create third user: %v", err)
	}
	if userID3 == userID1 {
		t.Error("expected different user ID for different email")
	}

	// Verify the users were actually created/retrieved correctly
	users, err := ListUsers(db)
	if err != nil {
		t.Fatalf("list users: %v", err)
	}

	// Should have at least 2 users now
	if len(users) < 2 {
		t.Errorf("expected at least 2 users, got %d", len(users))
	}

	// Find the user with email "new@example.com"
	found := false
	for _, user := range users {
		if user.GetEmail() == "new@example.com" {
			found = true
			if user.GetUsername() != "newuser" {
				t.Errorf("expected username 'newuser', got %s", user.GetUsername())
			}
			break
		}
	}
	if !found {
		t.Error("user with email 'new@example.com' not found in user list")
	}
}

// TestValidateAPIKeyWithMultipleUsers tests API key validation with multiple users.
func TestValidateAPIKeyWithMultipleUsers(t *testing.T) {
	db := testutil.GetTestDB(t)
	defer db.Close()

	// Create multiple users
	if err := CreateUser(db, "user1", "pass1", "user1@example.com", "user"); err != nil {
		t.Fatalf("create user1: %v", err)
	}
	if err := CreateUser(db, "user2", "pass2", "user2@example.com", "admin"); err != nil {
		t.Fatalf("create user2: %v", err)
	}

	// Generate API keys for both users
	apiKeyObj1, err := GenerateAPIKey(db, 1)
	if err != nil {
		t.Fatalf("generate API key for user1: %v", err)
	}

	apiKeyObj2, err := GenerateAPIKey(db, 2)
	if err != nil {
		t.Fatalf("generate API key for user2: %v", err)
	}

	// Validate both API keys
	validatedAPIKey1, err := ValidateAPIKey(db, apiKeyObj1.GetId())
	if err != nil {
		t.Fatalf("validate API key 1: %v", err)
	}
	if userIdStr := validatedAPIKey1.GetUserId(); userIdStr != "1" {
		t.Errorf("expected user ID '1', got '%s'", userIdStr)
	}

	validatedAPIKey2, err := ValidateAPIKey(db, apiKeyObj2.GetId())
	if err != nil {
		t.Fatalf("validate API key 2: %v", err)
	}
	if userIdStr := validatedAPIKey2.GetUserId(); userIdStr != "2" {
		t.Errorf("expected user ID '2', got '%s'", userIdStr)
	}

	// Ensure keys are different
	if apiKeyObj1.GetId() == apiKeyObj2.GetId() {
		t.Error("expected different API keys for different users")
	}
}
