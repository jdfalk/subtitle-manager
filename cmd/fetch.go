// file: cmd/fetch.go
package cmd

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"subtitle-manager/pkg/logging"
	"subtitle-manager/pkg/providers"
)

// fetchCmd downloads subtitles for a media file using a provider.
var fetchCmd = &cobra.Command{
	Use:   "fetch [provider] [media] [lang] [output]",
	Short: "Download subtitles using a provider",
	Args:  cobra.ExactArgs(4),
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logging.GetLogger("fetch")
		name, media, lang, out := args[0], args[1], args[2], args[3]
		key := viper.GetString("opensubtitles.api_key")
		p, err := providers.Get(name, key)
		if err != nil {
			return err
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
