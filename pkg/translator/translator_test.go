package translator

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"google.golang.org/grpc"
	pb "subtitle-manager/pkg/translatorpb/proto"
)

func TestGoogleTranslate(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		r.Body.Close()
		s := string(body)
		if !(strings.Contains(s, "q=hello") && strings.Contains(s, "target=es") && strings.Contains(s, "key=test")) {
			t.Fatalf("unexpected request body: %s", s)
		}
		fmt.Fprint(w, `{"data":{"translations":[{"translatedText":"hola"}]}}`)
	}))
	defer srv.Close()

	origURL := googleAPIURL
	SetGoogleAPIURL(srv.URL)
	defer SetGoogleAPIURL(origURL)

	got, err := GoogleTranslate("hello", "es", "test")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "hola" {
		t.Fatalf("expected hola, got %s", got)
	}
}

func TestUnsupportedServiceError(t *testing.T) {
	if ErrUnsupportedService.Error() == "" {
		t.Fatal("error string empty")
	}
}

func TestTranslate(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"data":{"translations":[{"translatedText":"hola"}]}}`)
	}))
	defer srv.Close()
	origURL := googleAPIURL
	SetGoogleAPIURL(srv.URL)
	defer SetGoogleAPIURL(origURL)

	got, err := Translate("google", "hello", "es", "test", "", "")
	if err != nil {
		t.Fatalf("translate: %v", err)
	}
	if got != "hola" {
		t.Fatalf("expected hola, got %s", got)
	}
}

func TestGRPCTranslate(t *testing.T) {
	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTranslatorServer(s, &mockServer{})
	go s.Serve(lis)
	defer s.Stop()

	got, err := GRPCTranslate("hello", "es", lis.Addr().String())
	if err != nil {
		t.Fatalf("grpc translate: %v", err)
	}
	if got != "hola" {
		t.Fatalf("expected hola, got %s", got)
	}
}

func TestTranslateGRPCProvider(t *testing.T) {
	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTranslatorServer(s, &mockServer{})
	go s.Serve(lis)
	defer s.Stop()

	got, err := Translate("grpc", "hello", "es", "", "", lis.Addr().String())
	if err != nil {
		t.Fatalf("translate grpc: %v", err)
	}
	if got != "hola" {
		t.Fatalf("expected hola, got %s", got)
	}
}

type mockServer struct {
	pb.UnimplementedTranslatorServer
}

func (mockServer) Translate(ctx context.Context, req *pb.TranslateRequest) (*pb.TranslateResponse, error) {
	return &pb.TranslateResponse{TranslatedText: "hola"}, nil
}
