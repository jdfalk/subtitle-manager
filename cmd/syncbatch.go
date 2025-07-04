package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/jdfalk/subtitle-manager/pkg/security"
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
		cfgPath, err := security.SanitizePath(syncBatchConfig)
		if err != nil {
			return err
		}
		f, err := os.Open(string(cfgPath))
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
		logger.Infof("starting batch synchronization of %d items", len(req.Items))
		start := time.Now()

		errs := syncer.SyncBatch(req.Items, req.Options)

		var successCount, failureCount int
		var failedItems []string

		for i, it := range req.Items {
			if err := errs[i]; err != nil {
				logger.Warnf("sync failed for %s: %v", it.Subtitle, err)
				failureCount++
				failedItems = append(failedItems, it.Subtitle)
			} else {
				outPath := it.Output
				if outPath == "" {
					outPath = it.Subtitle
				}
				logger.Infof("✅ synchronized %s -> %s", it.Subtitle, outPath)
				successCount++
			}
		}

		duration := time.Since(start)

		// Print detailed summary
		logger.Infof("=== BATCH SYNC SUMMARY ===")
		logger.Infof("Total items: %d", len(req.Items))
		logger.Infof("Successful: %d", successCount)
		logger.Infof("Failed: %d", failureCount)
		logger.Infof("Duration: %v", duration)

		if failureCount > 0 {
			logger.Warnf("Failed items:")
			for _, item := range failedItems {
				logger.Warnf("  - %s", item)
			}
			return fmt.Errorf("batch synchronization completed with %d failures out of %d items", failureCount, len(req.Items))
		}

		logger.Infof("🎉 All %d items synchronized successfully in %v", successCount, duration)
		return nil
	},
}

func init() {
	syncBatchCmd.Flags().StringVarP(&syncBatchConfig, "config", "c", "", "JSON configuration file")
	_ = syncBatchCmd.MarkFlagRequired("config")
	rootCmd.AddCommand(syncBatchCmd)
}
