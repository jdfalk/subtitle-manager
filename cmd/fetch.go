// file: cmd/fetch.go
package cmd

import (
	"context"
	"fmt"
	"os"

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
)

// fetchCmd downloads subtitles for a media file using a provider.
var fetchCmd = &cobra.Command{
	Use:   "fetch [provider] [media] [lang] [output]",
	Short: "Download subtitles using a provider",
	Args:  cobra.ExactArgs(4),
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logging.GetLogger("fetch")
		name, media, lang, out := args[0], args[1], args[2], args[3]
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
		data, err := p.Fetch(context.Background(), media, lang)
		if err != nil {
			return err
		}
		if err := os.WriteFile(out, data, 0644); err != nil {
			return err
		}
		logger.Infof("downloaded subtitle to %s", out)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
}
