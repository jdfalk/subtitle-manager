// file: pkg/tasks/tasks_test.go
package tasks

import (
	"context"
	"errors"
	"testing"
	"time"
)

// reset clears the global task map for isolated tests.
func reset() {
	mu.Lock()
	tasks = map[string]*Task{}
	mu.Unlock()
}

// waitUntilFinished polls task status until it is not running or timeout occurs.
func waitUntilFinished(t *Task) {
	for i := 0; i < 10; i++ {
		if t.Status != "running" {
			return
		}
		time.Sleep(10 * time.Millisecond)
	}
}

// TestStartSuccess verifies that a task records completion when the function succeeds.
func TestStartSuccess(t *testing.T) {
	reset()
	task := Start(context.Background(), "id", func(context.Context) error { return nil })
	waitUntilFinished(task)
	if task.Status != "completed" {
		t.Fatalf("expected status completed, got %s", task.Status)
	}
	if task.Progress != 100 {
		t.Fatalf("expected progress 100, got %d", task.Progress)
	}
	if task.CompletedAt.IsZero() {
		t.Fatalf("completed time not set")
	}
}

// TestStartFailure verifies that errors returned from the function mark the task as failed.
func TestStartFailure(t *testing.T) {
	reset()
	task := Start(context.Background(), "id", func(context.Context) error { return errors.New("boom") })
	waitUntilFinished(task)
	if task.Status != "failed" {
		t.Fatalf("expected status failed, got %s", task.Status)
	}
	if task.Error != "boom" {
		t.Fatalf("unexpected error %s", task.Error)
	}
}

// TestUpdateProgress verifies that Update sets the progress for an existing task.
func TestUpdateProgress(t *testing.T) {
	reset()
	mu.Lock()
	tasks["a"] = &Task{ID: "a"}
	mu.Unlock()
	Update("a", 42)
	mu.Lock()
	p := tasks["a"].Progress
	mu.Unlock()
	if p != 42 {
		t.Fatalf("expected progress 42, got %d", p)
	}
}

// TestListReturnsCopy verifies that List returns a copy of the task map.
func TestListReturnsCopy(t *testing.T) {
	reset()
	mu.Lock()
	tasks["x"] = &Task{ID: "x", Status: "running"}
	mu.Unlock()
	copyMap := List()
	copyMap["x"].Status = "changed"
	mu.Lock()
	status := tasks["x"].Status
	mu.Unlock()
	if status == "changed" {
		t.Fatalf("modifying list result affected original map")
	}
}
