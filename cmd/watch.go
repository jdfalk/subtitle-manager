package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"subtitle-manager/pkg/logging"
	"subtitle-manager/pkg/providers"
	"subtitle-manager/pkg/providers/addic7ed"
	"subtitle-manager/pkg/providers/betaseries"
	"subtitle-manager/pkg/providers/bsplayer"
	"subtitle-manager/pkg/providers/greeksubs"
	"subtitle-manager/pkg/providers/legendasdivx"
	"subtitle-manager/pkg/providers/opensubtitles"
	"subtitle-manager/pkg/providers/podnapisi"
	"subtitle-manager/pkg/providers/subscene"
	"subtitle-manager/pkg/providers/titlovi"
	"subtitle-manager/pkg/providers/tvsubtitles"
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
		var p providers.Provider
		switch name {
		case "opensubtitles":
			key := viper.GetString("opensubtitles.api_key")
			p = opensubtitles.New(key)
		case "subscene":
			p = subscene.New()
		case "addic7ed":
			p = addic7ed.New()
		case "betaseries":
			p = betaseries.New()
		case "bsplayer":
			p = bsplayer.New()
		case "podnapisi":
			p = podnapisi.New()
		case "tvsubtitles":
			p = tvsubtitles.New()
		case "titlovi":
			p = titlovi.New()
		case "legendasdivx":
			p = legendasdivx.New()
		case "greeksubs":
			p = greeksubs.New()
		default:
			return fmt.Errorf("unknown provider %s", name)
		}
		logger.Infof("watching %s", dir)
		ctx := context.Background()
		if recursive {
			return watcher.WatchDirectoryRecursive(ctx, dir, lang, p)
		}
		return watcher.WatchDirectory(ctx, dir, lang, p)
	},
}

func init() {
	watchCmd.Flags().BoolVarP(&recursive, "recursive", "r", false, "watch directories recursively")
	rootCmd.AddCommand(watchCmd)
}
