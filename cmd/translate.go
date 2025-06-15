package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/jdfalk/subtitle-manager/pkg/subtitles"
)

var translateCmd = &cobra.Command{
	Use:   "translate [input] [output] [lang]",
	Short: "Translate subtitle",
	Args:  cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logging.GetLogger("translate")
		in, out, lang := args[0], args[1], args[2]
		service := viper.GetString("translate_service")
		gKey := viper.GetString("google_api_key")
		gptKey := viper.GetString("openai_api_key")
		grpcAddr := viper.GetString("grpc_addr")
		if err := subtitles.TranslateFileToSRT(in, out, lang, service, gKey, gptKey, grpcAddr); err != nil {
			return err
		}
		if dbPath := viper.GetString("db_path"); dbPath != "" {
			backend := viper.GetString("db_backend")
			if store, err := database.OpenStore(dbPath, backend); err == nil {
				_ = store.InsertSubtitle(&database.SubtitleRecord{File: in, Language: lang, Service: service})
				store.Close()
			} else {
				logger.Warnf("db open: %v", err)
			}
		}
		logger.Infof("Translated %s to %s in %s", in, lang, out)
		return nil
	},
}

func init() {
	translateCmd.Flags().String("service", "google", "translation service: google, gpt or grpc")
	viper.BindPFlag("translate_service", translateCmd.Flags().Lookup("service"))
	translateCmd.Flags().String("grpc", "", "use remote gRPC translator at host:port")
	viper.BindPFlag("grpc_addr", translateCmd.Flags().Lookup("grpc"))
}
