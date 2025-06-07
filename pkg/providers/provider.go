// file: pkg/providers/provider.go
package providers

import "context"

// Provider downloads subtitles for a media file in the given language.
type Provider interface {
	// Fetch returns the subtitle bytes for the specified media file and language.
	Fetch(ctx context.Context, mediaPath, lang string) ([]byte, error)
}
