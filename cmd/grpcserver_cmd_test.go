package cmd

import (
	"context"
	"net"
	"testing"

	"google.golang.org/grpc"
	"subtitle-manager/pkg/translator"
	pb "subtitle-manager/pkg/translatorpb/proto"
)

// TestServerTranslate verifies the gRPC Translate method.
func TestServerTranslate(t *testing.T) {
	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("listen: %v", err)
	}
	// Replace the google provider with a stub to avoid network calls.
	orig := translator.GoogleTranslate
	translator.GoogleTranslate = func(text, lang, key string) (string, error) { return "ok", nil }
	defer func() { translator.GoogleTranslate = orig }()

	s := grpc.NewServer()
	pb.RegisterTranslatorServer(s, &server{})
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
