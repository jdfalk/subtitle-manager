package subtitles

import (
	"sort"

	"github.com/asticode/go-astisub"
)

// MergeTracks merges two subtitle item slices sorted by start time.
// The resulting slice is sorted and returned.
func MergeTracks(a, b []*astisub.Item) []*astisub.Item {
	items := append(a, b...)
	sort.Slice(items, func(i, j int) bool {
		return items[i].StartAt < items[j].StartAt
	})
	return items
}
