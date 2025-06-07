package watcher

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"
)

type fakeProvider struct{}

func (fakeProvider) Fetch(ctx context.Context, mediaPath, lang string) ([]byte, error) {
	return []byte("sub"), nil
}

func TestWatchDirectory(t *testing.T) {
	dir := t.TempDir()
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})
	go func() {
		if err := WatchDirectory(ctx, dir, "en", fakeProvider{}); err != context.Canceled {
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
