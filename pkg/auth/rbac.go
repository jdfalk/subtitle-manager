package auth

import "database/sql"

// CheckPermission returns true if the role has the given permission.
func CheckPermission(db *sql.DB, userID int64, permission string) (bool, error) {
	var role string
	row := db.QueryRow(`SELECT role FROM users WHERE id = ?`, userID)
	if err := row.Scan(&role); err != nil {
		return false, err
	}
	var count int
	row = db.QueryRow(`SELECT COUNT(1) FROM permissions WHERE role = ? AND permission = ?`, role, permission)
	if err := row.Scan(&count); err != nil {
		return false, err
	}
	return count > 0, nil
}
