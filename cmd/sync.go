package cmd

import (
	"os"
	"time"

	"github.com/asticode/go-astisub"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/jdfalk/subtitle-manager/pkg/security"
	"github.com/jdfalk/subtitle-manager/pkg/syncer"
)

var (
	syncUseAudio         bool
	syncUseEmbedded      bool
	syncAudioTrack       int
	syncSubtitleTracks   []int
	syncAudioWeight      float64
	syncTranslate        bool
	syncTranslateLang    string
	syncTranslateService string
)

// syncCmd aligns an external subtitle file with a media file.
var syncCmd = &cobra.Command{
	Use:   "sync [media] [subtitle] [output]",
	Short: "Synchronize subtitle with media using audio transcription and/or embedded subtitles",
	Long: `Synchronize subtitle timing with media file using automatic analysis.

The sync command can use multiple reference sources for timing alignment:
- Audio transcription via Whisper API (--use-audio)
- Embedded subtitle tracks from the media file (--use-embedded)
- Combination of both with weighted averaging (--audio-weight)

Examples:
  # Sync using embedded subtitles only
  subtitle-manager sync movie.mkv subs.srt output.srt --use-embedded

  # Sync using audio transcription only
  subtitle-manager sync movie.mkv subs.srt output.srt --use-audio

  # Sync using both with 70% audio, 30% embedded weighting
  subtitle-manager sync movie.mkv subs.srt output.srt --use-audio --use-embedded --audio-weight 0.7

  # Sync with translation to Spanish
  subtitle-manager sync movie.mkv subs.srt output.srt --use-audio --translate --translate-lang es`,
	Args: cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logging.GetLogger("sync")
		media, err := security.SanitizePath(args[0])
		if err != nil {
			return err
		}
		subPath, err := security.SanitizePath(args[1])
		if err != nil {
			return err
		}
		out, err := security.SanitizePath(args[2])
		if err != nil {
			return err
		}

		// Log configuration details
		logger.Infof("starting subtitle synchronization")
		logger.Infof("media file: %s", media)
		logger.Infof("subtitle file: %s", subPath)
		logger.Infof("output file: %s", out)

		opts := syncer.Options{
			UseAudio:         syncUseAudio,
			UseEmbedded:      syncUseEmbedded,
			AudioTrack:       syncAudioTrack,
			SubtitleTracks:   syncSubtitleTracks,
			WhisperKey:       viper.GetString("openai_api_key"),
			AudioWeight:      syncAudioWeight,
			Translate:        syncTranslate,
			TranslateLang:    syncTranslateLang,
			TranslateService: syncTranslateService,
			GoogleAPIKey:     viper.GetString("google_api_key"),
			GPTAPIKey:        viper.GetString("openai_api_key"),
			GRPCAddr:         viper.GetString("grpc_addr"),
			TargetLang:       viper.GetString("sync_language"),
			Service:          viper.GetString("translate_service"),
			GoogleKey:        viper.GetString("google_api_key"),
			GPTKey:           viper.GetString("openai_api_key"),
		}

		// Log sync method configuration
		if opts.UseAudio && opts.UseEmbedded {
			logger.Infof("sync method: combined (audio + embedded, weight=%.1f)", opts.AudioWeight)
		} else if opts.UseAudio {
			logger.Infof("sync method: audio transcription only")
		} else if opts.UseEmbedded {
			logger.Infof("sync method: embedded subtitles only")
		}

		if opts.Translate {
			logger.Infof("translation enabled: %s via %s", opts.TranslateLang, opts.TranslateService)
		}

		// Default to embedded if no sync method specified
		if !opts.UseAudio && !opts.UseEmbedded {
			opts.UseEmbedded = true
			logger.Info("no sync method specified, defaulting to embedded subtitles")
		}

		start := time.Now()
		items, err := syncer.Sync(string(media), string(subPath), opts)
		if err != nil {
			logger.Errorf("synchronization failed: %v", err)
			return err
		}
		syncDuration := time.Since(start)

		logger.Infof("synchronization completed in %v", syncDuration)
		logger.Infof("processed %d subtitle items", len(items))

		tmpSub := astisub.Subtitles{Items: items}
		f, err := os.Create(string(out))
		if err != nil {
			logger.Errorf("failed to create output file: %v", err)
			return err
		}
		defer f.Close()
		if err := tmpSub.WriteToSRT(f); err != nil {
			logger.Errorf("failed to write SRT file: %v", err)
			return err
		}

		totalDuration := time.Since(start)

		// Print final summary
		logger.Infof("=== SYNC SUMMARY ===")
		if opts.Translate {
			logger.Infof("✅ synchronized and translated %s -> %s (%s)", subPath, out, syncTranslateLang)
		} else {
			logger.Infof("✅ synchronized %s -> %s", subPath, out)
		}
		logger.Infof("subtitle items: %d", len(items))
		logger.Infof("sync duration: %v", syncDuration)
		logger.Infof("total duration: %v", totalDuration)
		return nil
	},
}

func init() {
	// Audio transcription options
	syncCmd.Flags().BoolVar(&syncUseAudio, "use-audio", false, "Use audio transcription for sync reference")
	syncCmd.Flags().IntVar(&syncAudioTrack, "audio-track", 0, "Audio track index to use for transcription (default: 0)")

	// Embedded subtitle options
	syncCmd.Flags().BoolVar(&syncUseEmbedded, "use-embedded", false, "Use embedded subtitles for sync reference")
	syncCmd.Flags().IntSliceVar(&syncSubtitleTracks, "subtitle-tracks", []int{0}, "Embedded subtitle track indices to use (default: [0])")

	// Weighting and advanced options
	syncCmd.Flags().Float64Var(&syncAudioWeight, "audio-weight", 0.7, "Weight for audio transcription vs embedded subtitles (0.0-1.0)")

	// Translation integration
	syncCmd.Flags().BoolVar(&syncTranslate, "translate", false, "Translate subtitles after synchronization")
	syncCmd.Flags().StringVar(&syncTranslateLang, "translate-lang", "", "Target language for translation (e.g., 'es', 'fr', 'de')")
	syncCmd.Flags().StringVar(&syncTranslateService, "translate-service", "google", "Translation service to use ('google', 'gpt', 'grpc')")

	rootCmd.AddCommand(syncCmd)
	syncCmd.Flags().String("lang", "", "translate subtitles to language before syncing")
	viper.BindPFlag("sync_language", syncCmd.Flags().Lookup("lang"))
	syncCmd.Flags().String("service", "", "translation service")
	viper.BindPFlag("translate_service", syncCmd.Flags().Lookup("service"))
	syncCmd.Flags().String("grpc", "", "use remote gRPC translator")
	viper.BindPFlag("grpc_addr", syncCmd.Flags().Lookup("grpc"))
}
