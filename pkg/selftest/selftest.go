// file: pkg/selftest/selftest.go
// version: 1.0.0
// guid: 0a1b2c3d-4e5f-6a7b-8c9d-0e1f2a3b4c5d
package selftest

import (
	"context"
	"database/sql"
	"os"
	"time"

	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/jdfalk/subtitle-manager/pkg/scheduler"
)

// ExitFunc is used for process termination.
// It can be overridden in tests.
var ExitFunc = os.Exit

// StartPeriodic runs self-tests at the given frequency. If frequency
// is empty or invalid, a default of 5 minutes is used. The only
// current check verifies database connectivity. On failure the process
// exits so external supervisors can restart it.
func StartPeriodic(ctx context.Context, db *sql.DB, freq string) {
	logger := logging.GetLogger("selftest")
	d, err := time.ParseDuration(freq)
	if err != nil || d <= 0 {
		d = 5 * time.Minute
	}
	go func() {
		_ = scheduler.Run(ctx, d, func(c context.Context) error {
			if err := db.PingContext(c); err != nil {
				logger.Errorf("selftest database ping failed: %v", err)
				ExitFunc(1)
				return err
			}
			logger.Debug("selftest passed")
			return nil
		})
	}()
}
