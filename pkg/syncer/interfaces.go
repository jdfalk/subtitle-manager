// file: pkg/syncer/interfaces.go

package syncer

import (
	"github.com/asticode/go-astisub"
)

// Transcriber interface defines the contract for audio transcription services.
// This interface allows for dependency injection and easier testing.
type Transcriber interface {
	// Transcribe converts audio from the specified file path to subtitle text.
	// Parameters:
	//   - path: file path to the audio/video file
	//   - lang: language code for transcription (empty for auto-detection)
	//   - apiKey: API key for the transcription service
	// Returns:
	//   - []byte: SRT subtitle content as bytes
	//   - error: any error that occurred during transcription
	Transcribe(path, lang, apiKey string) ([]byte, error)
}

// SubtitleExtractor interface defines the contract for extracting embedded subtitles.
// This interface allows for dependency injection and easier testing.
type SubtitleExtractor interface {
	// ExtractTrack extracts subtitle items from a specific track in the media file.
	// Parameters:
	//   - mediaPath: path to the media file containing subtitles
	//   - track: zero-based index of the subtitle track to extract
	// Returns:
	//   - []*astisub.Item: array of subtitle items with timing and text
	//   - error: any error that occurred during extraction
	ExtractTrack(mediaPath string, track int) ([]*astisub.Item, error)
}
