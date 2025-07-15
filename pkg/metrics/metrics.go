//go:build !gcommonmetrics

// file: pkg/metrics/metrics.go
// version: 1.0.0
// guid: a1b2c3d4-e5f6-7g8h-9i0j-k1l2m3n4o5p6

package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// ProviderRequests counts the number of requests made to subtitle providers
	ProviderRequests = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "subtitle_manager_provider_requests_total",
			Help: "The total number of requests made to subtitle providers",
		},
		[]string{"provider", "status"},
	)

	// TranslationRequests counts the number of translation requests processed
	TranslationRequests = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "subtitle_manager_translation_requests_total",
			Help: "The total number of translation requests processed",
		},
		[]string{"service", "target_language", "status"},
	)

	// APIRequests counts the number of API requests by endpoint
	APIRequests = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "subtitle_manager_api_requests_total",
			Help: "The total number of API requests by endpoint",
		},
		[]string{"endpoint", "method", "status_code"},
	)

	// RequestDuration tracks the duration of API requests
	RequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "subtitle_manager_request_duration_seconds",
			Help:    "The duration of API requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"endpoint", "method"},
	)

	// ActiveSessions tracks the number of active user sessions
	ActiveSessions = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "subtitle_manager_active_sessions",
			Help: "The number of active user sessions",
		},
	)

	// SubtitleDownloads counts successful subtitle downloads
	SubtitleDownloads = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "subtitle_manager_subtitle_downloads_total",
			Help: "The total number of subtitle downloads",
		},
		[]string{"provider", "language"},
	)
)
