package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/jdfalk/subtitle-manager/pkg/transcriber"
)

// transcribeCmd generates subtitles from audio using the Whisper API.
var transcribeCmd = &cobra.Command{
	Use:   "transcribe [media] [output] [lang]",
	Short: "Transcribe media to subtitles",
	Args:  cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logging.GetLogger("transcribe")
		media, out, lang := args[0], args[1], args[2]
		key := viper.GetString("openai_api_key")
		data, err := transcriber.WhisperTranscribe(media, lang, key)
		if err != nil {
			return err
		}
		if err := os.WriteFile(out, data, 0644); err != nil {
			return err
		}
		logger.Infof("transcribed %s to %s", media, out)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(transcribeCmd)
}
