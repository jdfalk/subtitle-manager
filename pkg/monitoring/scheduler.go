// file: pkg/monitoring/scheduler.go
// version: 1.0.0
// guid: 12345678-1234-1234-1234-123456789016

package monitoring

import (
	"context"
	"time"

	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/radarr"
	"github.com/jdfalk/subtitle-manager/pkg/scheduler"
	"github.com/jdfalk/subtitle-manager/pkg/sonarr"
)

// ScheduledMonitor manages periodic subtitle monitoring tasks.
type ScheduledMonitor struct {
	monitor *EpisodeMonitor
}

// NewScheduledMonitor creates a new scheduled monitoring service.
func NewScheduledMonitor(
	sonarrClient *sonarr.Client,
	radarrClient *radarr.Client,
	store database.SubtitleStore,
	maxRetries int,
	qualityCheck bool,
) *ScheduledMonitor {
	monitor := NewEpisodeMonitor(
		time.Hour, // Default interval, will be overridden by scheduler
		sonarrClient,
		radarrClient,
		store,
		maxRetries,
		qualityCheck,
	)

	return &ScheduledMonitor{
		monitor: monitor,
	}
}

// StartPeriodicMonitoring starts the monitoring daemon with a specified interval.
func (s *ScheduledMonitor) StartPeriodicMonitoring(ctx context.Context, interval time.Duration) error {
	s.monitor.interval = interval
	
	opts := scheduler.Options{
		Interval: interval,
		SkipFirst: false, // Run immediately on start
		Jitter:   time.Minute, // Add some jitter to avoid spikes
	}

	return scheduler.RunWithOptions(ctx, opts, s.monitor.checkForSubtitles)
}

// RunSyncTask performs a one-time sync from Sonarr/Radarr.
func (s *ScheduledMonitor) RunSyncTask(ctx context.Context, opts SyncOptions) error {
	// Sync from Sonarr if configured
	if err := s.monitor.SyncFromSonarr(ctx, opts); err != nil {
		s.monitor.logger.Errorf("Sonarr sync failed: %v", err)
	}

	// Sync from Radarr if configured
	if err := s.monitor.SyncFromRadarr(ctx, opts); err != nil {
		s.monitor.logger.Errorf("Radarr sync failed: %v", err)
	}

	return nil
}

// StartScheduledSync starts periodic synchronization with Sonarr/Radarr.
func (s *ScheduledMonitor) StartScheduledSync(ctx context.Context, interval time.Duration, opts SyncOptions) error {
	opts.ForceRefresh = false // Don't force refresh on scheduled syncs
	
	syncOpts := scheduler.Options{
		Interval: interval,
		SkipFirst: false,
		Jitter:   time.Minute * 5, // Longer jitter for sync tasks
	}

	return scheduler.RunWithOptions(ctx, syncOpts, func(ctx context.Context) error {
		return s.RunSyncTask(ctx, opts)
	})
}