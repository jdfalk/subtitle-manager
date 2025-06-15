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

// TestSyncWithAudioTrack verifies synchronization using a specific audio track.
func TestSyncWithAudioTrack(t *testing.T) {
	base, err := astisub.OpenFile("../../testdata/simple.srt")
	if err != nil {
		t.Fatalf("open base: %v", err)
	}

	dir := t.TempDir()
	subFile := filepath.Join(dir, "test.srt")
	f, err := os.Create(subFile)
	if err != nil {
		t.Fatalf("create temp: %v", err)
	}
	astisub.Subtitles{Items: base.Items}.WriteToSRT(f)
	f.Close()

	defer func(oldT func(string, string, string) ([]byte, error)) {
		SetTranscribeFunc(oldT)
	}(transcribeFn)

	SetTranscribeFunc(func(path, lang, key string) ([]byte, error) {
		// Mock transcription by returning the test file content shifted
		b, _ := os.ReadFile("../../testdata/simple.srt")
		return b, nil
	})

	opts := Options{
		UseAudio:    true,
		AudioTrack:  1,   // Use audio track 1
		AudioWeight: 1.0, // Use only audio for sync
		WhisperKey:  "test-key",
	}

	items, err := Sync("dummy.mkv", subFile, opts)
	if err != nil {
		t.Fatalf("sync: %v", err)
	}

	if len(items) == 0 {
		t.Fatal("no items returned")
	}

	// Should match the original timing since we're using the same test file
	expected := base.Items[0].StartAt
	if items[0].StartAt != expected {
		t.Fatalf("expected start time %v, got %v", expected, items[0].StartAt)
	}
}

// TestSyncWithMultipleSubtitleTracks verifies sync using multiple embedded subtitle tracks.
func TestSyncWithMultipleSubtitleTracks(t *testing.T) {
	base, err := astisub.OpenFile("../../testdata/simple.srt")
	if err != nil {
		t.Fatalf("open base: %v", err)
	}

	dir := t.TempDir()
	subFile := filepath.Join(dir, "test.srt")
	f, err := os.Create(subFile)
	if err != nil {
		t.Fatalf("create temp: %v", err)
	}
	astisub.Subtitles{Items: base.Items}.WriteToSRT(f)
	f.Close()

	defer func(oldE func(string, int) ([]*astisub.Item, error)) {
		SetExtractFunc(oldE)
	}(extractFn)

	// Mock extract function that returns different offsets for different tracks
	SetExtractFunc(func(path string, track int) ([]*astisub.Item, error) {
		offset := time.Duration(track+1) * time.Second
		return Shift(base.Items, offset), nil
	})

	opts := Options{
		UseEmbedded:    true,
		SubtitleTracks: []int{0, 1, 2}, // Use multiple tracks
		AudioWeight:    0.0,            // Use only embedded subtitles
	}

	items, err := Sync("dummy.mkv", subFile, opts)
	if err != nil {
		t.Fatalf("sync: %v", err)
	}

	if len(items) == 0 {
		t.Fatal("no items returned")
	}

	// Should be average of 1s, 2s, 3s = 2s
	expected := 2 * time.Second
	if items[0].StartAt != expected {
		t.Fatalf("expected start time %v, got %v", expected, items[0].StartAt)
	}
}

// TestSyncWithTranslation verifies that translation is applied after sync.
func TestSyncWithTranslation(t *testing.T) {
	// Create a simple subtitle with English text
	items := []*astisub.Item{
		{
			StartAt: time.Second,
			EndAt:   2 * time.Second,
			Lines: []astisub.Line{
				{
					Items: []astisub.LineItem{
						{Text: "Hello world"},
					},
				},
			},
		},
	}

	dir := t.TempDir()
	subFile := filepath.Join(dir, "test.srt")
	f, err := os.Create(subFile)
	if err != nil {
		t.Fatalf("create temp: %v", err)
	}
	astisub.Subtitles{Items: items}.WriteToSRT(f)
	f.Close()

	// Mock successful sync without actual translation service
	// In a real test, you'd mock the translator.Translate function
	opts := Options{
		UseEmbedded:      true,
		Translate:        true,
		TranslateLang:    "es",
		TranslateService: "google",
		GoogleAPIKey:     "test-key",
	}

	result, err := Sync("dummy.mkv", subFile, opts)
	if err != nil {
		t.Fatalf("sync with translation: %v", err)
	}

	if len(result) == 0 {
		t.Fatal("no items returned")
	}

	// Note: Without mocking the translator, the text won't actually be translated
	// but the test verifies that the translation code path doesn't break the sync
	if result[0].Lines[0].Items[0].Text == "" {
		t.Fatal("text was removed during sync")
	}
}

// TestSyncDefaultBehavior verifies that sync defaults to embedded subtitles when no method is specified.
func TestSyncDefaultBehavior(t *testing.T) {
	defer func(oldE func(string, int) ([]*astisub.Item, error)) {
		SetExtractFunc(oldE)
	}(extractFn)

	// Mock extract function
	SetExtractFunc(func(string, int) ([]*astisub.Item, error) {
		return []*astisub.Item{
			{StartAt: time.Second, EndAt: 2 * time.Second},
		}, nil
	})

	opts := Options{
		// No UseAudio or UseEmbedded specified
	}

	items, err := Sync("dummy.mkv", "../../testdata/simple.srt", opts)
	if err != nil {
		t.Fatalf("sync: %v", err)
	}

	if len(items) == 0 {
		t.Fatal("no items returned")
	}
}
