package cmd

import (
	"github.com/spf13/cobra"
	"subtitle-manager/pkg/database"
)

// migrateCmd migrates subtitle history from SQLite to PebbleDB.
var migrateCmd = &cobra.Command{
	Use:   "migrate [sqlite-file] [pebble-dir]",
	Short: "Migrate subtitles from SQLite to PebbleDB",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		sqlitePath := args[0]
		pebblePath := args[1]
		if err := database.MigrateToPebble(sqlitePath, pebblePath); err != nil {
			return err
		}
		cmd.Printf("Migrated subtitles from %s to %s\n", sqlitePath, pebblePath)
		return nil
	},
}

func init() { rootCmd.AddCommand(migrateCmd) }
