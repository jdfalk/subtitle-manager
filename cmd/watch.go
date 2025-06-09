package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"subtitle-manager/pkg/database"
		var store database.SubtitleStore
		if dbPath := viper.GetString("db_path"); dbPath != "" {
			backend := viper.GetString("db_backend")
			if s, err := database.OpenStore(dbPath, backend); err == nil {
				store = s
				defer s.Close()
			} else {
				logger.Warnf("db open: %v", err)
			}
		}
			return watcher.WatchDirectoryRecursive(ctx, dir, lang, name, p, store)
		return watcher.WatchDirectory(ctx, dir, lang, name, p, store)
	"subtitle-manager/pkg/providers"
	"subtitle-manager/pkg/watcher"
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
		if dbPath := viper.GetString("db_path"); dbPath != "" {
			backend := viper.GetString("db_backend")
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
