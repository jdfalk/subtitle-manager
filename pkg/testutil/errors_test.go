// file: pkg/testutil/errors_test.go
package testutil

import (
	"testing"
)

func TestMust(t *testing.T) {
	// Test successful case
	result := Must(t, "test operation", "success", nil)
	if result != "success" {
		t.Errorf("expected 'success', got %q", result)
	}
}

func TestMustNoError(t *testing.T) {
	// Test successful case
	MustNoError(t, "test operation", nil)
	// If we reach here, the test passed
}

func TestMustGet(t *testing.T) {
	// Test successful case
	result := MustGet(t, "test operation", func() (string, error) {
		return "success", nil
	})
	if result != "success" {
		t.Errorf("expected 'success', got %q", result)
	}
}

func TestMustEqual(t *testing.T) {
	// Test successful case
	MustEqual(t, "test equality", 42, 42)
	MustEqual(t, "test string equality", "hello", "hello")
}

func TestMustNotEqual(t *testing.T) {
	// Test successful case
	MustNotEqual(t, "test inequality", 42, 43)
	MustNotEqual(t, "test string inequality", "hello", "world")
}

func TestMustContain(t *testing.T) {
	// Test successful case
	MustContain(t, "test contains", "hello world", "world")
	MustContain(t, "test contains substring", "testing", "test")
}

func TestMustNotContain(t *testing.T) {
	// Test successful case
	MustNotContain(t, "test not contains", "hello world", "foo")
	MustNotContain(t, "test not contains", "testing", "xyz")
}

func TestContainsHelper(t *testing.T) {
	tests := []struct {
		haystack string
		needle   string
		expected bool
	}{
		{"hello world", "world", true},
		{"hello world", "foo", false},
		{"testing", "test", true},
		{"testing", "xyz", false},
		{"", "", true},
		{"hello", "", true},
		{"", "hello", false},
	}

	for _, test := range tests {
		result := contains(test.haystack, test.needle)
		if result != test.expected {
			t.Errorf("contains(%q, %q) = %v, expected %v",
				test.haystack, test.needle, result, test.expected)
		}
	}
}

// Example of how error cases would work (these would cause test failures if uncommented)
/*
func TestMustFailure(t *testing.T) {
	// This would cause the test to fail with t.Fatalf
	Must(t, "should fail", "", errors.New("test error"))
}

func TestMustNoErrorFailure(t *testing.T) {
	// This would cause the test to fail with t.Fatalf
	MustNoError(t, "should fail", errors.New("test error"))
}
*/
