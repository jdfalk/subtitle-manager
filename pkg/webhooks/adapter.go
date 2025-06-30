// file: pkg/webhooks/adapter.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174007

package webhooks

import (
	"context"

	"github.com/jdfalk/subtitle-manager/pkg/events"
)

// WebhookEventPublisher adapts webhook events to the events.EventPublisher interface.
type WebhookEventPublisher struct {
	manager *WebhookManager
}

// NewWebhookEventPublisher creates a new webhook event publisher.
func NewWebhookEventPublisher(manager *WebhookManager) *WebhookEventPublisher {
	return &WebhookEventPublisher{manager: manager}
}

// PublishSubtitleDownloaded publishes a subtitle downloaded event as a webhook.
func (w *WebhookEventPublisher) PublishSubtitleDownloaded(ctx context.Context, data events.SubtitleDownloadedData) {
	event := WebhookEvent{
		Type:      EventSubtitleDownloaded,
		Timestamp: data.Timestamp,
		Data:      SubtitleDownloadedData(data),
		Source:    "subtitle-manager",
	}
	w.manager.SendEvent(ctx, event)
}

// PublishSubtitleUpgraded publishes a subtitle upgraded event as a webhook.
func (w *WebhookEventPublisher) PublishSubtitleUpgraded(ctx context.Context, data events.SubtitleUpgradedData) {
	event := WebhookEvent{
		Type:      EventSubtitleUpgraded,
		Timestamp: data.Timestamp,
		Data:      SubtitleUpgradedData(data),
		Source:    "subtitle-manager",
	}
	w.manager.SendEvent(ctx, event)
}

// PublishSubtitleFailed publishes a subtitle failed event as a webhook.
func (w *WebhookEventPublisher) PublishSubtitleFailed(ctx context.Context, data events.SubtitleFailedData) {
	event := WebhookEvent{
		Type:      EventSubtitleFailed,
		Timestamp: data.Timestamp,
		Data:      SubtitleFailedData(data),
		Source:    "subtitle-manager",
	}
	w.manager.SendEvent(ctx, event)
}

// PublishSearchFailed publishes a search failed event as a webhook.
func (w *WebhookEventPublisher) PublishSearchFailed(ctx context.Context, data events.SearchFailedData) {
	event := WebhookEvent{
		Type:      EventSearchFailed,
		Timestamp: data.Timestamp,
		Data:      SearchFailedData(data),
		Source:    "subtitle-manager",
	}
	w.manager.SendEvent(ctx, event)
}