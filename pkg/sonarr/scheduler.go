package sonarr

import (
	"context"
	"time"

	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/scheduler"
)

// StartSync periodically syncs the Sonarr library using interval.
// The sync runs immediately then at each interval until ctx is cancelled.
func StartSync(ctx context.Context, interval time.Duration, c *Client, store database.SubtitleStore) {
	go scheduler.Run(ctx, interval, func(ctx context.Context) error {
		return Sync(ctx, c, store)
	})
}
