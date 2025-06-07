package subtitles

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/asticode/go-astisub"
)

// ExtractFromMedia extracts the first subtitle stream from the given media
// container using the `ffmpeg` command line tool. The resulting subtitle items
// are returned. The `ffmpeg` binary must be available in $PATH.
func ExtractFromMedia(mediaPath string) ([]*astisub.Item, error) {
	tmp, err := os.CreateTemp("", "subextract-*.srt")
	if err != nil {
		return nil, err
	}
	tmp.Close()
	defer os.Remove(tmp.Name())

	cmd := exec.CommandContext(context.Background(), "ffmpeg", "-y", "-i", mediaPath, "-map", "0:s:0", tmp.Name())
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
