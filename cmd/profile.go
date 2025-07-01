package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"runtime/pprof"
	"time"

	"github.com/asticode/go-astisub"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/jdfalk/subtitle-manager/pkg/subtitles"
)

var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Performance profiling and benchmarking tools",
}

var profileTranslateCmd = &cobra.Command{
	Use:   "translate [input] [lang]",
	Short: "Profile translation performance",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logging.GetLogger("profile-translate")
		in, lang := args[0], args[1]
		service := viper.GetString("translate_service")
		gKey := viper.GetString("google_api_key")
		gptKey := viper.GetString("openai_api_key")
		grpcAddr := viper.GetString("grpc_addr")
		
		// Get profiling options
		cpuProfile := viper.GetString("cpu_profile")
		memProfile := viper.GetString("mem_profile")
		iterations := viper.GetInt("iterations")

		logger.Infof("Profiling translation of %s to %s (%d iterations)", in, lang, iterations)

		// Start CPU profiling if requested
		if cpuProfile != "" {
			f, err := os.Create(cpuProfile)
			if err != nil {
				return fmt.Errorf("could not create CPU profile: %w", err)
			}
			defer f.Close()
			if err := pprof.StartCPUProfile(f); err != nil {
				return fmt.Errorf("could not start CPU profile: %w", err)
			}
			defer pprof.StopCPUProfile()
		}

		// Run translation benchmark
		start := time.Now()
		for i := 0; i < iterations; i++ {
			out := filepath.Join(os.TempDir(), fmt.Sprintf("profile_out_%d.srt", i))
			if err := subtitles.TranslateFileToSRT(in, out, lang, service, gKey, gptKey, grpcAddr); err != nil {
				return fmt.Errorf("translation failed: %w", err)
			}
			// Clean up temp file
			os.Remove(out)
		}
		elapsed := time.Since(start)

		// Write memory profile if requested
		if memProfile != "" {
			f, err := os.Create(memProfile)
			if err != nil {
				return fmt.Errorf("could not create memory profile: %w", err)
			}
			defer f.Close()
			runtime.GC() // get up-to-date statistics
			if err := pprof.WriteHeapProfile(f); err != nil {
				return fmt.Errorf("could not write memory profile: %w", err)
			}
		}

		// Report results
		avgTime := elapsed / time.Duration(iterations)
		logger.Infof("Translation benchmark results:")
		logger.Infof("  Total time: %v", elapsed)
		logger.Infof("  Average per translation: %v", avgTime)
		logger.Infof("  Iterations: %d", iterations)
		if cpuProfile != "" {
			logger.Infof("  CPU profile written to: %s", cpuProfile)
		}
		if memProfile != "" {
			logger.Infof("  Memory profile written to: %s", memProfile)
		}

		return nil
	},
}

var profileMergeCmd = &cobra.Command{
	Use:   "merge [sub1] [sub2]",
	Short: "Profile merge performance",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logging.GetLogger("profile-merge")
		sub1Path, sub2Path := args[0], args[1]
		
		// Get profiling options
		cpuProfile := viper.GetString("cpu_profile")
		memProfile := viper.GetString("mem_profile")
		iterations := viper.GetInt("iterations")

		logger.Infof("Profiling merge of %s and %s (%d iterations)", sub1Path, sub2Path, iterations)

		// Load subtitles once
		sub1, err := astisub.OpenFile(sub1Path)
		if err != nil {
			return fmt.Errorf("failed to open %s: %w", sub1Path, err)
		}
		sub2, err := astisub.OpenFile(sub2Path)
		if err != nil {
			return fmt.Errorf("failed to open %s: %w", sub2Path, err)
		}

		// Start CPU profiling if requested
		if cpuProfile != "" {
			f, err := os.Create(cpuProfile)
			if err != nil {
				return fmt.Errorf("could not create CPU profile: %w", err)
			}
			defer f.Close()
			if err := pprof.StartCPUProfile(f); err != nil {
				return fmt.Errorf("could not start CPU profile: %w", err)
			}
			defer pprof.StopCPUProfile()
		}

		// Run merge benchmark
		start := time.Now()
		for i := 0; i < iterations; i++ {
			_ = subtitles.MergeTracks(sub1.Items, sub2.Items)
		}
		elapsed := time.Since(start)

		// Write memory profile if requested
		if memProfile != "" {
			f, err := os.Create(memProfile)
			if err != nil {
				return fmt.Errorf("could not create memory profile: %w", err)
			}
			defer f.Close()
			runtime.GC() // get up-to-date statistics
			if err := pprof.WriteHeapProfile(f); err != nil {
				return fmt.Errorf("could not write memory profile: %w", err)
			}
		}

		// Report results
		avgTime := elapsed / time.Duration(iterations)
		logger.Infof("Merge benchmark results:")
		logger.Infof("  Total time: %v", elapsed)
		logger.Infof("  Average per merge: %v", avgTime)
		logger.Infof("  Iterations: %d", iterations)
		logger.Infof("  Sub1 items: %d", len(sub1.Items))
		logger.Infof("  Sub2 items: %d", len(sub2.Items))
		if cpuProfile != "" {
			logger.Infof("  CPU profile written to: %s", cpuProfile)
		}
		if memProfile != "" {
			logger.Infof("  Memory profile written to: %s", memProfile)
		}

		return nil
	},
}

func init() {
	// Add profile subcommands
	profileCmd.AddCommand(profileTranslateCmd)
	profileCmd.AddCommand(profileMergeCmd)
	
	// Add profiling flags
	profileCmd.PersistentFlags().String("cpu-profile", "", "write CPU profile to file")
	viper.BindPFlag("cpu_profile", profileCmd.PersistentFlags().Lookup("cpu-profile"))
	
	profileCmd.PersistentFlags().String("mem-profile", "", "write memory profile to file")
	viper.BindPFlag("mem_profile", profileCmd.PersistentFlags().Lookup("mem-profile"))
	
	profileCmd.PersistentFlags().Int("iterations", 10, "number of iterations for benchmarking")
	viper.BindPFlag("iterations", profileCmd.PersistentFlags().Lookup("iterations"))

	// Add translate service flags to profile translate
	profileTranslateCmd.Flags().String("service", "google", "translation service: google, gpt or grpc")
	viper.BindPFlag("translate_service", profileTranslateCmd.Flags().Lookup("service"))
	profileTranslateCmd.Flags().String("grpc", "", "use remote gRPC translator at host:port")
	viper.BindPFlag("grpc_addr", profileTranslateCmd.Flags().Lookup("grpc"))
}