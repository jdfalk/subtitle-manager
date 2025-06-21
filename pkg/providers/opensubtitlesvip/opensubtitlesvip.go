// file: pkg/providers/opensubtitlesvip/opensubtitlesvip.go
package opensubtitlesvip

import (
	"context"
	"fmt"
	"io"

	"github.com/oz/osdb"
)

// Client implements the providers.Provider interface for Opensubtitlesvip.
// It uses the osdb SDK to search and download subtitles.
type Client struct {
	api api
}

type api interface {
	LogIn(string, string, string) error
	FileSearch(string, []string) (osdb.Subtitles, error)
	DownloadSubtitles(osdb.Subtitles) ([]osdb.SubtitleFile, error)
}

// New returns a Client configured with reasonable defaults.
func New() *Client {
	c, err := osdb.NewClient()
	if err != nil {
		// Fallback to nil API which will cause errors on use
		return &Client{}
	}
	return &Client{api: c}
}

// Fetch downloads the subtitle for mediaPath in lang.
// It returns the subtitle bytes or an error.
func (c *Client) Fetch(ctx context.Context, mediaPath, lang string) ([]byte, error) {
	if c.api == nil {
		return nil, fmt.Errorf("client not initialized")
	}
	if err := c.api.LogIn("", "", lang); err != nil {
		return nil, err
	}
	subs, err := c.api.FileSearch(mediaPath, []string{lang})
	if err != nil {
		return nil, err
	}
	if len(subs) == 0 {
		return nil, fmt.Errorf("no subtitles found")
	}
	files, err := c.api.DownloadSubtitles(subs[:1])
	if err != nil {
		return nil, err
	}
	if len(files) == 0 {
		return nil, fmt.Errorf("no subtitle file returned")
	}
	r, err := files[0].Reader()
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return io.ReadAll(r)
}
