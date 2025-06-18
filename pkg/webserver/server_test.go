//go:build !ci
// +build !ci

//nolint:errcheck // Test file - ignoring error checks for brevity
package webserver

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/jdfalk/subtitle-manager/pkg/auth"
	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/testutil"
	"github.com/jdfalk/subtitle-manager/pkg/translator"

	"github.com/spf13/viper"
)

// TestHandler verifies that the handler serves index.html at root.
func TestHandler(t *testing.T) {
	db, err := database.Open(":memory:")
	testutil.MustNoError(t, "open db", err)
	defer db.Close()

	// create test user and api key
	testutil.MustNoError(t, "create user", auth.CreateUser(db, "test", "pass", "", "admin"))
	key, err := auth.GenerateAPIKey(db, 1)
	testutil.MustNoError(t, "generate key", err)

	h, err := Handler(db)
	testutil.MustNoError(t, "create handler", err)
	srv := httptest.NewServer(h)
	defer srv.Close()
	req, _ := http.NewRequest("GET", srv.URL+"/", nil)
	req.Header.Set("X-API-Key", key)
	resp, err := srv.Client().Do(req)
	if err != nil {
		t.Fatalf("http get: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("unexpected status: %d", resp.StatusCode)
	}
}

// TestSPAIndexFallback verifies that unknown paths return index.html so client
// side routing works on refresh.
func TestSPAIndexFallback(t *testing.T) {
	db, err := database.Open(":memory:")
	testutil.MustNoError(t, "open db", err)
	defer db.Close()

	testutil.MustNoError(t, "create user", auth.CreateUser(db, "test", "pass", "", "admin"))
	key, err := auth.GenerateAPIKey(db, 1)
	testutil.MustNoError(t, "generate key", err)

	h, err := Handler(db)
	testutil.MustNoError(t, "create handler", err)
	srv := httptest.NewServer(h)
	defer srv.Close()

	req, _ := http.NewRequest("GET", srv.URL+"/settings", nil)
	req.Header.Set("X-API-Key", key)
	resp, err := srv.Client().Do(req)
	if err != nil {
		t.Fatalf("http get: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status: %d", resp.StatusCode)
	}
}

// TestBaseURL verifies that the handler respects the configured base URL.
func TestBaseURL(t *testing.T) {
	db, err := database.Open(":memory:")
	testutil.MustNoError(t, "open db", err)
	defer db.Close()

	testutil.MustNoError(t, "create user", auth.CreateUser(db, "test", "pass", "", "admin"))
	key := testutil.MustGet(t, "key", func() (string, error) { return auth.GenerateAPIKey(db, 1) })

	viper.Set("base_url", "sub")
	defer viper.Reset()

	h, err := Handler(db)
	testutil.MustNoError(t, "handler", err)
	srv := httptest.NewServer(h)
	defer srv.Close()

	req, _ := http.NewRequest("GET", srv.URL+"/sub/", nil)
	req.Header.Set("X-API-Key", key)
	resp := testutil.MustGet(t, "request", func() (*http.Response, error) { return srv.Client().Do(req) })
	testutil.MustEqual(t, "status", http.StatusOK, resp.StatusCode)

	req2, _ := http.NewRequest("GET", srv.URL+"/", nil)
	req2.Header.Set("X-API-Key", key)
	resp2 := testutil.MustGet(t, "request2", func() (*http.Response, error) { return srv.Client().Do(req2) })
	if resp2.StatusCode == http.StatusOK {
		t.Fatalf("root path should not be served when base_url set")
	}
}

// TestRBAC verifies that permissions are enforced for protected routes.
func TestRBAC(t *testing.T) {
	db, err := database.Open(":memory:")
	testutil.MustNoError(t, "open db", err)
	defer db.Close()

	// viewer role should not access /api/config
	testutil.MustNoError(t, "create viewer", auth.CreateUser(db, "viewer", "p", "", "viewer"))
	vkey, err := auth.GenerateAPIKey(db, 1)
	testutil.MustNoError(t, "viewer key", err)

	// admin role can access
	testutil.MustNoError(t, "create admin", auth.CreateUser(db, "admin", "p", "", "admin"))
	akey, err := auth.GenerateAPIKey(db, 2)
	testutil.MustNoError(t, "admin key", err)

	h, err := Handler(db)
	testutil.MustNoError(t, "create handler", err)
	srv := httptest.NewServer(h)
	defer srv.Close()

	req, _ := http.NewRequest("GET", srv.URL+"/api/config", nil)
	req.Header.Set("X-API-Key", vkey)
	resp, err := srv.Client().Do(req)
	if err != nil {
		t.Fatalf("http get: %v", err)
	}
	if resp.StatusCode != http.StatusForbidden {
		t.Fatalf("viewer status %d", resp.StatusCode)
	}

	req2, _ := http.NewRequest("GET", srv.URL+"/api/config", nil)
	req2.Header.Set("X-API-Key", akey)
	resp2, err := srv.Client().Do(req2)
	if err != nil {
		t.Fatalf("admin get: %v", err)
	}
	if resp2.StatusCode != http.StatusOK {
		t.Fatalf("admin status %d", resp2.StatusCode)
	}
}

// TestConfigUpdate verifies that POST /api/config updates and persists values.
func TestConfigUpdate(t *testing.T) {
	db, err := database.Open(":memory:")
	testutil.MustNoError(t, "open db", err)
	defer db.Close()

	testutil.MustNoError(t, "create admin", auth.CreateUser(db, "admin", "p", "", "admin"))
	akey, err := auth.GenerateAPIKey(db, 1)
	testutil.MustNoError(t, "admin key", err)

	tmp := filepath.Join(t.TempDir(), "cfg.yaml")
	viper.SetConfigFile(tmp)
	viper.Set("test_key", "old")
	testutil.MustNoError(t, "write config", viper.WriteConfig())
	defer viper.Reset()

	h, err := Handler(db)
	testutil.MustNoError(t, "create handler", err)
	srv := httptest.NewServer(h)
	defer srv.Close()

	body := strings.NewReader(`{"test_key":"new"}`)
	req, _ := http.NewRequest("POST", srv.URL+"/api/config", body)
	req.Header.Set("X-API-Key", akey)
	resp, err := srv.Client().Do(req)
	if err != nil {
		t.Fatalf("post: %v", err)
	}
	if resp.StatusCode != http.StatusNoContent {
		t.Fatalf("status %d", resp.StatusCode)
	}

	if viper.GetString("test_key") != "new" {
		t.Fatalf("viper not updated")
	}
	data := testutil.MustGet(t, "read config", func() ([]byte, error) { return os.ReadFile(tmp) })
	if !strings.Contains(string(data), "test_key: new") {
		t.Fatalf("config not written")
	}
}

// TestScanHandlers verifies /api/scan and /api/scan/status.
func TestScanHandlers(t *testing.T) {
	db, err := database.Open(":memory:")
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	defer db.Close()
	if err := auth.CreateUser(db, "admin", "p", "", "admin"); err != nil {
		t.Fatalf("create admin: %v", err)
	}
	key, err := auth.GenerateAPIKey(db, 1)
	if err != nil {
		t.Fatalf("api key: %v", err)
	}
	// fake subtitle server
	subSrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("sub")) // nolint:errcheck
	}))
	defer subSrv.Close()
	viper.Set("providers.generic.api_url", subSrv.URL)
	defer viper.Reset()
	h, err := Handler(db)
	if err != nil {
		t.Fatalf("handler: %v", err)
	}
	srv := httptest.NewServer(h)
	defer srv.Close()
	// create video file
	dir := t.TempDir()
	vid := filepath.Join(dir, "file.mkv")
	_ = os.WriteFile(vid, []byte("x"), 0644) // nolint:errcheck
	// trigger scan
	body := strings.NewReader(`{"provider":"generic","directory":"` + dir + `","lang":"en"}`)
	req, _ := http.NewRequest("POST", srv.URL+"/api/scan", body)
	req.Header.Set("X-API-Key", key)
	resp, err := srv.Client().Do(req)
	if err != nil || resp.StatusCode != http.StatusAccepted {
		t.Fatalf("scan start: %v %d", err, resp.StatusCode)
	}
	// poll status until not running
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		req2, _ := http.NewRequest("GET", srv.URL+"/api/scan/status", nil)
		req2.Header.Set("X-API-Key", key)
		r2, err := srv.Client().Do(req2)
		if err != nil {
			t.Fatalf("status: %v", err)
		}
		var s struct {
			Running   bool `json:"running"`
			Completed int  `json:"completed"`
		}
		json.NewDecoder(r2.Body).Decode(&s)
		r2.Body.Close()
		if !s.Running {
			if s.Completed != 1 {
				t.Fatalf("completed %d", s.Completed)
			}
			return
		}
	}
	t.Fatalf("scan did not finish")
}

