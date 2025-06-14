package bazarr

import "testing"

// TestMapSettings verifies mapping of Bazarr settings.
func TestMapSettings(t *testing.T) {
	in := Settings{
		"general": map[string]any{
			"ip":       "0.0.0.0", // Updated to match actual Bazarr API
			"port":     6767,
			"base_url": "/bazarr", // Updated to match actual Bazarr API
		},
		"auth": map[string]any{
			"apikey":   "k", // Updated to match actual Bazarr API
			"username": "u", // Updated to match actual Bazarr API
			"password": "p",
		},
		"providers": map[string]any{
			"opensubtitles": map[string]any{"enabled": true},
			"embedded":      map[string]any{"enabled": false},
		},
		"languages": []any{
			map[string]any{"code": "en", "enabled": true},
			map[string]any{"code": "fr", "enabled": false},
		},
	}
	out := MapSettings(in)
	if out["web.host"] != "0.0.0.0" {
		t.Fatalf("web.host unexpected: %v", out["web.host"])
	}
	if out["web.port"].(int) != 6767 {
		t.Fatalf("web.port unexpected: %v", out["web.port"])
	}
	if out["web.base_url"] != "/bazarr" {
		t.Fatalf("web.base_url unexpected: %v", out["web.base_url"])
	}
	if out["auth.api_key"] != "k" {
		t.Fatalf("auth.api_key unexpected: %v", out["auth.api_key"])
	}
}

