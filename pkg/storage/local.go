// file: pkg/storage/local.go
// version: 1.0.0
// guid: 34567890-12cd-ef34-5678-901234567890

package storage

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jdfalk/subtitle-manager/pkg/security"
)

// LocalProvider implements StorageProvider for local filesystem storage.
type LocalProvider struct {
	basePath string
}

// NewLocalProvider creates a new local filesystem storage provider.
func NewLocalProvider(config StorageConfig) (*LocalProvider, error) {
	basePath := config.LocalPath
	if basePath == "" {
		basePath = "subtitles"
	}

	sanitizedBase, err := security.SanitizePath(basePath)
	if err != nil {
		return nil, fmt.Errorf("invalid base path: %w", err)
	}

	if err := os.MkdirAll(string(sanitizedBase), 0755); err != nil {
		return nil, fmt.Errorf("failed to create storage directory: %w", err)
	}

	return &LocalProvider{
		basePath: string(sanitizedBase),
	}, nil
}

// Store saves content to the local filesystem.
func (lp *LocalProvider) Store(ctx context.Context, key string, content io.Reader, contentType string) error {
	if key == "" {
		return ErrInvalidKey
	}

	// Sanitize the key to prevent directory traversal
	sanitized, err := security.SanitizeRelativePath(key)
	if err != nil {
		return ErrInvalidKey
	}
	key = sanitized

	fullPath := filepath.Join(lp.basePath, key)

	// Create directory if needed
	dir := filepath.Dir(fullPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Create and write file
	file, err := os.Create(fullPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	_, err = io.Copy(file, content)
	if err != nil {
		return fmt.Errorf("failed to write content: %w", err)
	}

	return nil
}

// Retrieve reads content from the local filesystem.
func (lp *LocalProvider) Retrieve(ctx context.Context, key string) (io.ReadCloser, error) {
	if key == "" {
		return nil, ErrInvalidKey
	}

	// Sanitize the key to prevent directory traversal
	sanitized, err := security.SanitizeRelativePath(key)
	if err != nil {
		return nil, ErrInvalidKey
	}
	key = sanitized

	fullPath := filepath.Join(lp.basePath, key)

	file, err := os.Open(fullPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	return file, nil
}

// Delete removes content from the local filesystem.
func (lp *LocalProvider) Delete(ctx context.Context, key string) error {
	if key == "" {
		return ErrInvalidKey
	}

	// Sanitize the key to prevent directory traversal
	sanitized, err := security.SanitizeRelativePath(key)
	if err != nil {
		return ErrInvalidKey
	}
	key = sanitized

	fullPath := filepath.Join(lp.basePath, key)

	err = os.Remove(fullPath)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to delete file: %w", err)
	}

	return nil
}

// Exists checks if content exists in the local filesystem.
func (lp *LocalProvider) Exists(ctx context.Context, key string) (bool, error) {
	if key == "" {
		return false, ErrInvalidKey
	}

	// Sanitize the key to prevent directory traversal
	sanitized, err := security.SanitizeRelativePath(key)
	if err != nil {
		return false, ErrInvalidKey
	}
	key = sanitized

	fullPath := filepath.Join(lp.basePath, key)

	_, err = os.Stat(fullPath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, fmt.Errorf("failed to check file existence: %w", err)
	}

	return true, nil
}

// List returns keys matching the given prefix in the local filesystem.
func (lp *LocalProvider) List(ctx context.Context, prefix string) ([]string, error) {
	var keys []string

	err := filepath.Walk(lp.basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil // Skip errors
		}

		if info.IsDir() {
			return nil
		}

		// Get relative path from base
		relPath, err := filepath.Rel(lp.basePath, path)
		if err != nil {
			return nil
		}

		// Convert to forward slashes for consistency
		relPath = filepath.ToSlash(relPath)

		// Check if it matches the prefix
		if strings.HasPrefix(relPath, prefix) || prefix == "" {
			keys = append(keys, relPath)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to list files: %w", err)
	}

	return keys, nil
}

// GetURL returns a file:// URL for local files (not very useful but implements interface).
func (lp *LocalProvider) GetURL(ctx context.Context, key string, expiry time.Duration) (string, error) {
	if key == "" {
		return "", ErrInvalidKey
	}

	// Sanitize the key to prevent directory traversal
	sanitized, err := security.SanitizeRelativePath(key)
	if err != nil {
		return "", ErrInvalidKey
	}
	key = sanitized

	fullPath := filepath.Join(lp.basePath, key)
	absPath, err := filepath.Abs(fullPath)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute path: %w", err)
	}

	return "file://" + absPath, nil
}

// Close is a no-op for local filesystem.
func (lp *LocalProvider) Close() error {
	return nil
}
