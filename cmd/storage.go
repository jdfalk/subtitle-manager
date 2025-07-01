// file: cmd/storage.go
// version: 1.0.0
// guid: 90123456-78f3-901c-def0-123456789def

// Package cmd implements storage management commands for subtitle-manager.
package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/jdfalk/subtitle-manager/pkg/storage"
)

var storageCmd = &cobra.Command{
	Use:   "storage",
	Short: "Manage cloud storage for subtitles",
	Long:  `Commands for managing cloud storage providers and subtitle files`,
}

var storageTestCmd = &cobra.Command{
	Use:   "test",
	Short: "Test cloud storage connection",
	Long:  `Test the configured cloud storage provider connection`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := logging.GetLogger("storage")
		
		config := storage.GetConfigFromViper()
		if config.Provider == "" {
			config.Provider = "local"
		}
		
		logger.WithFields(logrus.Fields{
			"provider": config.Provider,
		}).Info("Testing storage provider connection")
		
		provider, err := storage.NewProvider(config)
		if err != nil {
			logger.WithError(err).Fatal("Failed to create storage provider")
		}
		defer provider.Close()
		
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		
		// Test basic operations
		testKey := "test/connection-test.txt"
		testContent := fmt.Sprintf("Connection test at %s", time.Now().Format(time.RFC3339))
		
		// Test store
		logger.Info("Testing store operation")
		err = provider.Store(ctx, testKey, strings.NewReader(testContent), "text/plain")
		if err != nil {
			logger.WithError(err).Fatal("Failed to store test file")
		}
		
		// Test exists
		logger.Info("Testing exists operation")
		exists, err := provider.Exists(ctx, testKey)
		if err != nil {
			logger.WithError(err).Fatal("Failed to check file existence")
		}
		if !exists {
			logger.Fatal("Test file should exist after storing")
		}
		
		// Test retrieve
		logger.Info("Testing retrieve operation")
		reader, err := provider.Retrieve(ctx, testKey)
		if err != nil {
			logger.WithError(err).Fatal("Failed to retrieve test file")
		}
		reader.Close()
		
		// Test list
		logger.Info("Testing list operation")
		keys, err := provider.List(ctx, "test/")
		if err != nil {
			logger.WithError(err).Fatal("Failed to list files")
		}
		logger.WithField("count", len(keys)).Info("Listed files")
		
		// Test delete
		logger.Info("Testing delete operation")
		err = provider.Delete(ctx, testKey)
		if err != nil {
			logger.WithError(err).Fatal("Failed to delete test file")
		}
		
		logger.Info("✅ All storage operations successful!")
	},
}

var storageListCmd = &cobra.Command{
	Use:   "list [prefix]",
	Short: "List files in cloud storage",
	Long:  `List subtitle files stored in the configured cloud storage provider`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := logging.GetLogger("storage")
		
		config := storage.GetConfigFromViper()
		if config.Provider == "" {
			config.Provider = "local"
		}
		
		provider, err := storage.NewProvider(config)
		if err != nil {
			logger.WithError(err).Fatal("Failed to create storage provider")
		}
		defer provider.Close()
		
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		
		prefix := ""
		if len(args) > 0 {
			prefix = args[0]
		}
		
		logger.WithFields(logrus.Fields{
			"provider": config.Provider,
			"prefix":   prefix,
		}).Info("Listing files")
		
		keys, err := provider.List(ctx, prefix)
		if err != nil {
			logger.WithError(err).Fatal("Failed to list files")
		}
		
		fmt.Printf("Found %d files:\n", len(keys))
		for _, key := range keys {
			fmt.Printf("  %s\n", key)
		}
	},
}

var storageUploadCmd = &cobra.Command{
	Use:   "upload <local-file> <storage-key>",
	Short: "Upload a file to cloud storage",
	Long:  `Upload a local subtitle file to the configured cloud storage provider`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		logger := logging.GetLogger("storage")
		
		localFile := args[0]
		storageKey := args[1]
		
		config := storage.GetConfigFromViper()
		if config.Provider == "" {
			config.Provider = "local"
		}
		
		provider, err := storage.NewProvider(config)
		if err != nil {
			logger.WithError(err).Fatal("Failed to create storage provider")
		}
		defer provider.Close()
		
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
		defer cancel()
		
		// Open local file
		file, err := os.Open(localFile)
		if err != nil {
			logger.WithError(err).Fatal("Failed to open local file")
		}
		defer file.Close()
		
		logger.WithFields(logrus.Fields{
			"local_file":  localFile,
			"storage_key": storageKey,
			"provider":    config.Provider,
		}).Info("Uploading file")
		
		err = provider.Store(ctx, storageKey, file, "text/plain")
		if err != nil {
			logger.WithError(err).Fatal("Failed to upload file")
		}
		
		logger.Info("✅ File uploaded successfully!")
	},
}

var storageDownloadCmd = &cobra.Command{
	Use:   "download <storage-key> <local-file>",
	Short: "Download a file from cloud storage",
	Long:  `Download a subtitle file from the configured cloud storage provider`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		logger := logging.GetLogger("storage")
		
		storageKey := args[0]
		localFile := args[1]
		
		config := storage.GetConfigFromViper()
		if config.Provider == "" {
			config.Provider = "local"
		}
		
		provider, err := storage.NewProvider(config)
		if err != nil {
			logger.WithError(err).Fatal("Failed to create storage provider")
		}
		defer provider.Close()
		
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
		defer cancel()
		
		logger.WithFields(logrus.Fields{
			"storage_key": storageKey,
			"local_file":  localFile,
			"provider":    config.Provider,
		}).Info("Downloading file")
		
		// Retrieve from storage
		reader, err := provider.Retrieve(ctx, storageKey)
		if err != nil {
			logger.WithError(err).Fatal("Failed to download file")
		}
		defer reader.Close()
		
		// Create local file
		file, err := os.Create(localFile)
		if err != nil {
			logger.WithError(err).Fatal("Failed to create local file")
		}
		defer file.Close()
		
		// Copy content
		written, err := file.ReadFrom(reader)
		if err != nil {
			logger.WithError(err).Fatal("Failed to write file content")
		}
		
		logger.WithField("bytes", written).Info("✅ File downloaded successfully!")
	},
}

func init() {
	storageCmd.AddCommand(storageTestCmd)
	storageCmd.AddCommand(storageListCmd)
	storageCmd.AddCommand(storageUploadCmd)
	storageCmd.AddCommand(storageDownloadCmd)
	
	rootCmd.AddCommand(storageCmd)
}