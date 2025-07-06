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
	setTitle     string
	setGroup     string
	setAlt       string
	setLocks     string
	fetchID      int
	fetchYear    int
	fetchSeason  int
	fetchEpisode int
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

var metadataFetchCmd = &cobra.Command{
	Use:   "fetch [title]",
	Short: "Fetch metadata with languages and rating",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		tmdbKey := viper.GetString("tmdb_api_key")
		omdbKey := viper.GetString("omdb_api_key")
		ctx := context.Background()
		var info *metadata.MediaInfo
		var err error
		if fetchID > 0 {
			if fetchSeason > 0 {
				info, err = metadata.FetchEpisodeMetadataByID(ctx, fetchID, fetchSeason, fetchEpisode, tmdbKey, omdbKey)
			} else {
				info, err = metadata.FetchMovieMetadataByID(ctx, fetchID, tmdbKey, omdbKey)
			}
		} else {
			if len(args) == 0 {
				return fmt.Errorf("title or --id required")
			}
			title := args[0]
			if fetchSeason > 0 {
				info, err = metadata.FetchEpisodeMetadata(ctx, title, fetchSeason, fetchEpisode, tmdbKey, omdbKey)
			} else {
				info, err = metadata.FetchMovieMetadata(ctx, title, fetchYear, tmdbKey, omdbKey)
			}
		}
		if err != nil {
			return err
		}
		fmt.Printf("%s (%d) id=%d\n", info.Title, info.Year, info.TMDBID)
		if len(info.Languages) > 0 {
			fmt.Printf("Languages: %s\n", strings.Join(info.Languages, ", "))
		}
		if info.Rating > 0 {
			fmt.Printf("Rating: %.1f\n", info.Rating)
		}
		return nil
	},
}

func init() {
	metadataUpdateCmd.Flags().StringVar(&setTitle, "title", "", "new title")
	metadataUpdateCmd.Flags().StringVar(&setGroup, "release-group", "", "release group")
	metadataUpdateCmd.Flags().StringVar(&setAlt, "alt", "", "comma separated alternate titles")
	metadataUpdateCmd.Flags().StringVar(&setLocks, "lock", "", "comma separated locked fields")
	metadataFetchCmd.Flags().IntVar(&fetchID, "id", 0, "TMDB identifier")
	metadataFetchCmd.Flags().IntVar(&fetchYear, "year", 0, "release year for movie")
	metadataFetchCmd.Flags().IntVar(&fetchSeason, "season", 0, "season number for episode")
	metadataFetchCmd.Flags().IntVar(&fetchEpisode, "episode", 0, "episode number")
	metadataCmd.AddCommand(metadataSearchCmd)
	metadataCmd.AddCommand(metadataUpdateCmd)
	metadataCmd.AddCommand(metadataFetchCmd)
	rootCmd.AddCommand(metadataCmd)
}
