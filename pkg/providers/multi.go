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
	insts := Instances()
	if len(insts) == 0 {
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

	delay := time.Second
	for i, inst := range insts {
		if IsInBackoff(inst.ID) {
			continue
		}
		p, err := Get(inst.Name, key)
		if err != nil {
			continue
		}
		c, cancel := context.WithTimeout(ctx, 15*time.Second)
		data, err := p.Fetch(c, mediaPath, lang)
		cancel()
		if err == nil {
			SetBackoff(inst.ID, 0)
			return data, inst.ID, nil
		}
		if ctx.Err() != nil {
			return nil, "", ctx.Err()
		}
		SetBackoff(inst.ID, time.Duration(i+1)*delay)
		select {
		case <-time.After(time.Duration(i+1) * delay):
		case <-ctx.Done():
			return nil, "", ctx.Err()
		}
	}
	return nil, "", fmt.Errorf("no subtitle found")
}

// FetchFromTagged limits provider attempts to instances matching all tags.
func FetchFromTagged(ctx context.Context, mediaPath, lang, key string, tags []string, tm interface {
	FilterByTags(string, []string) ([]string, error)
}) ([]byte, string, error) {
	insts, err := InstancesByTags(tm, tags)
	if err != nil {
		return nil, "", err
	}
	if len(insts) == 0 {
		return nil, "", fmt.Errorf("no subtitle found")
	}
	return fetchFromInstances(ctx, insts, mediaPath, lang, key)
}

func fetchFromInstances(ctx context.Context, insts []Instance, mediaPath, lang, key string) ([]byte, string, error) {
	delay := time.Second
	for i, inst := range insts {
		if IsInBackoff(inst.ID) {
			continue
		}
		p, err := Get(inst.Name, key)
		if err != nil {
			continue
		}
		c, cancel := context.WithTimeout(ctx, 15*time.Second)
		data, err := p.Fetch(c, mediaPath, lang)
		cancel()
		if err == nil {
			SetBackoff(inst.ID, 0)
			return data, inst.ID, nil
		}
		if ctx.Err() != nil {
			return nil, "", ctx.Err()
		}
		SetBackoff(inst.ID, time.Duration(i+1)*delay)
		select {
		case <-time.After(time.Duration(i+1) * delay):
		case <-ctx.Done():
			return nil, "", ctx.Err()
		}
	}
	return nil, "", fmt.Errorf("no subtitle found")
}
