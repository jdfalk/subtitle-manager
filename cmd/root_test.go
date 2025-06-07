package cmd

import (
	"os"
	"testing"

	"github.com/spf13/viper"
)

func TestInitConfigEnv(t *testing.T) {
	os.Setenv("SM_GOOGLE_API_KEY", "testkey")
	defer os.Unsetenv("SM_GOOGLE_API_KEY")
	initConfig()
	if got := viper.GetString("google_api_key"); got != "testkey" {
		t.Fatalf("expected google_api_key=%s, got %s", "testkey", got)
	}
}
