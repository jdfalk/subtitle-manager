package transcriber

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

// TranscriptionMethod represents the method used for transcription
type TranscriptionMethod string

const (
	// MethodOpenAI uses the OpenAI Whisper API
	MethodOpenAI TranscriptionMethod = "openai"
	// MethodDocker uses a local Docker container
	MethodDocker TranscriptionMethod = "docker"
)

// whisperModel is the OpenAI model used for transcriptions.
var whisperModel = openai.Whisper1

// baseURL points to the OpenAI-compatible API. Tests and configuration may override it.
var baseURL = "https://api.openai.com/v1"

// SetWhisperModel overrides the default Whisper model for testing purposes.
func SetWhisperModel(m string) {
	whisperModel = m
}

// SetBaseURL overrides the OpenAI API base URL used for transcription requests.
func SetBaseURL(u string) {
	baseURL = u
}

// TranscribeWithMethod transcribes a media file using the specified method.
// This function provides a unified interface for both OpenAI API and Docker-based transcription.
func TranscribeWithMethod(ctx context.Context, method TranscriptionMethod, path, lang, apiKey string, dockerConfig *DockerTranscriberConfig) ([]byte, error) {
	switch method {
	case MethodOpenAI:
		return WhisperTranscribe(path, lang, apiKey)
	case MethodDocker:
		transcriber, err := NewDockerTranscriber(dockerConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to create Docker transcriber: %w", err)
		}
		defer transcriber.Close()

		if !transcriber.IsAvailable(ctx) {
			return nil, fmt.Errorf("Docker is not available")
		}

		return transcriber.TranscribeFile(ctx, path)
	default:
		return nil, fmt.Errorf("unsupported transcription method: %s", method)
	}
}

// WhisperTranscribe transcribes the media file at path using the Whisper API.
// The language code may be empty to enable auto detection. The API key is
// required. It returns the SRT subtitle bytes produced by the service.
func WhisperTranscribe(path, lang, apiKey string) ([]byte, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("api key required")
	}
	cfg := openai.DefaultConfig(apiKey)
	cfg.BaseURL = baseURL
	client := openai.NewClientWithConfig(cfg)
	resp, err := client.CreateTranscription(context.Background(), openai.AudioRequest{
		Model:    whisperModel,
		FilePath: path,
		Language: lang,
		Format:   openai.AudioResponseFormatSRT,
	})
	if err != nil {
		return nil, err
	}
	return []byte(resp.Text), nil
}
