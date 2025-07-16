// file: pkg/gcommonauth/auth.go
// version: 1.0.0
// guid: ec8f6814-42cc-41ab-956a-a0ee798664a4
package gcommonauth

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser inserts a new user with a hashed password.
// CreateUser inserts a new user with a hashed password and automatically
// generates an API key for the user. The generated API key is stored in the
// database but not returned to the caller.
func CreateUser(db *sql.DB, username, password, email, role string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	res, err := db.Exec(`INSERT INTO users (username, password_hash, email, role, created_at) VALUES (?, ?, ?, ?, ?)`,
		username, string(hash), email, role, time.Now())
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	_, err = GenerateAPIKey(db, id)
	return err
}

// AuthenticateUser verifies credentials and returns the user ID.
func AuthenticateUser(db *sql.DB, username, password string) (int64, error) {
	var id int64
	var hash string
	row := db.QueryRow(`SELECT id, password_hash FROM users WHERE username = ?`, username)
	if err := row.Scan(&id, &hash); err != nil {
		return 0, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return 0, err
	}
	return id, nil
}

// GenerateSession creates a new session token for the user.
func GenerateSession(db *sql.DB, userID int64, duration time.Duration) (string, error) {
	token := uuid.NewString()
	expires := time.Now().Add(duration)
	_, err := db.Exec(`INSERT INTO sessions (user_id, token, expires_at, created_at) VALUES (?, ?, ?, ?)`,
		userID, token, expires, time.Now())
	if err != nil {
		return "", err
	}
	return token, nil
}

// ValidateSession returns the user ID for a valid session token.
func ValidateSession(db *sql.DB, token string) (int64, error) {
	var userID int64
	var expires time.Time
	row := db.QueryRow(`SELECT user_id, expires_at FROM sessions WHERE token = ?`, token)
	if err := row.Scan(&userID, &expires); err != nil {
		return 0, err
	}
	if time.Now().After(expires) {
		return 0, sql.ErrNoRows
	}
	return userID, nil
}

// GenerateAPIKey creates a new API key for the user.
func GenerateAPIKey(db *sql.DB, userID int64) (string, error) {
	key := uuid.NewString()
	_, err := db.Exec(`INSERT INTO api_keys (user_id, key, created_at) VALUES (?, ?, ?)`,
		userID, key, time.Now())
	if err != nil {
		return "", err
	}
	return key, nil
}

// ValidateAPIKey returns the user ID associated with an API key.
func ValidateAPIKey(db *sql.DB, key string) (int64, error) {
	var userID int64
	row := db.QueryRow(`SELECT user_id FROM api_keys WHERE key = ?`, key)
	if err := row.Scan(&userID); err != nil {
		return 0, err
	}
	return userID, nil
}

// ResetPassword generates a new password for the specified user ID, updates the
// stored hash and returns the plaintext password. A new API key is also
// generated and returned.
func ResetPassword(db *sql.DB, userID int64) (string, string, error) {
	pass := uuid.NewString()[:12]
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", "", err
	}
	if _, err = db.Exec(`UPDATE users SET password_hash = ? WHERE id = ?`, string(hash), userID); err != nil {
		return "", "", err
	}
	key, err := GenerateAPIKey(db, userID)
	if err != nil {
		return "", "", err
	}
	return pass, key, nil
}

// GetOrCreateUser returns the existing user ID for email or inserts a new user
// with the provided username, email and role if none exists. The password is
// left empty for OAuth2 users.
func GetOrCreateUser(db *sql.DB, username, email, role string) (int64, error) {
	var id int64
	row := db.QueryRow(`SELECT id FROM users WHERE email = ?`, email)
	err := row.Scan(&id)
	if err == nil {
		return id, nil
	}
	if err != sql.ErrNoRows {
		return 0, err
	}
	res, err := db.Exec(`INSERT INTO users (username, password_hash, email, role, created_at) VALUES (?, '', ?, ?, ?)`,
		username, email, role, time.Now())
	if err != nil {
		return 0, err
	}
	id, err = res.LastInsertId()
	return id, err
}

// SetUserRole updates the role for the specified username.
func SetUserRole(db *sql.DB, username, role string) error {
	_, err := db.Exec(`UPDATE users SET role = ? WHERE username = ?`, role, username)
	return err
}

// GenerateOneTimeToken creates a single-use login token for the user.
// The token expires after the provided duration.
func GenerateOneTimeToken(db *sql.DB, userID int64, duration time.Duration) (string, error) {
	token := uuid.NewString()
	expires := time.Now().Add(duration)
	_, err := db.Exec(`INSERT INTO login_tokens (user_id, token, expires_at, used, created_at) VALUES (?, ?, ?, 0, ?)`,
		userID, token, expires, time.Now())
	if err != nil {
		return "", err
	}
	return token, nil
}

// ConsumeOneTimeToken validates and marks the token as used.
// It returns the associated user ID when successful.
func ConsumeOneTimeToken(db *sql.DB, token string) (int64, error) {
	var userID int64
	var expires time.Time
	var used int
	row := db.QueryRow(`SELECT user_id, expires_at, used FROM login_tokens WHERE token = ?`, token)
	if err := row.Scan(&userID, &expires, &used); err != nil {
		return 0, err
	}
	if used == 1 || time.Now().After(expires) {
		return 0, sql.ErrNoRows
	}
	if _, err := db.Exec(`UPDATE login_tokens SET used = 1 WHERE token = ?`, token); err != nil {
		return 0, err
	}
	return userID, nil
}

// InvalidateSession removes a session token from the database, effectively logging out the user.
func InvalidateSession(db *sql.DB, token string) error {
	_, err := db.Exec(`DELETE FROM sessions WHERE token = ?`, token)
	return err
}

// InvalidateUserSessions removes all session tokens for a specific user.
// This is useful for force-logout scenarios or when a user's permissions change.
func InvalidateUserSessions(db *sql.DB, userID int64) error {
	_, err := db.Exec(`DELETE FROM sessions WHERE user_id = ?`, userID)
	return err
}

// CleanupExpiredSessions removes all expired session tokens from the database.
// This should be called periodically to prevent the sessions table from growing indefinitely.
func CleanupExpiredSessions(db *sql.DB) error {
	_, err := db.Exec(`DELETE FROM sessions WHERE expires_at < ?`, time.Now())
	return err
}

// User represents an account in the system. ID is stored as a string for
// convenience when printing. The JSON struct tags ensure the API returns fields
// in lowercase so the frontend can parse them correctly.
type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

// ListUsers returns all users ordered by ID.
func ListUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query(`SELECT id, username, email, role, created_at FROM users ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []User
	for rows.Next() {
		var u User
		var id int64
		if err := rows.Scan(&id, &u.Username, &u.Email, &u.Role, &u.CreatedAt); err != nil {
			return nil, err
		}
		u.ID = strconv.FormatInt(id, 10)
		out = append(out, u)
	}
	return out, rows.Err()
}
