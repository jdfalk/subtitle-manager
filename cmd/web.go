package cmd

import (
	"os"

	"github.com/jdfalk/subtitle-manager/pkg/webserver"
	"github.com/spf13/cobra"
)

var addr string

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Run web UI server",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Check for SM_HTTP_PORT environment variable if addr wasn't explicitly set
		if addr == ":8080" { // Default value, check if env var overrides it
			if envPort := os.Getenv("SM_HTTP_PORT"); envPort != "" {
				addr = ":" + envPort
			}
		}
		return webserver.StartServer(addr)
	},
}

func init() {
	webCmd.Flags().StringVar(&addr, "addr", ":8080", "listen address")
	rootCmd.AddCommand(webCmd)
	v, _, _ := GetVersionInfo()
	webserver.SetVersion(v)
}
