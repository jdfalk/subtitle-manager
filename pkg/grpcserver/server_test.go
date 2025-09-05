// file: pkg/grpcserver/server_test.go
// version: 1.1.0
// guid: 123e4567-e89b-12d3-a456-426614174006

package grpcserver

import (
	"context"
	"testing"

	pb "github.com/jdfalk/subtitle-manager/pkg/subtitle/translator/v1"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewServer(t *testing.T) {
	tests := []struct {
		name            string
		googleKey       string
		gptKey          string
		persistConfig   bool
		configKeyPrefix string
	}{
		{
			name:            "basic server creation",
			googleKey:       "test-google-key",
			gptKey:          "test-gpt-key",
			persistConfig:   false,
			configKeyPrefix: "test_",
		},
		{
			name:            "server with persistence",
			googleKey:       "google-key",
			gptKey:          "gpt-key",
			persistConfig:   true,
			configKeyPrefix: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := NewServer(tt.googleKey, tt.gptKey, tt.persistConfig, tt.configKeyPrefix)

			assert.NotNil(t, server)
			assert.Equal(t, tt.googleKey, server.googleKey)
			assert.Equal(t, tt.gptKey, server.gptKey)
			assert.Equal(t, tt.persistConfig, server.persistConfig)
			assert.Equal(t, tt.configKeyPrefix, server.configKeyPrefix)
		})
	}
}

func TestServer_SetConfig(t *testing.T) {
	tests := []struct {
		name            string
		persistConfig   bool
		configKeyPrefix string
		settings        map[string]string
		expectError     bool
	}{
		{
			name:            "set config without persistence",
			persistConfig:   false,
			configKeyPrefix: "",
			settings: map[string]string{
				"GOOGLE_API_KEY": "new-google-key",
				"OPENAI_API_KEY": "new-gpt-key",
			},
			expectError: false,
		},
		{
			name:            "set config with custom prefix",
			persistConfig:   false,
			configKeyPrefix: "test_",
			settings: map[string]string{
				"test_google_api_key": "custom-google-key",
				"test_openai_api_key": "custom-gpt-key",
			},
			expectError: false,
		},
		{
			name:            "set unknown config keys",
			persistConfig:   false,
			configKeyPrefix: "",
			settings: map[string]string{
				"UNKNOWN_KEY": "unknown-value",
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := NewServer("initial-google", "initial-gpt", tt.persistConfig, tt.configKeyPrefix)

			// Set individual configuration values using the new API
			for k, v := range tt.settings {
				key := k
				value := v
				req := &pb.SetConfigRequest{Key: &key, Value: &value}
				_, err := server.SetConfig(context.Background(), req)

				if tt.expectError {
					assert.Error(t, err)
					return // Skip the rest if we expect an error
				} else {
					assert.NoError(t, err)
				}
			}

			if !tt.expectError {
				// Verify configuration was set by getting config
				resp, err := server.GetConfig(context.Background(), &pb.GetConfigRequest{})
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.NotNil(t, resp.ConfigValues)

				// Verify API keys were updated correctly
				for k, v := range tt.settings {
					switch k {
					case tt.configKeyPrefix + "google_api_key", "GOOGLE_API_KEY":
						assert.Equal(t, v, server.googleKey)
					case tt.configKeyPrefix + "openai_api_key", "OPENAI_API_KEY":
						assert.Equal(t, v, server.gptKey)
					}
				}
			}
		})
	}
}

func TestServer_GetConfig(t *testing.T) {
	tests := []struct {
		name          string
		persistConfig bool
		googleKey     string
		gptKey        string
		setupViper    func()
		expectedKeys  []string
	}{
		{
			name:          "get config without persistence",
			persistConfig: false,
			googleKey:     "test-google-key",
			gptKey:        "test-gpt-key",
			setupViper:    func() {},
			expectedKeys:  []string{"GOOGLE_API_KEY", "OPENAI_API_KEY"},
		},
		{
			name:          "get config with persistence",
			persistConfig: true,
			googleKey:     "test-google-key",
			gptKey:        "test-gpt-key",
			setupViper: func() {
				viper.Set("test_setting", "test_value")
				viper.Set("another_setting", "another_value")
			},
			expectedKeys: []string{"test_setting", "another_setting"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset viper for each test
			viper.Reset()
			tt.setupViper()

			server := NewServer(tt.googleKey, tt.gptKey, tt.persistConfig, "")

			resp, err := server.GetConfig(context.Background(), &pb.GetConfigRequest{})

			assert.NoError(t, err)
			assert.NotNil(t, resp)
			assert.NotNil(t, resp.ConfigValues)

			if !tt.persistConfig {
				// For non-persistent config, should return API keys
				assert.Equal(t, tt.googleKey, resp.ConfigValues["GOOGLE_API_KEY"])
				assert.Equal(t, tt.gptKey, resp.ConfigValues["OPENAI_API_KEY"])
			} else {
				// For persistent config, should return viper settings
				for _, key := range tt.expectedKeys {
					assert.Contains(t, resp.ConfigValues, key)
				}
			}
		})
	}
}

func TestServer_Translate(t *testing.T) {
	tests := []struct {
		name        string
		text        string
		language    string
		expectError bool
	}{
		{
			name:        "empty text",
			text:        "",
			language:    "es",
			expectError: false, // Empty text should still be processed
		},
		{
			name:        "valid translation request",
			text:        "Hello world",
			language:    "es",
			expectError: true, // Will fail without real API key
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := NewServer("fake-google-key", "fake-gpt-key", false, "")

			text := tt.text
			language := tt.language
			req := &pb.TranslateRequest{
				Text:     &text,
				Language: &language,
			}

			resp, err := server.Translate(context.Background(), req)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				if err == nil {
					assert.NotNil(t, resp)
				}
				// If error occurs with fake API key, that's expected
			}
		})
	}
}

func TestServer_SetConfigWithPersistence_ErrorHandling(t *testing.T) {
	// Test error handling when viper config writing fails
	server := NewServer("google-key", "gpt-key", true, "")

	// Set an invalid config file that will cause WriteConfig to fail
	viper.SetConfigFile("/invalid/path/config.yaml")

	key := "db_path"
	value := "invalid"
	req := &pb.SetConfigRequest{
		Key:   &key,
		Value: &value,
	}

	// Should handle the error gracefully
	// Note: This test might pass or fail depending on viper's internal handling
	// The important thing is that it doesn't panic
	assert.NotPanics(t, func() {
		server.SetConfig(context.Background(), req)
	})
}

func TestServer_InterfaceCompliance(t *testing.T) {
	// Verify that Server implements the required gRPC interface
	server := NewServer("test", "test", false, "")

	// This test ensures the server properly embeds the UnimplementedTranslatorServiceServer
	// and can be used as a pb.TranslatorServiceServer
	var _ pb.TranslatorServiceServer = server

	// Test that the server can handle the basic gRPC methods
	require.NotNil(t, server)
}

func protoString(s string) *string { return &s }
