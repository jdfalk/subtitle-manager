package subtitles

import (
	"fmt"

	"github.com/asticode/go-astisub"
)

// ExtractFromMedia extracts subtitle streams from the given media container.
// Currently not implemented and returns an error.
func ExtractFromMedia(mediaPath string) ([]astisub.Item, error) {
	return nil, fmt.Errorf("extraction not implemented")
}
