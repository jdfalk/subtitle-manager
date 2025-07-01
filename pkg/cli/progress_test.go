// file: pkg/cli/progress_test.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174003

package cli

import (
	"strings"
	"testing"
	"time"
)

func TestNewProgressBar(t *testing.T) {
	pb := NewProgressBar(100, "Test")
	if pb.total != 100 {
		t.Errorf("Expected total 100, got %d", pb.total)
	}
	if pb.prefix != "Test" {
		t.Errorf("Expected prefix 'Test', got %s", pb.prefix)
	}
	if pb.current != 0 {
		t.Errorf("Expected current 0, got %d", pb.current)
	}
}

func TestProgressBarUpdate(t *testing.T) {
	pb := NewProgressBar(3, "Test")

	// First update
	pb.Update("file1.txt")
	if pb.current != 1 {
		t.Errorf("Expected current 1, got %d", pb.current)
	}

	// Second update
	pb.Update("file2.txt")
	if pb.current != 2 {
		t.Errorf("Expected current 2, got %d", pb.current)
	}

	// Third update
	pb.Update("file3.txt")
	if pb.current != 3 {
		t.Errorf("Expected current 3, got %d", pb.current)
	}
}

func TestProgressBarFinish(t *testing.T) {
	pb := NewProgressBar(10, "Test")
	pb.current = 5

	pb.Finish()
	if pb.current != pb.total {
		t.Errorf("Expected current to equal total after Finish, got current=%d, total=%d", pb.current, pb.total)
	}
}

func TestFileCounter(t *testing.T) {
	fc := NewFileCounter()
	if fc.Total() != 0 {
		t.Errorf("Expected total 0 for new FileCounter, got %d", fc.Total())
	}

	fc.Add("file1.txt")
	fc.Add("file2.txt")
	if fc.Total() != 2 {
		t.Errorf("Expected total 2 after adding 2 files, got %d", fc.Total())
	}

	fc.SetTotal(5)
	if fc.Total() != 5 {
		t.Errorf("Expected total 5 after SetTotal(5), got %d", fc.Total())
	}
}

func TestProgressBarDraw(t *testing.T) {
	pb := NewProgressBar(4, "Test")
	pb.current = 2

	// Test that draw doesn't panic with various inputs
	pb.draw("test-file.txt")
	pb.draw("")
	pb.draw(strings.Repeat("a", 100)) // Long filename

	// Test zero total
	pb.total = 0
	pb.draw("test")
}

func TestProgressBarThrottling(t *testing.T) {
	pb := NewProgressBar(1000, "Test")

	start := time.Now()
	// Update many times quickly
	for i := 0; i < 100; i++ {
		pb.Update("file.txt")
	}
	elapsed := time.Since(start)

	// Should complete quickly due to throttling
	if elapsed > time.Second {
		t.Errorf("Updates took too long: %v", elapsed)
	}

	// Current should still be accurate
	if pb.current != 100 {
		t.Errorf("Expected current 100, got %d", pb.current)
	}
}
