package transcriber

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

// whisperModel is the OpenAI model used for transcriptions.
var whisperModel = openai.Whisper1

// baseURL points to the OpenAI API. Tests may override it.
var baseURL = "https://api.openai.com/v1"

// SetWhisperModel overrides the default Whisper model for testing purposes.
func SetWhisperModel(m string) {
	whisperModel = m
}

// SetBaseURL overrides the OpenAI API base URL.
func SetBaseURL(u string) {
	baseURL = u
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
