// file: pkg/notifications/notifications.go
package notifications

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Service sends notifications to external services.
type Service struct {
	DiscordWebhook string
	TelegramToken  string
	TelegramChatID string
	EmailURL       string
	client         *http.Client
}

// New creates a Service with the provided endpoints.
func New(discordURL, telegramToken, telegramChatID, emailURL string) *Service {
	return &Service{
		DiscordWebhook: discordURL,
		TelegramToken:  telegramToken,
		TelegramChatID: telegramChatID,
		EmailURL:       emailURL,
		client:         &http.Client{Timeout: 10 * time.Second},
	}
}

// Send dispatches the message to all configured notification targets.
func (s *Service) Send(ctx context.Context, msg string) error {
	if s.DiscordWebhook != "" {
		body, _ := json.Marshal(map[string]string{"content": msg})
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, s.DiscordWebhook, bytes.NewReader(body))
		if err == nil {
			req.Header.Set("Content-Type", "application/json")
			_, err = s.client.Do(req)
		}
		if err != nil {
			return err
		}
	}
	if s.TelegramToken != "" && s.TelegramChatID != "" {
		u := "https://api.telegram.org/bot" + s.TelegramToken + "/sendMessage"
		body := url.Values{"chat_id": {s.TelegramChatID}, "text": {msg}}
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, strings.NewReader(body.Encode()))
		if err == nil {
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			_, err = s.client.Do(req)
		}
		if err != nil {
			return err
		}
	}
	if s.EmailURL != "" {
		body, _ := json.Marshal(map[string]string{"message": msg})
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, s.EmailURL, bytes.NewReader(body))
		if err == nil {
			req.Header.Set("Content-Type", "application/json")
			_, err = s.client.Do(req)
		}
		if err != nil {
			return err
		}
	}
	return nil
}
