package subtitles

import (
	"testing"
	"time"

	"github.com/asticode/go-astisub"
)

func TestMergeTracks(t *testing.T) {
	a := []*astisub.Item{
		{StartAt: 1},
		{StartAt: 3},
	}
	b := []*astisub.Item{
		{StartAt: 2},
	}
	out := MergeTracks(a, b)
	if len(out) != 3 {
		t.Fatalf("expected 3 items, got %d", len(out))
	}
	if out[0].StartAt != 1 || out[1].StartAt != 2 || out[2].StartAt != 3 {
		t.Fatalf("unexpected order: %+v", out)
	}
}

func TestMergeTracksSorted(t *testing.T) {
	// Test that sorted inputs use the optimized merge path
	a := []*astisub.Item{
		{StartAt: 1 * time.Second},
		{StartAt: 3 * time.Second},
		{StartAt: 5 * time.Second},
	}
	b := []*astisub.Item{
		{StartAt: 2 * time.Second},
		{StartAt: 4 * time.Second},
		{StartAt: 6 * time.Second},
	}

	result := MergeTracks(a, b)

	// Verify correct merge order
	expected := []time.Duration{1, 2, 3, 4, 5, 6}
	if len(result) != len(expected) {
		t.Fatalf("expected %d items, got %d", len(expected), len(result))
	}

	for i, item := range result {
		if item.StartAt != expected[i]*time.Second {
			t.Errorf("item %d: expected %v, got %v", i, expected[i]*time.Second, item.StartAt)
		}
	}
}

func TestMergeTracksUnsorted(t *testing.T) {
	// Test that unsorted inputs still work correctly
	a := []*astisub.Item{
		{StartAt: 5 * time.Second},
		{StartAt: 1 * time.Second},
		{StartAt: 3 * time.Second},
	}
	b := []*astisub.Item{
		{StartAt: 6 * time.Second},
		{StartAt: 2 * time.Second},
		{StartAt: 4 * time.Second},
	}

	result := MergeTracks(a, b)

	// Verify correct sort order
	expected := []time.Duration{1, 2, 3, 4, 5, 6}
	if len(result) != len(expected) {
		t.Fatalf("expected %d items, got %d", len(expected), len(result))
	}

	for i, item := range result {
		if item.StartAt != expected[i]*time.Second {
			t.Errorf("item %d: expected %v, got %v", i, expected[i]*time.Second, item.StartAt)
		}
	}
}

func TestIsSorted(t *testing.T) {
	tests := []struct {
		name   string
		items  []*astisub.Item
		sorted bool
	}{
		{
			name:   "empty",
			items:  []*astisub.Item{},
			sorted: true,
		},
		{
			name:   "single item",
			items:  []*astisub.Item{{StartAt: 1 * time.Second}},
			sorted: true,
		},
		{
			name: "sorted",
			items: []*astisub.Item{
				{StartAt: 1 * time.Second},
				{StartAt: 2 * time.Second},
				{StartAt: 3 * time.Second},
			},
			sorted: true,
		},
		{
			name: "equal timestamps",
			items: []*astisub.Item{
				{StartAt: 1 * time.Second},
				{StartAt: 1 * time.Second},
				{StartAt: 2 * time.Second},
			},
			sorted: true,
		},
		{
			name: "unsorted",
			items: []*astisub.Item{
				{StartAt: 2 * time.Second},
				{StartAt: 1 * time.Second},
				{StartAt: 3 * time.Second},
			},
			sorted: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isSorted(tt.items)
			if result != tt.sorted {
				t.Errorf("isSorted() = %v, want %v", result, tt.sorted)
			}
		})
	}
}

func TestMergeSorted(t *testing.T) {
	a := []*astisub.Item{
		{StartAt: 1 * time.Second},
		{StartAt: 3 * time.Second},
		{StartAt: 5 * time.Second},
	}
	b := []*astisub.Item{
		{StartAt: 2 * time.Second},
		{StartAt: 4 * time.Second},
		{StartAt: 6 * time.Second},
	}

	result := mergeSorted(a, b)

	// Verify correct merge
	expected := []time.Duration{1, 2, 3, 4, 5, 6}
	if len(result) != len(expected) {
		t.Fatalf("expected %d items, got %d", len(expected), len(result))
	}

	for i, item := range result {
		if item.StartAt != expected[i]*time.Second {
			t.Errorf("item %d: expected %v, got %v", i, expected[i]*time.Second, item.StartAt)
		}
	}
}

func TestMergeSortedEdgeCases(t *testing.T) {
	// Test empty slices
	result := mergeSorted([]*astisub.Item{}, []*astisub.Item{})
	if len(result) != 0 {
		t.Errorf("expected empty result, got %d items", len(result))
	}

	// Test one empty slice
	a := []*astisub.Item{{StartAt: 1 * time.Second}}
	result = mergeSorted(a, []*astisub.Item{})
	if len(result) != 1 || result[0].StartAt != 1*time.Second {
		t.Errorf("merge with empty slice failed")
	}

	result = mergeSorted([]*astisub.Item{}, a)
	if len(result) != 1 || result[0].StartAt != 1*time.Second {
		t.Errorf("merge with empty slice failed")
	}
}
