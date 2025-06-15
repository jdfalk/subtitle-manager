// file: pkg/audio/extract.go
package audio

import (
	"context"
	"fmt"
	"os"
	"os/exec"
)

// ffmpegPath is the path to the ffmpeg binary used for audio extraction.
var ffmpegPath = "ffmpeg"

// SetFFmpegPath overrides the ffmpeg binary path.
func SetFFmpegPath(path string) { ffmpegPath = path }

// ExtractTrack extracts the specified audio track from mediaPath to a temporary
// WAV file and returns the file path. The caller is responsible for deleting
// the file when done.
func ExtractTrack(mediaPath string, track int) (string, error) {
	tmp, err := os.CreateTemp("", "audioextract-*.wav")
	if err != nil {
		return "", err
	}
	tmp.Close()

	args := []string{"-y", "-i", mediaPath, "-map", fmt.Sprintf("0:a:%d", track), "-vn", "-acodec", "pcm_s16le", "-ar", "16000", "-ac", "1", tmp.Name()}
	cmd := exec.CommandContext(context.Background(), ffmpegPath, args...)
	if out, err := cmd.CombinedOutput(); err != nil {
		os.Remove(tmp.Name())
		return "", fmt.Errorf("ffmpeg: %v: %s", err, out)
	}
	return tmp.Name(), nil
}
