// file: pkg/audio/audio_test.go
package audio

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"
)

// TestExtractTrackWithFile verifies that audio track extraction works correctly.
// This test requires ffmpeg to be available in the PATH.
func TestExtractTrackWithFile(t *testing.T) {
	// Skip test if ffmpeg is not available
	if !isFFmpegAvailable() {
		t.Skip("ffmpeg not available, skipping audio extraction test")
	}

	// Create a minimal test media file (not implemented for this test)
	// In practice, you'd use a test fixture
	mediaPath := "../../testdata/sample.mkv"
	if _, err := os.Stat(mediaPath); os.IsNotExist(err) {
		t.Skip("test media file not available")
	}

	audioFile, err := ExtractTrack(mediaPath, 0)
	if err != nil {
		t.Fatalf("ExtractTrack failed: %v", err)
	}
	defer os.Remove(audioFile)

	// Verify the output file exists and has content
	info, err := os.Stat(audioFile)
	if err != nil {
		t.Fatalf("extracted audio file does not exist: %v", err)
	}

	if info.Size() == 0 {
		t.Fatal("extracted audio file is empty")
	}

	t.Logf("successfully extracted audio track to %s (%d bytes)", audioFile, info.Size())
}

// TestExtractTrackWithDuration verifies duration-limited audio extraction.
func TestExtractTrackWithDuration(t *testing.T) {
	if !isFFmpegAvailable() {
		t.Skip("ffmpeg not available, skipping audio extraction test")
	}

	mediaPath := "../../testdata/sample.mkv"
	if _, err := os.Stat(mediaPath); os.IsNotExist(err) {
		t.Skip("test media file not available")
	}

	audioFile, err := ExtractTrackWithDuration(mediaPath, 0, 10*time.Second, 30*time.Second)
	if err != nil {
		t.Fatalf("ExtractTrackWithDuration failed: %v", err)
	}
	defer os.Remove(audioFile)

	info, err := os.Stat(audioFile)
	if err != nil {
		t.Fatalf("extracted audio file does not exist: %v", err)
	}

	if info.Size() == 0 {
		t.Fatal("extracted audio file is empty")
	}

	t.Logf("successfully extracted audio segment to %s (%d bytes)", audioFile, info.Size())
}

// TestGetAudioTracks verifies that audio track information can be retrieved.
func TestGetAudioTracks(t *testing.T) {
	if !isFFmpegAvailable() {
		t.Skip("ffprobe not available, skipping audio track info test")
	}

	mediaPath := "../../testdata/sample.mkv"
	if _, err := os.Stat(mediaPath); os.IsNotExist(err) {
		t.Skip("test media file not available")
	}

	tracks, err := GetAudioTracks(mediaPath)
	if err != nil {
		t.Fatalf("GetAudioTracks failed: %v", err)
	}

	if len(tracks) == 0 {
		t.Fatal("no audio tracks found")
	}

	t.Logf("found %d audio tracks", len(tracks))
	for i, track := range tracks {
		t.Logf("track %d: %+v", i, track)
	}
}

// TestSetFFmpegPath verifies that the ffmpeg path can be set and retrieved.
func TestSetFFmpegPath(t *testing.T) {
	originalPath := ffmpegPath
	defer func() { ffmpegPath = originalPath }() // Restore original path

	customPath := "/custom/path/to/ffmpeg"
	SetFFmpegPath(customPath)

	if ffmpegPath != customPath {
		t.Errorf("expected ffmpeg path to be %s, got %s", customPath, ffmpegPath)
	}
}

// TestSplitLines verifies the splitLines utility function.
func TestSplitLines(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "single line",
			input:    "line1",
			expected: []string{"line1"},
		},
		{
			name:     "empty string",
			input:    "",
			expected: []string{},
		},
		{
			name:     "whitespace only",
			input:    "   ",
			expected: []string{"   "},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := splitLines(tt.input)
			if len(result) != len(tt.expected) {
				t.Errorf("expected %d lines, got %d", len(tt.expected), len(result))
				return
			}
			for i, line := range result {
				if line != tt.expected[i] {
					t.Errorf("line %d: expected %q, got %q", i, tt.expected[i], line)
				}
			}
		})
	}
}

// TestExtractTrackErrorHandling tests error conditions for ExtractTrack.
func TestExtractTrackErrorHandling(t *testing.T) {
	// Test with non-existent file
	_, err := ExtractTrack("nonexistent.mkv", 0)
	if err == nil {
		t.Error("expected error for non-existent file, got nil")
	}

	// Test with invalid track index (requires ffmpeg to be available)
	if isFFmpegAvailable() {
		// Create a temporary empty file to test with
		tmpFile, err := os.CreateTemp("", "test-*.mp4")
		if err != nil {
			t.Fatalf("failed to create temp file: %v", err)
		}
		defer os.Remove(tmpFile.Name())
		tmpFile.Close()

		_, err = ExtractTrack(tmpFile.Name(), 999) // Invalid track index
		if err == nil {
			t.Error("expected error for invalid track index, got nil")
		}
	}
}

// TestExtractTrackWithDurationErrorHandling tests error conditions.
func TestExtractTrackWithDurationErrorHandling(t *testing.T) {
	// Test with non-existent file
	_, err := ExtractTrackWithDuration("nonexistent.mkv", 0, 0, 10*time.Second)
	if err == nil {
		t.Error("expected error for non-existent file, got nil")
	}
}

// TestGetAudioTracksErrorHandling tests error handling for GetAudioTracks.
func TestGetAudioTracksErrorHandling(t *testing.T) {
	// Test with non-existent file
	_, err := GetAudioTracks("nonexistent.mkv")
	if err == nil {
		t.Error("expected error for non-existent file, got nil")
	}
}

// TestGetAudioTracksWithMockData tests GetAudioTracks with predictable data.
func TestGetAudioTracksWithMockData(t *testing.T) {
	dir := t.TempDir()
	script := filepath.Join(dir, "ffprobe")
	data := `#!/bin/sh
echo '{"streams":[{"index":0,"codec_name":"aac","channels":2,"tags":{"language":"eng"}},{"index":1,"codec_name":"aac","channels":6,"tags":{"language":"spa"}}]}'
`
	if err := os.WriteFile(script, []byte(data), 0755); err != nil {
		t.Fatalf("write script: %v", err)
	}
	old := os.Getenv("PATH")
	os.Setenv("PATH", dir+":"+old)
	defer os.Setenv("PATH", old)

	tracks, err := GetAudioTracks("dummy.mkv")
	if err != nil {
		t.Fatalf("GetAudioTracks returned error: %v", err)
	}
	if len(tracks) != 2 {
		t.Fatalf("expected 2 tracks, got %d", len(tracks))
	}
	if tracks[0]["language"] != "eng" || tracks[1]["language"] != "spa" {
		t.Errorf("unexpected languages: %v", tracks)
	}
}

// isFFmpegAvailable checks if ffmpeg is available in the system PATH.
func isFFmpegAvailable() bool {
	_, err := os.Stat(ffmpegPath)
	if err == nil {
		return true
	}

	// Try to find ffmpeg in PATH
	_, err = exec.LookPath("ffmpeg")
	return err == nil
}
