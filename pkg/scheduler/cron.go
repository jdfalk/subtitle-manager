// file: pkg/scheduler/cron.go
package scheduler

import (
	"context"

	"github.com/robfig/cron/v3"

	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/jdfalk/subtitle-manager/pkg/providers"
	"github.com/jdfalk/subtitle-manager/pkg/scanner"
)

// RunCron executes fn immediately and then according to spec until ctx is canceled.
// Spec may be a standard cron expression or an "@every" duration. Errors after the
// first run are logged and do not stop subsequent executions. The error from the
// first run is returned to the caller.
func RunCron(ctx context.Context, spec string, fn func(context.Context) error) error {
	logger := logging.GetLogger("scheduler")
	if err := fn(ctx); err != nil {
		return err
	}
	parser := cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)
	c := cron.New(cron.WithParser(parser))
	if _, err := c.AddFunc(spec, func() {
		if err := fn(ctx); err != nil && ctx.Err() == nil {
			logger.Warnf("scheduled run: %v", err)
		}
	}); err != nil {
		return err
	}
	c.Start()
	<-ctx.Done()
	c.Stop()
	return ctx.Err()
}

// ScheduleScanDirectoryCron schedules a directory scan using spec until ctx is canceled.
// The scan runs immediately and then on the cron schedule.
func ScheduleScanDirectoryCron(ctx context.Context, spec, dir, lang, providerName string, p providers.Provider, upgrade bool, workers int, store database.SubtitleStore) error {
	return RunCron(ctx, spec, func(ctx context.Context) error {
		return scanner.ScanDirectory(ctx, dir, lang, providerName, p, upgrade, workers, store)
	})
}
