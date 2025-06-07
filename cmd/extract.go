package cmd

import (
	"os"

	"github.com/asticode/go-astisub"
	"github.com/spf13/cobra"

	"subtitle-manager/pkg/logging"
	"subtitle-manager/pkg/subtitles"
)

var extractCmd = &cobra.Command{
	Use:   "extract [media] [output]",
	Short: "Extract subtitles from media",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logging.GetLogger("extract")
		media, out := args[0], args[1]
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
		logger.Infof("extracted subtitles from %s to %s", media, out)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(extractCmd)
}
