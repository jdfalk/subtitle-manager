// file: pkg/security/path_test.go
package security

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
)

func TestValidateAndSanitizePath(t *testing.T) {
	dir := t.TempDir()
	viper.Set("media_directory", dir)
	defer viper.Reset()

	inside := filepath.Join(dir, "sub")
	if err := os.Mkdir(inside, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}

	p, err := ValidateAndSanitizePath(inside)
	if err != nil || p != inside {
		t.Fatalf("expected valid path, got %v %v", p, err)
	}

	// Test path traversal outside temp directory (should fail)
	if _, err := ValidateAndSanitizePath("/etc/passwd"); err == nil {
		t.Fatalf("expected traversal error for /etc/passwd")
	}

	if _, err := ValidateAndSanitizePath("relative/path"); err == nil {
		t.Fatalf("expected absolute path error")
	}

	if _, err := ValidateAndSanitizePath("/etc"); err == nil {
		t.Fatalf("expected disallowed path error")
	}
}

func TestValidateLanguageCode(t *testing.T) {
	tests := []struct {
		name    string
		lang    string
		wantErr bool
	}{
		{"valid short", "en", false},
		{"valid long", "english", false},
		{"valid mixed case", "EnUs", false},
		{"valid with numbers", "en2", false},
		{"empty string", "", true},
		{"too long", "verylonglanguagecode", true},
		{"with slash", "en/us", true},
		{"with dot", "en.us", true},
		{"with dash", "en-us", true},
		{"with underscore", "en_us", true},
		{"with space", "en us", true},
		{"path traversal", "../en", true},
		{"null byte", "en\x00", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateLanguageCode(tt.lang)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateLanguageCode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateProviderName(t *testing.T) {
	tests := []struct {
		name     string
		provider string
		wantErr  bool
	}{
		{"valid simple", "opensubtitles", false},
		{"valid with underscore", "open_subtitles", false},
		{"valid with dash", "open-subtitles", false},
		{"valid with numbers", "provider2", false},
		{"empty string", "", false}, // Empty is allowed
		{"too long", "verylongprovidernamethatexceedsthelimitof50characters", true},
		{"with slash", "open/subtitles", true},
		{"with dot", "open.subtitles", true},
		{"with space", "open subtitles", true},
		{"path traversal", "../provider", true},
		{"null byte", "provider\x00", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateProviderName(tt.provider)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateProviderName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateSubtitleOutputPath(t *testing.T) {
	dir := t.TempDir()
	viper.Set("media_directory", dir)
	defer viper.Reset()

	videoPath := filepath.Join(dir, "movie.mkv")
	if err := os.WriteFile(videoPath, []byte("test"), 0644); err != nil {
		t.Fatalf("create video file: %v", err)
	}

	// Valid case
	output, err := ValidateSubtitleOutputPath(videoPath, "en")
	if err != nil {
		t.Fatalf("expected valid output path, got error: %v", err)
	}
	expectedOutput := filepath.Join(dir, "movie.en.srt")
	if output != expectedOutput {
		t.Errorf("expected %s, got %s", expectedOutput, output)
	}

	// Invalid language
	_, err = ValidateSubtitleOutputPath(videoPath, "../en")
	if err == nil {
		t.Error("expected error for invalid language")
	}

	// Invalid video path
	_, err = ValidateSubtitleOutputPath("/invalid/path", "en")
	if err == nil {
		t.Error("expected error for invalid video path")
	}

	// Video path with traversal to a system directory (should fail)
	_, err = ValidateSubtitleOutputPath("/etc/passwd", "en")
	if err == nil {
		t.Error("expected error for path traversal in video path")
	}
}
