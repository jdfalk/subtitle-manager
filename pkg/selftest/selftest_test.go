// file: pkg/selftest/selftest_test.go
// version: 1.0.0
// guid: f1e2d3c4-b5a6-789b-c0d1-e2f3a4b5c6d7
package selftest

import (
	"context"
	"testing"
	"time"

	"github.com/jdfalk/subtitle-manager/pkg/testutil"
)

// TestStartPeriodicSuccess ensures the self test runs without exiting when the database is healthy.
func TestStartPeriodicSuccess(t *testing.T) {
	db := testutil.GetTestDB(t)
	defer db.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	exitCalled := false
	orig := ExitFunc
	ExitFunc = func(code int) { exitCalled = true }
	defer func() { ExitFunc = orig }()

	StartPeriodic(ctx, db, "10ms")
	time.Sleep(20 * time.Millisecond)
	if exitCalled {
		t.Fatalf("exit should not be called on healthy db")
	}
}

// TestStartPeriodicFailure ensures the process exits when the database ping fails.
func TestStartPeriodicFailure(t *testing.T) {
	db := testutil.GetTestDB(t)
	db.Close() // force ping failure

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	called := make(chan int, 1)
	orig := ExitFunc
	ExitFunc = func(code int) { called <- code }
	defer func() { ExitFunc = orig }()

	StartPeriodic(ctx, db, "10ms")
	select {
	case <-called:
	case <-time.After(50 * time.Millisecond):
		t.Fatalf("exit not called on failure")
	}
}
