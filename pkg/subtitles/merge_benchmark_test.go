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
