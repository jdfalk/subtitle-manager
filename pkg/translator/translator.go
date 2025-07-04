package translator

import (
	"context"
	"errors"
	"fmt"
	"strings"

	translate "cloud.google.com/go/translate"
	"golang.org/x/text/language"
	"google.golang.org/api/option"

	openai "github.com/sashabaranov/go-openai"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/jdfalk/subtitle-manager/pkg/translatorpb/proto"
)

var ErrUnsupportedService = errors.New("unsupported translation service")

var googleAPIURL = "https://translation.googleapis.com/language/translate/v2"
var openAIModel = openai.GPT3Dot5Turbo

// SetGoogleAPIURL overrides the Google Translate API URL (useful for testing).
func SetGoogleAPIURL(u string) {
	googleAPIURL = u
}

// SetOpenAIModel overrides the default ChatGPT model. The parameter m is the
// model identifier to use for future requests.
func SetOpenAIModel(m string) {
	openAIModel = m
}

// TranslateFunc defines the function signature for translation services.
type TranslateFunc func(text, targetLang, apiKey string) (string, error)

// GoogleClient wraps the methods used from the Google Translate SDK.
// It allows tests to mock the SDK without real credentials.
//
//go:generate go run github.com/vektra/mockery/v2 --name=GoogleClient --output=mocks --outpkg=mocks --filename=google_client.go
type GoogleClient interface {
	Translate(ctx context.Context, src []string, target language.Tag, opts *translate.Options) ([]translate.Translation, error)
	Close() error
}

// defaultGoogleClient instantiates the real Google Translate client.
var defaultGoogleClient = func(ctx context.Context, apiKey string) (GoogleClient, error) {
	return translate.NewClient(ctx,
		option.WithAPIKey(apiKey),
		option.WithEndpoint(googleAPIURL+"/"),
	)
}

// newGoogleClient is the factory used by GoogleTranslate to create clients.
// It can be replaced in tests via SetGoogleClientFactory.
var newGoogleClient = defaultGoogleClient

// OpenAIClient wraps the methods used from the OpenAI SDK.
// It allows tests to mock the ChatGPT API without real credentials.
//
//go:generate go run github.com/vektra/mockery/v2 --name=OpenAIClient --output=mocks --outpkg=mocks --filename=openai_client.go
type OpenAIClient interface {
	CreateChatCompletion(ctx context.Context, req openai.ChatCompletionRequest) (openai.ChatCompletionResponse, error)
}

// defaultOpenAIClient instantiates the real OpenAI client.
var defaultOpenAIClient = func(apiKey string) OpenAIClient {
	return openai.NewClient(apiKey)
}

// newOpenAIClient is the factory used by GPTTranslate to create clients.
// It can be replaced in tests via SetOpenAIClientFactory.
var newOpenAIClient = defaultOpenAIClient

// SetGoogleClientFactory replaces the Google client constructor.
// Tests can use this to inject a mock implementation.
func SetGoogleClientFactory(fn func(ctx context.Context, apiKey string) (GoogleClient, error)) {
	newGoogleClient = fn
}

// ResetGoogleClientFactory restores the default Google client constructor.
func ResetGoogleClientFactory() {
	newGoogleClient = defaultGoogleClient
}

// SetOpenAIClientFactory replaces the OpenAI client constructor. The fn
// parameter should return an implementation that satisfies OpenAIClient and will
// be used by GPTTranslate. Tests can use this to inject a mock client.
func SetOpenAIClientFactory(fn func(apiKey string) OpenAIClient) {
	newOpenAIClient = fn
}

// ResetOpenAIClientFactory restores the default OpenAI client constructor so the
// real OpenAI client is used again.
func ResetOpenAIClientFactory() {
	newOpenAIClient = defaultOpenAIClient
}

// GoogleTranslate translates text using Google Translate API.
func GoogleTranslate(text, targetLang, apiKey string) (string, error) {
	ctx := context.Background()
	client, err := newGoogleClient(ctx, apiKey)
	if err != nil {
		return "", err
	}
	defer client.Close()

	ts, err := client.Translate(ctx, []string{text}, language.Make(targetLang), nil)
	if err != nil {
		return "", err
	}
	if len(ts) == 0 {
		return "", fmt.Errorf("no translations")
	}
	return ts[0].Text, nil
}

// GPTTranslate translates text using the ChatGPT API.
func GPTTranslate(text, targetLang, apiKey string) (string, error) {
	client := newOpenAIClient(apiKey)
	req := openai.ChatCompletionRequest{
		Model: openAIModel,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: fmt.Sprintf("Translate the following text to %s. Return only the translated text.", targetLang),
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: text,
			},
		},
	}
	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return "", err
	}
	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no translation returned")
	}
	return strings.TrimSpace(resp.Choices[0].Message.Content), nil
}

// GRPCTranslate translates text using a remote gRPC translation service.
// The addr parameter specifies the server address (host:port).
// It returns the translated text provided by the service.
func GRPCTranslate(text, targetLang, addr string) (string, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return "", err
	}
	defer conn.Close()
	client := pb.NewTranslatorClient(conn)
	resp, err := client.Translate(context.Background(), &pb.TranslateRequest{
		Text:     text,
		Language: targetLang,
	})
	if err != nil {
		return "", err
	}
	return resp.TranslatedText, nil
}

// GRPCSetConfig sends configuration key/value pairs to a remote gRPC server.
// The addr parameter specifies the server address (host:port).
func GRPCSetConfig(settings map[string]string, addr string) error {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()
	client := pb.NewTranslatorClient(conn)
	_, err = client.SetConfig(context.Background(), &pb.ConfigRequest{Settings: settings})
	return err
}

var providers = map[string]TranslateFunc{
	"google":  GoogleTranslate,
	"gpt":     GPTTranslate,
	"chatgpt": GPTTranslate,
	"grpc":    GRPCTranslate,
}

// Translate selects a provider and performs translation.
// googleKey and gptKey are used depending on the provider.
// Translate selects the provider identified by service and performs the
// translation using the given credentials. googleKey, gptKey and grpcAddr are
// used depending on the provider.
func Translate(service, text, targetLang, googleKey, gptKey, grpcAddr string) (string, error) {
	fn, ok := providers[service]
	if !ok {
		return "", ErrUnsupportedService
	}
	key := googleKey
	if service == "gpt" || service == "chatgpt" {
		key = gptKey
	} else if service == "grpc" {
		key = grpcAddr
	}
	return fn(text, targetLang, key)
}
