// file: pkg/notifications/telegram_test.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174013

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

func TestTelegramNotifier_Notify_Success(t *testing.T) {
	expectedMessage := "Test Telegram notification"
	expectedChatID := "123456789"
	expectedBotToken := "bot-token-123"
	var receivedPayload map[string]string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify method
		assert.Equal(t, http.MethodPost, r.Method)

		// Verify URL path
		expectedPath := "/bot" + expectedBotToken + "/sendMessage"
		assert.Equal(t, expectedPath, r.URL.Path)

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

	notifier := TelegramNotifier{
		BotToken: expectedBotToken,
		ChatID:   expectedChatID,
		Client:   server.Client(),
		APIBase:  server.URL,
	}

	err := notifier.Notify(context.Background(), expectedMessage)

	assert.NoError(t, err)
	assert.Equal(t, expectedMessage, receivedPayload["text"])
	assert.Equal(t, expectedChatID, receivedPayload["chat_id"])
}

func TestTelegramNotifier_Notify_MissingBotToken(t *testing.T) {
	notifier := TelegramNotifier{
		BotToken: "", // Missing
		ChatID:   "123456789",
	}

	err := notifier.Notify(context.Background(), "test message")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "token and chat ID required")
}

func TestTelegramNotifier_Notify_MissingChatID(t *testing.T) {
	notifier := TelegramNotifier{
		BotToken: "bot-token-123",
		ChatID:   "", // Missing
	}

	err := notifier.Notify(context.Background(), "test message")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "token and chat ID required")
}

func TestTelegramNotifier_Notify_HTTPError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer server.Close()

	notifier := TelegramNotifier{
		BotToken: "bot-token-123",
		ChatID:   "123456789",
		Client:   server.Client(),
		APIBase:  server.URL,
	}

	err := notifier.Notify(context.Background(), "test message")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "status 400")
}

func TestTelegramNotifier_Notify_NetworkError(t *testing.T) {
	notifier := TelegramNotifier{
		BotToken: "bot-token-123",
		ChatID:   "123456789",
		APIBase:  "http://invalid-url-that-does-not-exist.local",
	}

	err := notifier.Notify(context.Background(), "test message")

	assert.Error(t, err)
}

func TestTelegramNotifier_Notify_ContextCanceled(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// This should not be reached due to context cancellation
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	notifier := TelegramNotifier{
		BotToken: "bot-token-123",
		ChatID:   "123456789",
		Client:   server.Client(),
		APIBase:  server.URL,
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	err := notifier.Notify(ctx, "test message")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "context canceled")
}

func TestTelegramNotifier_Notify_DefaultClient(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	notifier := TelegramNotifier{
		BotToken: "bot-token-123",
		ChatID:   "123456789",
		Client:   nil, // Use default client
		APIBase:  server.URL,
	}

	// This should work with the default client
	assert.NotPanics(t, func() {
		notifier.Notify(context.Background(), "test message")
	})
}

func TestTelegramNotifier_Notify_DefaultAPIBase(t *testing.T) {
	notifier := TelegramNotifier{
		BotToken: "bot-token-123",
		ChatID:   "123456789",
		APIBase:  "", // Use default
	}

	// This will likely fail with a network error, but shouldn't panic
	// and should use the default Telegram API base URL
	assert.NotPanics(t, func() {
		notifier.Notify(context.Background(), "test message")
	})
}

func TestTelegramNotifier_Notify_LongMessage(t *testing.T) {
	longMessage := make([]byte, 4000) // 4KB message
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

	notifier := TelegramNotifier{
		BotToken: "bot-token-123",
		ChatID:   "123456789",
		Client:   server.Client(),
		APIBase:  server.URL,
	}

	err := notifier.Notify(context.Background(), string(longMessage))

	assert.NoError(t, err)
	assert.Equal(t, string(longMessage), receivedPayload["text"])
}

func TestTelegramNotifier_Notify_EmptyMessage(t *testing.T) {
	var receivedPayload map[string]string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		require.NoError(t, err)

		err = json.Unmarshal(body, &receivedPayload)
		require.NoError(t, err)

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	notifier := TelegramNotifier{
		BotToken: "bot-token-123",
		ChatID:   "123456789",
		Client:   server.Client(),
		APIBase:  server.URL,
	}

	err := notifier.Notify(context.Background(), "")

	assert.NoError(t, err)
	assert.Equal(t, "", receivedPayload["text"])
	assert.Equal(t, "123456789", receivedPayload["chat_id"])
}

func TestTelegramNotifier_Notify_SpecialCharacters(t *testing.T) {
	specialMessage := "Hello! 🎉 Testing with emojis and special chars: @#$%^&*()"
	var receivedPayload map[string]string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		require.NoError(t, err)

		err = json.Unmarshal(body, &receivedPayload)
		require.NoError(t, err)

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	notifier := TelegramNotifier{
		BotToken: "bot-token-123",
		ChatID:   "123456789",
		Client:   server.Client(),
		APIBase:  server.URL,
	}

	err := notifier.Notify(context.Background(), specialMessage)

	assert.NoError(t, err)
	assert.Equal(t, specialMessage, receivedPayload["text"])
}
