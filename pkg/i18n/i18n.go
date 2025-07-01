// file: pkg/i18n/i18n.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174001

package i18n

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Localizer handles internationalization for the application
type Localizer struct {
	currentLang language.Tag
	messages    map[string]map[string]string
	printer     *message.Printer
	mu          sync.RWMutex
}

var (
	globalLocalizer *Localizer
	once            sync.Once
)

// Initialize sets up the global localizer with default language
func Initialize() {
	once.Do(func() {
		globalLocalizer = &Localizer{
			currentLang: language.English,
			messages:    make(map[string]map[string]string),
		}
		globalLocalizer.loadEmbeddedTranslations()
		globalLocalizer.printer = message.NewPrinter(globalLocalizer.currentLang)
	})
}

// SetLanguage changes the current language
func SetLanguage(langCode string) error {
	if globalLocalizer == nil {
		Initialize()
	}

	tag, err := language.Parse(langCode)
	if err != nil {
		return fmt.Errorf("invalid language code: %s", langCode)
	}

	globalLocalizer.mu.Lock()
	defer globalLocalizer.mu.Unlock()

	globalLocalizer.currentLang = tag
	globalLocalizer.printer = message.NewPrinter(tag)
	return nil
}

// T translates a message key to the current language
func T(key string, args ...interface{}) string {
	if globalLocalizer == nil {
		Initialize()
	}

	globalLocalizer.mu.RLock()
	defer globalLocalizer.mu.RUnlock()

	langCode := globalLocalizer.currentLang.String()

	// Try current language first
	if translations, exists := globalLocalizer.messages[langCode]; exists {
		if message, exists := translations[key]; exists {
			if len(args) > 0 {
				return fmt.Sprintf(message, args...)
			}
			return message
		}
	}

	// Fallback to English
	if langCode != "en" {
		if translations, exists := globalLocalizer.messages["en"]; exists {
			if message, exists := translations[key]; exists {
				if len(args) > 0 {
					return fmt.Sprintf(message, args...)
				}
				return message
			}
		}
	}

	// Fallback to key itself if no translation found
	if len(args) > 0 {
		return fmt.Sprintf(key, args...)
	}
	return key
}

// GetAvailableLanguages returns a list of available language codes
func GetAvailableLanguages() []string {
	if globalLocalizer == nil {
		Initialize()
	}

	globalLocalizer.mu.RLock()
	defer globalLocalizer.mu.RUnlock()

	languages := make([]string, 0, len(globalLocalizer.messages))
	for lang := range globalLocalizer.messages {
		languages = append(languages, lang)
	}
	return languages
}

