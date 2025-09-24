package watcher

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	providersmocks "github.com/jdfalk/subtitle-manager/pkg/providers/mocks"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/mock"
)

func TestWatchDirectory(t *testing.T) {
	dir := t.TempDir()
	viper.Set("media_directory", dir)
	defer viper.Reset()
	ctx, cancel := context.WithCancel(context.Background())
	m := providersmocks.NewMockProvider(t)
	m.On("Fetch", mock.Anything, mock.Anything, "en").Return([]byte("sub"), nil)
	done := make(chan struct{})
	go func() {
		if err := WatchDirectory(ctx, dir, "en", "test", m, nil); err != context.Canceled {
			t.Errorf("watch: %v", err)
		}
		close(done)
	}()
	time.Sleep(100 * time.Millisecond)

	f := filepath.Join(dir, "video.mkv")
	if err := os.WriteFile(f, []byte("x"), 0644); err != nil {
		t.Fatalf("create video: %v", err)
	}
	out := filepath.Join(dir, "video.en.srt")
	for i := 0; i < 10; i++ {
		if _, err := os.Stat(out); err == nil {
			break
		}
		time.Sleep(50 * time.Millisecond)
	}
	if _, err := os.Stat(out); err != nil {
		t.Fatalf("subtitle not downloaded: %v", err)
	}
	cancel()
	<-done
}

func TestWatchDirectoryRecursive(t *testing.T) {
	dir := t.TempDir()
	subdir := filepath.Join(dir, "a", "b")
	if err := os.MkdirAll(subdir, 0755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	viper.Set("media_directory", dir)
	defer viper.Reset()
	ctx, cancel := context.WithCancel(context.Background())
	m := providersmocks.NewMockProvider(t)
	m.On("Fetch", mock.Anything, mock.Anything, "en").Return([]byte("sub"), nil)
	done := make(chan struct{})
	go func() {
		if err := WatchDirectoryRecursive(ctx, dir, "en", "test", m, nil); err != context.Canceled {
			t.Errorf("watch recursive: %v", err)
		}
		close(done)
	}()
	time.Sleep(100 * time.Millisecond)

	f := filepath.Join(subdir, "video.mkv")
	if err := os.WriteFile(f, []byte("x"), 0644); err != nil {
		t.Fatalf("create video: %v", err)
	}
	out := filepath.Join(subdir, "video.en.srt")
	for i := 0; i < 10; i++ {
		if _, err := os.Stat(out); err == nil {
			break
		}
		time.Sleep(50 * time.Millisecond)
	}
	if _, err := os.Stat(out); err != nil {
		t.Fatalf("subtitle not downloaded: %v", err)
	}
	cancel()
	<-done
}
