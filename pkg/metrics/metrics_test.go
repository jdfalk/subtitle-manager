// file: pkg/metrics/metrics_test.go
// version: 1.0.0
// guid: b2c3d4e5-f6g7-8h9i-0j1k-l2m3n4o5p6q7

package metrics

import (
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
)

func TestProviderRequestsMetric(t *testing.T) {
	// Test that the provider requests metric is properly registered
	if ProviderRequests == nil {
		t.Fatal("ProviderRequests metric is nil")
	}

	// Reset the metric for testing
	ProviderRequests.Reset()

	// Increment the metric
	ProviderRequests.WithLabelValues("opensubtitles", "success").Inc()
	ProviderRequests.WithLabelValues("opensubtitles", "error").Inc()
	ProviderRequests.WithLabelValues("opensubtitles", "error").Inc()

	// Check the metric values
	expected := `
		# HELP subtitle_manager_provider_requests_total The total number of requests made to subtitle providers
		# TYPE subtitle_manager_provider_requests_total counter
		subtitle_manager_provider_requests_total{provider="opensubtitles",status="error"} 2
		subtitle_manager_provider_requests_total{provider="opensubtitles",status="success"} 1
	`
	if err := testutil.CollectAndCompare(ProviderRequests, strings.NewReader(expected)); err != nil {
		t.Errorf("unexpected metric output: %v", err)
	}
}

func TestTranslationRequestsMetric(t *testing.T) {
	// Test that the translation requests metric is properly registered
	if TranslationRequests == nil {
		t.Fatal("TranslationRequests metric is nil")
	}

	// Reset the metric for testing
	TranslationRequests.Reset()

	// Increment the metric
	TranslationRequests.WithLabelValues("google", "en", "success").Inc()
	TranslationRequests.WithLabelValues("openai", "fr", "error").Inc()

	// Check the metric values
	expected := `
		# HELP subtitle_manager_translation_requests_total The total number of translation requests processed
		# TYPE subtitle_manager_translation_requests_total counter
		subtitle_manager_translation_requests_total{service="google",status="success",target_language="en"} 1
		subtitle_manager_translation_requests_total{service="openai",status="error",target_language="fr"} 1
	`
	if err := testutil.CollectAndCompare(TranslationRequests, strings.NewReader(expected)); err != nil {
		t.Errorf("unexpected metric output: %v", err)
	}
}

func TestSubtitleDownloadsMetric(t *testing.T) {
	// Test that the subtitle downloads metric is properly registered
	if SubtitleDownloads == nil {
		t.Fatal("SubtitleDownloads metric is nil")
	}

	// Reset the metric for testing
	SubtitleDownloads.Reset()

	// Increment the metric
	SubtitleDownloads.WithLabelValues("opensubtitles", "en").Inc()
	SubtitleDownloads.WithLabelValues("opensubtitles", "en").Inc()
	SubtitleDownloads.WithLabelValues("subscene", "fr").Inc()

	// Check the metric values
	expected := `
		# HELP subtitle_manager_subtitle_downloads_total The total number of subtitle downloads
		# TYPE subtitle_manager_subtitle_downloads_total counter
		subtitle_manager_subtitle_downloads_total{language="en",provider="opensubtitles"} 2
		subtitle_manager_subtitle_downloads_total{language="fr",provider="subscene"} 1
	`
	if err := testutil.CollectAndCompare(SubtitleDownloads, strings.NewReader(expected)); err != nil {
		t.Errorf("unexpected metric output: %v", err)
	}
}

func TestAllMetricsRegistered(t *testing.T) {
	// Check that all metrics are properly defined
	metrics := []prometheus.Collector{
		ProviderRequests,
		TranslationRequests,
		APIRequests,
		RequestDuration,
		ActiveSessions,
		SubtitleDownloads,
	}

	for i, metric := range metrics {
		if metric == nil {
			t.Errorf("metric %d is nil", i)
		}
	}
}
