// file: cmd/whisper.go
package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/jdfalk/subtitle-manager/pkg/transcriber"
)

var whisperCmd = &cobra.Command{
	Use:   "whisper",
	Short: "Manage Whisper ASR container",
}

var whisperStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show Whisper container status",
	RunE: func(cmd *cobra.Command, args []string) error {
		wc, err := transcriber.NewWhisperContainer()
		if err != nil {
			return err
		}
		defer wc.Close()

		state, err := wc.GetContainerStatus(context.Background())
		if err != nil {
			return err
		}
		fmt.Println(state)
		return nil
	},
}

var whisperStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start Whisper container",
	RunE: func(cmd *cobra.Command, args []string) error {
		wc, err := transcriber.NewWhisperContainer()
		if err != nil {
			return err
		}
		defer wc.Close()

		if err := wc.StartContainer(context.Background()); err != nil {
			return err
		}
		fmt.Println("Whisper container started")
		return nil
	},
}

var whisperStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop Whisper container",
	RunE: func(cmd *cobra.Command, args []string) error {
		wc, err := transcriber.NewWhisperContainer()
		if err != nil {
			return err
		}
		defer wc.Close()

		if err := wc.StopContainer(context.Background()); err != nil {
			return err
		}
		fmt.Println("Whisper container stopped")
		return nil
	},
}

func init() {
	whisperCmd.AddCommand(whisperStatusCmd)
	whisperCmd.AddCommand(whisperStartCmd)
	whisperCmd.AddCommand(whisperStopCmd)
	rootCmd.AddCommand(whisperCmd)
}
