package subtitles

import (
	"testing"
	"time"

	"github.com/asticode/go-astisub"
)

func BenchmarkMergeTracks(b *testing.B) {
	a := make([]*astisub.Item, 1000)
	c := make([]*astisub.Item, 1000)
	for i := 0; i < 1000; i++ {
		a[i] = &astisub.Item{StartAt: time.Duration(i*2) * time.Second}
		c[i] = &astisub.Item{StartAt: time.Duration(i*2+1) * time.Second}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MergeTracks(a, c)
	}
}

func BenchmarkMergeTracksSorted(b *testing.B) {
	// Test with already-sorted inputs to measure optimization
	a := make([]*astisub.Item, 1000)
	c := make([]*astisub.Item, 1000)
	for i := 0; i < 1000; i++ {
		a[i] = &astisub.Item{StartAt: time.Duration(i*2) * time.Second}
		c[i] = &astisub.Item{StartAt: time.Duration(i*2+1) * time.Second}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MergeTracks(a, c)
	}
}

func BenchmarkMergeTracksUnsorted(b *testing.B) {
	// Test with randomly ordered inputs
	a := make([]*astisub.Item, 1000)
	c := make([]*astisub.Item, 1000)
	for i := 0; i < 1000; i++ {
		// Mix up the order to force full sort
		a[i] = &astisub.Item{StartAt: time.Duration((1000-i)*2) * time.Second}
		c[i] = &astisub.Item{StartAt: time.Duration((1000-i)*2+1) * time.Second}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MergeTracks(a, c)
	}
}

func BenchmarkMergeTracksLarge(b *testing.B) {
	// Test with larger datasets
	a := make([]*astisub.Item, 10000)
	c := make([]*astisub.Item, 10000)
	for i := 0; i < 10000; i++ {
		a[i] = &astisub.Item{StartAt: time.Duration(i*2) * time.Second}
		c[i] = &astisub.Item{StartAt: time.Duration(i*2+1) * time.Second}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MergeTracks(a, c)
	}
}
