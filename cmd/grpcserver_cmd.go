package cmd

import (
	"net"

	"github.com/jdfalk/subtitle-manager/pkg/grpcserver"
	"github.com/jdfalk/subtitle-manager/pkg/logging"
	pb "github.com/jdfalk/subtitle-manager/pkg/translatorpb"
	"github.com/jdfalk/subtitle-manager/pkg/webserver"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var grpcAddr string

// grpcServerCmd runs a gRPC translation server using the configured API keys.
var grpcServerCmd = &cobra.Command{
	Use:   "grpc-server",
	Short: "Run translation gRPC server",
	RunE: func(cmd *cobra.Command, args []string) error {
		s := grpc.NewServer()

		// Create server with persistent config enabled (uses Viper)
		server := grpcserver.NewServer(
			viper.GetString("google_api_key"),
			viper.GetString("openai_api_key"),
			true, // persistConfig = true
			"",   // no prefix for Viper keys
		)

		pb.RegisterTranslatorServiceServer(s, server)

		if err := webserver.InitializeHealth(""); err == nil {
			if provider := webserver.GetHealthProvider(); provider != nil {
				// ghealth.NewGRPCServer(provider).Register(s) // TODO: Replace with simple gRPC health
			}
		}
		lis, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			return err
		}
		logging.GetLogger("grpc-server").Infof("listening on %s", grpcAddr)
		return s.Serve(lis)
	},
}

func init() {
	grpcServerCmd.Flags().StringVar(&grpcAddr, "addr", ":50051", "listen address")
	rootCmd.AddCommand(grpcServerCmd)
}
