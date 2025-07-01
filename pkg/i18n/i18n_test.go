// file: pkg/i18n/i18n_test.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174002

package i18n

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitialize(t *testing.T) {
	// Reset global state for testing
	globalLocalizer = nil
	once = sync.Once{}

	Initialize()

	assert.NotNil(t, globalLocalizer)
	assert.Equal(t, "en", globalLocalizer.currentLang.String())
	assert.NotEmpty(t, globalLocalizer.messages)
}

func TestSetLanguage(t *testing.T) {
	// Reset global state for testing
	globalLocalizer = nil
	once = sync.Once{}
	Initialize()

	tests := []struct {
		langCode    string
		expectError bool
	}{
		{"en", false},
		{"es", false},
		{"fr", false},
		{"invalid", true},
		{"", true},
	}

	for _, tt := range tests {
		t.Run(tt.langCode, func(t *testing.T) {
			err := SetLanguage(tt.langCode)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestTranslation(t *testing.T) {
	// Reset global state for testing
	globalLocalizer = nil
	once = sync.Once{}
	Initialize()

	// Test English (default)
	assert.Equal(t, "Scan directory and download subtitles", T("cli.scan.short"))

	// Test Spanish
	SetLanguage("es")
	assert.Equal(t, "Escanear directorio y descargar subtítulos", T("cli.scan.short"))

	// Test French
	SetLanguage("fr")
	assert.Equal(t, "Analyser le répertoire et télécharger les sous-titres", T("cli.scan.short"))

	// Reset to English and test fallback
	SetLanguage("en")
	assert.Equal(t, "Scan directory and download subtitles", T("cli.scan.short"))

	// Test fallback to key for completely missing key
	assert.Equal(t, "missing.key", T("missing.key"))
}

func TestGetAvailableLanguages(t *testing.T) {
	// Reset global state for testing
	globalLocalizer = nil
	once = sync.Once{}
	Initialize()

	languages := GetAvailableLanguages()

	assert.Contains(t, languages, "en")
	assert.Contains(t, languages, "es")
	assert.Contains(t, languages, "fr")
	assert.True(t, len(languages) >= 3)
}
