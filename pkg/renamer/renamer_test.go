package renamer

import (
	"os"
	"path/filepath"
	"testing"
)

// TestRename ensures subtitle files are renamed to match the video file.
func TestRename(t *testing.T) {
	dir := t.TempDir()
	video := filepath.Join(dir, "new.mkv")
	if err := os.WriteFile(video, []byte("x"), 0644); err != nil {
		t.Fatalf("write video: %v", err)
	}
	subOld := filepath.Join(dir, "old.en.srt")
	if err := os.WriteFile(subOld, []byte("y"), 0644); err != nil {
		t.Fatalf("write sub: %v", err)
	}
	if err := Rename(video, "en"); err != nil {
		t.Fatalf("rename: %v", err)
	}
	newSub := filepath.Join(dir, "new.en.srt")
	if _, err := os.Stat(newSub); err != nil {
		t.Fatalf("expected %s: %v", newSub, err)
	}
	if _, err := os.Stat(subOld); !os.IsNotExist(err) {
		t.Fatalf("old subtitle still exists")
	}
}

// TestRenameMultiple ensures an error is returned when more than one subtitle matches.
func TestRenameMultiple(t *testing.T) {
	dir := t.TempDir()
	video := filepath.Join(dir, "movie.mkv")
	if err := os.WriteFile(video, []byte("x"), 0644); err != nil {
		t.Fatalf("write video: %v", err)
	}
	for _, name := range []string{"one.en.srt", "two.en.srt"} {
		if err := os.WriteFile(filepath.Join(dir, name), []byte("y"), 0644); err != nil {
			t.Fatalf("write sub: %v", err)
		}
	}
	if err := Rename(video, "en"); err == nil {
		t.Fatal("expected error for multiple subtitles")
	}
}
