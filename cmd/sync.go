package cmd

import (
	"os"

	"github.com/asticode/go-astisub"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jdfalk/subtitle-manager/pkg/logging"
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
		media, subPath, out := args[0], args[1], args[2]

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

		// Default to embedded if no sync method specified
		if !opts.UseAudio && !opts.UseEmbedded {
			opts.UseEmbedded = true
			logger.Info("no sync method specified, defaulting to embedded subtitles")
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

		if opts.Translate {
			logger.Infof("synchronized and translated %s -> %s (%s)", subPath, out, syncTranslateLang)
		} else {
			logger.Infof("synchronized %s -> %s", subPath, out)
		}
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
