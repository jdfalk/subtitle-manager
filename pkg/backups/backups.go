// file: pkg/backups/backups.go
package backups

import (
	"sync"
	"time"
)

// Backup represents a database backup on disk.
type Backup struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

var (
	mu      sync.Mutex
	history []Backup
)

// Create simulates creating a new backup file and stores metadata in memory.
func Create() Backup {
	mu.Lock()
	defer mu.Unlock()
	b := Backup{Name: time.Now().Format("20060102-150405") + ".bak", CreatedAt: time.Now()}
	history = append(history, b)
	return b
}

// List returns all known backups.
func List() []Backup {
	mu.Lock()
	defer mu.Unlock()
	out := make([]Backup, len(history))
	copy(out, history)
	return out
}

// Restore pretends to restore the latest backup and returns its name.
func Restore() string {
	mu.Lock()
	defer mu.Unlock()
	if len(history) == 0 {
		return ""
	}
	return history[len(history)-1].Name
}
