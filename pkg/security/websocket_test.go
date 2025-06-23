// file: pkg/security/websocket_test.go
package security

import (
	"testing"
)

func TestValidateWebSocketOrigin(t *testing.T) {
	tests := []struct {
		name     string
		origin   string
		host     string
		expected bool
	}{
		{
			name:     "empty origin should be allowed",
			origin:   "",
			host:     "localhost:8080",
			expected: true,
		},
		{
			name:     "same origin should be allowed",
			origin:   "http://localhost:8080",
			host:     "localhost:8080",
			expected: true,
		},
		{
			name:     "localhost variations should be allowed",
			origin:   "http://127.0.0.1:3000",
			host:     "localhost:8080",
			expected: true,
		},
		{
			name:     "localhost to localhost different ports should be allowed",
			origin:   "http://localhost:3000",
			host:     "localhost:8080",
			expected: true,
		},
		{
			name:     "external origin should be blocked",
			origin:   "http://evil.com",
			host:     "localhost:8080",
			expected: false,
		},
		{
			name:     "invalid origin should be blocked",
			origin:   "not-a-url",
			host:     "localhost:8080",
			expected: false,
		},
		{
			name:     "HTTPS same origin should be allowed",
			origin:   "https://localhost:8080",
			host:     "localhost:8080",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateWebSocketOrigin(tt.origin, tt.host)
			if result != tt.expected {
				t.Errorf("ValidateWebSocketOrigin(%q, %q) = %v, want %v",
					tt.origin, tt.host, result, tt.expected)
			}
		})
	}
}
