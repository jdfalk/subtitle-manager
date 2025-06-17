package providers

import (
	"context"
	"fmt"
	"time"
)

// FetchFromAll tries each known provider in order until one returns a subtitle.
// It uses an increasing delay between provider attempts to avoid rapid retries.
// The provider API key is reused when applicable. The name of the provider that
// succeeded is returned along with the subtitle bytes.
func FetchFromAll(ctx context.Context, mediaPath, lang, key string) ([]byte, string, error) {
	names := All()
	delay := time.Second
	for i, name := range names {
		p, err := Get(name, key)
		if err != nil {
			continue
		}
		c, cancel := context.WithTimeout(ctx, 15*time.Second)
		data, err := p.Fetch(c, mediaPath, lang)
		cancel()
		if err == nil {
			return data, name, nil
		}
		if ctx.Err() != nil {
			return nil, "", ctx.Err()
		}
		time.Sleep(time.Duration(i+1) * delay)
	}
	return nil, "", fmt.Errorf("no subtitle found")
}
