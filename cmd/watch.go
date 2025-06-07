package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"subtitle-manager/pkg/logging"
	"subtitle-manager/pkg/providers"
	"subtitle-manager/pkg/providers/opensubtitles"
	"subtitle-manager/pkg/watcher"
)

var watchCmd = &cobra.Command{
	Use:   "watch [provider] [directory] [lang]",
	Short: "Watch directory and auto-download subtitles",
	Args:  cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logging.GetLogger("watch")
		name, dir, lang := args[0], args[1], args[2]
		var p providers.Provider
		switch name {
		case "opensubtitles":
			key := viper.GetString("opensubtitles.api_key")
			p = opensubtitles.New(key)
		default:
			return fmt.Errorf("unknown provider %s", name)
		}
		logger.Infof("watching %s", dir)
		ctx := context.Background()
		return watcher.WatchDirectory(ctx, dir, lang, p)
	},
}

func init() {
	rootCmd.AddCommand(watchCmd)
}
