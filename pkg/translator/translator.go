package translator

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

var ErrUnsupportedService = errors.New("unsupported translation service")

var googleAPIURL = "https://translation.googleapis.com/language/translate/v2"

// SetGoogleAPIURL overrides the Google Translate API URL (useful for testing).
func SetGoogleAPIURL(u string) {
	googleAPIURL = u
}

// TranslateFunc defines the function signature for translation services.
type TranslateFunc func(text, targetLang, apiKey string) (string, error)

// GoogleTranslate translates text using Google Translate API.
func GoogleTranslate(text, targetLang, apiKey string) (string, error) {
	form := url.Values{}
	form.Set("q", text)
	form.Set("target", targetLang)
	form.Set("key", apiKey)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, googleAPIURL, strings.NewReader(form.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var gr struct {
		Data struct {
			Translations []struct {
				TranslatedText string `json:"translatedText"`
			} `json:"translations"`
		} `json:"data"`
	}
	if err := json.Unmarshal(body, &gr); err != nil {
		return "", err
	}
	if len(gr.Data.Translations) == 0 {
		return "", fmt.Errorf("no translations")
	}
	return gr.Data.Translations[0].TranslatedText, nil
}

// GPTTranslate translates text using the ChatGPT API.
func GPTTranslate(text, targetLang, apiKey string) (string, error) {
	client := openai.NewClient(apiKey)
	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: fmt.Sprintf("Translate the following text to %s. Return only the translated text.", targetLang),
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: text,
			},
		},
	}
	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return "", err
	}
	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no translation returned")
	}
	return strings.TrimSpace(resp.Choices[0].Message.Content), nil
}

var providers = map[string]TranslateFunc{
	"google":  GoogleTranslate,
	"gpt":     GPTTranslate,
	"chatgpt": GPTTranslate,
}

// Translate selects a provider and performs translation.
// googleKey and gptKey are used depending on the provider.
func Translate(service, text, targetLang, googleKey, gptKey string) (string, error) {
	fn, ok := providers[service]
	if !ok {
		return "", ErrUnsupportedService
	}
	key := googleKey
	if service == "gpt" || service == "chatgpt" {
		key = gptKey
	}
	return fn(text, targetLang, key)
}
