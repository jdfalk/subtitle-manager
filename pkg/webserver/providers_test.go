// file: pkg/webserver/providers_test.go
package webserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/spf13/viper"
)

// TestProvidersHandlerGet verifies that the handler lists providers
func TestProvidersHandlerGet(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/providers", nil)
	rr := httptest.NewRecorder()
	providersHandler().ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}

	var resp []ProviderInfo
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if len(resp) == 0 {
		t.Fatalf("expected providers list")
	}
}

// TestProvidersHandlerInvalidJSON verifies bad JSON returns 400
func TestProvidersHandlerInvalidJSON(t *testing.T) {
	req := httptest.NewRequest("POST", "/api/providers", bytes.NewBufferString("bad"))
	rr := httptest.NewRecorder()
	providersHandler().ServeHTTP(rr, req)
	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rr.Code)
	}
}

// TestProvidersHandlerUpdate verifies POST updates configuration
func TestProvidersHandlerUpdate(t *testing.T) {
	tmp := t.TempDir() + "/config.yaml"
	viper.SetConfigFile(tmp)
	defer viper.Reset()

	body := `{"name":"generic","enabled":true,"config":{"api_url":"http://example.com"}}`
	req := httptest.NewRequest("POST", "/api/providers", bytes.NewBufferString(body))
	rr := httptest.NewRecorder()
	providersHandler().ServeHTTP(rr, req)

	if rr.Code != http.StatusNoContent {
		t.Fatalf("expected status 204, got %d", rr.Code)
	}

	if !viper.GetBool("providers.generic.enabled") {
		t.Fatalf("config not updated")
	}
	if viper.GetString("providers.generic.config.api_url") != "http://example.com" {
		t.Fatalf("value not written")
	}
}
