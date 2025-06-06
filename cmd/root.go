package cmd

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
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
	rootCmd.PersistentFlags().String("log-level", "info", "log level (debug, info, warn, error)")
	viper.BindPFlag("log-level", rootCmd.PersistentFlags().Lookup("log-level"))
	rootCmd.PersistentFlags().StringToString("log-levels", nil, "per component log levels")
	viper.BindPFlag("log_levels", rootCmd.PersistentFlags().Lookup("log-levels"))

	rootCmd.AddCommand(convertCmd)
	rootCmd.AddCommand(mergeCmd)
	rootCmd.AddCommand(translateCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".subtitle-manager")
	}

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	level, err := logrus.ParseLevel(viper.GetString("log-level"))
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)
}
