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
		items, err := syncer.Sync(media, subPath, syncer.Options{})
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
	rootCmd.AddCommand(syncCmd)
}
