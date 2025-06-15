package cmd

import (
	"os"

	"github.com/asticode/go-astisub"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/jdfalk/subtitle-manager/pkg/syncer"
)

// syncCmd aligns an external subtitle file with a media file.
var syncCmd = &cobra.Command{
	Use:   "sync [media] [subtitle] [output]",
	Short: "Synchronize subtitle with media",
	Args:  cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logging.GetLogger("sync")
		media, subPath, out := args[0], args[1], args[2]
		opts := syncer.Options{
			UseAudio:       viper.GetBool("sync_use_audio"),
			UseEmbedded:    viper.GetBool("sync_use_embedded"),
			AudioTrack:     viper.GetInt("sync_audio_track"),
			SubtitleTracks: viper.GetIntSlice("sync_sub_tracks"),
			OpenAIKey:      viper.GetString("openai_api_key"),
		}
		items, err := syncer.Sync(media, subPath, opts)
		if err != nil {
			return err
		}
		tmpSub := astisub.Subtitles{Items: items}
		f, err := os.Create(out)
		if err != nil {
			return err
		}
		defer f.Close()
		if err := tmpSub.WriteToSRT(f); err != nil {
			return err
		}
		logger.Infof("synchronized %s -> %s", subPath, out)
		return nil
	},
}

func init() {
	syncCmd.Flags().Bool("use-audio", false, "use audio track for syncing")
	syncCmd.Flags().Bool("use-embedded", false, "use embedded subtitles for syncing")
	syncCmd.Flags().Int("audio-track", 0, "audio track index")
	syncCmd.Flags().IntSlice("sub-tracks", []int{0}, "embedded subtitle track indices")
	viper.BindPFlag("sync_use_audio", syncCmd.Flags().Lookup("use-audio"))
	viper.BindPFlag("sync_use_embedded", syncCmd.Flags().Lookup("use-embedded"))
	viper.BindPFlag("sync_audio_track", syncCmd.Flags().Lookup("audio-track"))
	viper.BindPFlag("sync_sub_tracks", syncCmd.Flags().Lookup("sub-tracks"))
	rootCmd.AddCommand(syncCmd)
}
