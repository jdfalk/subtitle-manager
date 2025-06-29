// file: pkg/providers/provider.go
package providers

import "context"

// Provider downloads subtitles for a media file in the given language.
type Provider interface {
	// Fetch returns the subtitle bytes for the specified media file and language.
	Fetch(ctx context.Context, mediaPath, lang string) ([]byte, error)
}

// Mock generation is now handled by the root generate.go file

// Searcher optionally exposes subtitle search functionality.
type Searcher interface {
	// Search returns download URLs for matching subtitles without fetching them.
	Search(ctx context.Context, mediaPath, lang string) ([]string, error)
}
