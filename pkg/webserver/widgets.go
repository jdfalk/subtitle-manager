// file: pkg/webserver/widgets.go
package webserver

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/jdfalk/subtitle-manager/pkg/database"
)

// dashboardLayoutHandler returns a handler for per-user dashboard layout
// preferences.
//
// Parameters:
//
//	db - application database handle
//
// The handler supports GET and POST methods and expects the user ID in the
// request context. GET responds with the stored layout JSON. POST updates it.
// Unsupported methods result in 405 Method Not Allowed.
func dashboardLayoutHandler(db *sql.DB) http.Handler {
	type response struct {
		Layout string `json:"layout"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, _ := r.Context().Value(userIDContextKey).(int64)
		switch r.Method {
		case http.MethodGet:
			layout, err := database.GetDashboardLayout(db, id)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(response{Layout: layout})
		case http.MethodPost:
			var in response
			if err := json.NewDecoder(r.Body).Decode(&in); err != nil || in.Layout == "" {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if err := database.SetDashboardLayout(db, id, in.Layout); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusNoContent)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}

// Widget describes a dashboard widget available to users.
//
// Fields:
//
//	ID          - unique identifier for the widget
//	Name        - display name presented in the UI
//	Description - short explanation of the widget purpose
type Widget struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// availableWidgets defines the widgets shipped with the application.
var availableWidgets = []Widget{
	{ID: "scan", Name: "Scan", Description: "Subtitle scanning controls"},
	{ID: "system_info", Name: "System Info", Description: "System information"},
	{ID: "quick_links", Name: "Quick Links", Description: "Navigation shortcuts"},
}

// widgetsListHandler returns a handler exposing the list of available widgets.
// Only GET is supported.
func widgetsListHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(availableWidgets)
	})
}
