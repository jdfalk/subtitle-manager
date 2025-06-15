package cmd

import (
	"os"

	"github.com/asticode/go-astisub"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/jdfalk/subtitle-manager/pkg/syncer"
)

// syncCmd aligns an external subtitle file with a media file.
var syncCmd = &cobra.Command{
	Use:   "sync [media] [subtitle] [output]",
	Short: "Synchronize subtitle with media",
	Args:  cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logging.GetLogger("sync")
		media, subPath, out := args[0], args[1], args[2]
		opts := syncer.Options{
			TargetLang: viper.GetString("sync_language"),
			Service:    viper.GetString("translate_service"),
			GoogleKey:  viper.GetString("google_api_key"),
			GPTKey:     viper.GetString("openai_api_key"),
			GRPCAddr:   viper.GetString("grpc_addr"),
		}
		items, err := syncer.Sync(media, subPath, opts)
		if err != nil {
			return err
		}
		tmpSub := astisub.Subtitles{Items: items}
		f, err := os.Create(out)
		if err != nil {
			return err
		}
		defer f.Close()
		if err := tmpSub.WriteToSRT(f); err != nil {
			return err
		}
		logger.Infof("synchronized %s -> %s", subPath, out)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
	syncCmd.Flags().String("lang", "", "translate subtitles to language before syncing")
	viper.BindPFlag("sync_language", syncCmd.Flags().Lookup("lang"))
	syncCmd.Flags().String("service", "", "translation service")
	viper.BindPFlag("translate_service", syncCmd.Flags().Lookup("service"))
	syncCmd.Flags().String("grpc", "", "use remote gRPC translator")
	viper.BindPFlag("grpc_addr", syncCmd.Flags().Lookup("grpc"))
}
