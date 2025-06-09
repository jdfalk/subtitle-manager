package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"subtitle-manager/pkg/database"
	"subtitle-manager/pkg/logging"
	"subtitle-manager/pkg/providers"
	"subtitle-manager/pkg/scanner"
)

var upgrade bool

// scanCmd scans a directory for video files and downloads subtitles.
var scanCmd = &cobra.Command{
	Use:   "scan [provider] [directory] [lang]",
	Short: "Scan directory and download subtitles",
	Args:  cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logging.GetLogger("scan")
		name, dir, lang := args[0], args[1], args[2]
		key := viper.GetString("opensubtitles.api_key")
		p, err := providers.Get(name, key)
		if err != nil {
			return err
		}
		logger.Infof("scanning %s", dir)
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
		return scanner.ScanDirectory(ctx, dir, lang, name, p, upgrade, workers, store)
	},
}

func init() {
	scanCmd.Flags().BoolVarP(&upgrade, "upgrade", "u", false, "replace existing subtitles")
	rootCmd.AddCommand(scanCmd)
}
