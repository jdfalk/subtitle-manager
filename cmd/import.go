// file: cmd/import.go
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"subtitle-manager/pkg/bazarr"
)

// importBazarrCmd imports configuration from a running Bazarr instance.
var importBazarrCmd = &cobra.Command{
	Use:   "import-bazarr [url] [api-key]",
	Short: "Import configuration from Bazarr",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		url, key := args[0], args[1]
		settings, err := bazarr.FetchSettings(url, key)
		if err != nil {
			return err
		}
		mapped := bazarr.MapSettings(settings)
		for k, v := range mapped {
			viper.Set(k, v)
		}
		if cfg := viper.ConfigFileUsed(); cfg != "" {
			if err := viper.WriteConfig(); err != nil {
				return err
			}
		}
		cmd.Println("Bazarr settings imported")
		return nil
	},
}

func init() { rootCmd.AddCommand(importBazarrCmd) }
