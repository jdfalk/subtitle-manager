package cmd

import (
	"encoding/json"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/jdfalk/subtitle-manager/pkg/syncer"
)

var syncBatchConfig string

// syncBatchCmd synchronizes multiple subtitle files using a JSON configuration.
// The configuration file should contain an array of objects with media, subtitle
// and output fields along with optional sync options.
var syncBatchCmd = &cobra.Command{
	Use:   "syncbatch",
	Short: "Synchronize multiple subtitles via configuration file",
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logging.GetLogger("syncbatch")
		f, err := os.Open(syncBatchConfig)
		if err != nil {
			return err
		}
		defer f.Close()
		var req struct {
			Items   []syncer.BatchItem `json:"items"`
			Options syncer.Options     `json:"options"`
		}
		if err := json.NewDecoder(f).Decode(&req); err != nil {
			return err
		}
		// Fill API keys from config when not provided
		if req.Options.WhisperKey == "" {
			req.Options.WhisperKey = viper.GetString("openai_api_key")
		}
		if req.Options.GoogleAPIKey == "" {
			req.Options.GoogleAPIKey = viper.GetString("google_api_key")
		}
		if req.Options.GPTAPIKey == "" {
			req.Options.GPTAPIKey = viper.GetString("openai_api_key")
		}
		if req.Options.GRPCAddr == "" {
			req.Options.GRPCAddr = viper.GetString("grpc_addr")
		}
		errs := syncer.SyncBatch(req.Items, req.Options)
		var hasErrors bool
		for i, it := range req.Items {
			if err := errs[i]; err != nil {
				logger.Warnf("sync %s: %v", it.Subtitle, err)
				hasErrors = true
			} else {
				logger.Infof("synchronized %s -> %s", it.Subtitle, it.Output)
			}
		}
		if hasErrors {
			return fmt.Errorf("some items failed during synchronization")
		}
		return nil
	},
}

func init() {
	syncBatchCmd.Flags().StringVarP(&syncBatchConfig, "config", "c", "", "JSON configuration file")
	syncBatchCmd.MarkFlagRequired("config")
	rootCmd.AddCommand(syncBatchCmd)
}
