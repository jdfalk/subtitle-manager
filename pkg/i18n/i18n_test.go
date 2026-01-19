// file: pkg/i18n/i18n_test.go
// version: 1.0.0
// guid: 6b2e8c74-6c8f-4f1e-9f47-7c2a94f4dcaa

package i18n

import (
	"os"
	"path/filepath"
	"sync"
	"testing"

	"golang.org/x/text/language"
)

func resetLocalizer(t *testing.T) {
	t.Helper()
	globalLocalizer = nil
	once = sync.Once{}
}

func containsString(values []string, target string) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}

func TestLocalizer_Initialize_DefaultsLoaded(t *testing.T) {
	// Arrange
	resetLocalizer(t)

	// Act
	Initialize()

	// Assert
	if globalLocalizer == nil {
		t.Fatal("expected global localizer to be initialized")
	}
	if globalLocalizer.currentLang != language.English {
		t.Fatalf("expected default language %q, got %q", language.English, globalLocalizer.currentLang)
	}
	if _, ok := globalLocalizer.messages["en"]; !ok {
		t.Fatal("expected English translations to be loaded")
	}
	if _, ok := globalLocalizer.messages["es"]; !ok {
		t.Fatal("expected Spanish translations to be loaded")
	}
	if _, ok := globalLocalizer.messages["fr"]; !ok {
		t.Fatal("expected French translations to be loaded")
	}
}

func TestLocalizer_SetLanguage_InvalidCode_ReturnsError(t *testing.T) {
	// Arrange
	resetLocalizer(t)
	Initialize()
	startingLang := globalLocalizer.currentLang

	// Act
	err := SetLanguage("invalid@lang")

	// Assert
	if err == nil {
		t.Fatal("expected error for invalid language code")
	}
	if globalLocalizer.currentLang != startingLang {
		t.Fatalf("expected language to remain %q, got %q", startingLang, globalLocalizer.currentLang)
	}
}

func TestLocalizer_SetLanguage_ValidCode_UpdatesLanguage(t *testing.T) {
	// Arrange
	resetLocalizer(t)
	Initialize()

	// Act
	err := SetLanguage("es")

	// Assert
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if globalLocalizer.currentLang.String() != "es" {
		t.Fatalf("expected language to be set to es, got %q", globalLocalizer.currentLang)
	}
	if translated := T("cli.scan.short"); translated != "Escanear directorio y descargar subtítulos" {
		t.Fatalf("expected Spanish translation, got %q", translated)
	}
}

func TestLocalizer_Translate_FallbacksToEnglish(t *testing.T) {
	// Arrange
	resetLocalizer(t)
	Initialize()

	globalLocalizer.mu.Lock()
	globalLocalizer.messages["en"]["only.english"] = "English only"
	globalLocalizer.mu.Unlock()

	if err := SetLanguage("es"); err != nil {
		t.Fatalf("expected language set to succeed, got %v", err)
	}

	// Act
	translated := T("only.english")

	// Assert
	if translated != "English only" {
		t.Fatalf("expected English fallback, got %q", translated)
	}
}

func TestLocalizer_Translate_FallbacksToKey(t *testing.T) {
	// Arrange
	resetLocalizer(t)
	Initialize()

	// Act
	translated := T("missing %s", "translation")

	// Assert
	if translated != "missing translation" {
		t.Fatalf("expected key fallback with formatting, got %q", translated)
	}
}

func TestLocalizer_GetAvailableLanguages_ReturnsLoadedLanguages(t *testing.T) {
	// Arrange
	resetLocalizer(t)
	Initialize()

	// Act
	languages := GetAvailableLanguages()

	// Assert
	if !containsString(languages, "en") {
		t.Fatal("expected English to be available")
	}
	if !containsString(languages, "es") {
		t.Fatal("expected Spanish to be available")
	}
	if !containsString(languages, "fr") {
		t.Fatal("expected French to be available")
	}
}

func TestLocalizer_LoadTranslationsFromFile_LoadsJSON(t *testing.T) {
	// Arrange
	resetLocalizer(t)
	Initialize()

	tempDir := t.TempDir()
	translations := []byte(`{"greeting": "Hallo %s"}`)
	translationsPath := filepath.Join(tempDir, "de.json")

	if err := os.WriteFile(translationsPath, translations, 0o600); err != nil {
		t.Fatalf("failed to write translations file: %v", err)
	}

	// Act
	err := LoadTranslationsFromFile(tempDir)

	// Assert
	if err != nil {
		t.Fatalf("expected no error loading translations, got %v", err)
	}
	if err := SetLanguage("de"); err != nil {
		t.Fatalf("expected to set language to de, got %v", err)
	}
	key := "greeting"
	if translated := T(key, "Welt"); translated != "Hallo Welt" {
		t.Fatalf("expected loaded translation, got %q", translated)
	}
}
