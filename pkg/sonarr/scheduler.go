// Package sonarr provides a client and utilities for interacting with the Sonarr API.
// It enables subtitle synchronization, library management, and automation for Sonarr servers.
//
// This package is used by subtitle-manager to keep Sonarr libraries in sync and manage subtitles.
//
// See: https://sonarr.tv/

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
