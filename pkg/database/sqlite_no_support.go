//go:build !sqlite
// +build !sqlite

// file: pkg/database/sqlite_no_support.go
// version: 1.0.0
// guid: 159e7e12-6a16-4782-bf87-5bd01f3814b9

package database

// HasSQLite reports whether the binary was built with SQLite support.
func HasSQLite() bool {
	return false
}
