// file: pkg/gcommonauth/session_test.go
// version: 1.1.0
// guid: 619fb89c-869d-43ba-baf7-167d4345cc8d
package gcommonauth

import (
	"strconv"
	"testing"
	"time"

	"github.com/jdfalk/subtitle-manager/pkg/testutil"
)

// TestSessionLifecycle tests the complete session lifecycle including invalidation.
func TestSessionLifecycle(t *testing.T) {
	db := testutil.GetTestDB(t)
	defer db.Close()

	// Create a test user
	err := CreateUser(db, "testuser", "password123", "test@example.com", "basic")
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	// Authenticate the user
	userID, err := AuthenticateUser(db, "testuser", "password123")
	if err != nil {
		t.Fatalf("failed to authenticate user: %v", err)
	}

	// Generate a session
	sessionObj, err := GenerateSession(db, userID, 24*time.Hour)
	if err != nil {
		t.Fatalf("failed to generate session: %v", err)
	}

	// Validate the session
	validatedSession, err := ValidateSession(db, sessionObj.GetId())
	if err != nil {
		t.Fatalf("failed to validate session: %v", err)
	}
	if userIdStr := validatedSession.GetUserId(); userIdStr != strconv.FormatInt(userID, 10) {
		t.Fatalf("session validation returned wrong user ID: expected %d, got %s", userID, userIdStr)
	}

	// Invalidate the session
	err = InvalidateSession(db, sessionObj.GetId())
	if err != nil {
		t.Fatalf("failed to invalidate session: %v", err)
	}

	// Try to validate the invalidated session (should fail)
	_, err = ValidateSession(db, sessionObj.GetId())
	if err == nil {
		t.Fatal("session validation should have failed after invalidation")
	}
}

// TestInvalidateUserSessions tests invalidating all sessions for a user.
func TestInvalidateUserSessions(t *testing.T) {
	db := testutil.GetTestDB(t)
	defer db.Close()

	// Create a test user
	err := CreateUser(db, "testuser", "password123", "test@example.com", "basic")
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	// Authenticate the user
	userID, err := AuthenticateUser(db, "testuser", "password123")
	if err != nil {
		t.Fatalf("failed to authenticate user: %v", err)
	}

	// Generate multiple sessions
	// Generate multiple sessions
	sess1, err := GenerateSession(db, userID, 24*time.Hour)
	if err != nil {
		t.Fatalf("failed to generate session 1: %v", err)
	}

	sess2, err := GenerateSession(db, userID, 24*time.Hour)
	if err != nil {
		t.Fatalf("failed to generate session 2: %v", err)
	}

	// Validate both sessions work
	_, err = ValidateSession(db, sess1.GetId())
	if err != nil {
		t.Fatalf("session 1 should be valid: %v", err)
	}

	_, err = ValidateSession(db, sess2.GetId())
	if err != nil {
		t.Fatalf("session 2 should be valid: %v", err)
	}

	// Invalidate all user sessions
	err = InvalidateUserSessions(db, userID)
	if err != nil {
		t.Fatalf("failed to invalidate user sessions: %v", err)
	}

	// Both sessions should now be invalid
	_, err = ValidateSession(db, sess1.GetId())
	if err == nil {
		t.Fatal("session 1 should be invalid after user session invalidation")
	}

	_, err = ValidateSession(db, sess2.GetId())
	if err == nil {
		t.Fatal("session 2 should be invalid after user session invalidation")
	}
}

// TestCleanupExpiredSessions tests the cleanup of expired sessions.
func TestCleanupExpiredSessions(t *testing.T) {
	db := testutil.GetTestDB(t)
	defer db.Close()

	// Create a test user
	err := CreateUser(db, "testuser", "password123", "test@example.com", "basic")
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	// Authenticate the user
	userID, err := AuthenticateUser(db, "testuser", "password123")
	if err != nil {
		t.Fatalf("failed to authenticate user: %v", err)
	}

	// Generate an expired session (expires 1 second ago)
	expiredSess, err := GenerateSession(db, userID, -1*time.Second)
	if err != nil {
		t.Fatalf("failed to generate expired session: %v", err)
	}

	// Generate a valid session
	validSess, err := GenerateSession(db, userID, 24*time.Hour)
	if err != nil {
		t.Fatalf("failed to generate valid session: %v", err)
	}

	// Validate that the expired session is indeed invalid
	_, err = ValidateSession(db, expiredSess.GetId())
	if err == nil {
		t.Fatal("expired session should not validate")
	}

	// Validate that the valid session works
	_, err = ValidateSession(db, validSess.GetId())
	if err != nil {
		t.Fatalf("valid session should validate: %v", err)
	}

	// Run cleanup
	err = CleanupExpiredSessions(db)
	if err != nil {
		t.Fatalf("failed to cleanup expired sessions: %v", err)
	}

	// Valid session should still work
	_, err = ValidateSession(db, validSess.GetId())
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
