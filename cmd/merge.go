package cmd

import (
	"os"

	"github.com/asticode/go-astisub"
	"github.com/spf13/cobra"

	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/jdfalk/subtitle-manager/pkg/security"
	"github.com/jdfalk/subtitle-manager/pkg/subtitles"
)

var mergeCmd = &cobra.Command{
	Use:   "merge [sub1] [sub2] [output]",
	Short: "Merge two subtitles into one",
	Args:  cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logging.GetLogger("merge")
		sub1Path, err := security.SanitizePath(args[0])
		if err != nil {
			return err
		}
		sub1, err := astisub.OpenFile(string(sub1Path))
		if err != nil {
			return err
		}
		sub2Path, err := security.SanitizePath(args[1])
		if err != nil {
			return err
		}
		sub2, err := astisub.OpenFile(string(sub2Path))
		if err != nil {
			return err
		}
		sub1.Items = subtitles.MergeTracks(sub1.Items, sub2.Items)
		outPath, err := security.SanitizePath(args[2])
		if err != nil {
			return err
		}
		f, err := os.Create(string(outPath))
		if err != nil {
			return err
		}
		defer f.Close()
		if err := sub1.WriteToSRT(f); err != nil {
			return err
		}
		logger.Infof("Merged %s and %s into %s", args[0], args[1], args[2])
		return nil
	},
}
