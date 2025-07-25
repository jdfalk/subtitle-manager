// file: cmd/fetch.go
package cmd

import (
	"context"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/jdfalk/subtitle-manager/pkg/providers"
	"github.com/jdfalk/subtitle-manager/pkg/tagging"
)

var tags string
var useProfile bool

var fetchCmd = &cobra.Command{
	Use:   "fetch [media] [lang] [output]",
	Short: "Download subtitles using all providers",
	Long: `Download subtitles using all providers.

When --profile is specified, the language parameter is ignored and the system 
uses language preferences from the media item's assigned language profile.`,
	Args: cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logging.GetLogger("fetch")
		media, lang, out := args[0], args[1], args[2]
		key := viper.GetString("opensubtitles.api_key")
		tagNames := []string{}
		if tags != "" {
			tagNames = strings.Split(tags, ",")
		}

		var data []byte
		var name string
		var actualLang string
		var err error

		// Get database connection for profile-based fetching
		if useProfile || len(tagNames) > 0 {
			dbPath := database.GetDatabasePath()
			store, errStore := database.OpenSQLStore(dbPath)
			if errStore != nil {
				return errStore
			}
			defer store.Close()

			if useProfile && len(tagNames) > 0 {
				// Use both profiles and tags
				tm := tagging.NewTagManager(store.DB())
				data, name, actualLang, err = providers.FetchWithProfileTagged(context.Background(), store.DB(), media, key, tagNames, tm)
			} else if useProfile {
				// Use profiles only
				data, name, actualLang, err = providers.FetchWithProfile(context.Background(), store.DB(), media, key)
			} else {
				// Use tags only (existing behavior)
				tm := tagging.NewTagManager(store.DB())
				data, name, err = providers.FetchFromTagged(context.Background(), media, lang, key, tagNames, tm)
				actualLang = lang
			}
		} else {
			// Standard fetch without profiles or tags
			data, name, err = providers.FetchFromAll(context.Background(), media, lang, key)
			actualLang = lang
		}
		if err != nil {
			return err
		}
		if err := os.WriteFile(out, data, 0644); err != nil {
			return err
		}
		if dbPath := viper.GetString("db_path"); dbPath != "" {
			backend := viper.GetString("db_backend")
			if store, err := database.OpenStore(dbPath, backend); err == nil {
				_ = store.InsertDownload(&database.DownloadRecord{File: out, VideoFile: media, Provider: name, Language: actualLang})
				store.Close()
			} else {
				logger.Warnf("db open: %v", err)
			}
		}
		if useProfile {
			logger.Infof("downloaded %s subtitle using profile-based search to %s", actualLang, out)
		} else {
			logger.Infof("downloaded subtitle to %s", out)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
	fetchCmd.Flags().StringVar(&tags, "tags", "", "comma separated provider tags")
	fetchCmd.Flags().BoolVar(&useProfile, "profile", false, "use language profile preferences for the media item")
}
