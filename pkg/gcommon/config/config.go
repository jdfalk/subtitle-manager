// file: pkg/gcommon/config/config.go
// version: 2.0.0
// guid: 9f89692c-72ac-4bc6-b7be-7ff20bbf12e3

package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/jdfalk/gcommon/sdks/go/v1/common"

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

// ToProto converts current viper settings to gcommon config requests.
// Returns a map of key-value pairs that can be used with SetConfigRequest.
func ToProto() map[string]*common.ConfigValue {
	configMap := make(map[string]*common.ConfigValue)
	
	// Create config values for each setting
	if dbPath := viper.GetString("db_path"); dbPath != "" {
		configMap["db_path"] = &common.ConfigValue{}
		configMap["db_path"].SetStringValue(dbPath)
	}
	
	if dbBackend := viper.GetString("db_backend"); dbBackend != "" {
		configMap["db_backend"] = &common.ConfigValue{}
		configMap["db_backend"].SetStringValue(dbBackend)
	}
	
	if logFile := viper.GetString("log_file"); logFile != "" {
		configMap["log_file"] = &common.ConfigValue{}
		configMap["log_file"].SetStringValue(logFile)
	}
	
	if googleKey := viper.GetString("google_api_key"); googleKey != "" {
		configMap["google_api_key"] = &common.ConfigValue{}
		configMap["google_api_key"].SetStringValue(googleKey)
	}
	
	if openaiKey := viper.GetString("openai_api_key"); openaiKey != "" {
		configMap["openai_api_key"] = &common.ConfigValue{}
		configMap["openai_api_key"].SetStringValue(openaiKey)
	}
	
	return configMap
}

// ApplyProto sets viper values from gcommon config values.
func ApplyProto(configMap map[string]*common.ConfigValue) {
	if configMap == nil {
		return
	}
	
	if dbPath, ok := configMap["db_path"]; ok && dbPath.GetStringValue() != "" {
		viper.Set("db_path", dbPath.GetStringValue())
	}
	
	if dbBackend, ok := configMap["db_backend"]; ok && dbBackend.GetStringValue() != "" {
		viper.Set("db_backend", dbBackend.GetStringValue())
	}
	
	if logFile, ok := configMap["log_file"]; ok && logFile.GetStringValue() != "" {
		viper.Set("log_file", logFile.GetStringValue())
	}
	
	if googleKey, ok := configMap["google_api_key"]; ok && googleKey.GetStringValue() != "" {
		viper.Set("google_api_key", googleKey.GetStringValue())
	}
	
	if openaiKey, ok := configMap["openai_api_key"]; ok && openaiKey.GetStringValue() != "" {
		viper.Set("openai_api_key", openaiKey.GetStringValue())
	}
}


