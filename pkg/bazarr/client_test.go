// file: pkg/bazarr/client_test.go
package bazarr

import (
	"testing"
)

// TestValidateBaseURL tests the URL validation function for security
func TestValidateBaseURL(t *testing.T) {
	tests := []struct {
		name        string
		url         string
		shouldError bool
		errorText   string
	}{
		// Valid URLs that should be allowed
		{
			name:        "valid HTTP localhost",
			url:         "http://localhost:6767",
			shouldError: false,
		},
		{
			name:        "valid HTTPS with domain",
			url:         "https://mybazarr.example.com",
			shouldError: false,
		},
		{
			name:        "valid HTTP with IP",
			url:         "http://192.168.1.100:6767",
			shouldError: false,
		},
		{
			name:        "valid with port",
			url:         "http://bazarr.local:8080",
			shouldError: false,
		},

		// Invalid URLs that should be blocked
		{
			name:        "invalid scheme FTP",
			url:         "ftp://localhost:6767",
			shouldError: true,
			errorText:   "only HTTP and HTTPS schemes are allowed",
		},
		{
			name:        "invalid scheme file",
			url:         "file:///etc/passwd",
			shouldError: true,
			errorText:   "only HTTP and HTTPS schemes are allowed",
		},
		{
			name:        "AWS metadata service",
			url:         "http://169.254.169.254/latest/meta-data/",
			shouldError: true,
			errorText:   "hostname 169.254.169.254 is not allowed",
		},
		{
			name:        "GCP metadata service",
			url:         "http://metadata.google.internal/computeMetadata/v1/",
			shouldError: true,
			errorText:   "hostname metadata.google.internal is not allowed",
		},
		{
			name:        "SSH port blocked",
			url:         "http://target.com:22",
			shouldError: true,
			errorText:   "port 22 is not allowed",
		},
		{
			name:        "RDP port blocked",
			url:         "http://target.com:3389",
			shouldError: true,
			errorText:   "port 3389 is not allowed",
		},
		{
			name:        "empty hostname",
			url:         "http://",
			shouldError: true,
			errorText:   "hostname cannot be empty",
		},
		{
			name:        "malformed URL",
			url:         "not-a-url",
			shouldError: true,
			errorText:   "only HTTP and HTTPS schemes are allowed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateBaseURL(tt.url)

			if tt.shouldError {
				if err == nil {
					t.Errorf("Expected error for URL %s, but got none", tt.url)
				} else if tt.errorText != "" && !containsString(err.Error(), tt.errorText) {
					t.Errorf("Expected error containing '%s', got '%s'", tt.errorText, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error for URL %s, but got: %v", tt.url, err)
				}
			}
		})
	}
}

// Helper function to check if a string contains another string
func containsString(haystack, needle string) bool {
	return len(haystack) >= len(needle) &&
		(needle == "" || haystack == needle ||
			len(haystack) > len(needle) &&
				(haystack[:len(needle)] == needle ||
					haystack[len(haystack)-len(needle):] == needle ||
					indexOfString(haystack, needle) >= 0))
}

// Simple string search function
func indexOfString(haystack, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

// TestFetchSettingsSSRFProtection tests that FetchSettings rejects malicious URLs
func TestFetchSettingsSSRFProtection(t *testing.T) {
	// Test that malicious URLs are rejected before making any HTTP request
	maliciousURLs := []string{
		"http://169.254.169.254/latest/meta-data/",
		"http://metadata.google.internal/computeMetadata/v1/",
		"ftp://malicious.com/file",
		"file:///etc/passwd",
		"http://target.com:22",
	}

	for _, url := range maliciousURLs {
		t.Run("blocks_"+url, func(t *testing.T) {
			_, err := FetchSettings(url, "test-api-key")
			if err == nil {
				t.Errorf("Expected FetchSettings to reject malicious URL %s, but it didn't", url)
			}
			if !containsString(err.Error(), "invalid baseURL") {
				t.Errorf("Expected error to mention 'invalid baseURL', got: %v", err)
			}
		})
	}
}
