// file: pkg/renamer/renamer_test.go
// version: 1.23.1
// guid: 3b4a1244-0d0c-4f05-9f10-d62d74d3117b

package renamer

import (
	"os"
	"path/filepath"
	"testing"
)

// TestRename exercises subtitle rename scenarios to cover error and success paths.
func TestRename(t *testing.T) {
	tests := []struct {
		name       string
		lang       string
		setupFiles []string
		wantErr    bool
		wantRename bool
	}{
		{
			name:       "no subtitle files",
			lang:       "en",
			setupFiles: nil,
			wantErr:    false,
			wantRename: false,
		},
		{
			name:       "already matching name",
			lang:       "en",
			setupFiles: []string{"movie.en.srt"},
			wantErr:    false,
			wantRename: false,
		},
		{
			name:       "rename single subtitle",
			lang:       "en",
			setupFiles: []string{"other.en.srt"},
			wantErr:    false,
			wantRename: true,
		},
		{
			name:       "multiple subtitle files",
			lang:       "en",
			setupFiles: []string{"one.en.srt", "two.en.srt"},
			wantErr:    true,
			wantRename: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange.
			tmpDir := t.TempDir()
			videoPath := filepath.Join(tmpDir, "movie.mkv")
			for _, name := range tt.setupFiles {
				path := filepath.Join(tmpDir, name)
				if err := os.WriteFile(path, []byte("subtitle"), 0o644); err != nil {
					t.Fatalf("create subtitle file: %v", err)
				}
			}

			// Act.
			err := Rename(videoPath, tt.lang)

			// Assert.
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if tt.wantRename {
				newPath := filepath.Join(tmpDir, "movie."+tt.lang+".srt")
				if _, statErr := os.Stat(newPath); statErr != nil {
					t.Fatalf("expected renamed subtitle at %s: %v", newPath, statErr)
				}
			}
		})
	}
}

// TestRenameReturnsRenameError validates rename failures propagate to callers.
func TestRenameReturnsRenameError(t *testing.T) {
	// Arrange.
	tmpDir := t.TempDir()
	videoPath := filepath.Join(tmpDir, "movie.mkv")
	sourcePath := filepath.Join(tmpDir, "other.en.srt")
	destPath := filepath.Join(tmpDir, "movie.en.srt")

	if err := os.WriteFile(sourcePath, []byte("subtitle"), 0o644); err != nil {
		t.Fatalf("create subtitle file: %v", err)
	}
	if err := os.Mkdir(destPath, 0o755); err != nil {
		t.Fatalf("create destination directory: %v", err)
	}

	// Act.
	err := Rename(videoPath, "en")

	// Assert.
	if err == nil {
		t.Fatalf("expected error for rename failure")
	}
}
