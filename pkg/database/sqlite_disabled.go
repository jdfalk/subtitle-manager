//go:build !sqlite
// +build !sqlite

// file: pkg/database/sqlite_disabled.go
// version: 1.0.0
// guid: 8f9e7d6c-5b4a-3f2e-1d0c-9b8a7c6f5e4d

package database

import (
	"database/sql"
	"fmt"
)

// OpenSQLStore returns an error when SQLite support is not compiled in.
// This function is only a stub when building without the 'sqlite' tag.
func OpenSQLStore(path string) (*SQLStore, error) {
	return nil, fmt.Errorf("SQLite support not available - application was built without CGO/SQLite support. Use Pebble database backend instead")
}

// Open returns an error when SQLite support is not compiled in.
// This function is only a stub when building without the 'sqlite' tag.
func Open(path string) (*sql.DB, error) {
	return nil, fmt.Errorf("SQLite support not available - application was built without CGO/SQLite support. Use Pebble database backend instead")
}
