// Package cmd implements the CLI commands for subtitle-manager.
// It provides the root command and subcommands for all user-facing operations.
//
// This package is the entry point for the application's command-line interface.
// It allows users to convert subtitles to different formats, manage subtitle files,
// and perform other related tasks from the command line.

package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/jdfalk/subtitle-manager/pkg/subtitles"
)

var convertCmd = &cobra.Command{
	Use:   "convert [input] [output]",
	Short: "Convert subtitle to SRT",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logging.GetLogger("convert")
		in, out := args[0], args[1]
		data, err := subtitles.ConvertToSRT(in)
		if err != nil {
			return err
		}
		f, err := os.Create(out)
		if err != nil {
			return err
		}
		defer f.Close()
		if _, err := f.Write(data); err != nil {
			return err
		}
		logger.Infof("Converted %s to %s", in, out)
		return nil
	},
}
