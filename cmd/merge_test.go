// file: cmd/merge_test.go
// version: 1.0.0
// guid: 6fd62945-5bb6-4b7d-81c5-6472913a6561
package cmd

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/asticode/go-astisub"
)

func TestMergeCommand_MergeSubtitles_WritesMergedFile(t *testing.T) {
	// Arrange
	tempDir := t.TempDir()
	firstPath := filepath.Join(tempDir, "first.srt")
	secondPath := filepath.Join(tempDir, "second.srt")
	outPath := filepath.Join(tempDir, "merged.srt")

	firstItems := []*astisub.Item{
		newSubtitleItem(1*time.Second, "First"),
		newSubtitleItem(5*time.Second, "Third"),
	}
	secondItems := []*astisub.Item{
		newSubtitleItem(3*time.Second, "Second"),
	}

	writeSubtitleFile(t, firstPath, firstItems)
	writeSubtitleFile(t, secondPath, secondItems)

	// Act
	err := mergeCmd.RunE(mergeCmd, []string{firstPath, secondPath, outPath})

	// Assert
	if err != nil {
		t.Fatalf("expected merge to succeed, got error: %v", err)
	}

	merged, err := astisub.OpenFile(outPath)
	if err != nil {
		t.Fatalf("expected to read merged subtitle file, got error: %v", err)
	}

	if len(merged.Items) != 3 {
		t.Fatalf("expected 3 merged items, got %d", len(merged.Items))
	}

	assertSubtitleItem(t, merged.Items[0], 1*time.Second, "First")
	assertSubtitleItem(t, merged.Items[1], 3*time.Second, "Second")
	assertSubtitleItem(t, merged.Items[2], 5*time.Second, "Third")
}

func TestMergeCommand_InvalidPath_ReturnsError(t *testing.T) {
	// Arrange
	args := []string{"relative1.srt", "relative2.srt", "relative3.srt"}

	// Act
	err := mergeCmd.RunE(mergeCmd, args)

	// Assert
	if err == nil {
		t.Fatal("expected error for invalid relative paths, got nil")
	}
}

func writeSubtitleFile(t *testing.T, path string, items []*astisub.Item) {
	t.Helper()

	file, err := os.Create(path)
	if err != nil {
		t.Fatalf("failed to create subtitle file: %v", err)
	}
	defer file.Close()

	subtitles := astisub.NewSubtitles()
	subtitles.Items = items

	if err := subtitles.WriteToSRT(file); err != nil {
		t.Fatalf("failed to write subtitle file: %v", err)
	}
}

func newSubtitleItem(start time.Duration, text string) *astisub.Item {
	return &astisub.Item{
		StartAt: start,
		EndAt:   start + time.Second,
		Lines: []astisub.Line{
			{
				Items: []astisub.LineItem{
					{Text: text},
				},
			},
		},
	}
}

func assertSubtitleItem(t *testing.T, item *astisub.Item, start time.Duration, text string) {
	t.Helper()

	if item.StartAt != start {
		t.Fatalf("expected start time %v, got %v", start, item.StartAt)
	}
	if len(item.Lines) != 1 || len(item.Lines[0].Items) != 1 {
		t.Fatalf("expected a single line item, got %+v", item.Lines)
	}
	if item.Lines[0].Items[0].Text != text {
		t.Fatalf("expected text %q, got %q", text, item.Lines[0].Items[0].Text)
	}
}
