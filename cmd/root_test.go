package cmd

import (
	"os"
	"testing"

	"github.com/spf13/viper"
)

func TestInitConfigEnv(t *testing.T) {
	os.Setenv("SM_GOOGLE_API_KEY", "testkey")
	os.Setenv("SM_OPENSUBTITLES_API_KEY", "osk")
	os.Setenv("SM_FFMPEG_PATH", "/usr/bin/ffmpeg")
	os.Setenv("SM_GOOGLE_API_URL", "http://api")
	defer os.Unsetenv("SM_GOOGLE_API_KEY")
	defer os.Unsetenv("SM_OPENSUBTITLES_API_KEY")
	defer os.Unsetenv("SM_FFMPEG_PATH")
	defer os.Unsetenv("SM_GOOGLE_API_URL")
	initConfig()
	if got := viper.GetString("google_api_key"); got != "testkey" {
		t.Fatalf("expected google_api_key=%s, got %s", "testkey", got)
	}
	if got := viper.GetString("opensubtitles.api_key"); got != "osk" {
		t.Fatalf("expected opensubtitles.api_key=%s, got %s", "osk", got)
	}
	if got := viper.GetString("ffmpeg_path"); got != "/usr/bin/ffmpeg" {
		t.Fatalf("expected ffmpeg_path=%s, got %s", "/usr/bin/ffmpeg", got)
	}
	if got := viper.GetString("google_api_url"); got != "http://api" {
		t.Fatalf("expected google_api_url=%s, got %s", "http://api", got)
	}
}
