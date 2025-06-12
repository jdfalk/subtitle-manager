// file: pkg/scheduler/scan.go
package scheduler

import (
	"context"
	"time"

	"subtitle-manager/pkg/database"
	"subtitle-manager/pkg/providers"
	"subtitle-manager/pkg/scanner"
)

// ScheduleScanDirectory periodically scans dir for subtitles using provider p.
// The scan runs immediately and then every interval until ctx is canceled.
func ScheduleScanDirectory(ctx context.Context, interval time.Duration, dir, lang, providerName string, p providers.Provider, upgrade bool, workers int, store database.SubtitleStore) error {
	return Run(ctx, interval, func(ctx context.Context) error {
		return scanner.ScanDirectory(ctx, dir, lang, providerName, p, upgrade, workers, store)
	})
}
