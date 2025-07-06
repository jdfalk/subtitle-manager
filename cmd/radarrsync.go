// file: cmd/radarrsync.go
// version: 1.0.0
// guid: a013d1b8-4cd3-4f59-8e4d-0d82cd9acae7

package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/jdfalk/subtitle-manager/pkg/radarr"
)

var radarrSyncCmd = &cobra.Command{
	Use:   "radarr-sync",
	Short: "Sync Radarr library once",
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logging.GetLogger("radarr-sync")
		url := viper.GetString("radarr_url")
		key := viper.GetString("radarr_api_key")
		if url == "" || key == "" {
			logger.Warn("radarr url or api key not configured")
			return nil
		}

		c := radarr.NewClient(url, key)
		store, err := database.OpenStoreWithConfig()
		if err != nil {
			return fmt.Errorf("open db: %w", err)
		}
		defer store.Close()

		ctx := context.Background()
		if err := radarr.Sync(ctx, c, store); err != nil {
			return err
		}
		logger.Info("radarr library sync complete")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(radarrSyncCmd)
}
