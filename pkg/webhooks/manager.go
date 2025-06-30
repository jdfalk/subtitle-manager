// file: pkg/webhooks/manager.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174001

package webhooks

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/jdfalk/subtitle-manager/pkg/logging"
)

// IncomingHandler processes incoming webhook events from external services.
type IncomingHandler interface {
	Handle(ctx context.Context, payload []byte, headers http.Header) error
	ValidateSignature(payload []byte, signature string) bool
	GetName() string
}

// OutgoingEndpoint represents a configured webhook destination.
type OutgoingEndpoint struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	URL         string            `json:"url"`
	Secret      string            `json:"secret,omitempty"`
	Events      []string          `json:"events"`
	Headers     map[string]string `json:"headers,omitempty"`
	Enabled     bool              `json:"enabled"`
	RetryCount  int               `json:"retry_count"`
	MaxRetries  int               `json:"max_retries"`
	LastAttempt time.Time         `json:"last_attempt"`
	LastSuccess time.Time         `json:"last_success"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

// WebhookEvent represents an event that can be sent via webhooks.
type WebhookEvent struct {
	Type      string      `json:"type"`
	Timestamp time.Time   `json:"timestamp"`
	Data      interface{} `json:"data"`
	Source    string      `json:"source,omitempty"`
	ID        string      `json:"id,omitempty"`
}

// WebhookManager handles both incoming and outgoing webhook operations.
type WebhookManager struct {
	incoming     map[string]IncomingHandler
	outgoing     []OutgoingEndpoint
	eventHistory []WebhookEvent
	rateLimiter  map[string]*RateLimiter
	whitelist    []string
	mu           sync.RWMutex
	client       *http.Client
	logger       *logrus.Entry
	testMode     bool // For testing - makes sending synchronous
}

// RateLimiter implements token bucket rate limiting for webhook requests.
type RateLimiter struct {
	tokens    int
	maxTokens int
	refillAt  time.Time
	interval  time.Duration
	mu        sync.Mutex
}

// NewWebhookManager creates a new webhook manager instance.
func NewWebhookManager() *WebhookManager {
	return &WebhookManager{
		incoming:     make(map[string]IncomingHandler),
		outgoing:     make([]OutgoingEndpoint, 0),
		eventHistory: make([]WebhookEvent, 0),
		rateLimiter:  make(map[string]*RateLimiter),
		whitelist:    make([]string, 0),
		client:       &http.Client{Timeout: 30 * time.Second},
		logger:       logging.GetLogger("webhook-manager"),
	}
}

// RegisterIncomingHandler registers a handler for incoming webhooks from a specific source.
func (wm *WebhookManager) RegisterIncomingHandler(source string, handler IncomingHandler) {
	wm.mu.Lock()
	defer wm.mu.Unlock()
	wm.incoming[source] = handler
	wm.logger.Infof("Registered incoming webhook handler for %s", source)
}

// AddOutgoingEndpoint adds a new outgoing webhook endpoint.
func (wm *WebhookManager) AddOutgoingEndpoint(endpoint OutgoingEndpoint) error {
	if err := validateWebhookURL(endpoint.URL); err != nil {
		return fmt.Errorf("invalid webhook URL: %v", err)
	}

	wm.mu.Lock()
	defer wm.mu.Unlock()
	
	endpoint.ID = generateID()
	endpoint.CreatedAt = time.Now()
	endpoint.UpdatedAt = time.Now()
	endpoint.Enabled = true
	endpoint.MaxRetries = 3
	
	wm.outgoing = append(wm.outgoing, endpoint)
	wm.logger.Infof("Added outgoing webhook endpoint: %s (%s)", endpoint.Name, endpoint.URL)
	return nil
}

// HandleIncomingWebhook processes an incoming webhook request.
func (wm *WebhookManager) HandleIncomingWebhook(source string, r *http.Request) error {
	// Check IP whitelist if configured
	if len(wm.whitelist) > 0 && !wm.isIPWhitelisted(r.RemoteAddr) {
		wm.logger.Warnf("Webhook request from non-whitelisted IP: %s", r.RemoteAddr)
		return fmt.Errorf("IP not whitelisted")
	}

	// Apply rate limiting
	if !wm.checkRateLimit(r.RemoteAddr) {
		wm.logger.Warnf("Rate limit exceeded for IP: %s", r.RemoteAddr)
		return fmt.Errorf("rate limit exceeded")
	}

	// Check payload size limit (1MB default)
	if r.ContentLength > 1024*1024 {
		wm.logger.Warnf("Payload too large: %d bytes", r.ContentLength)
		return fmt.Errorf("payload too large")
	}

	wm.mu.RLock()
	handler, exists := wm.incoming[source]
	wm.mu.RUnlock()

	if !exists {
		return fmt.Errorf("no handler registered for source: %s", source)
	}

	// Read and validate payload
	payload := make([]byte, r.ContentLength)
	if _, err := r.Body.Read(payload); err != nil {
		return fmt.Errorf("failed to read payload: %v", err)
	}

	// Validate HMAC signature if present
	if signature := r.Header.Get("X-Hub-Signature-256"); signature != "" {
		if !handler.ValidateSignature(payload, signature) {
			wm.logger.Warnf("Invalid signature for webhook from %s", source)
			return fmt.Errorf("invalid signature")
		}
	}

	return handler.Handle(r.Context(), payload, r.Header)
}

// SendEvent sends an event to all configured outgoing webhooks that subscribe to the event type.
func (wm *WebhookManager) SendEvent(ctx context.Context, event WebhookEvent) error {
	wm.mu.RLock()
	defer wm.mu.RUnlock()

	// Add to event history
	event.ID = generateID()
	event.Timestamp = time.Now()
	wm.eventHistory = append(wm.eventHistory, event)

	// Trim history to last 1000 events
	if len(wm.eventHistory) > 1000 {
		wm.eventHistory = wm.eventHistory[len(wm.eventHistory)-1000:]
	}

	var errs []error
	for i := range wm.outgoing {
		endpoint := &wm.outgoing[i]
		if !endpoint.Enabled {
			continue
		}

		// Check if endpoint subscribes to this event type
		if !contains(endpoint.Events, event.Type) && !contains(endpoint.Events, "*") {
			continue
		}

		if wm.testMode {
			// Synchronous for testing
			wm.sendToEndpoint(ctx, endpoint, event)
		} else {
			// Asynchronous for production
			go wm.sendToEndpoint(ctx, endpoint, event)
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("failed to send to some endpoints: %v", errs)
	}

	return nil
}

// sendToEndpoint sends an event to a specific endpoint with retry logic.
func (wm *WebhookManager) sendToEndpoint(ctx context.Context, endpoint *OutgoingEndpoint, event WebhookEvent) {
	payload, err := json.Marshal(event)
	if err != nil {
		wm.logger.Errorf("Failed to marshal event for endpoint %s: %v", endpoint.Name, err)
		return
	}

	for attempt := 0; attempt <= endpoint.MaxRetries; attempt++ {
		if attempt > 0 {
			// Exponential backoff: 2^attempt seconds
			delay := time.Duration(1<<uint(attempt)) * time.Second
			select {
			case <-ctx.Done():
				return
			case <-time.After(delay):
			}
		}

		endpoint.LastAttempt = time.Now()
		endpoint.RetryCount = attempt

		if err := wm.doHTTPRequest(ctx, endpoint, payload); err != nil {
			wm.logger.Warnf("Attempt %d failed for endpoint %s: %v", attempt+1, endpoint.Name, err)
			if attempt == endpoint.MaxRetries {
				wm.logger.Errorf("All retry attempts failed for endpoint %s", endpoint.Name)
			}
			continue
		}

		endpoint.LastSuccess = time.Now()
		endpoint.RetryCount = 0
		wm.logger.Debugf("Successfully sent event to endpoint %s", endpoint.Name)
		return
	}
}

// doHTTPRequest performs the actual HTTP request to the webhook endpoint.
func (wm *WebhookManager) doHTTPRequest(ctx context.Context, endpoint *OutgoingEndpoint, payload []byte) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint.URL, strings.NewReader(string(payload)))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "subtitle-manager-webhook/1.0")

	// Add custom headers
	for key, value := range endpoint.Headers {
		req.Header.Set(key, value)
	}

	// Add HMAC signature if secret is configured
	if endpoint.Secret != "" {
		signature := generateHMACSignature(payload, endpoint.Secret)
		req.Header.Set("X-Hub-Signature-256", "sha256="+signature)
	}

	resp, err := wm.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("webhook returned status %d", resp.StatusCode)
	}

	return nil
}

// GetEventHistory returns the recent webhook event history.
func (wm *WebhookManager) GetEventHistory(limit int) []WebhookEvent {
	wm.mu.RLock()
	defer wm.mu.RUnlock()

	if limit <= 0 || limit > len(wm.eventHistory) {
		limit = len(wm.eventHistory)
	}

	start := len(wm.eventHistory) - limit
	result := make([]WebhookEvent, limit)
	copy(result, wm.eventHistory[start:])
	return result
}

// GetOutgoingEndpoints returns all configured outgoing webhook endpoints.
func (wm *WebhookManager) GetOutgoingEndpoints() []OutgoingEndpoint {
	wm.mu.RLock()
	defer wm.mu.RUnlock()

	result := make([]OutgoingEndpoint, len(wm.outgoing))
	copy(result, wm.outgoing)
	return result
}

// SetIPWhitelist configures the IP whitelist for incoming webhooks.
func (wm *WebhookManager) SetIPWhitelist(ips []string) {
	wm.mu.Lock()
	defer wm.mu.Unlock()
	wm.whitelist = make([]string, len(ips))
	copy(wm.whitelist, ips)
	wm.logger.Infof("Updated IP whitelist with %d entries", len(ips))
}

// isIPWhitelisted checks if an IP address is in the whitelist.
func (wm *WebhookManager) isIPWhitelisted(remoteAddr string) bool {
	ip, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		ip = remoteAddr
	}

	for _, allowedIP := range wm.whitelist {
		if ip == allowedIP {
			return true
		}
		// Support CIDR notation
		if strings.Contains(allowedIP, "/") {
			if _, cidr, err := net.ParseCIDR(allowedIP); err == nil {
				if clientIP := net.ParseIP(ip); clientIP != nil {
					if cidr.Contains(clientIP) {
						return true
					}
				}
			}
		}
	}
	return false
}

// checkRateLimit applies rate limiting per IP address.
func (wm *WebhookManager) checkRateLimit(remoteAddr string) bool {
	ip, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		ip = remoteAddr
	}

	wm.mu.Lock()
	defer wm.mu.Unlock()

	limiter, exists := wm.rateLimiter[ip]
	if !exists {
		limiter = &RateLimiter{
			tokens:    10,
			maxTokens: 10,
			refillAt:  time.Now().Add(time.Minute),
			interval:  time.Minute,
		}
		wm.rateLimiter[ip] = limiter
	}

	return limiter.Allow()
}

// Allow checks if a request is allowed under the rate limit.
func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	if now.After(rl.refillAt) {
		rl.tokens = rl.maxTokens
		rl.refillAt = now.Add(rl.interval)
	}

	if rl.tokens > 0 {
		rl.tokens--
		return true
	}

	return false
}

// generateHMACSignature creates an HMAC-SHA256 signature for webhook validation.
func generateHMACSignature(payload []byte, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write(payload)
	return hex.EncodeToString(h.Sum(nil))
}

// generateID creates a simple unique identifier.
func generateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

// contains checks if a slice contains a specific string.
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}