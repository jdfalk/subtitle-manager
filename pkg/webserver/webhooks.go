// file: pkg/webserver/webhooks.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174005

package webserver

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jdfalk/subtitle-manager/pkg/webhooks"
)

// webhookConfigHandler handles webhook configuration requests.
func webhookConfigHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handleGetWebhookConfig(w, r)
		case http.MethodPost:
			handleCreateWebhook(w, r)
		case http.MethodPut:
			handleUpdateWebhook(w, r)
		case http.MethodDelete:
			handleDeleteWebhook(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}

// webhookTestHandler handles webhook testing requests.
func webhookTestHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		handleTestWebhook(w, r)
	})
}

// webhookHistoryHandler handles webhook history requests.
func webhookHistoryHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		handleGetWebhookHistory(w, r)
	})
}

// webhookEventTypesHandler returns available event types.
func webhookEventTypesHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		handleGetEventTypes(w, r)
	})
}

// handleGetWebhookConfig returns all configured webhook endpoints.
func handleGetWebhookConfig(w http.ResponseWriter, r *http.Request) {
	manager := webhooks.GetGlobalManager()
	endpoints := manager.GetOutgoingEndpoints()

	// Remove secrets from response for security
	for i := range endpoints {
		endpoints[i].Secret = ""
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"endpoints": endpoints,
	})
}

// CreateWebhookRequest represents a request to create a new webhook endpoint.
type CreateWebhookRequest struct {
	Name    string            `json:"name"`
	URL     string            `json:"url"`
	Secret  string            `json:"secret,omitempty"`
	Events  []string          `json:"events"`
	Headers map[string]string `json:"headers,omitempty"`
}

// handleCreateWebhook creates a new webhook endpoint with input validation.
// CodeQL: All inputs validated - name/URL required, events validated against whitelist
func handleCreateWebhook(w http.ResponseWriter, r *http.Request) {
	var req CreateWebhookRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate required fields to prevent empty/malformed requests
	if req.Name == "" || req.URL == "" {
		http.Error(w, "Name and URL are required", http.StatusBadRequest)
		return
	}

	if len(req.Events) == 0 {
		http.Error(w, "At least one event type must be selected", http.StatusBadRequest)
		return
	}

	// Validate event types against whitelist to prevent injection
	// CodeQL: Event types validated against predefined list in GetAvailableEventTypes()
	validEvents := webhooks.GetAvailableEventTypes()
	for _, event := range req.Events {
		if event != "*" && !contains(validEvents, event) {
			http.Error(w, "Invalid event type: "+event, http.StatusBadRequest)
			return
		}
	}

	// Create endpoint with validated inputs
	// CodeQL: URL will be validated in AddOutgoingEndpoint() via validateWebhookURL()
	endpoint := webhooks.OutgoingEndpoint{
		Name:    req.Name,
		URL:     req.URL,
		Secret:  req.Secret,
		Events:  req.Events,
		Headers: req.Headers,
	}

	manager := webhooks.GetGlobalManager()
	// URL validation and SSRF prevention occurs in AddOutgoingEndpoint()
	if err := manager.AddOutgoingEndpoint(endpoint); err != nil {
		http.Error(w, "Failed to create webhook: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Webhook endpoint created successfully",
	})
}