// TestExtract verifies that POST /api/extract returns subtitle items.
func TestExtract(t *testing.T) {
	db, err := database.Open(":memory:")
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	defer db.Close()

	if err := auth.CreateUser(db, "admin", "p", "", "admin"); err != nil {
		t.Fatalf("create admin: %v", err)
	}
	key, err := auth.GenerateAPIKey(db, 1)
	if err != nil {
		t.Fatalf("api key: %v", err)
	}

	dir := t.TempDir()
	script := filepath.Join(dir, "ffmpeg")
	data := "#!/bin/sh\ncp ../../testdata/simple.srt \"$6\"\n"
	if err := os.WriteFile(script, []byte(data), 0755); err != nil {
		t.Fatalf("write script: %v", err)
	}
	oldPath := os.Getenv("PATH")
	os.Setenv("PATH", dir+":"+oldPath)
	defer os.Setenv("PATH", oldPath)

	h, err := Handler(db)
	h = testutil.Must(t, "handler creation", h, err)
	srv := httptest.NewServer(h)
	defer srv.Close()

	body := strings.NewReader(`{"path":"dummy.mkv"}`)
	req, _ := http.NewRequest("POST", srv.URL+"/api/extract", body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", key)
	resp, err := srv.Client().Do(req)
	if err != nil {
		t.Fatalf("post: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	var items []any
	if err := json.NewDecoder(resp.Body).Decode(&items); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if len(items) == 0 {
		t.Fatalf("no items returned")
	}
}

// TestConvert verifies that POST /api/convert returns converted data.
func TestConvert(t *testing.T) {
	db, err := database.Open(":memory:")
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	defer db.Close()

	key := setupTestUser(t, db)

	h, err := Handler(db)
	if err != nil {
		t.Fatalf("handler error: %v", err)
	}
	srv := httptest.NewServer(h)
	defer srv.Close()

	data, err := os.ReadFile("../../testdata/simple.srt")
	if err != nil {
		t.Fatalf("read file: %v", err)
	}
	buf := &bytes.Buffer{}
	mw := multipart.NewWriter(buf)
	fw, err := mw.CreateFormFile("file", "simple.srt")
	if err != nil {
		t.Fatalf("form file: %v", err)
	}
	fw.Write(data)
	mw.Close()

	req, _ := http.NewRequest("POST", srv.URL+"/api/convert", buf)
	req.Header.Set("Content-Type", mw.FormDataContentType())
	req.Header.Set("X-API-Key", key)
	resp, err := srv.Client().Do(req)
	if err != nil {
		t.Fatalf("post: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		t.Fatalf("status %d, body: %s", resp.StatusCode, string(body))
	}
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Fatalf("read: %v", err)
	}
	if len(body) == 0 {
		t.Fatalf("no data returned")
	}
}

// TestTranslate verifies the subtitle translation API.
func TestTranslate(t *testing.T) {
	db, err := database.Open(":memory:")
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	defer db.Close()

	// Mock the Google Translate API
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{
			"data": {
				"translations": [
					{
						"translatedText": "1\n00:00:01,000 --> 00:00:04,000\nhola mundo\n\n"
					}
				]
			}
		}`))
	}))
	defer ts.Close()
	translator.SetGoogleAPIURL(ts.URL)
	defer translator.SetGoogleAPIURL("https://translation.googleapis.com/language/translate/v2")

	viper.Set("google_api_key", "test-key")
	defer viper.Reset()

	key := setupTestUser(t, db)

	h, err := Handler(db)
	if err != nil {
		t.Fatalf("handler error: %v", err)
	}
	srv := httptest.NewServer(h)
	defer srv.Close()

	b, _ := os.ReadFile("../../testdata/simple.srt")
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, _ := writer.CreateFormFile("file", "in.srt")
	fw.Write(b)
	writer.WriteField("lang", "es")
	writer.WriteField("service", "google")
	writer.Close()

	req, _ := http.NewRequest("POST", srv.URL+"/api/translate", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("X-API-Key", key)
	resp, err := srv.Client().Do(req)
	if err != nil {
		t.Fatalf("post: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status %d", resp.StatusCode)
	}
	out, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	if !bytes.Contains(out, []byte("hola")) {
		t.Fatalf("unexpected output: %s", out)
	}
}

// TestSetup verifies the initial setup workflow.
func TestSetup(t *testing.T) {
	db, err := database.Open(":memory:")
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	defer db.Close()

	h, err := Handler(db)
	if err != nil {
		t.Fatalf("handler error: %v", err)
	}
	srv := httptest.NewServer(h)
	defer srv.Close()

	// status should report setup needed
	resp, err := http.Get(srv.URL + "/api/setup/status")
	if err != nil {
		t.Fatalf("status: %v", err)
	}
	var st struct {
		Needed bool `json:"needed"`
	}
	json.NewDecoder(resp.Body).Decode(&st)
	resp.Body.Close()
	if !st.Needed {
		t.Fatalf("expected setup needed")
	}

	body := strings.NewReader(`{"server_name":"Test","admin_user":"a","admin_pass":"p"}`)
	resp2, err := http.Post(srv.URL+"/api/setup", "application/json", body)
	if err != nil {
		t.Fatalf("post: %v", err)
	}
	if resp2.StatusCode != http.StatusNoContent {
		t.Fatalf("status %d", resp2.StatusCode)
	}

	// now status should be false
	resp3, err := http.Get(srv.URL + "/api/setup/status")
	if err != nil {
		t.Fatalf("status2: %v", err)
	}
	json.NewDecoder(resp3.Body).Decode(&st)
	resp3.Body.Close()
	if st.Needed {
		t.Fatalf("setup still needed")
	}
}

// TestHistory verifies that /api/history returns history records and filters by language.
func TestHistory(t *testing.T) {
	db, err := database.Open(":memory:")
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	defer db.Close()
	if err := auth.CreateUser(db, "admin", "p", "", "admin"); err != nil {
		t.Fatalf("create admin: %v", err)
	}
	key, err := auth.GenerateAPIKey(db, 1)
	if err != nil {
		t.Fatalf("api key: %v", err)
	}
	_ = database.InsertSubtitle(db, "a.srt", "a.mkv", "en", "google", "", false)
	_ = database.InsertDownload(db, "b.srt", "b.mkv", "os", "en")
	h, err := Handler(db)
	if err != nil {
		t.Fatalf("handler: %v", err)
	}
	srv := httptest.NewServer(h)
	defer srv.Close()
	req, _ := http.NewRequest("GET", srv.URL+"/api/history", nil)
	req.Header.Set("X-API-Key", key)
	resp, err := srv.Client().Do(req)
	if err != nil {
		t.Fatalf("history: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status %d", resp.StatusCode)
	}
	var out struct {
		Translations []database.SubtitleRecord `json:"translations"`
		Downloads    []database.DownloadRecord `json:"downloads"`
	}
	json.NewDecoder(resp.Body).Decode(&out)
	resp.Body.Close()
	if len(out.Translations) != 1 || len(out.Downloads) != 1 {
		t.Fatalf("unexpected counts %d %d", len(out.Translations), len(out.Downloads))
	}
	req2, _ := http.NewRequest("GET", srv.URL+"/api/history?lang=fr", nil)
	req2.Header.Set("X-API-Key", key)
	resp2, _ := srv.Client().Do(req2)
	var out2 struct {
		Translations []database.SubtitleRecord `json:"translations"`
		Downloads    []database.DownloadRecord `json:"downloads"`
	}
	json.NewDecoder(resp2.Body).Decode(&out2)
	resp2.Body.Close()
	if len(out2.Translations) != 0 || len(out2.Downloads) != 0 {
		t.Fatalf("filter failed")
	}
}

// TestDownload verifies that POST /api/download fetches a subtitle and records history.
func TestDownload(t *testing.T) {
	db, err := database.Open(":memory:")
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	defer db.Close()

	if err := auth.CreateUser(db, "admin", "p", "", "admin"); err != nil {
		t.Fatalf("create admin: %v", err)
	}
	key, err := auth.GenerateAPIKey(db, 1)
	if err != nil {
		t.Fatalf("api key: %v", err)
	}

	// fake subtitle provider
	subSrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("sub"))
	}))
	defer subSrv.Close()
	viper.Set("providers.generic.api_url", subSrv.URL)
	defer viper.Reset()

	h, err := Handler(db)
	if err != nil {
		t.Fatalf("handler: %v", err)
	}
	srv := httptest.NewServer(h)
	defer srv.Close()

	dir := t.TempDir()
	vid := filepath.Join(dir, "video.mkv")
	os.WriteFile(vid, []byte("x"), 0644)

	body := strings.NewReader(`{"provider":"generic","path":"` + vid + `","lang":"en"}`)
	req, _ := http.NewRequest("POST", srv.URL+"/api/download", body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", key)
	resp, err := srv.Client().Do(req)
	if err != nil {
		t.Fatalf("post: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status %d", resp.StatusCode)
	}
	var res struct {
		File string `json:"file"`
	}
	json.NewDecoder(resp.Body).Decode(&res)
	resp.Body.Close()

	out := strings.TrimSuffix(vid, filepath.Ext(vid)) + ".en.srt"
	if res.File != out {
		t.Fatalf("returned %s", res.File)
	}
	if _, err := os.Stat(out); err != nil {
		t.Fatalf("subtitle not written: %v", err)
	}
	recs, err := database.ListDownloads(db)
	if err != nil || len(recs) != 1 {
		t.Fatalf("records %v %d", err, len(recs))
	}
	if recs[0].File != out {
		t.Fatalf("record file %s", recs[0].File)
	}
}

// TestWebhookSonarr verifies that Sonarr webhook triggers subtitle download.
func TestWebhookSonarr(t *testing.T) {
	db, err := database.Open(":memory:")
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	defer db.Close()

	h, err := Handler(db)
	if err != nil {
		t.Fatalf("handler: %v", err)
	}
	srv := httptest.NewServer(h)
	defer srv.Close()

	subSrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("sub")) }))
	defer subSrv.Close()
	viper.Set("providers.generic.api_url", subSrv.URL)
	defer viper.Reset()

	dir := t.TempDir()
	vid := filepath.Join(dir, "file.mkv")
	os.WriteFile(vid, []byte("x"), 0644)

	body := strings.NewReader(`{"provider":"generic","path":"` + vid + `","lang":"en"}`)
	req, _ := http.NewRequest("POST", srv.URL+"/api/webhooks/sonarr", body)
	req.Header.Set("Content-Type", "application/json")
	resp, err := srv.Client().Do(req)
	if err != nil {
		t.Fatalf("post: %v", err)
	}
	if resp.StatusCode != http.StatusNoContent {
		t.Fatalf("status %d", resp.StatusCode)
	}
	out := strings.TrimSuffix(vid, filepath.Ext(vid)) + ".en.srt"
	if _, err := os.Stat(out); err != nil {
		t.Fatalf("subtitle not written")
	}
}

// TestConvertUpload verifies that /api/convert returns an SRT file.
func TestConvertUpload(t *testing.T) {
	db, err := database.Open(":memory:")
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	defer db.Close()
	if err := auth.CreateUser(db, "admin", "p", "", "admin"); err != nil {
		t.Fatalf("create admin: %v", err)
	}
	key, err := auth.GenerateAPIKey(db, 1)
	if err != nil {
		t.Fatalf("api key: %v", err)
	}

	h, err := Handler(db)
	if err != nil {
		t.Fatalf("handler error: %v", err)
	}
	srv := httptest.NewServer(h)
	defer srv.Close()

	buf := &bytes.Buffer{}
	mw := multipart.NewWriter(buf)
	fw, _ := mw.CreateFormFile("file", "in.srt")
	data, _ := os.ReadFile("../../testdata/simple.srt")
	fw.Write(data)
	mw.Close()

	req, _ := http.NewRequest("POST", srv.URL+"/api/convert", buf)
	req.Header.Set("X-API-Key", key)
	req.Header.Set("Content-Type", mw.FormDataContentType())
	resp, err := srv.Client().Do(req)
	if err != nil {
		t.Fatalf("post: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		t.Fatalf("status %d, body: %s", resp.StatusCode, string(body))
	}
	b, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	if len(b) == 0 {
		t.Fatalf("no data returned")
	}
}

// TestTranslateUpload verifies that /api/translate performs translation using Google.
func TestTranslateUpload(t *testing.T) {
	// Mock the Google Translate API
	srvTrans := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"translations":[{"translatedText":"hola"}]}}`))
	}))
	defer srvTrans.Close()
	translator.SetGoogleAPIURL(srvTrans.URL)
	defer translator.SetGoogleAPIURL("https://translation.googleapis.com/language/translate/v2")

	db, err := database.Open(":memory:")
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	defer db.Close()
	if err := auth.CreateUser(db, "admin", "p", "", "admin"); err != nil {
		t.Fatalf("create admin: %v", err)
	}
	key, err := auth.GenerateAPIKey(db, 1)
	if err != nil {
		t.Fatalf("api key: %v", err)
	}

	viper.Set("translate_service", "google")
	viper.Set("google_api_key", "test-key")
	defer viper.Reset()

	h, err := Handler(db)
	if err != nil {
		t.Fatalf("handler error: %v", err)
	}
	srv := httptest.NewServer(h)
	defer srv.Close()

	buf := &bytes.Buffer{}
	mw := multipart.NewWriter(buf)
	fw, _ := mw.CreateFormFile("file", "in.srt")
	data, _ := os.ReadFile("../../testdata/simple.srt")
	fw.Write(data)
	mw.WriteField("lang", "es")
	mw.Close()

	req, _ := http.NewRequest("POST", srv.URL+"/api/translate", buf)
	req.Header.Set("X-API-Key", key)
	req.Header.Set("Content-Type", mw.FormDataContentType())
	resp, err := srv.Client().Do(req)
	if err != nil {
		t.Fatalf("post: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		t.Fatalf("status %d, body: %s", resp.StatusCode, string(body))
	}
	b, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	if !bytes.Contains(b, []byte("hola")) {
		t.Fatalf("translation missing: %s", string(b))
	}
}

