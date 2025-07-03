// file: pkg/storage/local_test.go
// version: 1.0.0
// guid: 89012345-67f2-890b-cdef-0123456789cd

package storage

import (
	"context"
	"strings"
	"testing"
	"time"
)

func TestLocalProvider(t *testing.T) {
	// Create temporary directory for testing
	tempDir := t.TempDir()

	config := StorageConfig{
		Provider:  "local",
		LocalPath: tempDir,
	}

	provider, err := NewLocalProvider(config)
	if err != nil {
		t.Fatalf("Failed to create local provider: %v", err)
	}
	defer provider.Close()

	ctx := context.Background()
	testKey := "test/file.srt"
	testContent := "Test subtitle content"

	// Test Store
	err = provider.Store(ctx, testKey, strings.NewReader(testContent), "text/plain")
	if err != nil {
		t.Fatalf("Failed to store file: %v", err)
	}

	// Test Exists
	exists, err := provider.Exists(ctx, testKey)
	if err != nil {
		t.Fatalf("Failed to check existence: %v", err)
	}
	if !exists {
		t.Fatalf("File should exist")
	}

	// Test Retrieve
	reader, err := provider.Retrieve(ctx, testKey)
	if err != nil {
		t.Fatalf("Failed to retrieve file: %v", err)
	}
	defer reader.Close()

	// Verify content
	buf := make([]byte, len(testContent))
	n, err := reader.Read(buf)
	if err != nil {
		t.Fatalf("Failed to read content: %v", err)
	}
	if string(buf[:n]) != testContent {
		t.Fatalf("Content mismatch: got %s, want %s", string(buf[:n]), testContent)
	}

	// Test List
	keys, err := provider.List(ctx, "test/")
	if err != nil {
		t.Fatalf("Failed to list files: %v", err)
	}
	if len(keys) != 1 || keys[0] != testKey {
		t.Fatalf("List result incorrect: got %v, want [%s]", keys, testKey)
	}

	// Test GetURL
	url, err := provider.GetURL(ctx, testKey, time.Hour)
	if err != nil {
		t.Fatalf("Failed to get URL: %v", err)
	}
	if !strings.HasPrefix(url, "file://") {
		t.Fatalf("URL should start with file://")
	}

	// Test Delete
	err = provider.Delete(ctx, testKey)
	if err != nil {
		t.Fatalf("Failed to delete file: %v", err)
	}

	// Test file no longer exists
	exists, err = provider.Exists(ctx, testKey)
	if err != nil {
		t.Fatalf("Failed to check existence after delete: %v", err)
	}
	if exists {
		t.Fatalf("File should not exist after delete")
	}
}

func TestLocalProvider_SecurityChecks(t *testing.T) {
	tempDir := t.TempDir()

	config := StorageConfig{
		Provider:  "local",
		LocalPath: tempDir,
	}

	provider, err := NewLocalProvider(config)
	if err != nil {
		t.Fatalf("Failed to create local provider: %v", err)
	}
	defer provider.Close()

	ctx := context.Background()

	// Test directory traversal protection
	maliciousKeys := []string{
		"../../../etc/passwd",
		"..\\..\\..\\windows\\system32\\config\\sam",
		"/etc/passwd",
		"\\windows\\system32\\config\\sam",
	}

	for _, key := range maliciousKeys {
		err := provider.Store(ctx, key, strings.NewReader("test"), "text/plain")
		if err != ErrInvalidKey {
			t.Fatalf("Expected ErrInvalidKey for malicious key %s, got %v", key, err)
		}
	}
}

func TestNewProvider(t *testing.T) {
	tests := []struct {
		name        string
		config      StorageConfig
		expectError bool
	}{
		{
			name: "local provider",
			config: StorageConfig{
				Provider:  "local",
				LocalPath: t.TempDir(),
			},
			expectError: false,
		},
		{
			name: "unsupported provider",
			config: StorageConfig{
				Provider: "unsupported",
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			provider, err := NewProvider(tt.config)
			if tt.expectError {
				if err == nil {
					t.Fatalf("Expected error but got none")
				}
				if provider != nil {
					t.Fatalf("Expected nil provider but got %v", provider)
				}
			} else {
				if err != nil {
					t.Fatalf("Unexpected error: %v", err)
				}
				if provider == nil {
					t.Fatalf("Expected provider but got nil")
				}
				provider.Close()
			}
		})
	}
}
