// file: pkg/webhooks/webhooks.go
package webhooks

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

// Dispatcher sends webhook events to a list of URLs.
type Dispatcher struct {
	URLs   []string
	client *http.Client
}

// New returns a new Dispatcher with the given URLs.
func New(urls []string) *Dispatcher {
	return &Dispatcher{URLs: urls, client: &http.Client{Timeout: 10 * time.Second}}
}

// Send delivers an event with optional payload to all configured URLs.
func (d *Dispatcher) Send(ctx context.Context, event string, payload any) error {
	body, err := json.Marshal(map[string]any{"event": event, "payload": payload})
	if err != nil {
		return err
	}
	for _, u := range d.URLs {
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, bytes.NewReader(body))
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/json")
		if _, err := d.client.Do(req); err != nil {
			return err
		}
	}
	return nil
}
