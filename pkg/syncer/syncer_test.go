package syncer

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/asticode/go-astisub"
	pb "github.com/jdfalk/subtitle-manager/pkg/translatorpb/proto"
	"google.golang.org/grpc"
)

// TestShift verifies that the Shift function offsets subtitles by the given duration.
func TestShift(t *testing.T) {
	items := []*astisub.Item{{StartAt: 0, EndAt: time.Second}}
	out := Shift(items, 2*time.Second)
	if out[0].StartAt != 2*time.Second || out[0].EndAt != 3*time.Second {
		t.Fatalf("unexpected values: %#v", out[0])
	}
}

// TestSync loads a subtitle file to ensure no error is returned.
func TestSync(t *testing.T) {
	items, err := Sync("dummy.mkv", "../../testdata/simple.srt", Options{})
	if err != nil {
		t.Fatalf("sync: %v", err)
	}
	if len(items) == 0 {
		t.Fatal("no items returned")
	}
}

// mockServer returns "hola" for any translation request.
type mockServer struct {
	pb.UnimplementedTranslatorServer
}

func (mockServer) Translate(ctx context.Context, req *pb.TranslateRequest) (*pb.TranslateResponse, error) {
	return &pb.TranslateResponse{TranslatedText: "hola"}, nil
}

// TestTranslate verifies that subtitle items are translated using a gRPC provider.
func TestTranslate(t *testing.T) {
	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTranslatorServer(s, &mockServer{})
	go s.Serve(lis)
	defer s.Stop()

	items := []*astisub.Item{{Lines: []astisub.Line{{Items: []astisub.LineItem{{Text: "hello"}}}}}}
	out, err := Translate(items, "es", "grpc", "", "", lis.Addr().String())
	if err != nil {
		t.Fatalf("translate: %v", err)
	}
	if out[0].String() != "hola" {
		t.Fatalf("expected hola, got %s", out[0].String())
	}
}

// TestSyncTranslate ensures Sync translates subtitles when options specify a language.
func TestSyncTranslate(t *testing.T) {
	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTranslatorServer(s, &mockServer{})
	go s.Serve(lis)
	defer s.Stop()

	opts := Options{TargetLang: "es", Service: "grpc", GRPCAddr: lis.Addr().String()}
	items, err := Sync("dummy.mkv", "../../testdata/simple.srt", opts)
	if err != nil {
		t.Fatalf("sync: %v", err)
	}
	if items[0].String() != "hola" {
		t.Fatalf("expected hola, got %s", items[0].String())
	}
}
