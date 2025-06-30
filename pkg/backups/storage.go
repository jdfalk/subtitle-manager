// file: pkg/backups/storage.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174001

package backups

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

// LocalStorage implements the Storage interface for local file system.
type LocalStorage struct {
	basePath string
}

// NewLocalStorage creates a new local storage instance.
func NewLocalStorage(basePath string) *LocalStorage {
	return &LocalStorage{
		basePath: basePath,
	}
}

// Store saves backup data to the local file system.
func (ls *LocalStorage) Store(ctx context.Context, data []byte, filename string) (string, error) {
	// Ensure base directory exists
	if err := os.MkdirAll(ls.basePath, 0755); err != nil {
		return "", fmt.Errorf("failed to create backup directory: %w", err)
	}

	// Create full file path with timestamp
	timestamp := time.Now().Format("20060102-150405")
	fullFilename := fmt.Sprintf("%s-%s", timestamp, filename)
	fullPath := filepath.Join(ls.basePath, fullFilename)

	// Write data to file
	file, err := os.Create(fullPath)
	if err != nil {
		return "", fmt.Errorf("failed to create backup file: %w", err)
	}
	defer file.Close()

	if _, err := file.Write(data); err != nil {
		return "", fmt.Errorf("failed to write backup data: %w", err)
	}

	return fullPath, nil
}

// Retrieve loads backup data from the local file system.
func (ls *LocalStorage) Retrieve(ctx context.Context, path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open backup file: %w", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read backup file: %w", err)
	}

	return data, nil
}

// Delete removes backup data from the local file system.
func (ls *LocalStorage) Delete(ctx context.Context, path string) error {
	if err := os.Remove(path); err != nil {
		return fmt.Errorf("failed to delete backup file: %w", err)
	}
	return nil
}

// List returns all backup files in the local storage directory.
func (ls *LocalStorage) List(ctx context.Context) ([]string, error) {
	entries, err := os.ReadDir(ls.basePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil
		}
		return nil, fmt.Errorf("failed to list backup directory: %w", err)
	}

	var files []string
	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".bak" {
			files = append(files, filepath.Join(ls.basePath, entry.Name()))
		}
	}

	return files, nil
}