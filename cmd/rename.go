package cmd

import (
	"github.com/spf13/cobra"

	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/jdfalk/subtitle-manager/pkg/renamer"
)

// renameCmd renames subtitle files to match a video filename.
var renameCmd = &cobra.Command{
	Use:   "rename [video] [lang]",
	Short: "Rename subtitle to match video file",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logging.GetLogger("rename")
		video, lang := args[0], args[1]
		if err := renamer.Rename(video, lang); err != nil {
			logger.Errorf("rename %s: %v", video, err)
			return err
		}
		logger.Infof("renamed subtitles for %s", video)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(renameCmd)
}
