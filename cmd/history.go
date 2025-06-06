package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"subtitle-manager/pkg/database"
	"subtitle-manager/pkg/logging"
)

var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Show translation history",
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logging.GetLogger("history")
		dbPath := viper.GetString("db_path")
		db, err := database.Open(dbPath)
		if err != nil {
			return err
		}
		defer db.Close()

		recs, err := database.ListSubtitles(db)
		if err != nil {
			return err
		}
		for _, r := range recs {
			fmt.Printf("%d\t%s\t%s\t%s\t%s\n", r.ID, r.File, r.Language, r.Service, r.CreatedAt.Format(time.RFC3339))
		}
		logger.Infof("listed %d records", len(recs))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(historyCmd)
}
