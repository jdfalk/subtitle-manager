// file: pkg/tasks/notify_test.go
package tasks

import (
	"context"
	"testing"
	"time"
)

// wait receives from ch with timeout.
func wait(ch chan TaskSnapshot) (TaskSnapshot, bool) {
	select {
	case t, ok := <-ch:
		return t, ok
	case <-time.After(100 * time.Millisecond):
		return TaskSnapshot{}, false
	}
}

func TestSubscribeReceivesStart(t *testing.T) {
	reset()
	ch := Subscribe()
	defer Unsubscribe(ch)
	Start(context.Background(), "a", func(context.Context) error { return nil })
	if _, ok := wait(ch); !ok {
		t.Fatalf("no update received")
	}
}

func TestUnsubscribe(t *testing.T) {
	reset()
	ch := Subscribe()
	Unsubscribe(ch)
	Start(context.Background(), "b", func(context.Context) error { return nil })
	if _, ok := wait(ch); ok {
		t.Fatalf("expected channel closed")
	}
}
