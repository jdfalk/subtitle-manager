package cmd

import (
	"context"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jdfalk/subtitle-manager/pkg/cli"
	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/jdfalk/subtitle-manager/pkg/metadata"
)

// scanLibCmd indexes video files in a directory into the media library.
var scanLibCmd = &cobra.Command{
	Use:   "scanlib [directory]",
	Short: "Scan media library and store metadata",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := args[0]
		logger := logging.GetLogger("scanlib")
		dbPath := viper.GetString("db_path")
		backend := viper.GetString("db_backend")
		store, err := database.OpenStore(dbPath, backend)
		if err != nil {
			return err
		}
		defer store.Close()
		logger.Infof("scanning %s", dir)

		// Count media files for progress tracking
		var mediaFiles []string
		err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				// Simple media file detection
				ext := filepath.Ext(path)
				if ext == ".mkv" || ext == ".mp4" || ext == ".avi" || ext == ".mov" {
					mediaFiles = append(mediaFiles, path)
				}
			}
			return nil
		})
		if err != nil {
			return err
		}

		// Create progress bar
		progress := cli.NewProgressBar(len(mediaFiles), "Scanning library")
		defer progress.Finish()

		// Progress callback
		progressCallback := func(file string) {
			progress.Update(file)
		}

		return metadata.ScanLibraryProgress(context.Background(), dir, store, progressCallback)
	},
}

func init() {
	rootCmd.AddCommand(scanLibCmd)
}
