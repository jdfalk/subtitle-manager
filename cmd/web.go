package cmd

import (
	"github.com/spf13/cobra"
	"subtitle-manager/pkg/webserver"
)

var addr string

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Run web UI server",
	RunE: func(cmd *cobra.Command, args []string) error {
		return webserver.StartServer(addr)
	},
}

func init() {
	webCmd.Flags().StringVar(&addr, "addr", ":8080", "listen address")
	rootCmd.AddCommand(webCmd)
}
