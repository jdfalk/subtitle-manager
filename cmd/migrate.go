package cmd

import (
	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/spf13/cobra"
)

// migrateCmd migrates subtitle history between database backends.
var migrateCmd = &cobra.Command{
	Use:   "migrate <src-backend> <src-path> <dest-backend> <dest-path>",
	Short: "Migrate subtitles between database backends",
	Args:  cobra.RangeArgs(2, 4),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 2 {
			// Backwards compatibility: sqlite to pebble
			return database.MigrateToPebble(args[0], args[1])
		}
		srcBackend := args[0]
		srcPath := args[1]
		destBackend := args[2]
		destPath := args[3]

		src, err := database.OpenStore(srcPath, srcBackend)
		if err != nil {
			return err
		}
		defer src.Close()

		dest, err := database.OpenStore(destPath, destBackend)
		if err != nil {
			return err
		}
		defer dest.Close()

		if err := database.Migrate(src, dest); err != nil {
			return err
		}

		cmd.Printf("Migrated data from %s:%s to %s:%s\n", srcBackend, srcPath, destBackend, destPath)
		return nil
	},
}

func init() { rootCmd.AddCommand(migrateCmd) }
