package cmd

import (
	"os"

	"github.com/asticode/go-astisub"
	"github.com/spf13/cobra"

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
		useEmb, _ := cmd.Flags().GetBool("use-embedded")
		tracks, _ := cmd.Flags().GetIntSlice("tracks")
		useAudio, _ := cmd.Flags().GetBool("use-audio")
		audioTrack, _ := cmd.Flags().GetInt("audio-track")
		opts := syncer.Options{UseEmbedded: useEmb, SubtitleTracks: tracks, UseAudio: useAudio, AudioTrack: audioTrack}
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
	syncCmd.Flags().Bool("use-embedded", false, "use embedded subtitle tracks")
	syncCmd.Flags().IntSlice("tracks", nil, "embedded subtitle tracks to analyse")
	syncCmd.Flags().Bool("use-audio", false, "use audio track (not implemented)")
	syncCmd.Flags().Int("audio-track", 0, "audio track index")
	rootCmd.AddCommand(syncCmd)
}
