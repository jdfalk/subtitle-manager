package cmd

import (
	"context"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/jdfalk/subtitle-manager/pkg/scheduler"
)

var interval time.Duration
var scheduleSpec string

// autoscanCmd periodically scans a directory for subtitles using all providers.
var autoscanCmd = &cobra.Command{
	Use:   "autoscan [directory] [lang]",
	Short: "Periodically scan directory and download subtitles",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logging.GetLogger("autoscan")
		dir, lang := args[0], args[1]
		ctx := context.Background()
		var store database.SubtitleStore
		if dbPath := viper.GetString("db_path"); dbPath != "" {
			backend := viper.GetString("db_backend")
			if s, err := database.OpenStore(dbPath, backend); err == nil {
				store = s
				defer s.Close()
			} else {
				logger.Warnf("db open: %v", err)
			}
		}
		workers := viper.GetInt("scan_workers")
		if scheduleSpec != "" {
			return scheduler.ScheduleScanDirectoryCron(ctx, scheduleSpec, dir, lang, "", nil, upgrade, workers, store)
		}
		return scheduler.ScheduleScanDirectory(ctx, interval, dir, lang, "", nil, upgrade, workers, store)
	},
}

func init() {
	autoscanCmd.Flags().DurationVarP(&interval, "interval", "i", time.Hour, "scan interval")
	autoscanCmd.Flags().StringVarP(&scheduleSpec, "schedule", "s", "", "cron schedule expression")
	autoscanCmd.Flags().BoolVarP(&upgrade, "upgrade", "u", false, "replace existing subtitles")
	rootCmd.AddCommand(autoscanCmd)
}
