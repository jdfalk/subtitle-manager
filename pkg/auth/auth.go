package auth

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser inserts a new user with a hashed password.
func CreateUser(db *sql.DB, username, password, email, role string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	_, err = db.Exec(`INSERT INTO users (username, password_hash, email, role, created_at) VALUES (?, ?, ?, ?, ?)`,
		username, string(hash), email, role, time.Now())
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
