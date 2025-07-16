package grpcserver

import (
	"context"
	"fmt"

	"github.com/jdfalk/subtitle-manager/pkg/translator"
	pb "github.com/jdfalk/subtitle-manager/pkg/translatorpb"
	"github.com/spf13/viper"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// Server implements the gRPC translator service with configurable persistence.
type Server struct {
	pb.UnimplementedTranslatorServer
	googleKey       string
	gptKey          string
	persistConfig   bool
	configKeyPrefix string
}

// NewServer creates a new gRPC translator server.
// If persistConfig is true, configuration changes will be saved using Viper.
// configKeyPrefix allows customization of config key names (e.g., "google_api_key" vs "GOOGLE_API_KEY").
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
func (s *Server) SetConfig(ctx context.Context, req *pb.ConfigRequest) (*emptypb.Empty, error) {
	for k, v := range req.Settings {
		if s.persistConfig {
			// Use Viper for persistent config
			viper.Set(k, v)
		}

		// Update in-memory values based on key mapping
		switch k {
		case s.configKeyPrefix + "google_api_key", "GOOGLE_API_KEY":
			s.googleKey = v
		case s.configKeyPrefix + "openai_api_key", "OPENAI_API_KEY":
			s.gptKey = v
		}
	}

	if s.persistConfig {
		if cfg := viper.ConfigFileUsed(); cfg != "" {
			if err := viper.WriteConfig(); err != nil {
				return nil, fmt.Errorf("failed to write config: %w", err)
			}
		}
	}

	return &emptypb.Empty{}, nil
}

// GetConfig returns current configuration values.
func (s *Server) GetConfig(ctx context.Context, _ *emptypb.Empty) (*pb.ConfigResponse, error) {
	settings := make(map[string]string)

	if s.persistConfig {
		// Return all Viper settings for persistent config
		for k, v := range viper.AllSettings() {
			settings[k] = fmt.Sprintf("%v", v)
		}
	} else {
		// Return only API keys for memory-only config
		settings["GOOGLE_API_KEY"] = s.googleKey
		settings["OPENAI_API_KEY"] = s.gptKey
	}

	return &pb.ConfigResponse{Settings: settings}, nil
}

// Translate translates text using the configured translation service.
func (s *Server) Translate(ctx context.Context, req *pb.TranslateRequest) (*pb.TranslateResponse, error) {
	text, err := translator.Translate("google", req.Text, req.Language, s.googleKey, s.gptKey, "")
	if err != nil {
		return nil, fmt.Errorf("translation failed: %w", err)
	}
	return &pb.TranslateResponse{TranslatedText: text}, nil
}
