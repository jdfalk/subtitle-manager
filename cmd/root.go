package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"subtitle-manager/pkg/translator"
)

var cfgFile string
var dbPath string
var dbBackend string
var rootCmd = &cobra.Command{
	Use:   "subtitle-manager",
	Short: "Subtitle Manager CLI",
	Long:  `A simple subtitle management tool built in Go`,
}

// Execute executes the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatalf("%v", err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.subtitle-manager.yaml)")
	rootCmd.PersistentFlags().StringVar(&dbPath, "db", "", "database file (default is $HOME/.subtitle-manager.db)")
	rootCmd.PersistentFlags().StringVar(&dbBackend, "db-backend", "sqlite", "database backend: sqlite or pebble")
	rootCmd.PersistentFlags().String("log-level", "info", "log level (debug, info, warn, error)")
	viper.BindPFlag("log-level", rootCmd.PersistentFlags().Lookup("log-level"))
	rootCmd.PersistentFlags().StringToString("log-levels", nil, "per component log levels")
	viper.BindPFlag("log_levels", rootCmd.PersistentFlags().Lookup("log-levels"))
	viper.BindPFlag("db_path", rootCmd.PersistentFlags().Lookup("db"))
	viper.BindPFlag("db_backend", rootCmd.PersistentFlags().Lookup("db-backend"))
	rootCmd.PersistentFlags().String("google-key", "", "Google Translate API key")
	viper.BindPFlag("google_api_key", rootCmd.PersistentFlags().Lookup("google-key"))
	rootCmd.PersistentFlags().String("openai-key", "", "OpenAI API key")
	viper.BindPFlag("openai_api_key", rootCmd.PersistentFlags().Lookup("openai-key"))
	rootCmd.PersistentFlags().String("opensubtitles-key", "", "OpenSubtitles API key")
	viper.BindPFlag("opensubtitles.api_key", rootCmd.PersistentFlags().Lookup("opensubtitles-key"))
	rootCmd.PersistentFlags().String("ffmpeg-path", "", "path to ffmpeg binary")
	viper.BindPFlag("ffmpeg_path", rootCmd.PersistentFlags().Lookup("ffmpeg-path"))
	rootCmd.PersistentFlags().Int("batch-workers", 4, "number of workers for batch translate")
	viper.BindPFlag("batch_workers", rootCmd.PersistentFlags().Lookup("batch-workers"))
	rootCmd.PersistentFlags().Int("scan-workers", 4, "number of workers for scanning")
	viper.BindPFlag("scan_workers", rootCmd.PersistentFlags().Lookup("scan-workers"))
	rootCmd.PersistentFlags().String("google-api-url", "", "override Google Translate API URL")
	viper.BindPFlag("google_api_url", rootCmd.PersistentFlags().Lookup("google-api-url"))
	rootCmd.PersistentFlags().String("openai-model", "gpt-3.5-turbo", "ChatGPT model")
	viper.BindPFlag("openai_model", rootCmd.PersistentFlags().Lookup("openai-model"))
	rootCmd.PersistentFlags().String("opensubtitles-api-url", "", "OpenSubtitles API base URL")
	viper.BindPFlag("opensubtitles.api_url", rootCmd.PersistentFlags().Lookup("opensubtitles-api-url"))
	rootCmd.PersistentFlags().String("opensubtitles-user-agent", "", "OpenSubtitles user agent")
	viper.BindPFlag("opensubtitles.user_agent", rootCmd.PersistentFlags().Lookup("opensubtitles-user-agent"))

	rootCmd.AddCommand(convertCmd)
	rootCmd.AddCommand(mergeCmd)
	rootCmd.AddCommand(translateCmd)
	rootCmd.AddCommand(fetchCmd)
	rootCmd.AddCommand(batchCmd)
	rootCmd.AddCommand(sonarrCmd)
	rootCmd.AddCommand(radarrCmd)
}

func initConfig() {
	viper.SetEnvPrefix("SM")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	viper.AutomaticEnv()
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".subtitle-manager")
		viper.SetDefault("db_path", filepath.Join(home, ".subtitle-manager.db"))
		viper.SetDefault("db_backend", "sqlite")
		viper.SetDefault("translate_service", "google")
		viper.SetDefault("google_api_key", "")
		viper.SetDefault("openai_api_key", "")
		viper.SetDefault("opensubtitles.api_key", "")
		viper.SetDefault("ffmpeg_path", "ffmpeg")
		viper.SetDefault("batch_workers", 4)
		viper.SetDefault("scan_workers", 4)
		viper.SetDefault("google_api_url", "https://translation.googleapis.com/language/translate/v2")
		viper.SetDefault("openai_model", "gpt-3.5-turbo")
		viper.SetDefault("opensubtitles.api_url", "https://rest.opensubtitles.org")
		viper.SetDefault("opensubtitles.user_agent", "subtitle-manager/0.1")
		viper.SetDefault("providers.generic.api_url", "")
		viper.SetDefault("providers.generic.username", "")
		viper.SetDefault("providers.generic.password", "")
		viper.SetDefault("providers.generic.api_key", "")
		viper.SetDefault("plex.url", "http://localhost:32400")
		viper.SetDefault("plex.token", "")
	}

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	level, err := logrus.ParseLevel(viper.GetString("log-level"))
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)
	if u := viper.GetString("google_api_url"); u != "" {
		translator.SetGoogleAPIURL(u)
	}
	if m := viper.GetString("openai_model"); m != "" {
		translator.SetOpenAIModel(m)
	}
}
