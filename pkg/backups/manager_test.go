// file: pkg/backups/manager_test.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174007

package backups

import (
	"context"
	"testing"
	"time"
)

func TestBackupManager_CreateBackup(t *testing.T) {
	storage := &noOpStorage{}
	compression := NewGzipCompression()
	encryption, err := NewAESEncryption([]byte("1234567890123456")) // 16 bytes for AES-128
	if err != nil {
		t.Fatalf("failed to create encryption: %v", err)
	}

	manager := NewBackupManager(storage, compression, encryption)

	ctx := context.Background()
	testData := []byte("test backup data")

	backup, err := manager.CreateBackup(ctx, BackupTypeDatabase, []string{"test"}, testData)
	if err != nil {
		t.Fatalf("failed to create backup: %v", err)
	}

	if backup.ID == "" {
		t.Error("backup ID should not be empty")
	}
	if backup.Type != BackupTypeDatabase {
		t.Errorf("expected backup type %s, got %s", BackupTypeDatabase, backup.Type)
	}
	if !backup.Compressed {
		t.Error("backup should be compressed")
	}
	if !backup.Encrypted {
		t.Error("backup should be encrypted")
	}
	if len(backup.Contents) != 1 || backup.Contents[0] != "test" {
		t.Errorf("expected contents [test], got %v", backup.Contents)
	}
}

func TestBackupManager_ListBackups(t *testing.T) {
	storage := &noOpStorage{}
	manager := NewBackupManager(storage, nil, nil)

	ctx := context.Background()
	testData := []byte("test data")

	// Create multiple backups
	backup1, err := manager.CreateBackup(ctx, BackupTypeDatabase, []string{"db"}, testData)
	if err != nil {
		t.Fatalf("failed to create backup1: %v", err)
	}

	backup2, err := manager.CreateBackup(ctx, BackupTypeConfiguration, []string{"config"}, testData)
	if err != nil {
		t.Fatalf("failed to create backup2: %v", err)
	}

	backups := manager.ListBackups()
	if len(backups) != 2 {
		t.Errorf("expected 2 backups, got %d", len(backups))
	}

	// Verify we can find both backups
	found1, found2 := false, false
	for _, backup := range backups {
		if backup.ID == backup1.ID {
			found1 = true
		}
		if backup.ID == backup2.ID {
			found2 = true
		}
	}

	if !found1 {
		t.Error("backup1 not found in list")
	}
	if !found2 {
		t.Error("backup2 not found in list")
	}
}

func TestBackupManager_RestoreBackup(t *testing.T) {
	storage := &mockStorage{data: make(map[string][]byte)}
	compression := NewGzipCompression()
	manager := NewBackupManager(storage, compression, nil)

	ctx := context.Background()
	originalData := []byte("test restore data")

	// Create backup
	backup, err := manager.CreateBackup(ctx, BackupTypeDatabase, []string{"test"}, originalData)
	if err != nil {
		t.Fatalf("failed to create backup: %v", err)
	}

	// Restore backup
	restoredData, err := manager.RestoreBackup(ctx, backup.ID)
	if err != nil {
		t.Fatalf("failed to restore backup: %v", err)
	}

	if string(restoredData) != string(originalData) {
		t.Errorf("restored data doesn't match original. Expected %q, got %q", string(originalData), string(restoredData))
	}
}

func TestBackupManager_DeleteBackup(t *testing.T) {
	storage := &mockStorage{data: make(map[string][]byte)}
	manager := NewBackupManager(storage, nil, nil)

	ctx := context.Background()
	testData := []byte("test data")

	// Create backup
	backup, err := manager.CreateBackup(ctx, BackupTypeDatabase, []string{"test"}, testData)
	if err != nil {
		t.Fatalf("failed to create backup: %v", err)
	}

	// Verify backup exists
	_, exists := manager.GetBackup(backup.ID)
	if !exists {
		t.Error("backup should exist")
	}

	// Delete backup
	err = manager.DeleteBackup(ctx, backup.ID)
	if err != nil {
		t.Fatalf("failed to delete backup: %v", err)
	}

	// Verify backup no longer exists
	_, exists = manager.GetBackup(backup.ID)
	if exists {
		t.Error("backup should not exist after deletion")
	}
}

// Mock storage for testing
type mockStorage struct {
	data map[string][]byte
}

func (m *mockStorage) Store(ctx context.Context, data []byte, filename string) (string, error) {
	path := "mock://" + filename + "-" + time.Now().Format("20060102150405")
	m.data[path] = make([]byte, len(data))
	copy(m.data[path], data)
	return path, nil
}

func (m *mockStorage) Retrieve(ctx context.Context, path string) ([]byte, error) {
	data, exists := m.data[path]
	if !exists {
		return nil, ErrBackupNotFound
	}
	result := make([]byte, len(data))
	copy(result, data)
	return result, nil
}

func (m *mockStorage) Delete(ctx context.Context, path string) error {
	delete(m.data, path)
	return nil
}

func (m *mockStorage) List(ctx context.Context) ([]string, error) {
	var paths []string
	for path := range m.data {
		paths = append(paths, path)
	}
	return paths, nil
}
