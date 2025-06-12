package transcriber

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

// TestWhisperTranscribe verifies that the request is made to the correct
// endpoint and the subtitle text is returned.
func TestWhisperTranscribe(t *testing.T) {
	var gotPath string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		fmt.Fprint(w, "1\n00:00:00,000 --> 00:00:01,000\ntext\n")
	}))
	defer srv.Close()

	SetBaseURL(srv.URL + "/v1")
	defer SetBaseURL("https://api.openai.com/v1")
	SetWhisperModel("whisper-1")

	dir := t.TempDir()
	file := filepath.Join(dir, "a.wav")
	if err := os.WriteFile(file, []byte("data"), 0644); err != nil {
		t.Fatalf("write file: %v", err)
	}

	b, err := WhisperTranscribe(file, "en", "k")
	if err != nil {
		t.Fatalf("transcribe: %v", err)
	}
	if gotPath != "/v1/audio/transcriptions" {
		t.Fatalf("unexpected path %s", gotPath)
	}
	if len(b) == 0 {
		t.Fatalf("empty result")
	}
}
