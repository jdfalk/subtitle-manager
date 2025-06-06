package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"subtitle-manager/pkg/logging"
	"subtitle-manager/pkg/subtitles"
)

var workers int

var batchCmd = &cobra.Command{
	Use:   "batch [lang] [files...]",
	Short: "Translate multiple subtitle files concurrently",
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logging.GetLogger("batch")
		lang := args[0]
		files := args[1:]
		service := viper.GetString("translate_service")
		gKey := viper.GetString("google_api_key")
		gptKey := viper.GetString("openai_api_key")
		grpcAddr := viper.GetString("grpc_addr")
		if err := subtitles.TranslateFilesToSRT(files, lang, service, gKey, gptKey, grpcAddr, workers); err != nil {
			return err
		}
		logger.Infof("translated %d files", len(files))
		return nil
	},
}

func init() {
	batchCmd.Flags().IntVar(&workers, "workers", 4, "number of concurrent workers")
	rootCmd.AddCommand(batchCmd)
}
