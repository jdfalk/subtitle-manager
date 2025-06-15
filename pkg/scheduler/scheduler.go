// file: pkg/scheduler/scheduler.go
package scheduler

import (
	"context"
	"math/rand"
	"time"

	"github.com/jdfalk/subtitle-manager/pkg/logging"
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

// RunWithOptions executes fn according to the provided Options. When SkipFirst
// is false the task runs immediately. Jitter is applied before each execution
// (except the initial run) and MaxRuns limits the total number of runs.
func RunWithOptions(ctx context.Context, opts Options, fn func(context.Context) error) error {
	logger := logging.GetLogger("scheduler")
	runs := 0
	if !opts.SkipFirst {
		if err := fn(ctx); err != nil {
			return err
		}
		runs++
		if opts.MaxRuns > 0 && runs >= opts.MaxRuns {
			return nil
		}
	}
	ticker := time.NewTicker(opts.Interval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			if opts.Jitter > 0 {
				d := time.Duration(rand.Int63n(int64(opts.Jitter)))
				select {
				case <-ctx.Done():
					return ctx.Err()
				case <-time.After(d):
				}
			}
			if err := fn(ctx); err != nil && ctx.Err() == nil {
				logger.Warnf("scheduled run: %v", err)
			}
			runs++
			if opts.MaxRuns > 0 && runs >= opts.MaxRuns {
				return nil
			}
		}
	}
}
