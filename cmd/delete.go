package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/logging"
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
			backend := viper.GetString("db_backend")
			if store, err := database.OpenStore(dbPath, backend); err == nil {
				_ = store.DeleteSubtitle(path)
				store.Close()
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
