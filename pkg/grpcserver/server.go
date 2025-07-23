package grpcserver

import (
	"context"
	"fmt"

	gconfig "github.com/jdfalk/subtitle-manager/pkg/gcommon/config"
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
func (s *Server) SetConfig(ctx context.Context, req *pb.SubtitleManagerConfig) (*emptypb.Empty, error) {
	if req.GoogleApiKey != nil {
		s.googleKey = req.GetGoogleApiKey()
		if s.persistConfig {
			viper.Set(s.configKeyPrefix+"google_api_key", s.googleKey)
		}
	}
	if req.OpenaiApiKey != nil {
		s.gptKey = req.GetOpenaiApiKey()
		if s.persistConfig {
			viper.Set(s.configKeyPrefix+"openai_api_key", s.gptKey)
		}
	}
	if req.DbPath != nil && s.persistConfig {
		viper.Set("db_path", req.GetDbPath())
	}
	if req.DbBackend != nil && s.persistConfig {
		viper.Set("db_backend", req.GetDbBackend())
	}
	if req.Sqlite3Filename != nil && s.persistConfig {
		viper.Set("sqlite3_filename", req.GetSqlite3Filename())
	}
	if req.LogFile != nil && s.persistConfig {
		viper.Set("log_file", req.GetLogFile())
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
func (s *Server) GetConfig(ctx context.Context, _ *emptypb.Empty) (*pb.SubtitleManagerConfig, error) {
	if s.persistConfig {
		cfg := gconfig.ToProto()
		if s.googleKey != "" {
			cfg.GoogleApiKey = &s.googleKey
		}
		if s.gptKey != "" {
			cfg.OpenaiApiKey = &s.gptKey
		}
		return cfg, nil
	}
	return &pb.SubtitleManagerConfig{
		GoogleApiKey: &s.googleKey,
		OpenaiApiKey: &s.gptKey,
	}, nil
}

// Translate translates text using the configured translation service.
func (s *Server) Translate(ctx context.Context, req *pb.TranslateRequest) (*pb.TranslateResponse, error) {
	text, err := translator.Translate("google", req.Text, req.Language, s.googleKey, s.gptKey, "")
	if err != nil {
		return nil, fmt.Errorf("translation failed: %w", err)
	}
	return &pb.TranslateResponse{TranslatedText: text}, nil
}
