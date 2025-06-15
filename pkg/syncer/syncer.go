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
	// WhisperKey provides the API key for audio transcription when
	// UseAudio is true.
	WhisperKey string
	// AudioWeight controls the influence of audio alignment versus embedded
	// subtitles. A value between 0 and 1 is expected. When zero, 0.7 is
	// used.
	AudioWeight float64
}

// transcribeFn wraps the audio transcription function. Tests may override it.
var transcribeFn = transcriber.WhisperTranscribe

// extractFn wraps subtitle track extraction. Tests may override it.
var extractFn = subtitles.ExtractSubtitleTrack

// SetTranscribeFunc overrides the default audio transcription function.
func SetTranscribeFunc(fn func(string, string, string) ([]byte, error)) {
	transcribeFn = fn
}

// SetExtractFunc overrides the default subtitle extraction function.
func SetExtractFunc(fn func(string, int) ([]*astisub.Item, error)) {
	extractFn = fn
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

	weight := opts.AudioWeight
	if weight == 0 {
		weight = 0.7
	}

	var total time.Duration
	var applied float64

	if opts.UseAudio {
		b, err := transcribeFn(mediaPath, "", opts.WhisperKey)
		if err == nil {
			refSub, err := astisub.ReadFromSRT(bytes.NewReader(b))
			if err == nil {
				off := computeOffset(refSub.Items, items)
				total += time.Duration(float64(off) * weight)
				applied += weight
			}
		}
	}

	if opts.UseEmbedded {
		tracks := opts.SubtitleTracks
		if len(tracks) == 0 {
			tracks = []int{0}
		}
		if len(tracks) > 0 {
			per := (1 - weight) / float64(len(tracks))
			for _, t := range tracks {
				refItems, err := extractFn(mediaPath, t)
				if err == nil {
					off := computeOffset(refItems, items)
					total += time.Duration(float64(off) * per)
					applied += per
				}
			}
		}
	}

	if applied > 0 {
		offset := time.Duration(float64(total) / applied)
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

// computeOffset returns the average difference between the start times of ref
// and target. Up to the first five items are compared.
func computeOffset(ref, target []*astisub.Item) time.Duration {
	n := len(ref)
	if len(target) < n {
		n = len(target)
	}
	if n > 5 {
		n = 5
	}
	if n == 0 {
		return 0
	}
	var sum time.Duration
	for i := 0; i < n; i++ {
		sum += ref[i].StartAt - target[i].StartAt
	}
	return sum / time.Duration(n)
}
