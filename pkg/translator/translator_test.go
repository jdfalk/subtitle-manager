// file: pkg/translator/translator_test.go
// version: 1.1.0
// guid: 8ae1f81d-0b31-49e8-bc2f-22e6b0a058d4

package translator

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	translate "cloud.google.com/go/translate"
	pb "github.com/jdfalk/subtitle-manager/pkg/translatorpb"
	openai "github.com/sashabaranov/go-openai"
	"github.com/stretchr/testify/mock"
	"golang.org/x/text/language"
	"google.golang.org/grpc"

	"github.com/jdfalk/subtitle-manager/pkg/cache"
	"github.com/jdfalk/subtitle-manager/pkg/translator/mocks"
)

// TestGoogleTranslate verifies that GoogleTranslate returns the translated text
// provided by the injected Google client.
func TestGoogleTranslate(t *testing.T) {
	m := mocks.NewGoogleClient(t)
	SetGoogleClientFactory(func(ctx context.Context, apiKey string) (GoogleClient, error) { return m, nil })
	defer ResetGoogleClientFactory()

	m.On("Translate", mock.Anything, []string{"hello"}, language.Make("es"), (*translate.Options)(nil)).Return([]translate.Translation{{Text: "hola"}}, nil)
	m.On("Close").Return(nil)

	got, err := GoogleTranslate("hello", "es", "test")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "hola" {
		t.Fatalf("expected hola, got %s", got)
	}
}

// TestGoogleTranslateBatch verifies that multiple strings can be translated in one call.
func TestGoogleTranslateBatch(t *testing.T) {
	m := mocks.NewGoogleClient(t)
	SetGoogleClientFactory(func(ctx context.Context, apiKey string) (GoogleClient, error) { return m, nil })
	defer ResetGoogleClientFactory()

	src := []string{"hello", "world"}
	expected := []translate.Translation{{Text: "hola"}, {Text: "mundo"}}
	m.On("Translate", mock.Anything, src, language.Make("es"), (*translate.Options)(nil)).Return(expected, nil)
	m.On("Close").Return(nil)

	got, err := GoogleTranslateBatch(src, "es", "k")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 2 || got[0] != "hola" || got[1] != "mundo" {
		t.Fatalf("unexpected result: %v", got)
	}
}

// TestUnsupportedServiceError ensures the unsupported service error is not
// empty.
func TestUnsupportedServiceError(t *testing.T) {
	if ErrUnsupportedService.Error() == "" {
		t.Fatal("error string empty")
	}
}

// TestTranslate validates that Translate delegates to the correct provider and
// returns the expected translation.
func TestTranslate(t *testing.T) {
	m := mocks.NewGoogleClient(t)
	SetGoogleClientFactory(func(ctx context.Context, apiKey string) (GoogleClient, error) { return m, nil })
	defer ResetGoogleClientFactory()

	m.On("Translate", mock.Anything, []string{"hello"}, language.Make("es"), (*translate.Options)(nil)).Return([]translate.Translation{{Text: "hola"}}, nil)
	m.On("Close").Return(nil)

	got, err := Translate("google", "hello", "es", "test", "", "")
	if err != nil {
		t.Fatalf("translate: %v", err)
	}
	if got != "hola" {
		t.Fatalf("expected hola, got %s", got)
	}
}

// TestGRPCTranslate confirms that GRPCTranslate communicates with the gRPC
// server and returns the translated text.
func TestGRPCTranslate(t *testing.T) {
	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTranslatorServer(s, &mockServer{})
	go func() {
		if err := s.Serve(lis); err != nil {
			t.Errorf("serve failed: %v", err)
		}
	}()
	defer s.Stop()

	got, err := GRPCTranslate("hello", "es", lis.Addr().String())
	if err != nil {
		t.Fatalf("grpc translate: %v", err)
	}
	if got != "hola" {
		t.Fatalf("expected hola, got %s", got)
	}
}

