package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jdfalk/subtitle-manager/pkg/captcha"
	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/translator"
)

var cfgFile string
var dbPath string
var dbBackend string
var sqliteFilename string
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

// GetDatabasePath returns the full database path using the database package helper
func GetDatabasePath() string {
	return database.GetDatabasePath()
}

// GetDatabaseBackend returns the configured database backend using the database package helper
func GetDatabaseBackend() string {
	return database.GetDatabaseBackend()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.subtitle-manager.yaml)")
	rootCmd.PersistentFlags().StringVar(&dbPath, "db-path", "", "database path (default is $HOME/.subtitle-manager/db for pebble, $HOME for sqlite)")
	rootCmd.PersistentFlags().StringVar(&dbBackend, "db-backend", "pebble", "database backend: sqlite, pebble, or postgres")
	rootCmd.PersistentFlags().StringVar(&sqliteFilename, "sqlite3-filename", "subtitle-manager.db", "SQLite database filename (only used when db-backend=sqlite)")
	rootCmd.PersistentFlags().String("log-level", "info", "log level (debug, info, warn, error)")
	viper.BindPFlag("log-level", rootCmd.PersistentFlags().Lookup("log-level"))
	rootCmd.PersistentFlags().StringToString("log-levels", nil, "per component log levels")
	viper.BindPFlag("log_levels", rootCmd.PersistentFlags().Lookup("log-levels"))
	viper.BindPFlag("db_path", rootCmd.PersistentFlags().Lookup("db-path"))
	viper.BindPFlag("db_backend", rootCmd.PersistentFlags().Lookup("db-backend"))
	viper.BindPFlag("sqlite3_filename", rootCmd.PersistentFlags().Lookup("sqlite3-filename"))
	rootCmd.PersistentFlags().String("admin-user", "", "admin username for automatic setup (Docker/env only)")
	viper.BindPFlag("admin_user", rootCmd.PersistentFlags().Lookup("admin-user"))
	rootCmd.PersistentFlags().String("admin-password", "", "admin password for automatic setup (Docker/env only)")
	viper.BindPFlag("admin_password", rootCmd.PersistentFlags().Lookup("admin-password"))
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

	rootCmd.PersistentFlags().String("anticaptcha-key", "", "Anti-Captcha API key")
	viper.BindPFlag("anticaptcha.api_key", rootCmd.PersistentFlags().Lookup("anticaptcha-key"))
	rootCmd.PersistentFlags().String("anticaptcha-api-url", "", "Anti-Captcha API base URL")
	viper.BindPFlag("anticaptcha.api_url", rootCmd.PersistentFlags().Lookup("anticaptcha-api-url"))

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

	// Check for config file from flag, environment variable, or default
	configFile := cfgFile
	if configFile == "" {
		configFile = viper.GetString("config_file") // This will get SM_CONFIG_FILE
	}

	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".subtitle-manager")
	}

	// Set defaults (needs to happen regardless of config file source)
	home, err := os.UserHomeDir()
	if err == nil {
		viper.SetDefault("db_path", filepath.Join(home, ".subtitle-manager", "db"))
	} else {
		viper.SetDefault("db_path", "/config/db")
	}
	viper.SetDefault("db_backend", "pebble")
	viper.SetDefault("sqlite3_filename", "subtitle-manager.db")
	viper.SetDefault("admin_user", "")
	viper.SetDefault("admin_password", "")
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
	viper.SetDefault("opensubtitles.user_agent", "github.com/jdfalk/subtitle-manager/0.1")
	viper.SetDefault("anticaptcha.api_key", "")
	viper.SetDefault("anticaptcha.api_url", "https://api.anti-captcha.com")
	viper.SetDefault("providers.generic.api_url", "")
	viper.SetDefault("providers.generic.username", "")
	viper.SetDefault("providers.generic.password", "")
	viper.SetDefault("providers.generic.api_key", "")
	// Enable embedded subtitle provider by default so users can start
	// extracting subtitles without additional configuration.
	viper.SetDefault("providers.embedded.enabled", true)
	viper.SetDefault("plex.url", "http://localhost:32400")
	viper.SetDefault("plex.token", "")
	viper.SetDefault("server_name", "Subtitle Manager")
	viper.SetDefault("base_url", "")
	viper.SetDefault("reverse_proxy", false)
	viper.SetDefault("integrations.sonarr.enabled", false)
	viper.SetDefault("integrations.radarr.enabled", false)
	viper.SetDefault("integrations.bazarr.import", false)
	viper.SetDefault("integrations.plex.enabled", false)
	viper.SetDefault("integrations.notifications.enabled", false)

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
	if u := viper.GetString("anticaptcha.api_url"); u != "" {
		captcha.SetAPIURL(u)
	}
}
