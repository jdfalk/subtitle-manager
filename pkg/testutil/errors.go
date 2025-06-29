// Package testutil provides testing utilities and helpers for subtitle-manager.
// It includes database setup, error handling, and common test fixtures.
//
// This package is used across test files to ensure consistent test environments.
// Developers can add custom test helpers here as needed.
//
// Usage:
//
//	import (
//		"github.com/jdfalk/subtitle-manager/pkg/testutil"
//		"testing"
//	)
//
//	func TestSomething(t *testing.T) {
//		db := testutil.GetTestDB(t)
//		defer db.Close()
//
//		// ... your test code ...
//	}
//
// For more complex testing needs, consider using table-driven tests or
// external testing frameworks as appropriate.
//
// file: pkg/testutil/errors.go
package testutil

import (
	"database/sql"
	"strings"
	"testing"

	"github.com/jdfalk/subtitle-manager/pkg/database"
)

// Must is a generic helper that checks for errors and calls t.Fatalf if one occurs.
// It returns the successful result, allowing for cleaner test code.
// Works when you already have separated return values.
//
// Usage:
//
//	value, err := someFunction()
//	result := testutil.Must(t, "operation desc", value, err)
//
// For direct function calls, use MustGet instead.
func Must[T any](t *testing.T, msg string, result T, err error) T {
	t.Helper()
	if err != nil {
		t.Fatalf("%s: %v", msg, err)
	}
	return result
}

// MustNoError checks for errors and calls t.Fatalf if one occurs.
// Used for functions that only return an error.
//
// Usage:
//
//	testutil.MustNoError(t, "create user", auth.CreateUser(db, "user", "pass", "", "admin"))
func MustNoError(t *testing.T, msg string, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("%s: %v", msg, err)
	}
}

// MustGet is a convenience helper for functions that return (T, error).
// It wraps the function call in a closure to work with Go's multiple return values.
//
// Usage:
//
//	data := testutil.MustGet(t, "read file", func() ([]byte, error) { return os.ReadFile(tmp) })
//	key := testutil.MustGet(t, "generate key", func() (string, error) { return auth.GenerateAPIKey(db, 1) })
func MustGet[T any](t *testing.T, msg string, fn func() (T, error)) T {
	t.Helper()
	result, err := fn()
	if err != nil {
		t.Fatalf("%s: %v", msg, err)
	}
	return result
}

// MustEqual checks if two values are equal and calls t.Fatalf if they differ.
// This is useful for common equality assertions in tests.
//
// Usage:
//
//	testutil.MustEqual(t, "status code", http.StatusOK, resp.StatusCode)
//	testutil.MustEqual(t, "user count", 1, len(users))
func MustEqual[T comparable](t *testing.T, msg string, expected, actual T) {
	t.Helper()
	if expected != actual {
		t.Fatalf("%s: expected %v, got %v", msg, expected, actual)
	}
}

// MustNotEqual checks if two values are different and calls t.Fatalf if they are equal.
//
// Usage:
//
//	testutil.MustNotEqual(t, "user ID", 0, user.ID)
func MustNotEqual[T comparable](t *testing.T, msg string, unexpected, actual T) {
	t.Helper()
	if unexpected == actual {
		t.Fatalf("%s: expected value to not equal %v, but it did", msg, unexpected)
	}
}

// MustContain checks if a string contains a substring and calls t.Fatalf if it doesn't.
//
// Usage:
//
//	testutil.MustContain(t, "config content", string(data), "test_key: new")
func MustContain(t *testing.T, msg string, haystack, needle string) {
	t.Helper()
	if !contains(haystack, needle) {
		t.Fatalf("%s: expected %q to contain %q", msg, haystack, needle)
	}
}

// MustNotContain checks if a string does not contain a substring and calls t.Fatalf if it does.
//
// Usage:
//
//	testutil.MustNotContain(t, "error message", err.Error(), "unexpected")
func MustNotContain(t *testing.T, msg string, haystack, needle string) {
	t.Helper()
	if contains(haystack, needle) {
		t.Fatalf("%s: expected %q to not contain %q", msg, haystack, needle)
	}
}

// contains is a helper function to check if a string contains a substring.
// We use our own implementation to avoid importing strings package.
func contains(s, substr string) bool {
	return len(substr) <= len(s) && (len(substr) == 0 || indexByte(s, substr) >= 0)
}

// indexByte is a simplified version of strings.Index for substring search.
func indexByte(s, substr string) int {
	if len(substr) == 0 {
		return 0
	}
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}

// GetTestDB creates an in-memory SQLite database for testing.
// The returned database is fully initialized with the schema and ready for use.
// If SQLite support is not available, the test will be skipped.
func GetTestDB(t *testing.T) *sql.DB {
	t.Helper()

	// Use the existing OpenSQLStore function to get a properly initialized database
	store, err := database.OpenSQLStore(":memory:")
	if err != nil {
		if strings.Contains(err.Error(), "SQLite support not available") {
			t.Skip("SQLite support not available - skipping test that requires SQLite")
		}
		t.Fatalf("Failed to create test database: %v", err)
	}

	// Return the underlying *sql.DB
	return store.DB()
}

// CheckSQLiteSupport checks if SQLite support is available.
// Returns an error if SQLite is not supported.
func CheckSQLiteSupport() error {
	db, err := database.Open(":memory:")
	if err != nil {
		return err
	}
	db.Close()
	return nil
}
