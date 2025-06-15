package subtitles

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/asticode/go-astisub"
)

// ffmpegPath is the name or path of the ffmpeg binary used for extraction.
var ffmpegPath = "ffmpeg"

// SetFFmpegPath allows tests or callers to override the ffmpeg binary path.
func SetFFmpegPath(path string) {
	ffmpegPath = path
}

// ExtractFromMedia extracts the first subtitle stream from the given media
// container using ffmpeg. It is a convenience wrapper around ExtractTrack with
// track index 0.
func ExtractFromMedia(mediaPath string) ([]*astisub.Item, error) {
	return ExtractTrack(mediaPath, 0)
}

// ExtractTrack extracts the specified subtitle track from the given media
// container using the `ffmpeg` command line tool. The resulting subtitle items
// are returned. The `ffmpeg` binary must be available in $PATH.
func ExtractTrack(mediaPath string, track int) ([]*astisub.Item, error) {
	tmp, err := os.CreateTemp("", "subextract-*.srt")
	if err != nil {
		return nil, err
	}
	tmp.Close()
	defer os.Remove(tmp.Name())

	cmd := exec.CommandContext(context.Background(), ffmpegPath, "-y", "-i", mediaPath, "-map", fmt.Sprintf("0:s:%d", track), tmp.Name())
	if out, err := cmd.CombinedOutput(); err != nil {
		return nil, fmt.Errorf("ffmpeg: %v: %s", err, out)
	}

	sub, err := astisub.OpenFile(tmp.Name())
	if err != nil {
		return nil, err
	}
	items := make([]*astisub.Item, len(sub.Items))
	copy(items, sub.Items)
	return items, nil
}
