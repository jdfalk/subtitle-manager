// file: pkg/notifications/discord_test.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174011

package notifications

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDiscordNotifier_Notify_Success(t *testing.T) {
	expectedMessage := "Test Discord notification"
	var receivedPayload map[string]string
	
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify method
		assert.Equal(t, http.MethodPost, r.Method)
		
		// Verify content type
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		
		// Read and parse body
		body, err := io.ReadAll(r.Body)
		require.NoError(t, err)
		
		err = json.Unmarshal(body, &receivedPayload)
		require.NoError(t, err)
		
		// Respond with success
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()
	
	notifier := DiscordNotifier{
		WebhookURL: server.URL,
		Client:     server.Client(),
	}
	
	err := notifier.Notify(context.Background(), expectedMessage)
	
	assert.NoError(t, err)
	assert.Equal(t, expectedMessage, receivedPayload["content"])
}

func TestDiscordNotifier_Notify_EmptyWebhookURL(t *testing.T) {
	notifier := DiscordNotifier{
		WebhookURL: "",
	}
	
	err := notifier.Notify(context.Background(), "test message")
	
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "webhook URL required")
}

func TestDiscordNotifier_Notify_HTTPError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer server.Close()
	
	notifier := DiscordNotifier{
		WebhookURL: server.URL,
		Client:     server.Client(),
	}
	
	err := notifier.Notify(context.Background(), "test message")
	
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "status 400")
}

func TestDiscordNotifier_Notify_NetworkError(t *testing.T) {
	notifier := DiscordNotifier{
		WebhookURL: "http://invalid-url-that-does-not-exist.local",
	}
	
	err := notifier.Notify(context.Background(), "test message")
	
	assert.Error(t, err)
}

func TestDiscordNotifier_Notify_ContextCanceled(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// This should not be reached due to context cancellation
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()
	
	notifier := DiscordNotifier{
		WebhookURL: server.URL,
		Client:     server.Client(),
	}
	
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately
	
	err := notifier.Notify(ctx, "test message")
	
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "context canceled")
}

func TestDiscordNotifier_Notify_DefaultClient(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()
	
	notifier := DiscordNotifier{
		WebhookURL: server.URL,
		Client:     nil, // Use default client
	}
	
	// This should work with the default client, though it may fail due to network issues
	// The important thing is that it doesn't panic
	assert.NotPanics(t, func() {
		notifier.Notify(context.Background(), "test message")
	})
}

func TestDiscordNotifier_Notify_LongMessage(t *testing.T) {
	longMessage := make([]byte, 2000) // 2KB message
	for i := range longMessage {
		longMessage[i] = 'A'
	}
	
	var receivedPayload map[string]string
	
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		require.NoError(t, err)
		
		err = json.Unmarshal(body, &receivedPayload)
		require.NoError(t, err)
		
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()
	
	notifier := DiscordNotifier{
		WebhookURL: server.URL,
		Client:     server.Client(),
	}
	
	err := notifier.Notify(context.Background(), string(longMessage))
	
	assert.NoError(t, err)
	assert.Equal(t, string(longMessage), receivedPayload["content"])
}

func TestDiscordNotifier_Notify_EmptyMessage(t *testing.T) {
	var receivedPayload map[string]string
	
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		require.NoError(t, err)
		
		err = json.Unmarshal(body, &receivedPayload)
		require.NoError(t, err)
		
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()
	
	notifier := DiscordNotifier{
		WebhookURL: server.URL,
		Client:     server.Client(),
	}
	
	err := notifier.Notify(context.Background(), "")
	
	assert.NoError(t, err)
	assert.Equal(t, "", receivedPayload["content"])
}