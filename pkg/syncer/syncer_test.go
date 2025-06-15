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

	"github.com/jdfalk/subtitle-manager/pkg/audio"
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

// TestSyncEmbedded verifies that embedded subtitle tracks are used to adjust
// timing.
func TestSyncEmbedded(t *testing.T) {
	dir := t.TempDir()
	script := filepath.Join(dir, "ffmpeg")
	data := `#!/bin/sh
last=$(eval echo \${$#}); cp ../../testdata/simple.srt "$last"
`
	if err := os.WriteFile(script, []byte(data), 0755); err != nil {
		t.Fatalf("write script: %v", err)
	}
	oldPath := os.Getenv("PATH")
	os.Setenv("PATH", dir+":"+oldPath)
	defer os.Setenv("PATH", oldPath)

	subtitles.SetFFmpegPath("ffmpeg")

	items, err := Sync("dummy.mkv", "../../testdata/simple_offset.srt", Options{UseEmbedded: true, SubtitleTracks: []int{0}})
	if err != nil {
		t.Fatalf("sync: %v", err)
	}
	if items[0].StartAt != time.Second {
		t.Fatalf("unexpected start %v", items[0].StartAt)
	}
}

// TestSyncAudio verifies that audio transcription is used when requested.
func TestSyncAudio(t *testing.T) {
	ffdir := t.TempDir()
	script := filepath.Join(ffdir, "ffmpeg")
	data := `#!/bin/sh
last=$(eval echo \${$#}); cp ../../testdata/simple.srt "$last"
`
	if err := os.WriteFile(script, []byte(data), 0755); err != nil {
		t.Fatalf("write script: %v", err)
	}
	oldPath := os.Getenv("PATH")
	os.Setenv("PATH", ffdir+":"+oldPath)
	defer os.Setenv("PATH", oldPath)
	audio.SetFFmpegPath("ffmpeg")

	var gotPath string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		fmt.Fprint(w, "1\n00:00:01,000 --> 00:00:02,000\nHello\n")
	}))
	defer srv.Close()
	transcriber.SetBaseURL(srv.URL + "/v1")
	defer transcriber.SetBaseURL("https://api.openai.com/v1")

	items, err := Sync("dummy.mkv", "../../testdata/simple_offset.srt", Options{UseAudio: true, AudioTrack: 0, OpenAIKey: "k"})
	if err != nil {
		t.Fatalf("sync: %v", err)
	}
	if gotPath != "/v1/audio/transcriptions" {
		t.Fatalf("unexpected path %s", gotPath)
	}
	if items[0].StartAt != time.Second {
		t.Fatalf("unexpected start %v", items[0].StartAt)
	}
}
