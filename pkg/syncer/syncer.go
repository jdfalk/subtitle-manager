package syncer

import (
	"time"

	"github.com/asticode/go-astisub"
	"github.com/jdfalk/subtitle-manager/pkg/translator"
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
	// TargetLang specifies the language to translate subtitles into. If empty
	// no translation is performed.
	TargetLang string
	// Service selects the translation provider when TargetLang is set. Valid
	// values are "google", "gpt" or "grpc".
	Service string
	// GoogleKey holds the API key for Google Translate when Service is
	// "google".
	GoogleKey string
	// GPTKey holds the OpenAI API key when Service is "gpt".
	GPTKey string
	// GRPCAddr specifies the gRPC translator address when Service is "grpc".
	GRPCAddr string
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
	if opts.TargetLang != "" {
		items, err = Translate(items, opts.TargetLang, opts.Service, opts.GoogleKey, opts.GPTKey, opts.GRPCAddr)
		if err != nil {
			return nil, err
		}
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

// Translate converts each subtitle item to lang using the selected service.
// googleKey, gptKey and grpcAddr are passed to the underlying translator
// depending on service. The returned slice contains translated items in the
// same order as the input.
func Translate(items []*astisub.Item, lang, service, googleKey, gptKey, grpcAddr string) ([]*astisub.Item, error) {
	out := make([]*astisub.Item, len(items))
	for i, it := range items {
		t, err := translator.Translate(service, it.String(), lang, googleKey, gptKey, grpcAddr)
		if err != nil {
			return nil, err
		}
		c := *it
		c.Lines = []astisub.Line{{Items: []astisub.LineItem{{Text: t}}}}
		out[i] = &c
	}
	return out, nil
}
