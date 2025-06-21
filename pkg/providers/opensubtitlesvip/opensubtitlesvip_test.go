package opensubtitlesvip

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"testing"

	"github.com/oz/osdb"
)

// TestClientFetch verifies that the client downloads subtitles using the expected URL.
type mockAPI struct {
	path string
	lang string
}

func (m *mockAPI) LogIn(user, pass, lang string) error {
	m.lang = lang
	return nil
}

func (m *mockAPI) FileSearch(path string, langs []string) (osdb.Subtitles, error) {
	m.path = path
	return osdb.Subtitles{{IDSubtitleFile: "1"}}, nil
}

func (m *mockAPI) DownloadSubtitles(subs osdb.Subtitles) ([]osdb.SubtitleFile, error) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	gz.Write([]byte("data"))
	gz.Close()
	return []osdb.SubtitleFile{{ID: subs[0].IDSubtitleFile, Data: base64.StdEncoding.EncodeToString(buf.Bytes())}}, nil
}

func TestClientFetch(t *testing.T) {
	m := &mockAPI{}
	c := &Client{api: m}
	b, err := c.Fetch(context.Background(), "/path/movie.mkv", "en")
	if err != nil {
		t.Fatalf("fetch: %v", err)
	}
	if string(b) != "data" {
		t.Fatalf("unexpected body: %s", b)
	}
	if m.path != "/path/movie.mkv" {
		t.Fatalf("unexpected path: %s", m.path)
	}
	if m.lang != "en" {
		t.Fatalf("unexpected lang: %s", m.lang)
	}
}
