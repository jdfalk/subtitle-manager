// file: pkg/gcommon/config/config.go
// version: 1.0.0
// guid: 9f89692c-72ac-4bc6-b7be-7ff20bbf12e3

package config

import (
	"os"
	"path/filepath"
	"strings"

	configpb "github.com/jdfalk/subtitle-manager/pkg/configpb"

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

// ToProto converts current viper settings to a protobuf config message.
func ToProto() *configpb.SubtitleManagerConfig {
	return &configpb.SubtitleManagerConfig{
		DbPath:          protoString(viper.GetString("db_path")),
		DbBackend:       protoString(viper.GetString("db_backend")),
		Sqlite3Filename: protoString(viper.GetString("sqlite3_filename")),
		LogFile:         protoString(viper.GetString("log_file")),
		GoogleApiKey:    protoString(viper.GetString("google_api_key")),
		OpenaiApiKey:    protoString(viper.GetString("openai_api_key")),
	}
}

// ApplyProto sets viper values from a protobuf config message.
func ApplyProto(cfg *configpb.SubtitleManagerConfig) {
	if cfg == nil {
		return
	}
	if cfg.DbPath != nil {
		viper.Set("db_path", cfg.GetDbPath())
	}
	if cfg.DbBackend != nil {
		viper.Set("db_backend", cfg.GetDbBackend())
	}
	if cfg.Sqlite3Filename != nil {
		viper.Set("sqlite3_filename", cfg.GetSqlite3Filename())
	}
	if cfg.LogFile != nil {
		viper.Set("log_file", cfg.GetLogFile())
	}
	if cfg.GoogleApiKey != nil {
		viper.Set("google_api_key", cfg.GetGoogleApiKey())
	}
	if cfg.OpenaiApiKey != nil {
		viper.Set("openai_api_key", cfg.GetOpenaiApiKey())
	}
}

func protoString(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
