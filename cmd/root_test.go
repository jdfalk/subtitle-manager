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
	os.Setenv("SM_OPENAI_API_URL", "http://local/v1")
	os.Setenv("SM_ANTICAPTCHA_API_KEY", "ac")
	defer os.Unsetenv("SM_GOOGLE_API_KEY")
	defer os.Unsetenv("SM_OPENSUBTITLES_API_KEY")
	defer os.Unsetenv("SM_FFMPEG_PATH")
	defer os.Unsetenv("SM_GOOGLE_API_URL")
	defer os.Unsetenv("SM_OPENAI_API_URL")
	defer os.Unsetenv("SM_ANTICAPTCHA_API_KEY")
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
	if got := viper.GetString("openai_api_url"); got != "http://local/v1" {
		t.Fatalf("expected openai_api_url=%s, got %s", "http://local/v1", got)
	}
	if got := viper.GetString("anticaptcha.api_key"); got != "ac" {
		t.Fatalf("expected anticaptcha.api_key=%s, got %s", "ac", got)
	}
}