// UpdateWebhookRequest represents a request to update a webhook endpoint.
type UpdateWebhookRequest struct {
	ID      string            `json:"id"`
	Name    string            `json:"name,omitempty"`
	URL     string            `json:"url,omitempty"`
	Secret  string            `json:"secret,omitempty"`
	Events  []string          `json:"events,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
	Enabled *bool             `json:"enabled,omitempty"`
}

// handleUpdateWebhook updates an existing webhook endpoint.
func handleUpdateWebhook(w http.ResponseWriter, r *http.Request) {
	var req UpdateWebhookRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.ID == "" {
		http.Error(w, "Webhook ID is required", http.StatusBadRequest)
		return
	}

	manager := webhooks.GetGlobalManager()
	endpoints := manager.GetOutgoingEndpoints()

	// Find the endpoint to update
	var targetIndex = -1
	for i, endpoint := range endpoints {
		if endpoint.ID == req.ID {
			targetIndex = i
			break
		}
	}

	if targetIndex == -1 {
		http.Error(w, "Webhook not found", http.StatusNotFound)
		return
	}

	// Update the endpoint
	manager.UpdateEndpoint(req.ID, func(endpoint *webhooks.OutgoingEndpoint) {
		if req.Name != "" {
			endpoint.Name = req.Name
		}
		if req.URL != "" {
			endpoint.URL = req.URL
		}
		if req.Secret != "" {
			endpoint.Secret = req.Secret
		}
		if len(req.Events) > 0 {
			endpoint.Events = req.Events
		}
		if req.Headers != nil {
			endpoint.Headers = req.Headers
		}
		if req.Enabled != nil {
			endpoint.Enabled = *req.Enabled
		}
	})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Webhook endpoint updated successfully",
	})
}

// handleDeleteWebhook deletes a webhook endpoint.
func handleDeleteWebhook(w http.ResponseWriter, r *http.Request) {
	webhookID := r.URL.Query().Get("id")
	if webhookID == "" {
		http.Error(w, "Webhook ID is required", http.StatusBadRequest)
		return
	}

	manager := webhooks.GetGlobalManager()
	if err := manager.RemoveEndpoint(webhookID); err != nil {
		http.Error(w, "Failed to delete webhook: "+err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Webhook endpoint deleted successfully",
	})
}

// TestWebhookRequest represents a request to test a webhook endpoint.
type TestWebhookRequest struct {
	URL     string            `json:"url"`
	Secret  string            `json:"secret,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
}

// handleTestWebhook sends a test event to a webhook endpoint.
func handleTestWebhook(w http.ResponseWriter, r *http.Request) {
	var req TestWebhookRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.URL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	// Create a test event
	testEvent := webhooks.WebhookEvent{
		Type: "test",
		Data: map[string]interface{}{
			"message":   "This is a test webhook from subtitle-manager",
			"timestamp": "2024-01-01T00:00:00Z",
			"test":      true,
		},
		Source: "subtitle-manager-test",
	}

	// Create a temporary endpoint for testing
	testEndpoint := webhooks.OutgoingEndpoint{
		ID:      "test",
		Name:    "Test Webhook",
		URL:     req.URL,
		Secret:  req.Secret,
		Events:  []string{"test"},
		Headers: req.Headers,
		Enabled: true,
	}

	manager := webhooks.GetGlobalManager()
	manager.TestEndpoint(context.Background(), &testEndpoint, testEvent)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Test webhook sent successfully",
	})
}

// handleGetWebhookHistory returns recent webhook events.
func handleGetWebhookHistory(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	limit := 50 // Default limit
	if limitStr != "" {
		if parsedLimit, err := strconv.Atoi(limitStr); err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	manager := webhooks.GetGlobalManager()
	history := manager.GetEventHistory(limit)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"events": history,
		"total":  len(history),
	})
}

// handleGetEventTypes returns available webhook event types.
func handleGetEventTypes(w http.ResponseWriter, r *http.Request) {
	eventTypes := webhooks.GetAvailableEventTypes()

	// Add descriptions for each event type
	eventTypeInfo := make([]map[string]interface{}, len(eventTypes))
	descriptions := map[string]string{
		webhooks.EventSubtitleDownloaded: "Triggered when a subtitle is successfully downloaded",
		webhooks.EventSubtitleUpgraded:   "Triggered when a subtitle is upgraded to a better version",
		webhooks.EventSubtitleFailed:     "Triggered when subtitle download fails",
		webhooks.EventSearchFailed:       "Triggered when subtitle search fails",
		webhooks.EventSystemStarted:      "Triggered when the system starts",
		webhooks.EventSystemStopped:      "Triggered when the system stops",
		webhooks.EventSystemError:        "Triggered when a system error occurs",
		webhooks.EventCustom:             "Custom events from external integrations",
	}

	for i, eventType := range eventTypes {
		eventTypeInfo[i] = map[string]interface{}{
			"type":        eventType,
			"description": descriptions[eventType],
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"event_types": eventTypeInfo,
	})
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
