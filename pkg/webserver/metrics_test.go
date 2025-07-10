// file: pkg/webserver/metrics_test.go
// version: 1.0.0
// guid: c3d4e5f6-g7h8-9i0j-1k2l-m3n4o5p6q7r8

package webserver

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jdfalk/subtitle-manager/pkg/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func TestMetricsEndpoint(t *testing.T) {
	// Create a test mux similar to the one in Handler()
	mux := http.NewServeMux()

	// Add the metrics endpoint like in the main Handler function
	mux.Handle("/metrics", promhttp.Handler())

	// Reset metrics to ensure clean state
	metrics.ProviderRequests.Reset()
	metrics.TranslationRequests.Reset()
	metrics.SubtitleDownloads.Reset()

	// Increment some metrics to verify they appear in the output
	metrics.ProviderRequests.WithLabelValues("test_provider", "success").Inc()
	metrics.TranslationRequests.WithLabelValues("google", "en", "success").Inc()
	metrics.SubtitleDownloads.WithLabelValues("test_provider", "en").Inc()

	// Create a test request
	req := httptest.NewRequest("GET", "/metrics", nil)
	w := httptest.NewRecorder()

	// Serve the request
	mux.ServeHTTP(w, req)

	// Check the response
	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	body := w.Body.String()

	// Check that our metrics appear in the output
	expectedMetrics := []string{
		"subtitle_manager_provider_requests_total",
		"subtitle_manager_translation_requests_total",
		"subtitle_manager_subtitle_downloads_total",
	}

	for _, metric := range expectedMetrics {
		if !strings.Contains(body, metric) {
			t.Errorf("expected metric %s not found in output", metric)
		}
	}

	// Check that we have the specific metric values we set
	if !strings.Contains(body, `subtitle_manager_provider_requests_total{provider="test_provider",status="success"} 1`) {
		t.Error("expected provider requests metric value not found")
	}

	if !strings.Contains(body, `subtitle_manager_translation_requests_total{service="google",status="success",target_language="en"} 1`) {
		t.Error("expected translation requests metric value not found")
	}
}

func TestMetricsEndpointContentType(t *testing.T) {
	// Create a test mux with metrics endpoint
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	// Create a test request
	req := httptest.NewRequest("GET", "/metrics", nil)
	w := httptest.NewRecorder()

	// Serve the request
	mux.ServeHTTP(w, req)

	// Check the content type
	contentType := w.Header().Get("Content-Type")
	expectedPrefix := "text/plain; version=0.0.4; charset=utf-8"

	if !strings.HasPrefix(contentType, expectedPrefix) {
		t.Errorf("expected content type starting with %s, got %s", expectedPrefix, contentType)
	}
}
