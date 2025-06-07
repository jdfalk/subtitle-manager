package cmd

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"subtitle-manager/pkg/logging"
	"subtitle-manager/pkg/providers"
	"subtitle-manager/pkg/scanner"
)

var sonarrCmd = &cobra.Command{
	Use:   "sonarr [provider] [lang]",
	Short: "Handle Sonarr download event",
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logging.GetLogger("sonarr")
		providerName, lang := args[0], args[1]
		path := os.Getenv("SONARR_EPISODEFILE_PATH")
		if p := viper.GetString("path"); p != "" {
			path = p
		}
		if path == "" {
			return cmd.Usage()
		}
		key := viper.GetString("opensubtitles.api_key")
		p, err := providers.Get(providerName, key)
		if err != nil {
			return err
		}
		ctx := context.Background()
		logger.Infof("processing %s", path)
		return scanner.ProcessFile(ctx, path, lang, p, true)
	},
}

func init() {
	sonarrCmd.Flags().String("path", "", "path to downloaded episode")
	rootCmd.AddCommand(sonarrCmd)
}
