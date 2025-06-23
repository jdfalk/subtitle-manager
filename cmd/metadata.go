// file: cmd/metadata.go
package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/metadata"
)

var metadataCmd = &cobra.Command{
	Use:   "metadata",
	Short: "Manage media metadata",
}

var metadataSearchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Search TMDB for a title",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		key := viper.GetString("tmdb_api_key")
		info, err := metadata.QueryMovie(context.Background(), args[0], 0, key)
		if err != nil {
			return err
		}
		fmt.Printf("%s (%d) id=%d\n", info.Title, info.Year, info.TMDBID)
		return nil
	},
}

var (
	setTitle string
	setGroup string
	setAlt   string
	setLocks string
)

var metadataUpdateCmd = &cobra.Command{
	Use:   "update [file]",
	Short: "Update metadata for a media item",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		path := args[0]
		dbPath := viper.GetString("db_path")
		backend := viper.GetString("db_backend")
		store, err := database.OpenStore(dbPath, backend)
		if err != nil {
			return err
		}
		defer store.Close()
		if setGroup != "" {
			if err := store.SetMediaReleaseGroup(path, setGroup); err != nil {
				return fmt.Errorf("failed to set release group: %w", err)
			}
		}
		if setAlt != "" {
			titles := strings.Split(setAlt, ",")
			for i := range titles {
				titles[i] = strings.TrimSpace(titles[i])
			}
			if err := store.SetMediaAltTitles(path, titles); err != nil {
				return fmt.Errorf("failed to set alternate titles: %w", err)
			}
		}
		if setLocks != "" {
			if err := store.SetMediaFieldLocks(path, setLocks); err != nil {
				return fmt.Errorf("failed to set field locks: %w", err)
			}
		}
		if setTitle != "" {
			if err := store.SetMediaTitle(path, setTitle); err != nil {
				return fmt.Errorf("failed to set title: %w", err)
			}
		}
		return nil
	},
}

func init() {
	metadataUpdateCmd.Flags().StringVar(&setTitle, "title", "", "new title")
	metadataUpdateCmd.Flags().StringVar(&setGroup, "release-group", "", "release group")
	metadataUpdateCmd.Flags().StringVar(&setAlt, "alt", "", "comma separated alternate titles")
	metadataUpdateCmd.Flags().StringVar(&setLocks, "lock", "", "comma separated locked fields")
	metadataCmd.AddCommand(metadataSearchCmd)
	metadataCmd.AddCommand(metadataUpdateCmd)
	rootCmd.AddCommand(metadataCmd)
}
