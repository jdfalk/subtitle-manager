package grpcserver

import (
	"context"

	translatorpb "github.com/jdfalk/subtitle-manager/pkg/subtitle/translator/v1"
	"github.com/jdfalk/subtitle-manager/pkg/translator"
	"github.com/spf13/viper"
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
func (s *Server) SetConfig(ctx context.Context, req *translatorpb.SetConfigRequest) (*translatorpb.SetConfigResponse, error) {
	key := req.GetKey()
	value := req.GetValue()

	if s.persistConfig {
		// Use Viper for persistent configuration
		configKey := key
		if s.configKeyPrefix != "" {
			configKey = s.configKeyPrefix + key
		}
		viper.Set(configKey, value)

		// Try to save config - this might fail if no config file is set
		if err := viper.WriteConfig(); err != nil {
			// If WriteConfig fails (e.g., no config file), try SafeWriteConfig
			if err := viper.SafeWriteConfig(); err != nil {
				// Log the error but continue - the value is still set in memory
				// This allows the server to work without a config file
			}
		}
	} else {
		// Update internal state for non-persistent config
		switch key {
		case "GOOGLE_API_KEY", s.configKeyPrefix + "google_api_key":
			s.googleKey = value
		case "OPENAI_API_KEY", s.configKeyPrefix + "openai_api_key":
			s.gptKey = value
		}
	}

	return &translatorpb.SetConfigResponse{}, nil
}

// GetConfig retrieves the current configuration.
func (s *Server) GetConfig(ctx context.Context, _ *translatorpb.GetConfigRequest) (*translatorpb.GetConfigResponse, error) {
	configValues := make(map[string]string)

	if s.persistConfig {
		// Return all Viper settings
		allSettings := viper.AllSettings()
		for k, v := range allSettings {
			if str, ok := v.(string); ok {
				configValues[k] = str
			}
		}
	} else {
		// Return current API keys
		configValues["GOOGLE_API_KEY"] = s.googleKey
		configValues["OPENAI_API_KEY"] = s.gptKey
	}

	response := &translatorpb.GetConfigResponse{}
	response.SetConfigValues(configValues)
	return response, nil
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
