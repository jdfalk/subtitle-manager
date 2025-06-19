// file: pkg/providers/instance.go
package providers

import (
	"sort"
	"sync"
	"time"
)

// Instance represents a configured provider with optional priority and tags.
type Instance struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Priority int      `json:"priority"`
	Tags     []string `json:"tags"`
	Enabled  bool     `json:"enabled"`
}

var (
	instancesMu sync.RWMutex
	instances   = map[string]Instance{}

	backoffMu  sync.RWMutex
	backoffMap = map[string]time.Time{}
)

// GetInstance retrieves a provider instance by ID.
func GetInstance(id string) (Instance, bool) {
	instancesMu.RLock()
	defer instancesMu.RUnlock()
	inst, ok := instances[id]
	return inst, ok
}

// RegisterInstance adds a provider instance to the registry or updates it if it already exists.
func RegisterInstance(inst Instance) {
	instancesMu.Lock()
	defer instancesMu.Unlock()
	instances[inst.ID] = inst
}

// Instances returns all registered provider instances ordered by priority.
func Instances() []Instance {
	instancesMu.RLock()
	defer instancesMu.RUnlock()
	out := make([]Instance, 0, len(instances))
	for _, inst := range instances {
		if inst.Enabled {
			out = append(out, inst)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Priority > out[j].Priority })
	return out
}

// SetBackoff records a backoff duration for a provider instance. A zero or negative duration clears the backoff.
func SetBackoff(id string, d time.Duration) {
	backoffMu.Lock()
	defer backoffMu.Unlock()
	if d <= 0 {
		delete(backoffMap, id)
		return
	}
	backoffMap[id] = time.Now().Add(d)
}

// inBackoff reports whether the given provider instance currently has a backoff delay.
func inBackoff(id string) bool {
	backoffMu.RLock()
	defer backoffMu.RUnlock()
	until, ok := backoffMap[id]
	return ok && time.Now().Before(until)
}
