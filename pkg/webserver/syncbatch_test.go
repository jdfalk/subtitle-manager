package webserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/jdfalk/subtitle-manager/pkg/syncer"
)

func TestSyncBatchHandler(t *testing.T) {
	dir := t.TempDir()
	data, err := os.ReadFile("../../testdata/simple.srt")
	if err != nil {
		t.Fatalf("read testdata: %v", err)
	}
	in1 := filepath.Join(dir, "a.srt")
	in2 := filepath.Join(dir, "b.srt")
	if err := os.WriteFile(in1, data, 0644); err != nil {
		t.Fatalf("write in1: %v", err)
	}
	if err := os.WriteFile(in2, data, 0644); err != nil {
		t.Fatalf("write in2: %v", err)
	}
	out1 := filepath.Join(dir, "out1.srt")
	out2 := filepath.Join(dir, "out2.srt")

	reqBody := struct {
		Items   []syncer.BatchItem `json:"items"`
		Options syncer.Options     `json:"options"`
	}{
		Items: []syncer.BatchItem{
			{Media: "dummy1.mkv", Subtitle: in1, Output: out1},
			{Media: "dummy2.mkv", Subtitle: in2, Output: out2},
		},
		Options: syncer.Options{},
	}
	b, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	handler := syncBatchHandler()
	req := httptest.NewRequest(http.MethodPost, "/api/sync/batch", bytes.NewReader(b))
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status: %d", w.Code)
	}
	var resp struct {
		Results []struct {
			Output string `json:"output"`
			Error  string `json:"error"`
		} `json:"results"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if len(resp.Results) != 2 {
		t.Fatalf("results len %d", len(resp.Results))
	}
	for i, p := range []string{out1, out2} {
		if resp.Results[i].Error != "" {
			t.Fatalf("item %d err %s", i, resp.Results[i].Error)
		}
		fi, err := os.Stat(p)
		if err != nil || fi.Size() == 0 {
			t.Fatalf("output %s missing", p)
		}
	}
}

func TestSyncBatchHandlerInvalidJSON(t *testing.T) {
	handler := syncBatchHandler()
	req := httptest.NewRequest(http.MethodPost, "/api/sync/batch", bytes.NewReader([]byte("invalid")))
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}
