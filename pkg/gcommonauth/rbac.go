// file: pkg/gcommonauth/rbac.go
// version: 1.0.0
// guid: 46f03c5a-8a80-4b44-bb7b-559a3cd93863
package gcommonauth

import "database/sql"

// CheckPermission returns true if the role has the given permission.
func CheckPermission(db *sql.DB, userID int64, permission string) (bool, error) {
	var perm string
	row := db.QueryRow(`SELECT p.permission FROM users u JOIN permissions p ON u.role = p.role WHERE u.id = ?`, userID)
	if err := row.Scan(&perm); err != nil {
		return false, err
	}
	levels := map[string]int{"read": 1, "basic": 2, "all": 3}
	if levels[perm] >= levels[permission] {
		return true, nil
	}
	return false, nil
}
