// file: pkg/providers/instance.go
package providers

import (
	"sort"
	"sync"
	"time"
)

// Instance represents a configured provider with optional priority and tags.
// Implements TaggedEntity interface for universal tagging support.
type Instance struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Priority int      `json:"priority"`
	Tags     []string `json:"tags"`
	Enabled  bool     `json:"enabled"`
}

// GetEntityType returns the entity type for universal tagging.
func (i *Instance) GetEntityType() string {
	return "provider"
}

// GetEntityID returns the entity ID for universal tagging.
func (i *Instance) GetEntityID() string {
	return i.ID
}

// GetTags returns the current tags assigned to this provider instance.
func (i *Instance) GetTags() []string {
	return i.Tags
}

// SetTags updates the tags assigned to this provider instance.
func (i *Instance) SetTags(tags []string) {
	i.Tags = tags
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

// ListInstances returns all registered provider instances ordered by priority (enabled and disabled).
func ListInstances() []Instance {
	instancesMu.RLock()
	defer instancesMu.RUnlock()
	out := make([]Instance, 0, len(instances))
	for _, inst := range instances {
		out = append(out, inst)
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

// IsInBackoff reports whether the given provider instance currently has a backoff delay.
func IsInBackoff(id string) bool {
	backoffMu.RLock()
	defer backoffMu.RUnlock()
	until, ok := backoffMap[id]
	return ok && time.Now().Before(until)
}

// ClearBackoff removes the backoff restriction for a provider instance.
func ClearBackoff(instanceID string) {
	backoffMu.Lock()
	defer backoffMu.Unlock()
	delete(backoffMap, instanceID)
}
