package cmd

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jdfalk/subtitle-manager/pkg/cli"
	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/i18n"
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
		// Set the short description with i18n after initialization
		cmd.Short = i18n.T("cli.scan.short")
		
		logger := logging.GetLogger("scan")
		dir, lang := args[0], args[1]

		// Validate inputs before processing
		if _, err := security.ValidateAndSanitizePath(dir); err != nil {
			logger.Errorf(i18n.T("common.error.invalid_path"), err)
			return err
		}

		if err := security.ValidateLanguageCode(lang); err != nil {
			logger.Errorf(i18n.T("common.error.invalid_lang"), err)
			return err
		}

		logger.Infof(i18n.T("cli.scan.scanning"), dir)
		ctx := context.Background()
		var store database.SubtitleStore
		if dbPath := viper.GetString("db_path"); dbPath != "" {
			backend := viper.GetString("db_backend")
			if s, err := database.OpenStore(dbPath, backend); err == nil {
				store = s
				defer s.Close()
			} else {
				logger.Warnf(i18n.T("common.error.db_open"), err)
			}
		}
		workers := viper.GetInt("scan_workers")

		// Count video files for progress tracking
		videoExtensions := []string{".mkv", ".mp4", ".avi", ".mov"}
		isVideoFile := func(path string) bool {
			ext := strings.ToLower(filepath.Ext(path))
			for _, e := range videoExtensions {
				if ext == e {
					return true
				}
			}
			return false
		}

		var videoFiles []string
		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && isVideoFile(path) {
				videoFiles = append(videoFiles, path)
			}
			return nil
		})
		if err != nil {
			return err
		}

		// Create progress bar
		progress := cli.NewProgressBar(len(videoFiles), "Scanning")
		defer progress.Finish()

		// Progress callback
		progressCallback := func(file string) {
			progress.Update(file)
		}

		return scanner.ScanDirectoryProgress(ctx, dir, lang, "", nil, upgrade, workers, store, progressCallback)
	},
}

func init() {
	scanCmd.Flags().BoolVarP(&upgrade, "upgrade", "u", false, "replace existing subtitles")
	rootCmd.AddCommand(scanCmd)
}
