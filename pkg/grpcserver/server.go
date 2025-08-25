package grpcserver

import (
	"context"
	"fmt"

	"github.com/jdfalk/subtitle-manager/pkg/translator"
	translatorpb "github.com/jdfalk/subtitle-manager/pkg/translatorpb"
)

// Server implements the gRPC translator service with configurable persistence.
type Server struct {
	translatorpb.UnimplementedTranslatorServiceServer
	googleKey       string
	gptKey          string
	persistConfig   bool
	configKeyPrefix string
}

// NewServer creates a new server instance.
func NewServer(googleKey, gptKey string, persistConfig bool, configKeyPrefix string) *Server {
	return &Server{
		googleKey:       googleKey,
		gptKey:          gptKey,
		persistConfig:   persistConfig,
		configKeyPrefix: configKeyPrefix,
	}
}

// SetConfig updates configuration values. If persistConfig is enabled, changes
// are saved using Viper. The configKeyPrefix determines how keys are mapped.
// TODO: Implement proper SetConfig when protobuf types are available
func (s *Server) SetConfig(ctx context.Context, req *translatorpb.SetConfigRequest) (*translatorpb.SetConfigResponse, error) {
	return nil, fmt.Errorf("SetConfig not yet implemented")
}

// GetConfig retrieves the current configuration.
// TODO: Implement proper GetConfig when protobuf types are available
func (s *Server) GetConfig(ctx context.Context, _ *translatorpb.GetConfigRequest) (*translatorpb.GetConfigResponse, error) {
	return nil, fmt.Errorf("GetConfig not yet implemented")
}

// Translate performs text translation using the configured translator.
func (s *Server) Translate(ctx context.Context, req *translatorpb.TranslateRequest) (*translatorpb.TranslateResponse, error) {
	service := "google" // Default service
	text := req.GetText()
	targetLang := req.GetLanguage()

	result, err := translator.Translate(service, text, targetLang, s.googleKey, s.gptKey, "")
	if err != nil {
		return nil, err
	}

	response := &translatorpb.TranslateResponse{}
	response.SetTranslatedText(result)
	return response, nil
}