// TestMapSettingsWithInfo verifies the new mapping function with detailed information.
func TestMapSettingsWithInfo(t *testing.T) {
	in := Settings{
		"general": map[string]any{
			"ip":                  "0.0.0.0",
			"port":                6767,
			"base_url":            "/bazarr",
			"theme":               "dark",
			"debug":               true,
			"minimum_score":       70,
			"minimum_score_movie": 80,
			"enabled_providers":   []any{"opensubtitles", "addic7ed"},
			"use_radarr":          true,
			"use_sonarr":          true,
		},
		"auth": map[string]any{
			"apikey":   "test-api-key",
			"username": "admin",
			"password": "secret",
		},
		"radarr": map[string]any{
			"ip":     "192.168.1.100",
			"port":   7878,
			"apikey": "radarr-key",
			"ssl":    false,
		},
		"sonarr": map[string]any{
			"ip":     "192.168.1.101",
			"port":   8989,
			"apikey": "sonarr-key",
			"ssl":    true,
		},
		"plex": map[string]any{
			"ip":             "192.168.1.102",
			"port":           32400,
			"apikey":         "plex-token",
			"ssl":            false,
			"movie_library":  "Movies",
			"series_library": "TV Shows",
		},
		"opensubtitles": map[string]any{
			"username": "user@example.com",
			"password": "pass123",
			"vip":      true,
			"ssl":      true,
		},
	}

	config, mappings := MapSettingsWithInfo(in)

	// Test that we get both config and mappings
	if len(config) == 0 {
		t.Fatal("Expected non-empty config map")
	}
	if len(mappings) == 0 {
		t.Fatal("Expected non-empty mappings slice")
	}

	// Test specific mappings
	expectedMappings := map[string]struct {
		section     string
		description string
	}{
		"web.host":                         {"Web Server", "Server bind address"},
		"web.port":                         {"Web Server", "Server port"},
		"web.base_url":                     {"Web Server", "Base URL path"},
		"ui.theme":                         {"User Interface", "UI theme preference"},
		"log.debug":                        {"Logging", "Enable debug logging"},
		"matching.minimum_score":           {"Subtitle Matching", "Minimum score for series subtitles"},
		"matching.minimum_score_movie":     {"Subtitle Matching", "Minimum score for movie subtitles"},
		"providers.enabled":                {"Subtitle Providers", "Enabled subtitle providers"},
		"integrations.radarr.enabled":      {"Integrations", "Enable Radarr integration"},
		"integrations.sonarr.enabled":      {"Integrations", "Enable Sonarr integration"},
		"auth.api_key":                     {"Authentication", "API key for external access"},
		"auth.username":                    {"Authentication", "Web UI username"},
		"auth.password":                    {"Authentication", "Web UI password"},
		"integrations.radarr.host":         {"Radarr Integration", "Radarr server IP address"},
		"integrations.radarr.port":         {"Radarr Integration", "Radarr server port"},
		"integrations.radarr.api_key":      {"Radarr Integration", "Radarr API key"},
		"integrations.radarr.ssl":          {"Radarr Integration", "Use SSL for Radarr connection"},
		"integrations.sonarr.host":         {"Sonarr Integration", "Sonarr server IP address"},
		"integrations.sonarr.port":         {"Sonarr Integration", "Sonarr server port"},
		"integrations.sonarr.api_key":      {"Sonarr Integration", "Sonarr API key"},
		"integrations.sonarr.ssl":          {"Sonarr Integration", "Use SSL for Sonarr connection"},
		"integrations.plex.host":           {"Plex Integration", "Plex server IP address"},
		"integrations.plex.port":           {"Plex Integration", "Plex server port"},
		"integrations.plex.token":          {"Plex Integration", "Plex authentication token"},
		"integrations.plex.ssl":            {"Plex Integration", "Use SSL for Plex connection"},
		"integrations.plex.movie_library":  {"Plex Integration", "Plex movie library name"},
		"integrations.plex.series_library": {"Plex Integration", "Plex series library name"},
		"providers.opensubtitles.username": {"OpenSubtitles Provider", "OpenSubtitles username"},
		"providers.opensubtitles.password": {"OpenSubtitles Provider", "OpenSubtitles password"},
		"providers.opensubtitles.vip":      {"OpenSubtitles Provider", "OpenSubtitles VIP status"},
		"providers.opensubtitles.ssl":      {"OpenSubtitles Provider", "OpenSubtitles use SSL"},
	}

	// Check that expected mappings exist and have correct metadata
	mappingByKey := make(map[string]MappingInfo)
	for _, mapping := range mappings {
		mappingByKey[mapping.Key] = mapping
	}

	for expectedKey, expected := range expectedMappings {
		mapping, exists := mappingByKey[expectedKey]
		if !exists {
			t.Errorf("Expected mapping for key %s not found", expectedKey)
			continue
		}
		if mapping.Section != expected.section {
			t.Errorf("Key %s: expected section %s, got %s", expectedKey, expected.section, mapping.Section)
		}
		if mapping.Description != expected.description {
			t.Errorf("Key %s: expected description %s, got %s", expectedKey, expected.description, mapping.Description)
		}
		if mapping.Value == nil {
			t.Errorf("Key %s: expected non-nil value", expectedKey)
		}
	}

	// Test that config contains the expected values
	if config["web.host"] != "0.0.0.0" {
		t.Errorf("Expected web.host to be '0.0.0.0', got %v", config["web.host"])
	}
	if config["web.port"] != 6767 {
		t.Errorf("Expected web.port to be 6767, got %v", config["web.port"])
	}
	if config["ui.theme"] != "dark" {
		t.Errorf("Expected ui.theme to be 'dark', got %v", config["ui.theme"])
	}
	if config["log.debug"] != true {
		t.Errorf("Expected log.debug to be true, got %v", config["log.debug"])
	}

	// Test provider settings
	if config["integrations.radarr.enabled"] != true {
		t.Errorf("Expected integrations.radarr.enabled to be true, got %v", config["integrations.radarr.enabled"])
	}
	if config["integrations.radarr.host"] != "192.168.1.100" {
		t.Errorf("Expected integrations.radarr.host to be '192.168.1.100', got %v", config["integrations.radarr.host"])
	}
	if config["providers.opensubtitles.vip"] != true {
		t.Errorf("Expected providers.opensubtitles.vip to be true, got %v", config["providers.opensubtitles.vip"])
	}
}
