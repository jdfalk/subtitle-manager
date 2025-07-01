package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/jdfalk/subtitle-manager/pkg/queue"
	"github.com/jdfalk/subtitle-manager/pkg/subtitles"
	"github.com/jdfalk/subtitle-manager/pkg/tasks"
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
		async := viper.GetBool("async")

		if async {
			// Use asynchronous queue
			q := queue.GetQueue()
			if !q.IsRunning() {
				if err := q.Start(); err != nil {
					return fmt.Errorf("failed to start translation queue: %w", err)
				}
				defer q.Stop() // Clean up when command finishes
			}

			job := queue.NewSingleFileJob(in, out, lang, service, gKey, gptKey, grpcAddr)
			taskID, err := q.Add(job)
			if err != nil {
				return fmt.Errorf("failed to queue translation job: %w", err)
			}

			logger.Infof("Translation job queued with ID: %s", taskID)
			logger.Info("Processing translation asynchronously...")
			
			// Wait for the job to complete since this is a CLI command
			for {
				taskList := tasks.List()
				if task, exists := taskList[taskID]; exists {
					if task.Status == "completed" {
						logger.Infof("Translation completed successfully")
						break
					} else if task.Status == "failed" {
						return fmt.Errorf("translation failed: %s", task.Error)
					}
				}
				time.Sleep(100 * time.Millisecond)
			}
			
			// Record in database like sync version
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
		}

		// Synchronous execution (existing behavior)
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
	translateCmd.Flags().Bool("async", false, "queue translation for asynchronous processing")
	viper.BindPFlag("async", translateCmd.Flags().Lookup("async"))
}
