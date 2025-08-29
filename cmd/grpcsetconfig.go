// file: cmd/grpcsetconfig.go
// version: 2.0.0
// guid: e252459e-8583-447a-81d5-1ac0eb51979c

package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/jdfalk/gcommon/sdks/go/v1/common"
	"github.com/jdfalk/gcommon/sdks/go/v1/config"
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
		client := pb.NewTranslatorServiceClient(conn)
		// Create a gcommon ConfigValue for the setting
		configValue := &common.ConfigValue{}
		configValue.SetStringValue(grpcConfigValue)
		
		// Create a SetConfigRequest
		req := &config.SetConfigRequest{}
		req.SetKey(grpcConfigKey)
		req.SetValue(configValue)
		
		// Note: The actual gRPC service call would depend on the service implementation
		// For now, this shows the pattern of using gcommon config types
		_ = req // Placeholder to use the variable
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
