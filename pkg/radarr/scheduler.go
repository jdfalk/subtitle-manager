package radarr

import (
	"context"
	"time"

	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/scheduler"
)

// StartSync periodically syncs the Radarr library using interval.
// The sync runs immediately then at each interval until ctx is cancelled.
func StartSync(ctx context.Context, interval time.Duration, c *Client, store database.SubtitleStore) {
	go scheduler.Run(ctx, interval, func(ctx context.Context) error {
		return Sync(ctx, c, store)
	})
}

// StartContinuousSync runs Sync on a fixed interval with optional jitter.
// It wraps scheduler.RunWithOptions to allow smoother scheduling when multiple
// services are enabled.
func StartContinuousSync(ctx context.Context, interval, jitter time.Duration, c *Client, store database.SubtitleStore) {
	opts := scheduler.Options{Interval: interval, Jitter: jitter}
	go scheduler.RunWithOptions(ctx, opts, func(ctx context.Context) error {
		return Sync(ctx, c, store)
	})
}
