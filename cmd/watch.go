package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/jdfalk/subtitle-manager/pkg/providers"
	"github.com/jdfalk/subtitle-manager/pkg/watcher"
)

var recursive bool

var watchCmd = &cobra.Command{
	Use:   "watch [provider] [directory] [lang]",
	Short: "Watch directory and auto-download subtitles",
	Args:  cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logging.GetLogger("watch")
		name, dir, lang := args[0], args[1], args[2]
		key := viper.GetString("opensubtitles.api_key")
		p, err := providers.Get(name, key)
		if err != nil {
			return err
		}
		logger.Infof("watching %s", dir)
		ctx := context.Background()
		var store database.SubtitleStore
		dbPath := database.GetDatabasePath()
		if dbPath != "" {
			backend := database.GetDatabaseBackend()
			if s, err := database.OpenStore(dbPath, backend); err == nil {
				store = s
				defer s.Close()
			} else {
				logger.Warnf("db open: %v", err)
			}
		}
		if recursive {
			return watcher.WatchDirectoryRecursive(ctx, dir, lang, name, p, store)
		}
		return watcher.WatchDirectory(ctx, dir, lang, name, p, store)
	},
}

func init() {
	watchCmd.Flags().BoolVarP(&recursive, "recursive", "r", false, "watch directories recursively")
	rootCmd.AddCommand(watchCmd)
}
