// file: pkg/tasks/tasks.go
package tasks

import (
	"context"
	"sync"
	"time"
)

// Task represents a background job run by the system.
type Task struct {
	mu          sync.RWMutex
	ID          string    `json:"id"`
	Status      string    `json:"status"`
	Progress    int       `json:"progress"`
	StartedAt   time.Time `json:"started_at"`
	CompletedAt time.Time `json:"completed_at"`
	Error       string    `json:"error"`
}

// GetStatus safely returns the current status of the task.
func (t *Task) GetStatus() string {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.Status
}

// GetProgress safely returns the current progress of the task.
func (t *Task) GetProgress() int {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.Progress
}

// GetError safely returns the error message of the task.
func (t *Task) GetError() string {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.Error
}

// GetCompletedAt safely returns the completion time of the task.
func (t *Task) GetCompletedAt() time.Time {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.CompletedAt
}

// GetSnapshot returns a copy of the task with all fields safely read.
func (t *Task) GetSnapshot() Task {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return Task{
		ID:          t.ID,
		Status:      t.Status,
		Progress:    t.Progress,
		StartedAt:   t.StartedAt,
		CompletedAt: t.CompletedAt,
		Error:       t.Error,
	}
}

// setStatus safely sets the status of the task.
func (t *Task) setStatus(status string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.Status = status
}

// setProgress safely sets the progress of the task.
func (t *Task) setProgress(progress int) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.Progress = progress
}

// setError safely sets the error message of the task.
func (t *Task) setError(err string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.Error = err
}

// setCompletedAt safely sets the completion time of the task.
func (t *Task) setCompletedAt(completedAt time.Time) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.CompletedAt = completedAt
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
	snapshot := t.GetSnapshot()
	broadcast(&snapshot)

	go func() {
		err := fn(ctx)
		if err != nil {
			t.setStatus("failed")
			t.setError(err.Error())
		} else {
			t.setStatus("completed")
		}
		t.setProgress(100)
		t.setCompletedAt(time.Now())
		snapshot := t.GetSnapshot()
		broadcast(&snapshot)
	}()
	return t
}

// List returns a copy of all known tasks keyed by ID.
func List() map[string]*Task {
	mu.Lock()
	defer mu.Unlock()
	out := make(map[string]*Task, len(tasks))
	for k, v := range tasks {
		snapshot := v.GetSnapshot()
		out[k] = &snapshot
	}
	return out
}

// Update sets the progress percentage for the task id.
func Update(id string, progress int) {
	mu.Lock()
	t, ok := tasks[id]
	mu.Unlock()
	if ok {
		t.setProgress(progress)
		snapshot := t.GetSnapshot()
		broadcast(&snapshot)
	}
}
