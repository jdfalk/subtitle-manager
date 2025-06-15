package syncer

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/asticode/go-astisub"
)

// TestShift verifies that the Shift function offsets subtitles by the given duration.
func TestShift(t *testing.T) {
	items := []*astisub.Item{{StartAt: 0, EndAt: time.Second}}
	out := Shift(items, 2*time.Second)
	if out[0].StartAt != 2*time.Second || out[0].EndAt != 3*time.Second {
		t.Fatalf("unexpected values: %#v", out[0])
	}
}

// TestSync loads a subtitle file to ensure no error is returned.
func TestSync(t *testing.T) {
	items, err := Sync("dummy.mkv", "../../testdata/simple.srt", Options{})
	if err != nil {
		t.Fatalf("sync: %v", err)
	}
	if len(items) == 0 {
		t.Fatal("no items returned")
	}
}

// TestSyncWithEmbedded verifies that embedded subtitles adjust timing.
func TestSyncWithEmbedded(t *testing.T) {
	dir := t.TempDir()
	script := filepath.Join(dir, "ffmpeg")
	data := "#!/bin/sh\ncp ../../testdata/simple.srt \"$6\"\n"
	if err := os.WriteFile(script, []byte(data), 0755); err != nil {
		t.Fatalf("write script: %v", err)
	}
	oldPath := os.Getenv("PATH")
	os.Setenv("PATH", dir+":"+oldPath)
	defer os.Setenv("PATH", oldPath)

	opts := Options{UseEmbedded: true, SubtitleTracks: []int{0}}
	items, err := Sync("dummy.mkv", "../../testdata/shifted.srt", opts)
	if err != nil {
		t.Fatalf("sync: %v", err)
	}
	if len(items) == 0 {
		t.Fatal("no items returned")
	}
	if items[0].StartAt != time.Second {
		t.Fatalf("expected start 1s got %v", items[0].StartAt)
	}
}
