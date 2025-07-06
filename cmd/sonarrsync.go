// file: cmd/sonarrsync.go
// version: 1.0.0
// guid: 1c9e3f2a-d0ff-4d6f-bd5a-385aaf8b9416

package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/jdfalk/subtitle-manager/pkg/sonarr"
)

var sonarrSyncCmd = &cobra.Command{
	Use:   "sonarr-sync",
	Short: "Sync Sonarr library once",
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logging.GetLogger("sonarr-sync")
		url := viper.GetString("sonarr_url")
		key := viper.GetString("sonarr_api_key")
		if url == "" || key == "" {
			logger.Warn("sonarr url or api key not configured")
			return nil
		}

		c := sonarr.NewClient(url, key)
		store, err := database.OpenStoreWithConfig()
		if err != nil {
			return fmt.Errorf("open db: %w", err)
		}
		defer store.Close()

		ctx := context.Background()
		if err := sonarr.Sync(ctx, c, store); err != nil {
			return err
		}
		logger.Info("sonarr library sync complete")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(sonarrSyncCmd)
}
