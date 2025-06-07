package subtitles

import (
	"github.com/asticode/go-astisub"
	"testing"
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
