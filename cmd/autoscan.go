package cmd

import (
	"context"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"subtitle-manager/pkg/database"
	"subtitle-manager/pkg/logging"
	"subtitle-manager/pkg/providers"
	"subtitle-manager/pkg/scheduler"
)

var interval time.Duration

// autoscanCmd periodically scans a directory for subtitles using a provider.
var autoscanCmd = &cobra.Command{
	Use:   "autoscan [provider] [directory] [lang]",
	Short: "Periodically scan directory and download subtitles",
	Args:  cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logging.GetLogger("autoscan")
		name, dir, lang := args[0], args[1], args[2]
		key := viper.GetString("opensubtitles.api_key")
		p, err := providers.Get(name, key)
		if err != nil {
			return err
		}
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
		return scheduler.ScheduleScanDirectory(ctx, interval, dir, lang, name, p, upgrade, workers, store)
	},
}

func init() {
	autoscanCmd.Flags().DurationVarP(&interval, "interval", "i", time.Hour, "scan interval")
	autoscanCmd.Flags().BoolVarP(&upgrade, "upgrade", "u", false, "replace existing subtitles")
	rootCmd.AddCommand(autoscanCmd)
}
