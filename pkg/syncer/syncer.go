package syncer

import (
	"bytes"
	"os"
	"time"

	"github.com/asticode/go-astisub"
	"github.com/jdfalk/subtitle-manager/pkg/audio"
	"github.com/jdfalk/subtitle-manager/pkg/subtitles"
	"github.com/jdfalk/subtitle-manager/pkg/transcriber"
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
	// WhisperKey provides the API key for audio transcription when
	// UseAudio is true.
	WhisperKey string
	// AudioWeight controls the influence of audio alignment versus embedded
	// subtitles. A value between 0 and 1 is expected. When zero, 0.7 is
	// used.
	AudioWeight float64
	// Translate indicates whether to translate subtitles after synchronization.
	Translate bool
	// TranslateLang specifies the target language for translation.
	TranslateLang string
	// TranslateService specifies the translation service ("google", "gpt", "grpc").
	TranslateService string
	// GoogleAPIKey provides the API key for Google Translate service.
	GoogleAPIKey string
	// GPTAPIKey provides the API key for ChatGPT translation service.
	GPTAPIKey string
	// GRPCAddr provides the address for gRPC translation service.
	GRPCAddr string

	// Transcriber is the service used for audio transcription.
	// If nil, the default transcriber.WhisperTranscribe will be used.
	Transcriber Transcriber
	// SubtitleExtractor is the service used for extracting embedded subtitles.
	// If nil, the default subtitles.ExtractSubtitleTrack will be used.
	SubtitleExtractor SubtitleExtractor
}

// defaultTranscriber wraps the transcriber.WhisperTranscribe function to implement the Transcriber interface.
type defaultTranscriber struct{}

// Transcribe implements the Transcriber interface using the default WhisperTranscribe function.
func (d defaultTranscriber) Transcribe(path, lang, apiKey string) ([]byte, error) {
	return transcriber.WhisperTranscribe(path, lang, apiKey)
}

// defaultSubtitleExtractor wraps the subtitles.ExtractSubtitleTrack function to implement the SubtitleExtractor interface.
type defaultSubtitleExtractor struct{}

// ExtractTrack implements the SubtitleExtractor interface using the default ExtractSubtitleTrack function.
func (d defaultSubtitleExtractor) ExtractTrack(mediaPath string, track int) ([]*astisub.Item, error) {
	return subtitles.ExtractSubtitleTrack(mediaPath, track)
}

// Sync attempts to synchronize the subtitle at subPath with the media file at
// mediaPath according to opts. The resulting subtitle items are returned.
//
// The function supports multiple synchronization methods:
// - Audio transcription via Whisper API for precise timing alignment
// - Embedded subtitle tracks for reference timing
// - Weighted combination of both methods for optimal results
// - Optional translation of synchronized subtitles
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

	// Use injected transcriber or fall back to default
	transcriber := opts.Transcriber
	if transcriber == nil {
		transcriber = defaultTranscriber{}
	}

	// Use injected extractor or fall back to default
	extractor := opts.SubtitleExtractor
	if extractor == nil {
		extractor = defaultSubtitleExtractor{}
	}

	if opts.UseAudio {
		// Extract specific audio track for transcription
		audioFile, err := audio.ExtractTrack(mediaPath, opts.AudioTrack)
		if err != nil {
			// Fall back to original media file if audio extraction fails
			audioFile = mediaPath
		} else {
			defer os.Remove(audioFile) // Clean up temporary audio file
		}

		audioLang := "" // Auto-detect language
		b, err := transcriber.Transcribe(audioFile, audioLang, opts.WhisperKey)
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
			// Calculate weight for embedded subtitles
			// If audio is also used, embedded gets (1 - weight)
			// If only embedded is used, it gets full weight (1.0)
			embeddedWeight := 1.0
			if opts.UseAudio {
				embeddedWeight = 1 - weight
			}
			per := embeddedWeight / float64(len(tracks))
			for _, t := range tracks {
				refItems, err := extractor.ExtractTrack(mediaPath, t)
				if err == nil {
					off := computeOffset(refItems, items)
					total += time.Duration(float64(off) * per)
					applied += per
				}
			}
		}
	}

	// Default behavior: use embedded subtitles if neither audio nor embedded is explicitly set
	if !opts.UseAudio && !opts.UseEmbedded {
		refItems, err := extractor.ExtractTrack(mediaPath, 0)
		if err == nil {
			offset := computeOffset(refItems, items)
			items = Shift(items, offset)
		}
	} else if applied > 0 {
		// Apply weighted offset if using audio and/or embedded methods
		offset := time.Duration(float64(total) / applied)
		items = Shift(items, offset)
	}

	// Apply translation if requested
	if opts.Translate && opts.TranslateLang != "" {
		// Use specified service or default to Google Translate
		service := opts.TranslateService
		if service == "" {
			service = "google"
		}

		for _, item := range items {
			for i, line := range item.Lines {
				for j, lineItem := range line.Items {
					if lineItem.Text != "" {
						translated, err := translator.Translate(service, lineItem.Text, opts.TranslateLang,
							opts.GoogleAPIKey, opts.GPTAPIKey, opts.GRPCAddr)
						if err == nil {
							lineItem.Text = translated
							line.Items[j] = lineItem
						}
						// Silently continue on translation errors to avoid breaking sync
					}
				}
				item.Lines[i] = line
			}
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
