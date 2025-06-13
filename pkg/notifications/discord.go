// file: pkg/notifications/discord.go
package notifications

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// DiscordNotifier sends messages to a Discord webhook.
type DiscordNotifier struct {
	// WebhookURL is the Discord webhook endpoint.
	WebhookURL string
	// Client is used to make HTTP requests. Defaults to http.DefaultClient.
	Client *http.Client
}

// Notify posts msg to the Discord webhook.
func (d DiscordNotifier) Notify(ctx context.Context, msg string) error {
	if d.WebhookURL == "" {
		return fmt.Errorf("webhook URL required")
	}
	c := d.Client
	if c == nil {
		c = http.DefaultClient
	}
	body, _ := json.Marshal(map[string]string{"content": msg})
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, d.WebhookURL, bytes.NewReader(body))
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
