package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	pb "github.com/jdfalk/subtitle-manager/pkg/translatorpb/proto"
)

var grpcConfigAddr string
var grpcConfigKey string
var grpcConfigValue string

var grpcSetConfigCmd = &cobra.Command{
	Use:   "grpc-set-config",
	Short: "Set configuration value via gRPC",
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := grpc.Dial(grpcConfigAddr, grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer conn.Close()
		client := pb.NewTranslatorClient(conn)
		_, err = client.SetConfig(context.Background(), &pb.ConfigRequest{Settings: map[string]string{grpcConfigKey: grpcConfigValue}})
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
