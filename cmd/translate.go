package cmd

import (
	"os"

	"github.com/asticode/go-astisub"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"subtitle-manager/pkg/database"
	"subtitle-manager/pkg/logging"
	"subtitle-manager/pkg/translator"
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
		sub, err := astisub.OpenFile(in)
		if err != nil {
			return err
		}
		for _, item := range sub.Items {
			text := item.String()
			t, err := translator.Translate(service, text, lang, gKey, gptKey, grpcAddr)
			if err != nil {
				return err
			}
			item.Lines = []astisub.Line{{Items: []astisub.LineItem{{Text: t}}}}
		}
		f, err := os.Create(out)
		if err != nil {
			return err
		}
		defer f.Close()
		if err := sub.WriteToSRT(f); err != nil {
			return err
		}
		if dbPath := viper.GetString("db_path"); dbPath != "" {
			if db, err := database.Open(dbPath); err == nil {
				_ = database.InsertSubtitle(db, in, lang, service)
				db.Close()
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
	translateCmd.Flags().String("google-key", "", "Google Translate API key")
	viper.BindPFlag("google_api_key", translateCmd.Flags().Lookup("google-key"))
	translateCmd.Flags().String("openai-key", "", "OpenAI API key")
	viper.BindPFlag("openai_api_key", translateCmd.Flags().Lookup("openai-key"))
	translateCmd.Flags().String("grpc", "", "use remote gRPC translator at host:port")
	viper.BindPFlag("grpc_addr", translateCmd.Flags().Lookup("grpc"))
}
