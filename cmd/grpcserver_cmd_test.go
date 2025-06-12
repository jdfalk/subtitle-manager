package cmd

import (
	"context"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
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
	// Replace the Google API endpoint with a stub server.
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"data":{"translations":[{"translatedText":"ok"}]}}`)
	}))
	defer ts.Close()
	translator.SetGoogleAPIURL(ts.URL)
	defer translator.SetGoogleAPIURL("https://translation.googleapis.com/language/translate/v2")

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
