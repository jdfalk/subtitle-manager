// file: pkg/webhooks/manager_test.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174004

package webhooks

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// TestWebhookManager_AddOutgoingEndpoint verifies endpoint addition and validation.
func TestWebhookManager_AddOutgoingEndpoint(t *testing.T) {
	wm := NewWebhookManager()

	endpoint := OutgoingEndpoint{
		Name:    "Test Endpoint",
		URL:     "https://example.com/webhook",
		Events:  []string{EventSubtitleDownloaded},
		Headers: map[string]string{"Authorization": "Bearer token"},
	}

	err := wm.AddOutgoingEndpoint(endpoint)
	if err != nil {
		t.Fatalf("Failed to add valid endpoint: %v", err)
	}

	endpoints := wm.GetOutgoingEndpoints()
	if len(endpoints) != 1 {
		t.Fatalf("Expected 1 endpoint, got %d", len(endpoints))
	}

	if endpoints[0].Name != "Test Endpoint" {
		t.Fatalf("Expected name 'Test Endpoint', got %s", endpoints[0].Name)
	}
}

// TestWebhookManager_AddInvalidEndpoint verifies URL validation.
func TestWebhookManager_AddInvalidEndpoint(t *testing.T) {
	wm := NewWebhookManager()

	invalidEndpoint := OutgoingEndpoint{
		Name: "Invalid Endpoint",
		URL:  "http://example.com/webhook", // HTTP not allowed
	}

	err := wm.AddOutgoingEndpoint(invalidEndpoint)
	if err == nil {
		t.Fatal("Expected error for HTTP URL, got nil")
	}

	endpoints := wm.GetOutgoingEndpoints()
	if len(endpoints) != 0 {
		t.Fatalf("Expected 0 endpoints after invalid addition, got %d", len(endpoints))
	}
}

