package database

import (
	"testing"
	"time"
)

func TestPebbleAuthFunctionality(t *testing.T) {
	// Test PebbleDB authentication features
	store, err := OpenPebble(t.TempDir())
	if err != nil {
		t.Fatal("Failed to open PebbleDB store:", err)
	}
	defer store.Close()

	// Test user creation
	userID, err := store.CreateUser("testuser", "hashedpassword", "test@example.com", "user")
	if err != nil {
		t.Fatal("Failed to create user:", err)
	}
	t.Logf("✅ Successfully created user with ID: %s", userID)

	// Test user retrieval
	user, err := store.GetUserByUsername("testuser")
	if err != nil {
		t.Fatal("Failed to get user by username:", err)
	}
	if user == nil || user.GetEmail() != "test@example.com" {
		t.Fatal("User data incorrect")
	}
	t.Log("✅ Successfully retrieved user by username")

	// Test session creation
	err = store.CreateSession(userID, "session-token-123", 24*time.Hour)
	if err != nil {
		t.Fatal("Failed to create session:", err)
	}
	t.Log("✅ Successfully created session")

	// Test session validation
	validatedUserID, err := store.ValidateSession("session-token-123")
	if err != nil {
		t.Fatal("Failed to validate session:", err)
	}
	if validatedUserID != userID {
		t.Fatal("Session validation returned wrong user ID")
	}
	t.Log("✅ Successfully validated session")

	// Test API key creation
	err = store.CreateAPIKey(userID, "api-key-123")
	if err != nil {
		t.Fatal("Failed to create API key:", err)
	}
	t.Log("✅ Successfully created API key")

	// Test API key validation
	validatedUserID, err = store.ValidateAPIKey("api-key-123")
	if err != nil {
		t.Fatal("Failed to validate API key:", err)
	}
	if validatedUserID != userID {
		t.Fatal("API key validation returned wrong user ID")
	}
	t.Log("✅ Successfully validated API key")

	// Test dashboard preferences
	err = store.SetDashboardLayout(userID, `{"layout": "grid"}`)
	if err != nil {
		t.Fatal("Failed to set dashboard layout:", err)
	}
	t.Log("✅ Successfully set dashboard layout")

	layout, err := store.GetDashboardLayout(userID)
	if err != nil {
		t.Fatal("Failed to get dashboard layout:", err)
	}
	if layout != `{"layout": "grid"}` {
		t.Fatal("Dashboard layout incorrect")
	}
	t.Log("✅ Successfully retrieved dashboard layout")

	// Test permissions
	err = store.InitializeDefaultPermissions()
	if err != nil {
		t.Fatal("Failed to initialize permissions:", err)
	}
	t.Log("✅ Successfully initialized default permissions")

	perms, err := store.GetPermissionsForRole("user")
	if err != nil {
		t.Fatal("Failed to get permissions for role:", err)
	}
	if len(perms) == 0 {
		t.Fatal("No permissions found for user role")
	}
	t.Logf("✅ Successfully retrieved %d permissions for user role", len(perms))

	t.Log("🎉 All PebbleDB authentication features working correctly!")
}
