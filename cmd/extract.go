package cmd

import (
	"os"

	"github.com/asticode/go-astisub"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/jdfalk/subtitle-manager/pkg/subtitles"
)

var extractCmd = &cobra.Command{
	Use:   "extract [media] [output]",
	Short: "Extract subtitles from media",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logging.GetLogger("extract")
		media, out := args[0], args[1]
		if ff := viper.GetString("ffmpeg_path"); ff != "" {
			subtitles.SetFFmpegPath(ff)
		}
		items, err := subtitles.ExtractFromMedia(media)
		if err != nil {
			return err
		}
		sub := astisub.NewSubtitles()
		sub.Items = items
		f, err := os.Create(out)
		if err != nil {
			return err
		}
		defer f.Close()
		if err := sub.WriteToSRT(f); err != nil {
			return err
		}
		if dbPath := viper.GetString("db_path"); dbPath != "" {
			backend := viper.GetString("db_backend")
			if store, err := database.OpenStore(dbPath, backend); err == nil {
				_ = store.InsertSubtitle(&database.SubtitleRecord{
					File:      out,
					VideoFile: media,
					Service:   "extract",
					Embedded:  true,
				})
				store.Close()
			} else {
				logger.Warnf("db open: %v", err)
			}
		}
		logger.Infof("extracted subtitles from %s to %s", media, out)
		return nil
	},
}

func init() {
	extractCmd.Flags().String("ffmpeg", "", "path to ffmpeg binary")
	viper.BindPFlag("ffmpeg_path", extractCmd.Flags().Lookup("ffmpeg"))
	rootCmd.AddCommand(extractCmd)
}
