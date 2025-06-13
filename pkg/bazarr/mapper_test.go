package bazarr

import "testing"

// TestMapSettings verifies mapping of Bazarr settings.
func TestMapSettings(t *testing.T) {
	in := Settings{
		"general": map[string]any{
			"bind_address": "0.0.0.0",
			"port":         6767,
			"url_base":     "/bazarr",
		},
		"auth": map[string]any{
			"api_key":  "k",
			"user":     "u",
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
	prov, ok := out["providers"].(map[string]bool)
	if !ok || !prov["opensubtitles"] || prov["embedded"] {
		t.Fatalf("provider mapping incorrect: %v", prov)
	}
	langs, ok := out["languages"].([]string)
	if !ok || len(langs) != 1 || langs[0] != "en" {
		t.Fatalf("language mapping incorrect: %v", langs)
	}
}