// TestProvidersDefault verifies that the providers endpoint only returns
// the embedded provider when no other providers are configured.
func TestProvidersDefault(t *testing.T) {
	db, err := database.Open(":memory:")
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	defer db.Close()

	key := setupTestUser(t, db)

	viper.Reset()
	viper.SetDefault("providers.embedded.enabled", true)

	h, err := Handler(db)
	if err != nil {
		t.Fatalf("handler error: %v", err)
	}
	srv := httptest.NewServer(h)
	defer srv.Close()

	req, _ := http.NewRequest("GET", srv.URL+"/api/providers", nil)
	req.Header.Set("X-API-Key", key)
	resp, err := srv.Client().Do(req)
	if err != nil {
		t.Fatalf("get providers: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status %d", resp.StatusCode)
	}
	var out []ProviderInfo
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		t.Fatalf("decode: %v", err)
	}
	resp.Body.Close()
	if len(out) != 1 {
		t.Fatalf("expected 1 provider, got %d", len(out))
	}
	if out[0].Name != "embedded" || !out[0].Enabled {
		t.Fatalf("unexpected provider %+v", out[0])
	}
}

// TestExtractLanguageFromFilename verifies default detection and pattern overrides.
func TestExtractLanguageFromFilename(t *testing.T) {
	ResetLanguagePatterns()

	cases := map[string]string{
		"movie.es.srt":      "Spanish",
		"movie.zh-cn.srt":   "Chinese (Simplified)",
		"movie.zh-hk.srt":   "Chinese (Hong Kong)",
		"movie.zh-hant.srt": "Chinese (Traditional)",
		"movie.zh-tw.srt":   "Chinese (Taiwan)",
	}
	for name, want := range cases {
		if lang := extractLanguageFromFilename(name); lang != want {
			t.Fatalf("%s expected %s, got %s", name, want, lang)
		}
	}

	SetLanguagePatterns(map[string]string{"xx": "TestLang"})
	defer ResetLanguagePatterns()
	if lang := extractLanguageFromFilename("movie.xx.srt"); lang != "TestLang" {
		t.Fatalf("override failed: %s", lang)
	}
}

// setupTestUser creates a test user with an API key and returns the key.
func setupTestUser(t *testing.T, db *sql.DB) string {
	testutil.MustNoError(t, "create admin", auth.CreateUser(db, "admin", "p", "", "admin"))
	return testutil.MustGet(t, "api key", func() (string, error) { return auth.GenerateAPIKey(db, 1) })
}

// TestSecurityHeaders ensures that the server sets common security headers.
func TestSecurityHeaders(t *testing.T) {
	db, err := database.Open(":memory:")
	testutil.MustNoError(t, "open db", err)
	defer db.Close()

	key := setupTestUser(t, db)

	h, err := Handler(db)
	testutil.MustNoError(t, "handler", err)
	srv := httptest.NewServer(h)
	defer srv.Close()

	req, _ := http.NewRequest("GET", srv.URL+"/api/history", nil)
	req.Header.Set("X-API-Key", key)
	resp := testutil.MustGet(t, "request", func() (*http.Response, error) { return srv.Client().Do(req) })

	headers := resp.Header
	for _, name := range []string{"X-Frame-Options", "X-Content-Type-Options", "Content-Security-Policy", "Strict-Transport-Security"} {
		if headers.Get(name) == "" {
			t.Fatalf("%s header missing", name)
		}
	}
}
