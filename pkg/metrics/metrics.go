//go:build !gcommonmetrics

// file: pkg/metrics/metrics.go
// version: 1.1.0
// guid: a1b2c3d4-e5f6-7g8h-9i0j-k1l2m3n4o5p6

package metrics

import (
	gmetrics "github.com/jdfalk/gcommon/pkg/metrics"
)

var (
	// Provider is the active metrics provider.
	Provider gmetrics.Provider

	// ProviderRequests counts the number of requests made to subtitle providers.
	ProviderRequests gmetrics.Counter

	// TranslationRequests counts the number of translation requests processed.
	TranslationRequests gmetrics.Counter

	// APIRequests counts the number of API requests by endpoint.
	APIRequests gmetrics.Counter

	// RequestDuration tracks the duration of API requests.
	RequestDuration gmetrics.Histogram

	// ActiveSessions tracks the number of active user sessions.
	ActiveSessions gmetrics.Gauge

	// SubtitleDownloads counts successful subtitle downloads.
	SubtitleDownloads gmetrics.Counter
)

// Initialize configures the metrics provider and registers application metrics.
func Initialize() error {
	if Provider != nil {
		return nil
	}

	cfg := gmetrics.Config{
		Enabled:   true,
		Provider:  "prometheus",
		Namespace: "subtitle_manager",
	}

	p, err := gmetrics.NewProvider(cfg)
	if err != nil {
		return err
	}
	Provider = p

	ProviderRequests = p.Counter("provider_requests_total",
		gmetrics.WithDescription("The total number of requests made to subtitle providers"),
		gmetrics.WithTags(gmetrics.Tag{Key: "provider"}, gmetrics.Tag{Key: "status"}),
	)

	TranslationRequests = p.Counter("translation_requests_total",
		gmetrics.WithDescription("The total number of translation requests processed"),
		gmetrics.WithTags(
			gmetrics.Tag{Key: "service"},
			gmetrics.Tag{Key: "target_language"},
			gmetrics.Tag{Key: "status"},
		),
	)

	APIRequests = p.Counter("api_requests_total",
		gmetrics.WithDescription("The total number of API requests by endpoint"),
		gmetrics.WithTags(
			gmetrics.Tag{Key: "endpoint"},
			gmetrics.Tag{Key: "method"},
			gmetrics.Tag{Key: "status_code"},
		),
	)

	RequestDuration = p.Histogram("request_duration_seconds",
		gmetrics.WithDescription("The duration of API requests in seconds"),
		gmetrics.WithBuckets(gmetrics.DefaultBuckets),
		gmetrics.WithTags(
			gmetrics.Tag{Key: "endpoint"},
			gmetrics.Tag{Key: "method"},
		),
	)

	ActiveSessions = p.Gauge("active_sessions",
		gmetrics.WithDescription("The number of active user sessions"),
	)

	SubtitleDownloads = p.Counter("subtitle_downloads_total",
		gmetrics.WithDescription("The total number of subtitle downloads"),
		gmetrics.WithTags(
			gmetrics.Tag{Key: "provider"},
			gmetrics.Tag{Key: "language"},
		),
	)

	return nil
}
