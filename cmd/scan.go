package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/jdfalk/subtitle-manager/pkg/scanner"
	"github.com/jdfalk/subtitle-manager/pkg/security"
)

var upgrade bool

// scanCmd scans a directory for video files and downloads subtitles.
var scanCmd = &cobra.Command{
	Use:   "scan [directory] [lang]",
	Short: "Scan directory and download subtitles",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logging.GetLogger("scan")
		dir, lang := args[0], args[1]

		// Validate inputs before processing
		if _, err := security.ValidateAndSanitizePath(dir); err != nil {
			logger.Errorf("invalid directory path: %v", err)
			return err
		}

		if err := security.ValidateLanguageCode(lang); err != nil {
			logger.Errorf("invalid language code: %v", err)
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
		return scanner.ScanDirectory(ctx, dir, lang, "", nil, upgrade, workers, store)
	},
}

func init() {
	scanCmd.Flags().BoolVarP(&upgrade, "upgrade", "u", false, "replace existing subtitles")
	rootCmd.AddCommand(scanCmd)
}
