package main

import (
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/jdfalk/subtitle-manager/pkg/grpcserver"
	"github.com/jdfalk/subtitle-manager/pkg/logging"
	pb "github.com/jdfalk/subtitle-manager/pkg/translatorpb"
)

func main() {
	addr := ":50051"
	if v := os.Getenv("TRANSLATOR_ADDR"); v != "" {
		addr = v
	}

	s := grpc.NewServer()

	// Create server with memory-only config (no persistence)
	server := grpcserver.NewServer(
		os.Getenv("GOOGLE_API_KEY"),
		os.Getenv("OPENAI_API_KEY"),
		false, // persistConfig = false
		"",    // no prefix needed
	)

	pb.RegisterTranslatorServer(s, server)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		logging.GetLogger("grpc-server").Fatalf("Failed to listen: %v", err)
	}
	logging.GetLogger("grpc-server").Infof("listening on %s", addr)
	if err := s.Serve(lis); err != nil {
		logging.GetLogger("grpc-server").Fatalf("Failed to serve: %v", err)
	}
}
