package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"

	"subtitle-manager/pkg/database"
	"subtitle-manager/pkg/logging"
)

// deleteCmd removes a subtitle file from disk and the database.
var deleteCmd = &cobra.Command{
	Use:   "delete [file]",
	Short: "Delete subtitle file and remove record",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logging.GetLogger("delete")
		path := args[0]
		if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
			return err
		}
		if dbPath := viper.GetString("db_path"); dbPath != "" {
			if db, err := database.Open(dbPath); err == nil {
				_ = database.DeleteSubtitle(db, path)
				db.Close()
			} else {
				logger.Warnf("db open: %v", err)
			}
		}
		logger.Infof("deleted %s", path)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
