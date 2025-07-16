// file: pkg/gcommon/config/config.go
// version: 1.0.0
// guid: 9f89692c-72ac-4bc6-b7be-7ff20bbf12e3

package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Load initializes Viper using the given Cobra command and optional config file path.
// Environment variables are automatically loaded using the SM_ prefix.
// Default values mirror the previous internal loader for compatibility.
func Load(cmd *cobra.Command, cfgFile string) error {
	viper.SetEnvPrefix("SM")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	viper.AutomaticEnv()

	configFile := cfgFile
	if configFile == "" {
		configFile = viper.GetString("config_file")
	}
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".subtitle-manager")
	}

	home, err := os.UserHomeDir()
	if err == nil {
		viper.SetDefault("db_path", filepath.Join(home, ".subtitle-manager", "db"))
	} else {
		viper.SetDefault("db_path", "/config/db")
	}
	viper.SetDefault("db_backend", "pebble")
	viper.SetDefault("sqlite3_filename", "subtitle-manager.db")
	viper.SetDefault("log_file", "/config/logs/subtitle-manager.log")

	if err := viper.ReadInConfig(); err == nil {
		cmd.Printf("Using config file: %s\n", viper.ConfigFileUsed())
	}
	return nil
}
