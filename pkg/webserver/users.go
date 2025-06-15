package webserver

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/spf13/viper"

	"github.com/jdfalk/subtitle-manager/pkg/auth"
	"github.com/jdfalk/subtitle-manager/pkg/notifications"
)

// usersHandler returns a list of all users.
func usersHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		users, err := auth.ListUsers(db)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(users)
	})
}

// userResetHandler resets the password for the specified user ID, generates a new
// API key and emails the credentials using the configured notification service.
func userResetHandler(db *sql.DB) http.Handler {
	type response struct {
		Password string `json:"password"`
		APIKey   string `json:"api_key"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/api/users/"), "/")
		if len(parts) != 2 || parts[1] != "reset" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		id, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil || id == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		pass, key, err := auth.ResetPassword(db, id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var email, username string
		row := db.QueryRow(`SELECT email, username FROM users WHERE id = ?`, id)
		_ = row.Scan(&email, &username)
		if email != "" {
			svc := notifications.New(
				viper.GetString("notifications.discord_webhook"),
				viper.GetString("notifications.telegram_token"),
				viper.GetString("notifications.telegram_chat_id"),
				viper.GetString("notifications.email_url"),
			)
			msg := fmt.Sprintf("Credentials for %s\nPassword: %s\nAPI Key: %s", username, pass, key)
			_ = svc.Send(context.Background(), msg)
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(response{Password: pass, APIKey: key})
	})
}
