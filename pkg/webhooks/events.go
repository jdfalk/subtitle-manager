// file: pkg/webhooks/events.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174003

package webhooks

import (
	"context"
	"time"

	"github.com/jdfalk/subtitle-manager/pkg/logging"
)

// Global webhook manager instance
var globalManager *WebhookManager

// Event type constants for outgoing webhooks
const (
	EventSubtitleDownloaded = "subtitle.downloaded"
	EventSubtitleUpgraded   = "subtitle.upgraded"
	EventSubtitleFailed     = "subtitle.failed"
	EventSearchFailed       = "search.failed"
	EventSystemStarted      = "system.started"
	EventSystemStopped      = "system.stopped"
	EventSystemError        = "system.error"
	EventCustom             = "custom"
)

// SubtitleDownloadedData represents data for subtitle download events.
type SubtitleDownloadedData struct {
	FilePath     string    `json:"file_path"`
	SubtitlePath string    `json:"subtitle_path"`
	Language     string    `json:"language"`
	Provider     string    `json:"provider"`
	Score        float64   `json:"score"`
	Size         int64     `json:"size"`
	Timestamp    time.Time `json:"timestamp"`
}

// SubtitleUpgradedData represents data for subtitle upgrade events.
type SubtitleUpgradedData struct {
	FilePath        string    `json:"file_path"`
	OldSubtitlePath string    `json:"old_subtitle_path"`
	NewSubtitlePath string    `json:"new_subtitle_path"`
	Language        string    `json:"language"`
	OldProvider     string    `json:"old_provider"`
	NewProvider     string    `json:"new_provider"`
	OldScore        float64   `json:"old_score"`
	NewScore        float64   `json:"new_score"`
	Timestamp       time.Time `json:"timestamp"`
}

// SubtitleFailedData represents data for subtitle failure events.
type SubtitleFailedData struct {
	FilePath  string    `json:"file_path"`
	Language  string    `json:"language"`
	Provider  string    `json:"provider,omitempty"`
	Error     string    `json:"error"`
	Timestamp time.Time `json:"timestamp"`
}

// SearchFailedData represents data for search failure events.
type SearchFailedData struct {
	Query     string    `json:"query"`
	Language  string    `json:"language"`
	Provider  string    `json:"provider,omitempty"`
	Error     string    `json:"error"`
	Timestamp time.Time `json:"timestamp"`
}

// SystemEventData represents data for system events.
type SystemEventData struct {
	Event     string    `json:"event"`
	Message   string    `json:"message,omitempty"`
	Details   string    `json:"details,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}

// InitializeGlobalManager initializes the global webhook manager.
func InitializeGlobalManager() {
	if globalManager == nil {
		globalManager = NewWebhookManager()
		logger := logging.GetLogger("webhook-events")
		logger.Info("Initialized global webhook manager")
	}
}

// GetGlobalManager returns the global webhook manager instance.
func GetGlobalManager() *WebhookManager {
	if globalManager == nil {
		InitializeGlobalManager()
	}
	return globalManager
}

// SendSubtitleDownloadedEvent sends a subtitle downloaded event.
func SendSubtitleDownloadedEvent(ctx context.Context, data SubtitleDownloadedData) error {
	event := WebhookEvent{
		Type:      EventSubtitleDownloaded,
		Timestamp: time.Now(),
		Data:      data,
		Source:    "subtitle-manager",
	}
	return GetGlobalManager().SendEvent(ctx, event)
}

// SendSubtitleUpgradedEvent sends a subtitle upgraded event.
func SendSubtitleUpgradedEvent(ctx context.Context, data SubtitleUpgradedData) error {
	event := WebhookEvent{
		Type:      EventSubtitleUpgraded,
		Timestamp: time.Now(),
		Data:      data,
		Source:    "subtitle-manager",
	}
	return GetGlobalManager().SendEvent(ctx, event)
}

// SendSubtitleFailedEvent sends a subtitle failed event.
func SendSubtitleFailedEvent(ctx context.Context, data SubtitleFailedData) error {
	event := WebhookEvent{
		Type:      EventSubtitleFailed,
		Timestamp: time.Now(),
		Data:      data,
		Source:    "subtitle-manager",
	}
	return GetGlobalManager().SendEvent(ctx, event)
}

// SendSearchFailedEvent sends a search failed event.
func SendSearchFailedEvent(ctx context.Context, data SearchFailedData) error {
	event := WebhookEvent{
		Type:      EventSearchFailed,
		Timestamp: time.Now(),
		Data:      data,
		Source:    "subtitle-manager",
	}
	return GetGlobalManager().SendEvent(ctx, event)
}

// SendSystemEvent sends a system event.
func SendSystemEvent(ctx context.Context, data SystemEventData) error {
	event := WebhookEvent{
		Type:      EventSystemStarted,
		Timestamp: time.Now(),
		Data:      data,
		Source:    "subtitle-manager",
	}
	if data.Event == "stopped" {
		event.Type = EventSystemStopped
	} else if data.Event == "error" {
		event.Type = EventSystemError
	}
	return GetGlobalManager().SendEvent(ctx, event)
}

// SendCustomEvent sends a custom event.
func SendCustomEvent(ctx context.Context, eventType string, data interface{}) error {
	event := WebhookEvent{
		Type:      eventType,
		Timestamp: time.Now(),
		Data:      data,
		Source:    "subtitle-manager",
	}
	return GetGlobalManager().SendEvent(ctx, event)
}

// GetAvailableEventTypes returns all available event types for webhook configuration.
func GetAvailableEventTypes() []string {
	return []string{
		EventSubtitleDownloaded,
		EventSubtitleUpgraded,
		EventSubtitleFailed,
		EventSearchFailed,
		EventSystemStarted,
		EventSystemStopped,
		EventSystemError,
		EventCustom,
	}
}
