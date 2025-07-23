// file: cmd/grpcsetconfig.go
// version: 1.2.0
// guid: e252459e-8583-447a-81d5-1ac0eb51979c

package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/jdfalk/subtitle-manager/pkg/translatorpb"
)

var grpcConfigAddr string
var grpcConfigKey string
var grpcConfigValue string

var grpcSetConfigCmd = &cobra.Command{
	Use:   "grpc-set-config",
	Short: "Set configuration value via gRPC",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		conn, err := grpc.NewClient(grpcConfigAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return err
		}
		defer conn.Close()
		client := pb.NewTranslatorClient(conn)
		cfg := &pb.SubtitleManagerConfig{}
		switch grpcConfigKey {
		case "google_api_key":
			cfg.GoogleApiKey = &grpcConfigValue
		case "openai_api_key":
			cfg.OpenaiApiKey = &grpcConfigValue
		case "db_path":
			cfg.DbPath = &grpcConfigValue
		case "db_backend":
			cfg.DbBackend = &grpcConfigValue
		case "sqlite3_filename":
			cfg.Sqlite3Filename = &grpcConfigValue
		case "log_file":
			cfg.LogFile = &grpcConfigValue
		default:
			return fmt.Errorf("unknown config key: %s", grpcConfigKey)
		}
		_, err = client.SetConfig(ctx, cfg)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	grpcSetConfigCmd.Flags().StringVar(&grpcConfigAddr, "addr", "localhost:50051", "gRPC server address")
	grpcSetConfigCmd.Flags().StringVar(&grpcConfigKey, "key", "", "configuration key")
	grpcSetConfigCmd.Flags().StringVar(&grpcConfigValue, "value", "", "configuration value")
	_ = grpcSetConfigCmd.MarkFlagRequired("key")
	_ = grpcSetConfigCmd.MarkFlagRequired("value")
	rootCmd.AddCommand(grpcSetConfigCmd)
}
