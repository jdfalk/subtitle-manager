// file: pkg/storage/manager_test.go
// version: 1.0.0
// guid: c7c1f9e4-1f42-4bb1-8b68-f94ff8e49e3a

package storage

import (
	"context"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/jdfalk/subtitle-manager/pkg/storage/mocks"
)

func TestStorageManager_LocalProviderWorkflow(t *testing.T) {
	tempDir := t.TempDir()
	config := StorageConfig{
		Provider:  "local",
		LocalPath: tempDir,
	}

	manager, err := NewStorageManager(config)
	if err != nil {
		t.Fatalf("create storage manager: %v", err)
	}
	t.Cleanup(func() {
		if closeErr := manager.Close(); closeErr != nil {
			t.Fatalf("close storage manager: %v", closeErr)
		}
	})

	ctx := context.Background()
	key := "movies/test.srt"
	content := "subtitle content"

	if err := manager.Store(ctx, key, strings.NewReader(content), "text/plain"); err != nil {
		t.Fatalf("store: %v", err)
	}

	exists, err := manager.Exists(ctx, key)
	if err != nil {
		t.Fatalf("exists: %v", err)
	}
	if !exists {
		t.Fatalf("expected stored file to exist")
	}

	reader, err := manager.Retrieve(ctx, key)
	if err != nil {
		t.Fatalf("retrieve: %v", err)
	}
	defer reader.Close()

	data, err := io.ReadAll(reader)
	if err != nil {
		t.Fatalf("read: %v", err)
	}
	if string(data) != content {
		t.Fatalf("unexpected content: got %q, want %q", string(data), content)
	}

	keys, err := manager.List(ctx, "movies/")
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if len(keys) != 1 || keys[0] != key {
		t.Fatalf("unexpected list result: %v", keys)
	}

	url, err := manager.GetURL(ctx, key, time.Minute)
	if err != nil {
		t.Fatalf("get URL: %v", err)
	}
	if !strings.HasPrefix(url, "file://") {
		t.Fatalf("expected file URL, got %q", url)
	}

	if err := manager.Delete(ctx, key); err != nil {
		t.Fatalf("delete: %v", err)
	}

	exists, err = manager.Exists(ctx, key)
	if err != nil {
		t.Fatalf("exists after delete: %v", err)
	}
	if exists {
		t.Fatalf("expected file to be deleted")
	}
}

func TestStorageManager_DeleteAndCloseBackupProviders(t *testing.T) {
	primary := mocks.NewMockStorageProvider(t)
	backup := mocks.NewMockStorageProvider(t)
	manager := &StorageManager{
		primary: primary,
		backup:  backup,
		config: StorageConfig{
			EnableBackup: true,
		},
	}

	ctx := context.Background()
	key := "backup/key"

	primary.EXPECT().Delete(ctx, key).Return(ErrNotFound)
	backup.EXPECT().Delete(ctx, key).Return(nil)

	if err := manager.Delete(ctx, key); err != ErrNotFound {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}

	primary.EXPECT().Close().Return(ErrConnectionFailed)
	backup.EXPECT().Close().Return(nil)

	if err := manager.Close(); err != ErrConnectionFailed {
		t.Fatalf("expected ErrConnectionFailed, got %v", err)
	}
}
