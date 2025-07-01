package cmd

import (
	"path/filepath"
	"strings"

	"github.com/sourcegraph/conc/pool"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jdfalk/subtitle-manager/pkg/cli"
	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/jdfalk/subtitle-manager/pkg/subtitles"
)

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
		workers := viper.GetInt("batch_workers")

		// Create progress bar
		progress := cli.NewProgressBar(len(files), "Translating")
		defer progress.Finish()

		// Use pool to translate files with progress tracking
		p := pool.New().WithErrors().WithMaxGoroutines(workers)
		for _, in := range files {
			in := in
			out := strings.TrimSuffix(in, filepath.Ext(in)) + "." + lang + ".srt"
			p.Go(func() error {
				err := subtitles.TranslateFileToSRT(in, out, lang, service, gKey, gptKey, grpcAddr)
				if err == nil {
					progress.Update(in)
				}
				return err
			})
		}

		if err := p.Wait(); err != nil {
			return err
		}
		logger.Infof("translated %d files", len(files))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(batchCmd)
}
