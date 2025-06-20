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
		if t.GetStatus() != "running" {
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
	if task.GetStatus() != "completed" {
		t.Fatalf("expected status completed, got %s", task.GetStatus())
	}
	if task.GetProgress() != 100 {
		t.Fatalf("expected progress 100, got %d", task.GetProgress())
	}
	if task.GetCompletedAt().IsZero() {
		t.Fatalf("completed time not set")
	}
}

// TestStartFailure verifies that errors returned from the function mark the task as failed.
func TestStartFailure(t *testing.T) {
	reset()
	task := Start(context.Background(), "id", func(context.Context) error { return errors.New("boom") })
	waitUntilFinished(task)
	if task.GetStatus() != "failed" {
		t.Fatalf("expected status failed, got %s", task.GetStatus())
	}
	if task.GetError() != "boom" {
		t.Fatalf("unexpected error %s", task.GetError())
	}
}

// TestUpdateProgress verifies that Update sets the progress for an existing task.
func TestUpdateProgress(t *testing.T) {
	reset()
	mu.Lock()
	task := &Task{ID: "a"}
	tasks["a"] = task
	mu.Unlock()
	Update("a", 42)
	if task.GetProgress() != 42 {
		t.Fatalf("expected progress 42, got %d", task.GetProgress())
	}
}

// TestListReturnsCopy verifies that List returns a copy of the task map.
func TestListReturnsCopy(t *testing.T) {
	reset()
	mu.Lock()
	originalTask := &Task{ID: "x", Status: "running"}
	tasks["x"] = originalTask
	mu.Unlock()
	copyMap := List()
	copyMap["x"].Status = "changed"
	if originalTask.GetStatus() == "changed" {
		t.Fatalf("modifying list result affected original map")
	}
}
