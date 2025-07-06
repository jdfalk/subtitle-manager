//go:build sqlite
// +build sqlite

// file: pkg/database/sqlite_support.go
// version: 1.0.0
// guid: 91c8823b-fd29-4ea0-9732-e3b75d95d7e9

package database

// HasSQLite reports whether the binary was built with SQLite support.
func HasSQLite() bool {
	return true
}
