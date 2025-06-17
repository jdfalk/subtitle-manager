// file: cmd/fetch.go
package cmd

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/jdfalk/subtitle-manager/pkg/providers"
)

var fetchCmd = &cobra.Command{
	Use:   "fetch [media] [lang] [output]",
	Short: "Download subtitles using all providers",
	Args:  cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logging.GetLogger("fetch")
		media, lang, out := args[0], args[1], args[2]
		key := viper.GetString("opensubtitles.api_key")
		data, name, err := providers.FetchFromAll(context.Background(), media, lang, key)
		if err != nil {
			return err
		}
		if err := os.WriteFile(out, data, 0644); err != nil {
			return err
		}
		if dbPath := viper.GetString("db_path"); dbPath != "" {
			backend := viper.GetString("db_backend")
			if store, err := database.OpenStore(dbPath, backend); err == nil {
				_ = store.InsertDownload(&database.DownloadRecord{File: out, VideoFile: media, Provider: name, Language: lang})
				store.Close()
			} else {
				logger.Warnf("db open: %v", err)
			}
		}
		logger.Infof("downloaded subtitle to %s", out)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
}
