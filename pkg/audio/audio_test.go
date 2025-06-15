// file: pkg/audio/audio_test.go
package audio

import (
	"os"
	"os/exec"
	"testing"
	"time"
)

// TestExtractTrack verifies that audio track extraction works correctly.
// This test requires ffmpeg to be available in the PATH.
func TestExtractTrack(t *testing.T) {
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
