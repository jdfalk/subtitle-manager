package cmd

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	"subtitle-manager/pkg/translator"
	pb "subtitle-manager/pkg/translatorpb/proto"
)

var grpcAddr string

// grpcServerCmd runs a gRPC translation server using the configured API keys.
var grpcServerCmd = &cobra.Command{
	Use:   "grpc-server",
	Short: "Run translation gRPC server",
	RunE: func(cmd *cobra.Command, args []string) error {
		s := grpc.NewServer()
		pb.RegisterTranslatorServer(s, &server{
			googleKey: viper.GetString("google_api_key"),
			gptKey:    viper.GetString("openai_api_key"),
		})
		lis, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			return err
		}
		log.Printf("listening on %s", grpcAddr)
		return s.Serve(lis)
	},
}

type server struct {
	pb.UnimplementedTranslatorServer
	googleKey string
	gptKey    string
}

func (s *server) GetConfig(ctx context.Context, _ *emptypb.Empty) (*pb.ConfigResponse, error) {
	out := make(map[string]string)
	for k, v := range viper.AllSettings() {
		out[k] = fmt.Sprintf("%v", v)
	}
	return &pb.ConfigResponse{Settings: out}, nil
}

func (s *server) Translate(ctx context.Context, req *pb.TranslateRequest) (*pb.TranslateResponse, error) {
	text, err := translator.Translate("google", req.Text, req.Language, s.googleKey, s.gptKey, "")
	if err != nil {
		return nil, err
	}
	return &pb.TranslateResponse{TranslatedText: text}, nil
}

func init() {
	grpcServerCmd.Flags().StringVar(&grpcAddr, "addr", ":50051", "listen address")
	rootCmd.AddCommand(grpcServerCmd)
}
