package syncer

import (
	"bytes"
	"time"

	"github.com/asticode/go-astisub"

	"github.com/jdfalk/subtitle-manager/pkg/subtitles"
	"github.com/jdfalk/subtitle-manager/pkg/transcriber"
)

// Options controls how the synchronization process behaves.
type Options struct {
	// UseAudio indicates whether the audio track should be analyzed.
	UseAudio bool
	// UseEmbedded indicates whether embedded subtitles should be used.
	UseEmbedded bool
	// AudioTrack selects the audio track index if UseAudio is true.
	AudioTrack int
	// SubtitleTracks selects embedded subtitle track indices when UseEmbedded is true.
	SubtitleTracks []int
	// WhisperAPIKey provides the key used when transcribing audio.
	WhisperAPIKey string
	// Language specifies the language code for transcription.
	Language string
}

// Sync attempts to synchronize the subtitle at subPath with the media file at
// mediaPath according to opts. The resulting subtitle items are returned.
//
// This is an early implementation that simply loads the subtitle file without
// performing actual alignment. Future versions will analyze the selected audio
// and subtitle tracks to automatically adjust timing.
func Sync(mediaPath, subPath string, opts Options) ([]*astisub.Item, error) {
	sub, err := astisub.OpenFile(subPath)
	if err != nil {
		return nil, err
	}

	var ref []*astisub.Item

	if opts.UseEmbedded {
		if emb, err := subtitles.ExtractFromMedia(mediaPath); err == nil && len(emb) > 0 {
			ref = emb
		}
	}

	if ref == nil && opts.UseAudio {
		if data, err := transcriber.WhisperTranscribe(mediaPath, opts.Language, opts.WhisperAPIKey); err == nil {
			if s, err := astisub.ReadFromSRT(bytes.NewReader(data)); err == nil {
				ref = s.Items
			}
		}
	}

	offset := time.Duration(0)
	if ref != nil && len(ref) > 0 && len(sub.Items) > 0 {
		offset = ref[0].StartAt - sub.Items[0].StartAt
	}

	items := Shift(sub.Items, offset)
	return items, nil
}

// Shift adjusts each subtitle item by offset and returns the updated slice.
func Shift(items []*astisub.Item, offset time.Duration) []*astisub.Item {
	out := make([]*astisub.Item, len(items))
	for i, it := range items {
		c := *it
		c.StartAt += offset
		c.EndAt += offset
		out[i] = &c
	}
	return out
}
