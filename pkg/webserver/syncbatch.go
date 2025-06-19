// file: pkg/webserver/syncbatch.go
package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/jdfalk/subtitle-manager/pkg/syncer"
)

// syncBatchHandler processes a batch of subtitle synchronization requests.
// Each request contains the media path, subtitle path and optional output path.
// Synchronization options mirror those of the single sync operation.
func syncBatchHandler() http.Handler {
	type item struct {
		Media    string `json:"media"`
		Subtitle string `json:"subtitle"`
		Output   string `json:"output"`
	}
	type req struct {
		Items   []item         `json:"items"`
		Options syncer.Options `json:"options"`
	}
	type result struct {
		Output string `json:"output"`
		Error  string `json:"error,omitempty"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		var q req
		if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		batch := make([]syncer.BatchItem, len(q.Items))
		for i, it := range q.Items {
			batch[i] = syncer.BatchItem{Media: it.Media, Subtitle: it.Subtitle, Output: it.Output}
		}
		errs := syncer.SyncBatch(batch, q.Options)
		resp := struct {
			Results []result `json:"results"`
		}{Results: make([]result, len(batch))}
		for i := range batch {
			resp.Results[i].Output = batch[i].Output
			if resp.Results[i].Output == "" {
				resp.Results[i].Output = batch[i].Subtitle
			}
			if errs[i] != nil {
				resp.Results[i].Error = errs[i].Error()
			}
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	})
}
