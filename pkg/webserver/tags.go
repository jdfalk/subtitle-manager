// file: pkg/webserver/tags.go
package webserver

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/jdfalk/subtitle-manager/pkg/database"
)

// tagsHandler manages global tags.
func tagsHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			tags, err := database.ListTags(db)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			_ = json.NewEncoder(w).Encode(tags)
		case http.MethodPost:
			var in struct {
				Name string `json:"name"`
			}
			if err := json.NewDecoder(r.Body).Decode(&in); err != nil || in.Name == "" {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if err := database.InsertTag(db, in.Name); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}

// tagDeleteHandler removes a tag by ID.
func tagDeleteHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		idStr := strings.TrimPrefix(r.URL.Path, "/api/tags/")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := database.DeleteTag(db, id); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})
}

// userTagsHandler manages tag assignments for a user.
func userTagsHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/api/users/"), "/")
		if len(parts) < 2 || parts[1] != "tags" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		id, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		switch r.Method {
		case http.MethodGet:
			tags, err := database.ListTagsForUser(db, id)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			_ = json.NewEncoder(w).Encode(tags)
		case http.MethodPost:
			var in struct {
				TagID int64 `json:"tag_id"`
			}
			if err := json.NewDecoder(r.Body).Decode(&in); err != nil || in.TagID == 0 {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if err := database.AssignTagToUser(db, id, in.TagID); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
		case http.MethodDelete:
			var in struct {
				TagID int64 `json:"tag_id"`
			}
			if err := json.NewDecoder(r.Body).Decode(&in); err != nil || in.TagID == 0 {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if err := database.RemoveTagFromUser(db, id, in.TagID); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusNoContent)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}

// mediaTagsHandler manages tag assignments for a media item.
func mediaTagsHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/api/media/"), "/")
		if len(parts) < 2 || parts[1] != "tags" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		id, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		switch r.Method {
		case http.MethodGet:
			tags, err := database.ListTagsForMedia(db, id)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			_ = json.NewEncoder(w).Encode(tags)
		case http.MethodPost:
			var in struct {
				TagID int64 `json:"tag_id"`
			}
			if err := json.NewDecoder(r.Body).Decode(&in); err != nil || in.TagID == 0 {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if err := database.AssignTagToMedia(db, id, in.TagID); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
		case http.MethodDelete:
			var in struct {
				TagID int64 `json:"tag_id"`
			}
			if err := json.NewDecoder(r.Body).Decode(&in); err != nil || in.TagID == 0 {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if err := database.RemoveTagFromMedia(db, id, in.TagID); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusNoContent)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}
