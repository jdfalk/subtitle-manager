// file: pkg/auth/session_test.go

package auth

import (
	"testing"
	"time"

	"github.com/jdfalk/subtitle-manager/pkg/database"
)

// TestSessionLifecycle tests the complete session lifecycle including invalidation.
func TestSessionLifecycle(t *testing.T) {
	db, err := database.Open(":memory:")
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	// Create a test user
	err = CreateUser(db, "testuser", "password123", "test@example.com", "basic")
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	// Authenticate the user
	userID, err := AuthenticateUser(db, "testuser", "password123")
	if err != nil {
		t.Fatalf("failed to authenticate user: %v", err)
	}

	// Generate a session
	token, err := GenerateSession(db, userID, 24*time.Hour)
	if err != nil {
		t.Fatalf("failed to generate session: %v", err)
	}

	// Validate the session
	validatedUserID, err := ValidateSession(db, token)
	if err != nil {
		t.Fatalf("failed to validate session: %v", err)
	}
	if validatedUserID != userID {
		t.Fatalf("session validation returned wrong user ID: expected %d, got %d", userID, validatedUserID)
	}

	// Invalidate the session
	err = InvalidateSession(db, token)
	if err != nil {
		t.Fatalf("failed to invalidate session: %v", err)
	}

	// Try to validate the invalidated session (should fail)
	_, err = ValidateSession(db, token)
	if err == nil {
		t.Fatal("session validation should have failed after invalidation")
	}
}

// TestInvalidateUserSessions tests invalidating all sessions for a user.
func TestInvalidateUserSessions(t *testing.T) {
	db, err := database.Open(":memory:")
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	// Create a test user
	err = CreateUser(db, "testuser", "password123", "test@example.com", "basic")
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	// Authenticate the user
	userID, err := AuthenticateUser(db, "testuser", "password123")
	if err != nil {
		t.Fatalf("failed to authenticate user: %v", err)
	}

	// Generate multiple sessions
	token1, err := GenerateSession(db, userID, 24*time.Hour)
	if err != nil {
		t.Fatalf("failed to generate session 1: %v", err)
	}

	token2, err := GenerateSession(db, userID, 24*time.Hour)
	if err != nil {
		t.Fatalf("failed to generate session 2: %v", err)
	}

	// Validate both sessions work
	_, err = ValidateSession(db, token1)
	if err != nil {
		t.Fatalf("session 1 should be valid: %v", err)
	}

	_, err = ValidateSession(db, token2)
	if err != nil {
		t.Fatalf("session 2 should be valid: %v", err)
	}

	// Invalidate all user sessions
	err = InvalidateUserSessions(db, userID)
	if err != nil {
		t.Fatalf("failed to invalidate user sessions: %v", err)
	}

	// Both sessions should now be invalid
	_, err = ValidateSession(db, token1)
	if err == nil {
		t.Fatal("session 1 should be invalid after user session invalidation")
	}

	_, err = ValidateSession(db, token2)
	if err == nil {
		t.Fatal("session 2 should be invalid after user session invalidation")
	}
}

// TestCleanupExpiredSessions tests the cleanup of expired sessions.
func TestCleanupExpiredSessions(t *testing.T) {
	db, err := database.Open(":memory:")
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	// Create a test user
	err = CreateUser(db, "testuser", "password123", "test@example.com", "basic")
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	// Authenticate the user
	userID, err := AuthenticateUser(db, "testuser", "password123")
	if err != nil {
		t.Fatalf("failed to authenticate user: %v", err)
	}

	// Generate an expired session (expires 1 second ago)
	expiredToken, err := GenerateSession(db, userID, -1*time.Second)
	if err != nil {
		t.Fatalf("failed to generate expired session: %v", err)
	}

	// Generate a valid session
	validToken, err := GenerateSession(db, userID, 24*time.Hour)
	if err != nil {
		t.Fatalf("failed to generate valid session: %v", err)
	}

	// Validate that the expired session is indeed invalid
	_, err = ValidateSession(db, expiredToken)
	if err == nil {
		t.Fatal("expired session should not validate")
	}

	// Validate that the valid session works
	_, err = ValidateSession(db, validToken)
	if err != nil {
		t.Fatalf("valid session should validate: %v", err)
	}

	// Run cleanup
	err = CleanupExpiredSessions(db)
	if err != nil {
		t.Fatalf("failed to cleanup expired sessions: %v", err)
	}

	// Valid session should still work
	_, err = ValidateSession(db, validToken)
	if err != nil {
		t.Fatalf("valid session should still work after cleanup: %v", err)
	}

	// Count sessions in the database to verify cleanup worked
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM sessions").Scan(&count)
	if err != nil {
		t.Fatalf("failed to count sessions: %v", err)
	}

	// Should have only 1 session (the valid one)
	if count != 1 {
		t.Fatalf("expected 1 session after cleanup, got %d", count)
	}
}
