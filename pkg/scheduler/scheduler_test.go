package scheduler

import (
	"context"
	"sync/atomic"
	"testing"
	"time"
)

// TestRun verifies that Run invokes the task repeatedly until the context is cancelled.
func TestRun(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	var n int32
	go func() {
		time.Sleep(120 * time.Millisecond)
		cancel()
	}()
	err := Run(ctx, 50*time.Millisecond, func(context.Context) error {
		atomic.AddInt32(&n, 1)
		return nil
	})
	if err != context.Canceled {
		t.Fatalf("expected context canceled, got %v", err)
	}
	if n < 2 {
		t.Fatalf("expected at least 2 executions, got %d", n)
	}
}
