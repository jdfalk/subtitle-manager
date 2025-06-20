// file: pkg/backups/backups_test.go
package backups

import (
	"testing"
	"time"

	"github.com/spf13/afero"
)

// resetHistory clears the package state between tests.
func resetHistory() {
	mu.Lock()
	history = nil
	mu.Unlock()
}

// TestCreate verifies that Create records a backup and the returned name can be used on a fake file system.
func TestCreate(t *testing.T) {
	resetHistory()
	fs := afero.NewMemMapFs()

	b := Create()
	if b.Name == "" {
		t.Fatalf("expected backup name")
	}
	// allow up to 3 seconds to avoid flakes on slow systems
	if time.Since(b.CreatedAt) > time.Second*3 {
		t.Fatalf("creation time too old: %v", b.CreatedAt)
	}
	list := List()
	if len(list) != 1 {
		t.Fatalf("expected 1 backup, got %d", len(list))
	}
	// attempt to create file using fake fs to ensure no real disk writes are needed
	f, err := fs.Create(b.Name)
	if err != nil {
		t.Fatalf("create backup file: %v", err)
	}
	f.Close()
	ok, err := afero.Exists(fs, b.Name)
	if err != nil || !ok {
		t.Fatalf("backup file not found in fake fs")
	}
}

// TestListReturnsCopy ensures List returns a copy of history so callers cannot modify internal state.
func TestListReturnsCopy(t *testing.T) {
	resetHistory()
	b1 := Create()
	_ = b1
	b2 := Create()
	_ = b2

	l1 := List()
	if len(l1) != 2 {
		t.Fatalf("expected 2 backups, got %d", len(l1))
	}
	l1[0].Name = "changed"
	l2 := List()
	if l2[0].Name == "changed" {
		t.Fatalf("List returned internal slice")
	}
}

// TestRestore verifies Restore returns the latest backup name or empty when none exist.
func TestRestore(t *testing.T) {
	resetHistory()
	if name := Restore(); name != "" {
		t.Fatalf("expected empty restore name, got %q", name)
	}
	b1 := Create()
	if name := Restore(); name != b1.Name {
		t.Fatalf("expected %q, got %q", b1.Name, name)
	}
	b2 := Create()
	if name := Restore(); name != b2.Name {
		t.Fatalf("expected %q, got %q", b2.Name, name)
	}
}
