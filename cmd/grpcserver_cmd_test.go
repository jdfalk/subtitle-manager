package cmd

import (
	"context"
	"net"
	"testing"
	"time"

	translate "cloud.google.com/go/translate"
	"github.com/jdfalk/subtitle-manager/pkg/grpcserver"
	pb "github.com/jdfalk/subtitle-manager/pkg/subtitle/translator/v1"
	"github.com/jdfalk/subtitle-manager/pkg/translator"
	translatormocks "github.com/jdfalk/subtitle-manager/pkg/translator/mocks"
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
	m := translatormocks.NewMockGoogleClient(t)
	translator.SetGoogleClientFactory(func(ctx context.Context, apiKey string) (translator.GoogleClient, error) { return m, nil })
	defer translator.ResetGoogleClientFactory()

	m.On("Translate", mock.Anything, []string{"x"}, language.Make("en"), (*translate.Options)(nil)).Return([]translate.Translation{{Text: "ok"}}, nil)
	m.On("Close").Return(nil)

	s := grpc.NewServer()
	server := grpcserver.NewServer("k", "", false, "")
	pb.RegisterTranslatorServiceServer(s, server)
	go s.Serve(lis)
	defer s.Stop()

	conn, err := grpc.NewClient(lis.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	defer conn.Close()
	c := pb.NewTranslatorServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	text := "x"
	language := "en"
	resp, err := c.Translate(ctx, &pb.TranslateRequest{Text: &text, Language: &language})
	if err != nil {
		t.Fatalf("translate: %v", err)
	}
	if resp.TranslatedText == nil || *resp.TranslatedText == "" {
		t.Fatal("empty response")
	}
}
