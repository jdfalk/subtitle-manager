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

func init() {
	whisperCmd.AddCommand(whisperStatusCmd)
	rootCmd.AddCommand(whisperCmd)
}