// TestWebhookManager_SendEvent verifies event sending to subscribed endpoints.
func TestWebhookManager_SendEvent(t *testing.T) {
	// Create test server to receive webhooks
	var receivedEvents []WebhookEvent
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var event WebhookEvent
		if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		receivedEvents = append(receivedEvents, event)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	wm := NewWebhookManager()
	wm.testMode = true // Enable synchronous mode for testing
	wm.client = server.Client() // Use test client that accepts self-signed certs

	// Add endpoint subscribed to subtitle events - directly add to bypass validation
	endpoint := OutgoingEndpoint{
		ID:         "test-1",
		Name:       "Test Endpoint",
		URL:        server.URL,
		Events:     []string{EventSubtitleDownloaded},
		Enabled:    true,
		MaxRetries: 3,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	wm.mu.Lock()
	wm.outgoing = append(wm.outgoing, endpoint)
	wm.mu.Unlock()

	// Send event
	event := WebhookEvent{
		Type: EventSubtitleDownloaded,
		Data: SubtitleDownloadedData{
			FilePath:  "/test/movie.mkv",
			Language:  "en",
			Provider:  "opensubtitles",
			Timestamp: time.Now(),
		},
	}

	err := wm.SendEvent(context.Background(), event)
	if err != nil {
		t.Fatalf("Failed to send event: %v", err)
	}

	// No sleep needed in synchronous mode

	if len(receivedEvents) != 1 {
		t.Fatalf("Expected 1 received event, got %d", len(receivedEvents))
	}

	if receivedEvents[0].Type != EventSubtitleDownloaded {
		t.Fatalf("Expected event type %s, got %s", EventSubtitleDownloaded, receivedEvents[0].Type)
	}
}

// TestWebhookManager_EventFiltering verifies events are only sent to subscribed endpoints.
func TestWebhookManager_EventFiltering(t *testing.T) {
	var receivedEvents []WebhookEvent
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var event WebhookEvent
		json.NewDecoder(r.Body).Decode(&event)
		receivedEvents = append(receivedEvents, event)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	wm := NewWebhookManager()
	wm.testMode = true // Enable synchronous mode for testing
	wm.client = server.Client()

	// Add endpoint subscribed only to download events - bypass validation
	endpoint := OutgoingEndpoint{
		ID:         "test-1",
		Name:       "Download Only",
		URL:        server.URL,
		Events:     []string{EventSubtitleDownloaded},
		Enabled:    true,
		MaxRetries: 3,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	wm.mu.Lock()
	wm.outgoing = append(wm.outgoing, endpoint)
	wm.mu.Unlock()

	// Send download event (should be received)
	downloadEvent := WebhookEvent{
		Type: EventSubtitleDownloaded,
		Data: SubtitleDownloadedData{FilePath: "/test/movie.mkv"},
	}
	wm.SendEvent(context.Background(), downloadEvent)

	// Send system event (should NOT be received)
	systemEvent := WebhookEvent{
		Type: EventSystemStarted,
		Data: SystemEventData{Event: "started"},
	}
	wm.SendEvent(context.Background(), systemEvent)

	// No sleep needed in synchronous mode

	if len(receivedEvents) != 1 {
		t.Fatalf("Expected 1 received event, got %d", len(receivedEvents))
	}

	if receivedEvents[0].Type != EventSubtitleDownloaded {
		t.Fatalf("Expected download event, got %s", receivedEvents[0].Type)
	}
}

// TestWebhookManager_RateLimit verifies rate limiting functionality.
func TestWebhookManager_RateLimit(t *testing.T) {
	wm := NewWebhookManager()

	// Test rate limiting for same IP
	remoteAddr := "192.168.1.100:12345"

	// First 10 requests should pass
	for i := 0; i < 10; i++ {
		if !wm.checkRateLimit(remoteAddr) {
			t.Fatalf("Request %d should have passed rate limit", i+1)
		}
	}

	// 11th request should be blocked
	if wm.checkRateLimit(remoteAddr) {
		t.Fatal("Request should have been rate limited")
	}
}

// TestWebhookManager_IPWhitelist verifies IP whitelisting functionality.
func TestWebhookManager_IPWhitelist(t *testing.T) {
	wm := NewWebhookManager()

	// Set whitelist
	wm.SetIPWhitelist([]string{"192.168.1.100", "10.0.0.0/8"})

	tests := []struct {
		ip       string
		expected bool
	}{
		{"192.168.1.100:12345", true},  // Exact match
		{"10.0.0.50:12345", true},      // CIDR match
		{"192.168.1.101:12345", false}, // Not whitelisted
		{"172.16.0.1:12345", false},    // Not whitelisted
	}

	for _, test := range tests {
		result := wm.isIPWhitelisted(test.ip)
		if result != test.expected {
			t.Errorf("IP %s: expected %v, got %v", test.ip, test.expected, result)
		}
	}
}

// TestWebhookManager_HandleIncomingWebhook verifies incoming webhook processing.
func TestWebhookManager_HandleIncomingWebhook(t *testing.T) {
	wm := NewWebhookManager()

	// Register test handler
	handler := &TestIncomingHandler{}
	wm.RegisterIncomingHandler("test", handler)

	// Create test request
	payload := `{"path": "/test/movie.mkv", "lang": "en"}`
	req := httptest.NewRequest("POST", "/webhook", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	req.RemoteAddr = "192.168.1.100:12345"

	err := wm.HandleIncomingWebhook("test", req)
	if err != nil {
		t.Fatalf("Failed to handle incoming webhook: %v", err)
	}

	if !handler.Called {
		t.Fatal("Handler was not called")
	}
}

// TestIncomingHandler is a test implementation of IncomingHandler.
type TestIncomingHandler struct {
	Called bool
}

func (h *TestIncomingHandler) Handle(ctx context.Context, payload []byte, headers http.Header) error {
	h.Called = true
	return nil
}

func (h *TestIncomingHandler) ValidateSignature(payload []byte, signature string) bool {
	return true
}

func (h *TestIncomingHandler) GetName() string {
	return "test"
}

// TestRateLimiter_Allow verifies rate limiter token bucket behavior.
func TestRateLimiter_Allow(t *testing.T) {
	rl := &RateLimiter{
		tokens:    5,
		maxTokens: 5,
		refillAt:  time.Now().Add(time.Minute),
		interval:  time.Minute,
	}

	// Use all tokens
	for i := 0; i < 5; i++ {
		if !rl.Allow() {
			t.Fatalf("Token %d should have been allowed", i+1)
		}
	}

	// Next request should be denied
	if rl.Allow() {
		t.Fatal("Request should have been denied - no tokens left")
	}

	// Set refill time to past to trigger refill
	rl.refillAt = time.Now().Add(-time.Second)

	// Should allow requests again after refill
	if !rl.Allow() {
		t.Fatal("Request should have been allowed after refill")
	}
}

// TestHMACSignature verifies HMAC signature generation and validation.
func TestHMACSignature(t *testing.T) {
	payload := []byte("test payload")
	secret := "test-secret"

	signature := generateHMACSignature(payload, secret)
	if signature == "" {
		t.Fatal("Signature should not be empty")
	}

	// Verify signature validates correctly
	handler := NewCustomHandler(secret)
	if !handler.ValidateSignature(payload, signature) {
		t.Fatal("Signature validation should pass")
	}

	// Verify wrong signature fails
	if handler.ValidateSignature(payload, "wrong-signature") {
		t.Fatal("Wrong signature should fail validation")
	}
}

// TestWebhookManager_EventHistory verifies event history tracking.
func TestWebhookManager_EventHistory(t *testing.T) {
	wm := NewWebhookManager()

	// Send multiple events
	for i := 0; i < 5; i++ {
		event := WebhookEvent{
			Type: EventSubtitleDownloaded,
			Data: SubtitleDownloadedData{FilePath: "/test/movie.mkv"},
		}
		wm.SendEvent(context.Background(), event)
	}

	history := wm.GetEventHistory(10)
	if len(history) != 5 {
		t.Fatalf("Expected 5 events in history, got %d", len(history))
	}

	// Test history limit
	limitedHistory := wm.GetEventHistory(3)
	if len(limitedHistory) != 3 {
		t.Fatalf("Expected 3 events with limit, got %d", len(limitedHistory))
	}
}