package logging

import (
    "sync"

    "github.com/sirupsen/logrus"
)

// MemoryHook stores recent log entries for retrieval via the web UI.
type MemoryHook struct {
    mu   sync.Mutex
    logs []string
    max  int
}

// Hook is the shared instance used by all loggers.
var Hook = NewMemoryHook(100)

func init() {
    logrus.AddHook(Hook)
}

// NewMemoryHook returns a hook keeping up to size log entries in memory.
func NewMemoryHook(size int) *MemoryHook {
    return &MemoryHook{max: size}
}

// Levels implements logrus.Hook.
func (h *MemoryHook) Levels() []logrus.Level { return logrus.AllLevels }

// Fire appends the formatted entry to the slice, trimming old entries when the
// buffer is full.
func (h *MemoryHook) Fire(e *logrus.Entry) error {
    line, err := e.String()
    if err != nil {
        return err
    }
    h.mu.Lock()
    defer h.mu.Unlock()
    if len(h.logs) >= h.max {
        copy(h.logs, h.logs[1:])
        h.logs[len(h.logs)-1] = line
    } else {
        h.logs = append(h.logs, line)
    }
    return nil
}

// Logs returns a copy of the collected log lines.
func (h *MemoryHook) Logs() []string {
    h.mu.Lock()
    defer h.mu.Unlock()
    out := make([]string, len(h.logs))
    copy(out, h.logs)
    return out
}
