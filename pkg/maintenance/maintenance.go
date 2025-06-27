package maintenance

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"
	"time"

	"github.com/jdfalk/subtitle-manager/pkg/auth"
	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/metadata"
	"github.com/jdfalk/subtitle-manager/pkg/scheduler"
)

// frequencyToDuration converts a textual frequency like "daily" to a
// time.Duration. Unknown values default to 24 hours. It mirrors the
// updater package frequency parsing but is local to avoid export.
func frequencyToDuration(freq string) time.Duration {
	switch freq {
	case "hourly":
		return time.Hour
	case "daily":
		return 24 * time.Hour
	case "weekly":
		return 7 * 24 * time.Hour
	case "monthly":
		return 30 * 24 * time.Hour
	default:
		d, err := time.ParseDuration(freq)
		if err != nil {
			return 24 * time.Hour
		}
		return d
	}
}

// CleanupDatabase removes expired sessions and performs VACUUM for
// SQLite databases.
//
// Parameters:
//   - ctx: cancellation context controlling task lifetime
//   - db:  database handle
//
// Returns an error if cleanup fails.
func CleanupDatabase(ctx context.Context, db *sql.DB) error {
	if err := auth.CleanupExpiredSessions(db); err != nil {
		return err
	}
	if database.GetDatabaseBackend() == "sqlite" {
		_, err := db.ExecContext(ctx, "VACUUM")
		return err
	}
	return nil
}

// StartDatabaseCleanup schedules CleanupDatabase to run periodically.
// The frequency string may be "hourly", "daily", "weekly" or any
// valid time.ParseDuration value.
func StartDatabaseCleanup(ctx context.Context, db *sql.DB, frequency string) {
	interval := frequencyToDuration(frequency)
	go scheduler.Run(ctx, interval, func(c context.Context) error {
		return CleanupDatabase(c, db)
	})
}

// RefreshMetadata updates stored media items with data from TMDB.
// It fetches metadata for each item in the store using the provided
// TMDB API key. Errors during individual lookups are ignored.
//
// Parameters:
//   - ctx:   cancellation context
//   - store: media database store
//   - apiKey: TMDB API key
//
// RefreshMetadata updates stored media items with data from TMDB and OMDb.
// Parameters:
//   - ctx:     cancellation context
//   - store:   media database store
//   - tmdbKey: TMDB API key
//   - omdbKey: OMDb API key
func RefreshMetadata(ctx context.Context, store database.SubtitleStore, tmdbKey, omdbKey string) error {
	items, err := store.ListMediaItems()
	if err != nil {
		return err
	}
	for _, it := range items {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		if it.Season > 0 {
			_, _ = metadata.FetchEpisodeMetadata(ctx, it.Title, it.Season, it.Episode, tmdbKey, omdbKey)
		} else {
			_, _ = metadata.FetchMovieMetadata(ctx, it.Title, 0, tmdbKey, omdbKey)
		}
	}
	return nil
}

// StartMetadataRefresh schedules RefreshMetadata periodically using the
// provided frequency string.
// StartMetadataRefresh schedules RefreshMetadata periodically using the provided frequency.
func StartMetadataRefresh(ctx context.Context, store database.SubtitleStore, tmdbKey, omdbKey, frequency string) {
	interval := frequencyToDuration(frequency)
	go scheduler.Run(ctx, interval, func(c context.Context) error {
		return RefreshMetadata(c, store, tmdbKey, omdbKey)
	})
}

// DiskScan calculates the total size of files under root. It returns the
// aggregated size in bytes.
//
// Parameters:
//   - ctx:  context controlling cancellation
//   - root: directory path to scan
//
// Returns the total size in bytes and any error encountered during scanning.
func DiskScan(ctx context.Context, root string) (int64, error) {
	var size int64
	err := filepath.Walk(root, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if ctx.Err() != nil {
			return ctx.Err()
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})
	return size, err
}

// StartDiskScan schedules DiskScan to run periodically and discards the
// result. The directory size can be logged or processed by callers as
// needed.
func StartDiskScan(ctx context.Context, root, frequency string) {
	interval := frequencyToDuration(frequency)
	go scheduler.Run(ctx, interval, func(c context.Context) error {
		_, err := DiskScan(c, root)
		return err
	})
}
