package syncer

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/asticode/go-astisub"

	"github.com/jdfalk/subtitle-manager/pkg/subtitles"
	"github.com/jdfalk/subtitle-manager/pkg/transcriber"
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

// TestSyncWithEmbedded verifies that embedded subtitles are used to compute an offset.
func TestSyncWithEmbedded(t *testing.T) {
	dir := t.TempDir()
	script := filepath.Join(dir, "ffmpeg")
	data := "#!/bin/sh\ncp ../../testdata/simple.srt \"$6\"\n"
	if err := os.WriteFile(script, []byte(data), 0755); err != nil {
		t.Fatalf("write script: %v", err)
	}
	subtitles.SetFFmpegPath(script)
	defer subtitles.SetFFmpegPath("ffmpeg")

	off := filepath.Join(dir, "off.srt")
	b := []byte("1\n00:00:03,000 --> 00:00:04,000\nHello\n")
	if err := os.WriteFile(off, b, 0644); err != nil {
		t.Fatalf("write off: %v", err)
	}

	items, err := Sync("dummy.mkv", off, Options{UseEmbedded: true})
	if err != nil {
		t.Fatalf("sync: %v", err)
	}
	if len(items) == 0 || items[0].StartAt != time.Second {
		t.Fatalf("unexpected start time: %v", items[0].StartAt)
	}
}

// TestSyncWithWhisper verifies that audio transcription can supply the reference timing.
func TestSyncWithWhisper(t *testing.T) {
	var gotPath string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		fmt.Fprint(w, "1\n00:00:01,000 --> 00:00:02,000\nHi\n")
	}))
	defer srv.Close()
	transcriber.SetBaseURL(srv.URL + "/v1")
	defer transcriber.SetBaseURL("https://api.openai.com/v1")

	dir := t.TempDir()
	off := filepath.Join(dir, "off.srt")
	b := []byte("1\n00:00:03,000 --> 00:00:04,000\nHi\n")
	if err := os.WriteFile(off, b, 0644); err != nil {
		t.Fatalf("write off: %v", err)
	}
	audio := filepath.Join(dir, "a.wav")
	if err := os.WriteFile(audio, []byte("data"), 0644); err != nil {
		t.Fatalf("write audio: %v", err)
	}

	items, err := Sync(audio, off, Options{UseAudio: true, WhisperAPIKey: "k"})
	if err != nil {
		t.Fatalf("sync: %v", err)
	}
	if gotPath != "/v1/audio/transcriptions" {
		t.Fatalf("unexpected path %s", gotPath)
	}
	if len(items) == 0 || items[0].StartAt != time.Second {
		t.Fatalf("unexpected start time: %v", items[0].StartAt)
	}
}
