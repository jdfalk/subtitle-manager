// file: pkg/media/filestorage_test.go
// version: 1.0.0
// guid: 9b6f7c2a-4d7a-4b1f-9e2d-2e5a6f7b8c9d

package media

import (
	"strings"
	"testing"
)

func TestFileStorage_RegisterAndGetFilePath(t *testing.T) {
	// Arrange: create storage and register a file ID/path.
	storage := NewFileStorage()
	fileID := "file-123"
	expectedPath := "/media/test-file.mkv"

	// Act: store and retrieve the path.
	storage.RegisterFile(fileID, expectedPath)
	path, err := storage.GetFilePath(fileID)

	// Assert: expect stored path and no error.
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if path != expectedPath {
		t.Fatalf("expected path %q, got %q", expectedPath, path)
	}
}

func TestFileStorage_GetFilePathMissing_ReturnsError(t *testing.T) {
	// Arrange: create storage with no entries.
	storage := NewFileStorage()

	// Act: attempt to fetch a missing file path.
	path, err := storage.GetFilePath("missing-id")

	// Assert: expect an error and empty path.
	if err == nil {
		t.Fatalf("expected an error for missing file ID")
	}
	if path != "" {
		t.Fatalf("expected empty path for missing file ID, got %q", path)
	}
}

func TestFileStorage_CreateTempFile_RegistersFile(t *testing.T) {
	// Arrange: create storage with prefix/suffix.
	storage := NewFileStorage()
	prefix := "subtitle"
	suffix := ".srt"

	// Act: create temp file mapping.
	fileID, filePath, err := storage.CreateTempFile(prefix, suffix)

	// Assert: expect a registered path with prefix/suffix and no error.
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if fileID == "" {
		t.Fatalf("expected non-empty file ID")
	}
	if !strings.HasPrefix(filePath, "/tmp/") {
		t.Fatalf("expected temp path in /tmp, got %q", filePath)
	}
	if !strings.Contains(filePath, prefix+"-") {
		t.Fatalf("expected temp path to include prefix %q, got %q", prefix, filePath)
	}
	if !strings.HasSuffix(filePath, suffix) {
		t.Fatalf("expected temp path to include suffix %q, got %q", suffix, filePath)
	}

	storedPath, err := storage.GetFilePath(fileID)
	if err != nil {
		t.Fatalf("expected no error retrieving stored path, got %v", err)
	}
	if storedPath != filePath {
		t.Fatalf("expected stored path %q, got %q", filePath, storedPath)
	}
}
