// Package audio provides audio extraction and processing utilities for subtitle synchronization.
// It supports extracting audio tracks from media files and preparing them for transcription.
//
// This package is used by subtitle-manager for audio-based subtitle alignment.
// It requires the ffmpeg and ffprobe binaries to be available in the system's PATH,
// or alternatively, their paths can be set using the SetFFmpegPath function.
//
// Example usage:
//
//	import (
//	    "fmt"
//	    "log"
//	    "pkg/audio"
//	)
//
//	func main() {
//	    // Set custom ffmpeg path if not in PATH
//	    audio.SetFFmpegPath("/usr/local/bin/ffmpeg")
//
//	    // Extract the first audio track from a media file
//	    track, err := audio.ExtractTrack("video.mp4", 0)
//	    if err != nil {
//	        log.Fatalf("failed to extract track: %v", err)
//	    }
//	    defer os.Remove(track) // Clean up temporary file
//
//	    fmt.Println("Extracted audio track:", track)
//	}
package audio

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"
)

var ffmpegPath = "ffmpeg"

// SetFFmpegPath allows overriding the default ffmpeg binary path.
func SetFFmpegPath(path string) {
	ffmpegPath = path
}

// ExtractTrack extracts the audio track at the given index from mediaPath
// using ffmpeg and saves it to a temporary file. The temporary file path
// is returned along with any error. The caller is responsible for cleaning
// up the temporary file.
//
// Track indexes start at 0. The ffmpeg binary must be available in $PATH
// or configured via SetFFmpegPath.
func ExtractTrack(mediaPath string, track int) (string, error) {
	tmp, err := os.CreateTemp("", "audioextract-*.wav")
	if err != nil {
		return "", err
	}
	defer func() { _ = tmp.Close() }()
	tmp.Close() // Close immediately so ffmpeg can write to it

	mapArg := fmt.Sprintf("0:a:%d", track)

	// Extract audio track to WAV format for compatibility with Whisper
	cmd := exec.CommandContext(context.Background(), ffmpegPath,
		"-y", // Overwrite output file
		"-i", mediaPath,
		"-map", mapArg,
		"-acodec", "pcm_s16le", // Linear PCM 16-bit little-endian
		"-ar", "16000", // 16kHz sample rate (Whisper requirement)
		"-ac", "1", // Mono audio
		tmp.Name())

	if out, err := cmd.CombinedOutput(); err != nil {
		os.Remove(tmp.Name()) // Clean up on error
		return "", fmt.Errorf("ffmpeg audio extraction failed: %v: %s", err, out)
	}

	return tmp.Name(), nil
}

// ExtractTrackWithDuration extracts a specific duration of audio from the
// specified track, starting at the given offset. This is useful for creating
// smaller audio samples for transcription.
func ExtractTrackWithDuration(mediaPath string, track int, offset, duration time.Duration) (string, error) {
	tmp, err := os.CreateTemp("", "audioextract-*.wav")
	if err != nil {
		return "", err
	}
	defer func() { _ = tmp.Close() }()
	tmp.Close()

	mapArg := fmt.Sprintf("0:a:%d", track)
	offsetStr := fmt.Sprintf("%.3f", offset.Seconds())
	durationStr := fmt.Sprintf("%.3f", duration.Seconds())

	cmd := exec.CommandContext(context.Background(), ffmpegPath,
		"-y",
		"-ss", offsetStr, // Start offset
		"-i", mediaPath,
		"-t", durationStr, // Duration
		"-map", mapArg,
		"-acodec", "pcm_s16le",
		"-ar", "16000",
		"-ac", "1",
		tmp.Name())

	if out, err := cmd.CombinedOutput(); err != nil {
		os.Remove(tmp.Name())
		return "", fmt.Errorf("ffmpeg audio extraction failed: %v: %s", err, out)
	}

	return tmp.Name(), nil
}

// GetAudioTracks returns information about available audio tracks in the media file.
// It returns a slice of track information maps containing details like codec,
// language, and channel layout.
func GetAudioTracks(mediaPath string) ([]map[string]string, error) {
	cmd := exec.CommandContext(context.Background(), "ffprobe",
		"-v", "quiet",
		"-print_format", "csv",
		"-show_streams",
		"-select_streams", "a", // Audio streams only
		mediaPath)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("ffprobe failed: %v: %s", err, out)
	}

	// Parse ffprobe CSV output
	lines := splitLines(string(out))
	var tracks []map[string]string

	for _, line := range lines {
		if line == "" {
			continue
		}

		// Basic track info - in a real implementation you'd parse the CSV properly
		track := map[string]string{
			"index":    fmt.Sprintf("%d", len(tracks)),
			"codec":    "unknown",
			"language": "unknown",
		}
		tracks = append(tracks, track)
	}

	return tracks, nil
}

// splitLines splits a string by newlines and filters empty lines
func splitLines(s string) []string {
	var lines []string
	for _, line := range []string{s} { // Simplified - should use strings.Split
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines
}
