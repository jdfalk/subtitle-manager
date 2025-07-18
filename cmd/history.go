package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/logging"
)

var historyVideo string

var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Show translation history",
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logging.GetLogger("history")
		dbPath := database.GetDatabasePath()
		backend := database.GetDatabaseBackend()
		store, err := database.OpenStore(dbPath, backend)
		if err != nil {
			return err
		}
		defer store.Close()

		var recs []database.SubtitleRecord
		if historyVideo != "" {
			recs, err = store.ListSubtitlesByVideo(historyVideo)
		} else {
			recs, err = store.ListSubtitles()
		}
		if err != nil {
			return err
		}
		for _, r := range recs {
			fmt.Printf("%s\t%s\t%s\t%s\t%s\n", r.ID, r.File, r.Language, r.Service, r.CreatedAt.Format(time.RFC3339))
		}
		logger.Infof("listed %d records", len(recs))
		return nil
	},
}

func init() {
	historyCmd.Flags().StringVar(&historyVideo, "video", "", "filter by video file")
	rootCmd.AddCommand(historyCmd)
}