// loadEmbeddedTranslations loads the built-in translation messages
func (l *Localizer) loadEmbeddedTranslations() {
	// English translations
	l.messages["en"] = map[string]string{
		// CLI Commands
		"cli.scan.short":        "Scan directory and download subtitles",
		"cli.scan.scanning":     "scanning %s",
		"cli.scan.flag.upgrade": "replace existing subtitles",
		"cli.convert.short":     "Convert subtitle to SRT",
		"cli.convert.converted": "Converted %s to %s",
		"cli.merge.short":       "Merge subtitle files",
		"cli.translate.short":   "Translate subtitle file",
		"cli.fetch.short":       "Fetch subtitles from online sources",
		"cli.batch.short":       "Batch process subtitle files",
		"cli.web.short":         "Start web interface",
		"cli.version.short":     "Show version information",

		// Common messages
		"common.error.invalid_path": "invalid directory path: %v",
		"common.error.invalid_lang": "invalid language code: %v",
		"common.error.db_open":      "db open: %v",
		"common.error.file_create":  "failed to create file: %v",
		"common.error.file_write":   "failed to write file: %v",

		// Web UI messages
		"web.title":         "Subtitle Manager",
		"web.nav.dashboard": "Dashboard",
		"web.nav.library":   "Media Library",
		"web.nav.convert":   "Convert",
		"web.nav.translate": "Translate",
		"web.nav.extract":   "Extract",
		"web.nav.history":   "History",
		"web.nav.wanted":    "Wanted",
		"web.nav.settings":  "Settings",
		"web.nav.system":    "System",

		// Settings
		"settings.language": "Language",
		"settings.theme":    "Theme",
		"settings.save":     "Save",
		"settings.cancel":   "Cancel",
	}

	// Spanish translations
	l.messages["es"] = map[string]string{
		// CLI Commands
		"cli.scan.short":        "Escanear directorio y descargar subtítulos",
		"cli.scan.scanning":     "escaneando %s",
		"cli.scan.flag.upgrade": "reemplazar subtítulos existentes",
		"cli.convert.short":     "Convertir subtítulo a SRT",
		"cli.convert.converted": "Convertido %s a %s",
		"cli.merge.short":       "Fusionar archivos de subtítulos",
		"cli.translate.short":   "Traducir archivo de subtítulos",
		"cli.fetch.short":       "Obtener subtítulos de fuentes en línea",
		"cli.batch.short":       "Procesar archivos de subtítulos por lotes",
		"cli.web.short":         "Iniciar interfaz web",
		"cli.version.short":     "Mostrar información de versión",

		// Common messages
		"common.error.invalid_path": "ruta de directorio inválida: %v",
		"common.error.invalid_lang": "código de idioma inválido: %v",
		"common.error.db_open":      "abrir bd: %v",
		"common.error.file_create":  "error al crear archivo: %v",
		"common.error.file_write":   "error al escribir archivo: %v",

		// Web UI messages
		"web.title":         "Gestor de Subtítulos",
		"web.nav.dashboard": "Panel",
		"web.nav.library":   "Biblioteca de Medios",
		"web.nav.convert":   "Convertir",
		"web.nav.translate": "Traducir",
		"web.nav.extract":   "Extraer",
		"web.nav.history":   "Historial",
		"web.nav.wanted":    "Buscados",
		"web.nav.settings":  "Configuración",
		"web.nav.system":    "Sistema",

		// Settings
		"settings.language": "Idioma",
		"settings.theme":    "Tema",
		"settings.save":     "Guardar",
		"settings.cancel":   "Cancelar",
	}

	// French translations
	l.messages["fr"] = map[string]string{
		// CLI Commands
		"cli.scan.short":        "Analyser le répertoire et télécharger les sous-titres",
		"cli.scan.scanning":     "analyse de %s",
		"cli.scan.flag.upgrade": "remplacer les sous-titres existants",
		"cli.convert.short":     "Convertir le sous-titre en SRT",
		"cli.convert.converted": "Converti %s en %s",
		"cli.merge.short":       "Fusionner les fichiers de sous-titres",
		"cli.translate.short":   "Traduire le fichier de sous-titres",
		"cli.fetch.short":       "Récupérer les sous-titres depuis des sources en ligne",
		"cli.batch.short":       "Traiter les fichiers de sous-titres par lots",
		"cli.web.short":         "Démarrer l'interface web",
		"cli.version.short":     "Afficher les informations de version",

		// Common messages
		"common.error.invalid_path": "chemin de répertoire invalide: %v",
		"common.error.invalid_lang": "code de langue invalide: %v",
		"common.error.db_open":      "ouverture bd: %v",
		"common.error.file_create":  "échec de création du fichier: %v",
		"common.error.file_write":   "échec d'écriture du fichier: %v",

		// Web UI messages
		"web.title":         "Gestionnaire de Sous-titres",
		"web.nav.dashboard": "Tableau de bord",
		"web.nav.library":   "Bibliothèque multimédia",
		"web.nav.convert":   "Convertir",
		"web.nav.translate": "Traduire",
		"web.nav.extract":   "Extraire",
		"web.nav.history":   "Historique",
		"web.nav.wanted":    "Recherchés",
		"web.nav.settings":  "Paramètres",
		"web.nav.system":    "Système",

		// Settings
		"settings.language": "Langue",
		"settings.theme":    "Thème",
		"settings.save":     "Enregistrer",
		"settings.cancel":   "Annuler",
	}
}

// LoadTranslationsFromFile loads translations from external JSON files
func LoadTranslationsFromFile(translationsDir string) error {
	if globalLocalizer == nil {
		Initialize()
	}

	return filepath.WalkDir(translationsDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && strings.HasSuffix(path, ".json") {
			lang := strings.TrimSuffix(d.Name(), ".json")

			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			var translations map[string]string
			if err := json.Unmarshal(data, &translations); err != nil {
				return err
			}

			globalLocalizer.mu.Lock()
			globalLocalizer.messages[lang] = translations
			globalLocalizer.mu.Unlock()
		}

		return nil
	})
}
