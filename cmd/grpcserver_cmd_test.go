package cmd

import (
	"context"
	"net"
	"testing"
	"time"

	translate "cloud.google.com/go/translate"
	"github.com/jdfalk/subtitle-manager/pkg/grpcserver"
	"github.com/jdfalk/subtitle-manager/pkg/translator"
	translatormocks "github.com/jdfalk/subtitle-manager/pkg/translator/mocks"
	pb "github.com/jdfalk/subtitle-manager/pkg/translatorpb/proto"
	"github.com/stretchr/testify/mock"
	"golang.org/x/text/language"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// TestServerTranslate verifies the gRPC Translate method.
func TestServerTranslate(t *testing.T) {
	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("listen: %v", err)
	}
	m := translatormocks.NewGoogleClient(t)
	translator.SetGoogleClientFactory(func(ctx context.Context, apiKey string) (translator.GoogleClient, error) { return m, nil })
	defer translator.ResetGoogleClientFactory()

	m.On("Translate", mock.Anything, []string{"x"}, language.Make("en"), (*translate.Options)(nil)).Return([]translate.Translation{{Text: "ok"}}, nil)
	m.On("Close").Return(nil)

	s := grpc.NewServer()
	server := grpcserver.NewServer("k", "", false, "")
	pb.RegisterTranslatorServer(s, server)
	go s.Serve(lis)
	defer s.Stop()

	conn, err := grpc.NewClient(lis.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	defer conn.Close()
	c := pb.NewTranslatorClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	resp, err := c.Translate(ctx, &pb.TranslateRequest{Text: "x", Language: "en"})
	if err != nil {
		t.Fatalf("translate: %v", err)
	}
	if resp.TranslatedText == "" {
		t.Fatal("empty response")
	}
}
