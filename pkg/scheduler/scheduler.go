// file: pkg/scheduler/scheduler.go
package scheduler

import (
	"context"
	"time"

	"subtitle-manager/pkg/logging"
)

// Run executes fn immediately and then at each interval until ctx is canceled.
// Errors returned by fn after the first execution are logged and do not stop
// subsequent runs. The initial error is returned to the caller.
func Run(ctx context.Context, interval time.Duration, fn func(context.Context) error) error {
	logger := logging.GetLogger("scheduler")
	if err := fn(ctx); err != nil {
		return err
	}
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			if err := fn(ctx); err != nil && ctx.Err() == nil {
				logger.Warnf("scheduled run: %v", err)
			}
		}
	}
}
