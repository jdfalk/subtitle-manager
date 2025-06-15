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

// TestComputeOffset verifies that computeOffset returns the expected duration.
func TestComputeOffset(t *testing.T) {
	ref := []*astisub.Item{{StartAt: 2 * time.Second}}
	target := []*astisub.Item{{StartAt: time.Second}}
	if d := computeOffset(ref, target); d != time.Second {
		t.Fatalf("unexpected offset %v", d)
	}
}

// TestSyncWeighted verifies synchronization using both audio and embedded
// subtitles with weighted averaging.
func TestSyncWeighted(t *testing.T) {
	base, err := astisub.OpenFile("../../testdata/simple.srt")
	if err != nil {
		t.Fatalf("open base: %v", err)
	}
	shifted := Shift(base.Items, -1*time.Second)
	dir := t.TempDir()
	subFile := filepath.Join(dir, "shifted.srt")
	f, err := os.Create(subFile)
	if err != nil {
		t.Fatalf("create temp: %v", err)
	}
	astisub.Subtitles{Items: shifted}.WriteToSRT(f)
	f.Close()

	defer func(oldT func(string, string, string) ([]byte, error), oldE func(string, int) ([]*astisub.Item, error)) {
		SetTranscribeFunc(oldT)
		SetExtractFunc(oldE)
	}(transcribeFn, extractFn)

	SetTranscribeFunc(func(string, string, string) ([]byte, error) {
		b, _ := os.ReadFile("../../testdata/simple.srt")
		return b, nil
	})

	SetExtractFunc(func(string, int) ([]*astisub.Item, error) {
		return Shift(base.Items, time.Second), nil
	})

	items, err := Sync("dummy.mkv", subFile, Options{UseAudio: true, UseEmbedded: true, AudioWeight: 0.7})
	if err != nil {
		t.Fatalf("sync: %v", err)
	}
	if len(items) == 0 {
		t.Fatal("no items returned")
	}
	exp := 1300 * time.Millisecond
	if items[0].StartAt != exp {
		t.Fatalf("unexpected start %v", items[0].StartAt)
	}
}
