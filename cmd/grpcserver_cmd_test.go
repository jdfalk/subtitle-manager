package cmd

import (
	"context"
	"net"
	"testing"

	translate "cloud.google.com/go/translate"
	"github.com/stretchr/testify/mock"
	"golang.org/x/text/language"
	"google.golang.org/grpc"
	"subtitle-manager/pkg/translator"
	translatormocks "subtitle-manager/pkg/translator/mocks"
	pb "subtitle-manager/pkg/translatorpb/proto"
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
	pb.RegisterTranslatorServer(s, &server{googleKey: "k"})
	go s.Serve(lis)
	defer s.Stop()

	conn, err := grpc.Dial(lis.Addr().String(), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	defer conn.Close()
	c := pb.NewTranslatorClient(conn)
	resp, err := c.Translate(context.Background(), &pb.TranslateRequest{Text: "x", Language: "en"})
	if err != nil {
		t.Fatalf("translate: %v", err)
	}
	if resp.TranslatedText == "" {
		t.Fatal("empty response")
	}
}
