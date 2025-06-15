// file: pkg/webhooks/webhooks.go
package webhooks

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/spf13/viper"

	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/jdfalk/subtitle-manager/pkg/providers"
	"github.com/jdfalk/subtitle-manager/pkg/scanner"
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

// event describes a webhook payload with the file path and subtitle parameters.
type event struct {
	Path     string `json:"path"`
	Lang     string `json:"lang"`
	Provider string `json:"provider"`
}

// handle processes a webhook event by fetching a subtitle for the provided file.
func handle(w http.ResponseWriter, r *http.Request, ev event) {
	if ev.Path == "" || ev.Lang == "" || ev.Provider == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	key := viper.GetString("opensubtitles.api_key")
	p, err := providers.Get(ev.Provider, key)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := scanner.ProcessFile(r.Context(), ev.Path, ev.Lang, ev.Provider, p, true, nil); err != nil {
		logging.GetLogger("webhook").Warnf("process %s: %v", ev.Path, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// SonarrHandler handles webhook events from Sonarr.
func SonarrHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		var ev event
		if err := json.NewDecoder(r.Body).Decode(&ev); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		handle(w, r, ev)
	})
}

// RadarrHandler handles webhook events from Radarr.
func RadarrHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		var ev event
		if err := json.NewDecoder(r.Body).Decode(&ev); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		handle(w, r, ev)
	})
}

// CustomHandler accepts generic webhook events with the same payload format.
func CustomHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		var ev event
		if err := json.NewDecoder(r.Body).Decode(&ev); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		handle(w, r, ev)
	})
}
