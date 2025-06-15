package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

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
		return metadata.ScanLibrary(context.Background(), dir, store)
	},
}

func init() {
	rootCmd.AddCommand(scanLibCmd)
}