// TestTranslateGRPCProvider checks that the gRPC provider works when selected
// via Translate.
func TestTranslateGRPCProvider(t *testing.T) {
	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTranslatorServer(s, &mockServer{})
	go func() {
		if err := s.Serve(lis); err != nil {
			t.Errorf("serve failed: %v", err)
		}
	}()
	defer s.Stop()

	got, err := Translate("grpc", "hello", "es", "", "", lis.Addr().String())
	if err != nil {
		t.Fatalf("translate grpc: %v", err)
	}
	if got != "hola" {
		t.Fatalf("expected hola, got %s", got)
	}
}

// TestGPTTranslate verifies that GPTTranslate uses the injected OpenAI client
// to obtain a translation.
func TestGPTTranslate(t *testing.T) {
	m := mocks.NewOpenAIClient(t)
	SetOpenAIClientFactory(func(apiKey string) OpenAIClient { return m })
	defer ResetOpenAIClientFactory()

	m.On("CreateChatCompletion", mock.Anything, mock.Anything).Return(openai.ChatCompletionResponse{
		Choices: []openai.ChatCompletionChoice{{Message: openai.ChatCompletionMessage{Content: "hola"}}},
	}, nil)

	got, err := GPTTranslate("hello", "es", "test")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "hola" {
		t.Fatalf("expected hola, got %s", got)
	}
}

// TestTranslateGPTProvider ensures the GPT provider works when selected via
// Translate.
func TestTranslateGPTProvider(t *testing.T) {
	m := mocks.NewOpenAIClient(t)
	SetOpenAIClientFactory(func(apiKey string) OpenAIClient { return m })
	defer ResetOpenAIClientFactory()

	m.On("CreateChatCompletion", mock.Anything, mock.Anything).Return(openai.ChatCompletionResponse{
		Choices: []openai.ChatCompletionChoice{{Message: openai.ChatCompletionMessage{Content: "hola"}}},
	}, nil)

	got, err := Translate("gpt", "hello", "es", "", "test", "")
	if err != nil {
		t.Fatalf("translate gpt: %v", err)
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

// TestSetOpenAIModel verifies that the model can be changed.
func TestSetOpenAIModel(t *testing.T) {
	orig := openAIModel
	SetOpenAIModel("test-model")
	if openAIModel != "test-model" {
		t.Fatalf("expected test-model, got %s", openAIModel)
	}
	SetOpenAIModel(orig)
}

// TestSupportedServices ensures the provider list is returned alphabetically.
func TestSupportedServices(t *testing.T) {
	expected := []string{"chatgpt", "google", "gpt", "grpc"}
	got := SupportedServices()
	if len(got) != len(expected) {
		t.Fatalf("expected %d services, got %d", len(expected), len(got))
	}
	for i, name := range expected {
		if got[i] != name {
			t.Fatalf("expected %s at index %d, got %s", name, i, got[i])
		}
	}
}

// TestTranslateCacheManager verifies that Translate uses the configured cache
// manager to avoid duplicate API calls.
func TestTranslateCacheManager(t *testing.T) {
	count := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		count++
		fmt.Fprint(w, `{"data":{"translations":[{"translatedText":"hola"}]}}`)
	}))
	defer srv.Close()

	SetGoogleAPIURL(srv.URL)
	defer SetGoogleAPIURL("https://translation.googleapis.com/language/translate/v2")

	manager, err := cache.NewManager(cache.DefaultConfig())
	if err != nil {
		t.Fatalf("cache init: %v", err)
	}
	defer manager.Close()
	SetCacheManager(manager)
	defer SetCacheManager(nil)

	got, err := Translate("google", "hello", "es", "k", "", "")
	if err != nil || got != "hola" {
		t.Fatalf("first translate failed: %v %s", err, got)
	}

	got, err = Translate("google", "hello", "es", "k", "", "")
	if err != nil || got != "hola" {
		t.Fatalf("second translate failed: %v %s", err, got)
	}
	if count != 1 {
		t.Fatalf("expected 1 API call, got %d", count)
	}
}
