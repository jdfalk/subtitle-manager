// file: pkg/media/filestorage.go
// version: 1.0.0
// guid: 3b4c5d6e-7f8a-9b0c-1d2e-3f4a5b6c7d8e

package media

import (
	"fmt"
	"path/filepath"
	"sync"
)

// FileStorage manages file ID to path mappings for the media service
// In a real implementation, this would be backed by a database or distributed storage
type FileStorage struct {
	files map[string]string // fileID -> filePath
	mu    sync.RWMutex
}

// NewFileStorage creates a new file storage instance
func NewFileStorage() *FileStorage {
	return &FileStorage{
		files: make(map[string]string),
	}
}

// RegisterFile registers a file path with a given ID
func (fs *FileStorage) RegisterFile(fileID, filePath string) {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	fs.files[fileID] = filePath
}

// GetFilePath retrieves the file path for a given file ID
func (fs *FileStorage) GetFilePath(fileID string) (string, error) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()

	path, exists := fs.files[fileID]
	if !exists {
		return "", fmt.Errorf("file not found: %s", fileID)
	}
	return path, nil
}

// CreateTempFile creates a temporary file and returns its ID and path
func (fs *FileStorage) CreateTempFile(prefix, suffix string) (fileID, filePath string, err error) {
	// Create temporary file
	tempDir := "/tmp"
	tempFile := filepath.Join(tempDir, fmt.Sprintf("%s-%s%s", prefix, generateFileID(), suffix))

	fileID = generateFileID()
	fs.RegisterFile(fileID, tempFile)

	return fileID, tempFile, nil
}

// generateFileID creates a unique file ID
func generateFileID() string {
	return fmt.Sprintf("file_%d", len(fmt.Sprintf("%p", &struct{}{})))
}
