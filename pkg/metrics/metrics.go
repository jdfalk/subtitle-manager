//go:build !gcommonmetrics

// file: pkg/metrics/metrics.go
// version: 1.2.0
// guid: a1b2c3d4-e5f6-7g8h-9i0j-k1l2m3n4o5p6

package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// ProviderRequests counts the number of requests made to subtitle providers.
	ProviderRequests *prometheus.CounterVec

	// TranslationRequests counts the number of translation requests processed.
	TranslationRequests *prometheus.CounterVec

	// APIRequests counts the number of API requests by endpoint.
	APIRequests *prometheus.CounterVec

	// RequestDuration tracks the duration of API requests.
	RequestDuration *prometheus.HistogramVec

	// ActiveSessions tracks the number of active user sessions.
	ActiveSessions prometheus.Gauge

	// SubtitleDownloads counts successful subtitle downloads.
	SubtitleDownloads *prometheus.CounterVec
)

// Initialize configures the metrics provider and registers application metrics.
func Initialize() error {
	if ProviderRequests != nil {
		return nil
	}

	ProviderRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "subtitle_manager",
			Name:      "provider_requests_total",
			Help:      "The total number of requests made to subtitle providers",
		},
		[]string{"provider", "status"},
	)

	TranslationRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "subtitle_manager",
			Name:      "translation_requests_total",
			Help:      "The total number of translation requests processed",
		},
		[]string{"service", "target_language", "status"},
	)

	APIRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "subtitle_manager",
			Name:      "api_requests_total",
			Help:      "The total number of API requests by endpoint",
		},
		[]string{"endpoint", "method", "status_code"},
	)

	RequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "subtitle_manager",
			Name:      "request_duration_seconds",
			Help:      "The duration of API requests in seconds",
		},
		[]string{"endpoint", "method"},
	)

	ActiveSessions = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "subtitle_manager",
		Name:      "active_sessions",
		Help:      "The number of active user sessions",
	})

	SubtitleDownloads = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "subtitle_manager",
			Name:      "subtitle_downloads_total",
			Help:      "The total number of subtitle downloads",
		},
		[]string{"provider", "language"},
	)

	prometheus.MustRegister(
		ProviderRequests,
		TranslationRequests,
		APIRequests,
		RequestDuration,
		ActiveSessions,
		SubtitleDownloads,
	)

	return nil
}
