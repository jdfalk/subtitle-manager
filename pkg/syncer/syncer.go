package syncer

import (
	"time"

	"github.com/asticode/go-astisub"
	"github.com/jdfalk/subtitle-manager/pkg/subtitles"
)

// Options controls how the synchronization process behaves.
type Options struct {
	// UseAudio indicates whether the audio track should be analyzed.
	UseAudio bool
	// UseEmbedded indicates whether embedded subtitles should be used.
	UseEmbedded bool
	// AudioTrack selects the audio track index if UseAudio is true.
	AudioTrack int
	// SubtitleTracks selects embedded subtitle track indices when UseEmbedded is true.
	SubtitleTracks []int
}

// Sync attempts to synchronize the subtitle at subPath with the media file at
// mediaPath according to opts. The resulting subtitle items are returned.
//
// When UseEmbedded is true, the specified subtitle tracks are extracted from the
// media container and compared with the external subtitle. A simple offset and
// scale are calculated using the first and last cues of each track. The average
// of all tracks is applied to the external subtitle items.
//
// Audio-based synchronization is not yet implemented.
func Sync(mediaPath, subPath string, opts Options) ([]*astisub.Item, error) {
	sub, err := astisub.OpenFile(subPath)
	if err != nil {
		return nil, err
	}
	items := make([]*astisub.Item, len(sub.Items))
	copy(items, sub.Items)

	if opts.UseEmbedded {
		tracks := opts.SubtitleTracks
		if len(tracks) == 0 {
			tracks = []int{0}
		}
		var offsets []time.Duration
		var scales []float64
		for _, tr := range tracks {
			emb, err := subtitles.ExtractFromMediaTrack(mediaPath, tr)
			if err != nil || len(emb) < 2 || len(items) < 2 {
				continue
			}
			off := emb[0].StartAt - items[0].StartAt
			offsets = append(offsets, off)
			embDur := emb[len(emb)-1].StartAt - emb[0].StartAt
			extDur := items[len(items)-1].StartAt - items[0].StartAt
			if extDur > 0 {
				scales = append(scales, float64(embDur)/float64(extDur))
			}
		}
		var offset time.Duration
		if len(offsets) > 0 {
			var sum time.Duration
			for _, o := range offsets {
				sum += o
			}
			offset = sum / time.Duration(len(offsets))
		}
		scale := 1.0
		if len(scales) > 0 {
			var sum float64
			for _, s := range scales {
				sum += s
			}
			scale = sum / float64(len(scales))
		}
		for _, it := range items {
			it.StartAt = time.Duration(float64(it.StartAt)*scale) + offset
			it.EndAt = time.Duration(float64(it.EndAt)*scale) + offset
		}
	}

	return items, nil
}

// Shift adjusts each subtitle item by offset and returns the updated slice.
func Shift(items []*astisub.Item, offset time.Duration) []*astisub.Item {
	out := make([]*astisub.Item, len(items))
	for i, it := range items {
		c := *it
		c.StartAt += offset
		c.EndAt += offset
		out[i] = &c
	}
	return out
}
