package logging

import (
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
)

// TestMemoryHookFire verifies that old entries are dropped when limit is reached.
func TestMemoryHookFire(t *testing.T) {
	hook := NewMemoryHook(2)
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{DisableTimestamp: true})
	logger.AddHook(hook)

	logger.Info("first")
	logger.Info("second")
	logger.Info("third")

	logs := hook.Logs()
	if len(logs) != 2 {
		t.Fatalf("expected 2 logs, got %d", len(logs))
	}
	if !strings.Contains(logs[0], "second") || !strings.Contains(logs[1], "third") {
		t.Fatalf("unexpected log contents: %v", logs)
	}
}

// TestMemoryHookLogsCopy ensures that Logs returns a copy and not internal slice.
func TestMemoryHookLogsCopy(t *testing.T) {
	hook := NewMemoryHook(1)
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{DisableTimestamp: true})
	logger.AddHook(hook)

	logger.Info("entry")

	logs := hook.Logs()
	if len(logs) != 1 {
		t.Fatalf("expected 1 log, got %d", len(logs))
	}
	logs[0] = "changed"
	if hook.Logs()[0] == "changed" {
		t.Fatalf("Logs returned slice is not a copy")
	}
}
