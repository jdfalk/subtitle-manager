package subtitles

import (
	"sort"

	"github.com/asticode/go-astisub"
)

// MergeTracks merges two subtitle item slices sorted by start time.
// The resulting slice is sorted and returned.
// This function is optimized for the case where inputs are already sorted.
func MergeTracks(a, b []*astisub.Item) []*astisub.Item {
	// Fast path: if both inputs are already sorted, use merge algorithm
	if isSorted(a) && isSorted(b) {
		return mergeSorted(a, b)
	}

	// Fallback: general case with full sort
	items := append(a, b...)
	sort.Slice(items, func(i, j int) bool {
		return items[i].StartAt < items[j].StartAt
	})
	return items
}

// isSorted checks if subtitle items are sorted by start time
func isSorted(items []*astisub.Item) bool {
	for i := 1; i < len(items); i++ {
		if items[i-1].StartAt > items[i].StartAt {
			return false
		}
	}
	return true
}

// mergeSorted merges two already-sorted subtitle slices in O(n+m) time
func mergeSorted(a, b []*astisub.Item) []*astisub.Item {
	result := make([]*astisub.Item, 0, len(a)+len(b))
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i].StartAt <= b[j].StartAt {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}

	// Append remaining items
	result = append(result, a[i:]...)
	result = append(result, b[j:]...)

	return result
}
