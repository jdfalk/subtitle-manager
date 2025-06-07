package subtitles

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExtractFromMedia(t *testing.T) {
	dir := t.TempDir()
	script := filepath.Join(dir, "ffmpeg")
	data := "#!/bin/sh\ncp ../../testdata/simple.srt \"$6\"\n"
	if err := os.WriteFile(script, []byte(data), 0755); err != nil {
		t.Fatalf("write script: %v", err)
	}
	oldPath := os.Getenv("PATH")
	os.Setenv("PATH", dir+":"+oldPath)
	defer os.Setenv("PATH", oldPath)

	items, err := ExtractFromMedia("dummy.mkv")
	if err != nil {
		t.Fatalf("extract: %v", err)
	}
	if len(items) == 0 {
		t.Fatal("no items extracted")
	}
}
