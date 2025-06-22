// file: pkg/security/url_test.go
package security

import "testing"

func TestValidateURL(t *testing.T) {
	tests := []struct {
		name        string
		url         string
		shouldError bool
		errorText   string
	}{
		{"valid http", "http://localhost:8000", false, ""},
		{"invalid scheme", "ftp://host", true, "only HTTP and HTTPS"},
		{"blocked host", "http://169.254.169.254/", true, "hostname 169.254.169.254"},
		{"blocked port", "http://example.com:22", true, "port 22"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ValidateURL(tt.url)
			if tt.shouldError {
				if err == nil {
					t.Fatalf("expected error for %s", tt.url)
				}
				if tt.errorText != "" && !contains(err.Error(), tt.errorText) {
					t.Fatalf("expected error to contain %q, got %v", tt.errorText, err)
				}
			} else if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func contains(s, substr string) bool {
	return len(substr) == 0 || (len(s) >= len(substr) && (s == substr || s[:len(substr)] == substr || len(s) > len(substr) && (s[len(s)-len(substr):] == substr)))
}
