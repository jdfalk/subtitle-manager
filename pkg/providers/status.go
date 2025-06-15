// file: pkg/providers/status.go
package providers

import (
	"context"
	"sync"
	"time"
)

// Status represents the availability of a provider.
type Status struct {
	Name      string    `json:"name"`
	Available bool      `json:"available"`
	CheckedAt time.Time `json:"checked_at"`
}

var (
	statusMu  sync.Mutex
	statusMap = map[string]Status{}
)

// Refresh simulates checking each provider and marks them as available.
func Refresh(ctx context.Context, names []string) {
	statusMu.Lock()
	defer statusMu.Unlock()
	for _, n := range names {
		statusMap[n] = Status{Name: n, Available: true, CheckedAt: time.Now()}
	}
}

// Reset clears all stored provider status information.
func Reset() {
	statusMu.Lock()
	defer statusMu.Unlock()
	statusMap = map[string]Status{}
}

// List returns a copy of the provider status map.
func List() map[string]Status {
	statusMu.Lock()
	defer statusMu.Unlock()
	out := make(map[string]Status, len(statusMap))
	for k, v := range statusMap {
		out[k] = v
	}
	return out
}
