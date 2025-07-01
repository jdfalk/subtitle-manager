// file: cmd/queue.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174004
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/jdfalk/subtitle-manager/pkg/queue"
	"github.com/jdfalk/subtitle-manager/pkg/tasks"
)

var queueCmd = &cobra.Command{
	Use:   "queue",
	Short: "Manage translation queue",
	Long:  "Commands to manage the asynchronous translation queue",
}

var queueStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show queue status",
	RunE: func(cmd *cobra.Command, args []string) error {
		q := queue.GetQueue()
		status := q.Status()

		output, _ := json.MarshalIndent(status, "", "  ")
		fmt.Println(string(output))

		// Also show active tasks
		taskList := tasks.List()
		if len(taskList) > 0 {
			fmt.Println("\nActive Tasks:")
			for id, task := range taskList {
				fmt.Printf("  %s: %s (%d%%)\n", id, task.Status, task.Progress)
				if task.Error != "" {
					fmt.Printf("    Error: %s\n", task.Error)
				}
			}
		} else {
			fmt.Println("\nNo active tasks")
		}

		fmt.Println("\nNote: Queue workers are automatically managed within long-running services.")
		fmt.Println("Use 'subtitle-manager translate --async' to queue translation jobs.")

		return nil
	},
}

var queueStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the translation queue (for testing)",
	Long:  "Start the translation queue in this process. Note: This is mainly for testing. In production, queues are automatically managed by long-running services.",
	RunE: func(cmd *cobra.Command, args []string) error {
		q := queue.GetQueue()
		if err := q.Start(); err != nil {
			return fmt.Errorf("failed to start queue: %w", err)
		}
		fmt.Println("Translation queue started (will stop when command exits)")

		// Keep the process running to demonstrate the queue
		fmt.Println("Press Ctrl+C to stop...")
		select {} // Block forever until interrupt
	},
}

var queueStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the translation queue (for testing)",
	Long:  "Stop the translation queue in this process. Note: This only works if the queue was started in the same process.",
	RunE: func(cmd *cobra.Command, args []string) error {
		q := queue.GetQueue()
		if err := q.Stop(); err != nil {
			return fmt.Errorf("failed to stop queue: %w", err)
		}
		fmt.Println("Translation queue stopped")
		return nil
	},
}

func init() {
	queueCmd.AddCommand(queueStatusCmd)
	queueCmd.AddCommand(queueStartCmd)
	queueCmd.AddCommand(queueStopCmd)
}
