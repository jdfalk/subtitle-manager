package audio

import (
	"os"
	"path/filepath"
	"testing"
)

// TestExtractTrack verifies that ExtractTrack invokes ffmpeg and returns a file path.
func TestExtractTrack(t *testing.T) {
	dir := t.TempDir()
	script := filepath.Join(dir, "ffmpeg")
	data := `#!/bin/sh
last=$(eval echo \${$#}); cp ../../testdata/simple.srt "$last"
`
	if err := os.WriteFile(script, []byte(data), 0755); err != nil {
		t.Fatalf("write script: %v", err)
	}
	old := os.Getenv("PATH")
	os.Setenv("PATH", dir+":"+old)
	defer os.Setenv("PATH", old)

	path, err := ExtractTrack("dummy.mkv", 0)
	if err != nil {
		t.Fatalf("extract: %v", err)
	}
	defer os.Remove(path)
	if fi, err := os.Stat(path); err != nil || fi.Size() == 0 {
		t.Fatalf("invalid file: %v", err)
	}
}
