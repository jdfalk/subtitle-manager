package webserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/jdfalk/subtitle-manager/pkg/database"
)

func TestSearchHistoryHandlerCycle(t *testing.T) {
	skipIfNoSQLite(t)
	db, err := database.Open(":memory:")
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	defer db.Close()

	handler := searchHistoryHandler(db)

	// Initial GET should return empty slice
	req, _ := http.NewRequest("GET", "/api/search/history", nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("get status %d", rr.Code)
	}
	var items []SearchHistoryItem
	if err := json.NewDecoder(rr.Body).Decode(&items); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if len(items) != 0 {
		t.Fatalf("expected empty history")
	}

	// Save a history item
	hist := SearchHistoryItem{
		Query:     SearchRequest{Providers: []string{"opensubtitles"}, MediaPath: "a.mkv", Language: "en"},
		Results:   2,
		Timestamp: time.Now(),
	}
	data, _ := json.Marshal(hist)
	req2, _ := http.NewRequest("POST", "/api/search/history", bytes.NewBuffer(data))
	req2.Header.Set("Content-Type", "application/json")
	rr2 := httptest.NewRecorder()
	handler.ServeHTTP(rr2, req2)
	if rr2.Code != http.StatusOK {
		t.Fatalf("post status %d", rr2.Code)
	}

	// Verify item saved
	req3, _ := http.NewRequest("GET", "/api/search/history", nil)
	rr3 := httptest.NewRecorder()
	handler.ServeHTTP(rr3, req3)
	if rr3.Code != http.StatusOK {
		t.Fatalf("get after post status %d", rr3.Code)
	}
	if err := json.NewDecoder(rr3.Body).Decode(&items); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if len(items) != 1 {
		t.Fatalf("expected 1 history item, got %d", len(items))
	}

	// Delete history
	req4, _ := http.NewRequest("DELETE", "/api/search/history", nil)
	rr4 := httptest.NewRecorder()
	handler.ServeHTTP(rr4, req4)
	if rr4.Code != http.StatusOK {
		t.Fatalf("delete status %d", rr4.Code)
	}

	req5, _ := http.NewRequest("GET", "/api/search/history", nil)
	rr5 := httptest.NewRecorder()
	handler.ServeHTTP(rr5, req5)
	if err := json.NewDecoder(rr5.Body).Decode(&items); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if len(items) != 0 {
		t.Fatalf("expected empty history after delete")
	}
}
