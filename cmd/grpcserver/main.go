package main

import (
	"context"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"

	"github.com/jdfalk/subtitle-manager/pkg/translator"
	pb "github.com/jdfalk/subtitle-manager/pkg/translatorpb/proto"
)

type server struct {
	pb.UnimplementedTranslatorServer
	googleKey string
	gptKey    string
}

func (s *server) GetConfig(ctx context.Context, _ *emptypb.Empty) (*pb.ConfigResponse, error) {
	out := map[string]string{
		"GOOGLE_API_KEY": s.googleKey,
		"OPENAI_API_KEY": s.gptKey,
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

func main() {
	addr := ":50051"
	if v := os.Getenv("TRANSLATOR_ADDR"); v != "" {
		addr = v
	}
	s := grpc.NewServer()
	pb.RegisterTranslatorServer(s, &server{
		googleKey: os.Getenv("GOOGLE_API_KEY"),
		gptKey:    os.Getenv("OPENAI_API_KEY"),
	})
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("listening on %s", addr)
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
