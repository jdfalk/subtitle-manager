// file: pkg/tasks/tasks.go
package tasks

import (
	"context"
	"sync"
	"time"
)

// Task represents a background job run by the system.
type Task struct {
	ID          string    `json:"id"`
	Status      string    `json:"status"`
	Progress    int       `json:"progress"`
	StartedAt   time.Time `json:"started_at"`
	CompletedAt time.Time `json:"completed_at"`
	Error       string    `json:"error"`
}

var (
	mu    sync.Mutex
	tasks = map[string]*Task{}
)

// Start launches fn as a goroutine and tracks its progress in the global task map.
// The returned Task pointer can be polled for updates.
func Start(ctx context.Context, id string, fn func(context.Context) error) *Task {
	mu.Lock()
	t := &Task{ID: id, Status: "running", StartedAt: time.Now()}
	tasks[id] = t
	mu.Unlock()

	go func() {
		err := fn(ctx)
		mu.Lock()
		defer mu.Unlock()
		if err != nil {
			t.Status = "failed"
			t.Error = err.Error()
		} else {
			t.Status = "completed"
		}
		t.Progress = 100
		t.CompletedAt = time.Now()
	}()
	return t
}

// List returns a copy of all known tasks keyed by ID.
func List() map[string]*Task {
	mu.Lock()
	defer mu.Unlock()
	out := make(map[string]*Task, len(tasks))
	for k, v := range tasks {
		cp := *v
		out[k] = &cp
	}
	return out
}

// Update sets the progress percentage for the task id.
func Update(id string, progress int) {
	mu.Lock()
	defer mu.Unlock()
	if t, ok := tasks[id]; ok {
		t.Progress = progress
	}
}
