package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/logging"
)

// downloadsCmd shows subtitle download history.
var downloadsVideo string

var downloadsCmd = &cobra.Command{
	Use:   "downloads",
	Short: "Show subtitle download history",
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logging.GetLogger("downloads")
		dbPath := viper.GetString("db_path")
		backend := viper.GetString("db_backend")
		store, err := database.OpenStore(dbPath, backend)
		if err != nil {
			return err
		}
		defer store.Close()

		var recs []database.DownloadRecord
		if downloadsVideo != "" {
			recs, err = store.ListDownloadsByVideo(downloadsVideo)
		} else {
			recs, err = store.ListDownloads()
		}
		if err != nil {
			return err
		}
		for _, r := range recs {
			fmt.Printf("%s\t%s\t%s\t%s\t%s\n", r.ID, r.VideoFile, r.Language, r.Provider, r.CreatedAt.Format(time.RFC3339))
		}
		logger.Infof("listed %d records", len(recs))
		return nil
	},
}

func init() {
	downloadsCmd.Flags().StringVar(&downloadsVideo, "video", "", "filter by video file")
	rootCmd.AddCommand(downloadsCmd)
}
