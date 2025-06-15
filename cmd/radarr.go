package cmd

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/jdfalk/subtitle-manager/pkg/providers"
	"github.com/jdfalk/subtitle-manager/pkg/scanner"
)

var radarrCmd = &cobra.Command{
	Use:   "radarr [provider] [lang]",
	Short: "Handle Radarr download event",
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logging.GetLogger("radarr")
		providerName, lang := args[0], args[1]
		path := os.Getenv("RADARR_MOVIEFILE_PATH")
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
		return scanner.ProcessFile(ctx, path, lang, providerName, p, true, store)
	},
}

func init() {
	radarrCmd.Flags().String("path", "", "path to downloaded movie")
	rootCmd.AddCommand(radarrCmd)
}
