package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"subtitle-manager/pkg/database"
	"subtitle-manager/pkg/logging"
	"subtitle-manager/pkg/plex"
)

var plexCmd = &cobra.Command{
	Use:   "plex",
	Short: "Interact with Plex server",
}

var plexLibraryCmd = &cobra.Command{
	Use:   "library",
	Short: "Sync Plex library into database",
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logging.GetLogger("plex")
		url := viper.GetString("plex.url")
		token := viper.GetString("plex.token")
		if url == "" || token == "" {
			return cmd.Usage()
		}
		c := plex.NewClient(url, token)
		ctx := context.Background()
		items, err := c.AllItems(ctx)
		if err != nil {
			return err
		}
		dbPath := viper.GetString("db_path")
		backend := viper.GetString("db_backend")
		store, err := database.OpenStore(dbPath, backend)
		if err != nil {
			return err
		}
		defer store.Close()
		for _, it := range items {
			rec := &database.MediaItem{Path: it.Path, Title: it.Title, Season: it.Season, Episode: it.Episode}
			if err := store.InsertMediaItem(rec); err != nil {
				logger.Warnf("insert %s: %v", it.Path, err)
			}
		}
		logger.Infof("synced %d items", len(items))
		return nil
	},
}

var refreshKey string

var plexRefreshCmd = &cobra.Command{
	Use:   "refresh",
	Short: "Trigger Plex to refresh the library",
	RunE: func(cmd *cobra.Command, args []string) error {
		url := viper.GetString("plex.url")
		token := viper.GetString("plex.token")
		if url == "" || token == "" {
			return cmd.Usage()
		}
		c := plex.NewClient(url, token)
		ctx := context.Background()
		if refreshKey != "" {
			return c.RefreshItem(ctx, refreshKey)
		}
		return c.RefreshLibrary(ctx)
	},
}

func init() {
	plexCmd.PersistentFlags().String("url", "", "Plex base URL")
	plexCmd.PersistentFlags().String("token", "", "Plex API token")
	viper.BindPFlag("plex.url", plexCmd.PersistentFlags().Lookup("url"))
	viper.BindPFlag("plex.token", plexCmd.PersistentFlags().Lookup("token"))

	plexRefreshCmd.Flags().StringVar(&refreshKey, "item", "", "ratingKey of item to refresh")
	plexCmd.AddCommand(plexLibraryCmd)
	plexCmd.AddCommand(plexRefreshCmd)
	rootCmd.AddCommand(plexCmd)
}
