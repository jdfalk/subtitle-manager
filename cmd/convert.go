package cmd

import (
	"os"

	"github.com/asticode/go-astisub"
	"github.com/spf13/cobra"

	"subtitle-manager/pkg/logging"
)

var convertCmd = &cobra.Command{
	Use:   "convert [input] [output]",
	Short: "Convert subtitle to SRT",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logging.GetLogger("convert")
		in, out := args[0], args[1]
		sub, err := astisub.OpenFile(in)
		if err != nil {
			return err
		}
		f, err := os.Create(out)
		if err != nil {
			return err
		}
		defer f.Close()
		if err := sub.WriteToSRT(f); err != nil {
			return err
		}
		logger.Infof("Converted %s to %s", in, out)
		return nil
	},
}
