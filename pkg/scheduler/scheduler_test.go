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

// TestRunWithOptions verifies advanced scheduling options.
func TestRunWithOptions(t *testing.T) {
	var n int32
	ctx := context.Background()
	err := RunWithOptions(ctx, Options{Interval: 10 * time.Millisecond, MaxRuns: 2}, func(context.Context) error {
		atomic.AddInt32(&n, 1)
		return nil
	})
	if err != nil {
		t.Fatalf("run: %v", err)
	}
	if n != 2 {
		t.Fatalf("expected 2 runs, got %d", n)
	}
}

// TestRunWithSkipFirst verifies SkipFirst prevents immediate execution.
func TestRunWithSkipFirst(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	var n int32
	go func() {
		time.Sleep(80 * time.Millisecond)
		cancel()
	}()
	err := RunWithOptions(ctx, Options{Interval: 50 * time.Millisecond, SkipFirst: true}, func(context.Context) error {
		atomic.AddInt32(&n, 1)
		return nil
	})
	if err != context.Canceled {
		t.Fatalf("expected context canceled, got %v", err)
	}
	if n != 1 {
		t.Fatalf("expected 1 run, got %d", n)
	}
}

// TestRunCron verifies that RunCron executes according to the provided schedule.
func TestRunCron(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	var n int32
	go func() {
		time.Sleep(1 * time.Second) // Give cron plenty of time to start and execute
		cancel()
	}()
	err := RunCron(ctx, "@every 250ms", func(context.Context) error {
		atomic.AddInt32(&n, 1)
		return nil
	})
	if err != context.Canceled {
		t.Fatalf("expected context canceled, got %v", err)
	}
	// Use atomic load to avoid race condition
	count := atomic.LoadInt32(&n)
	if count < 2 {
		t.Fatalf("expected at least 2 executions, got %d", count)
	}
}
