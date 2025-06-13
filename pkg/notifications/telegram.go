// file: pkg/notifications/telegram.go
package notifications

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// TelegramNotifier sends messages to a Telegram chat.
type TelegramNotifier struct {
	// BotToken is the API token for the bot.
	BotToken string
	// ChatID is the destination chat identifier.
	ChatID string
	// Client is used to make HTTP requests.
	Client *http.Client
	// APIBase overrides the Telegram API base URL for testing.
	APIBase string
}

// Notify posts msg to the Telegram bot API.
func (t TelegramNotifier) Notify(ctx context.Context, msg string) error {
	if t.BotToken == "" || t.ChatID == "" {
		return fmt.Errorf("token and chat ID required")
	}
	c := t.Client
	if c == nil {
		c = http.DefaultClient
	}
	body, _ := json.Marshal(map[string]string{"chat_id": t.ChatID, "text": msg})
	base := t.APIBase
	if base == "" {
		base = "https://api.telegram.org"
	}
	url := fmt.Sprintf("%s/bot%s/sendMessage", base, t.BotToken)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		return fmt.Errorf("status %d", resp.StatusCode)
	}
	return nil
}
