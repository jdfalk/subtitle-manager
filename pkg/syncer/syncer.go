package syncer

import (
	"bytes"
	"os"
	"time"

	"github.com/asticode/go-astisub"

	"github.com/jdfalk/subtitle-manager/pkg/audio"
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
	// OpenAIKey provides the API key used for Whisper transcription when
	// UseAudio is enabled.
	OpenAIKey string
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
	items := make([]*astisub.Item, len(sub.Items))
	copy(items, sub.Items)

	var ref []*astisub.Item

	if opts.UseEmbedded {
		for _, tr := range opts.SubtitleTracks {
			r, err := subtitles.ExtractTrack(mediaPath, tr)
			if err != nil {
				return nil, err
			}
			ref = r
			break
		}
	}

	if opts.UseAudio {
		aPath, err := audio.ExtractTrack(mediaPath, opts.AudioTrack)
		if err != nil {
			return nil, err
		}
		defer os.Remove(aPath)
		data, err := transcriber.WhisperTranscribe(aPath, "", opts.OpenAIKey)
		if err != nil {
			return nil, err
		}
		tsub, err := astisub.ReadFromSRT(bytes.NewReader(data))
		if err != nil {
			return nil, err
		}
		if len(ref) == 0 {
			ref = tsub.Items
		}
	}

	if len(ref) > 0 && len(items) > 0 {
		offset := ref[0].StartAt - items[0].StartAt
		items = Shift(items, offset)
	}

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
