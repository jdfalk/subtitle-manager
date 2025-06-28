//go:build !sqlite
// +build !sqlite

// file: pkg/database/drivers_nosqlite.go
// version: 1.0.0
// guid: 4b3c2d1e-9f8a-5d6c-7e5f-1a9b8c7d6e5f

// Package database provides database abstraction layer for subtitle-manager.
// This file is used for pure Go builds without SQLite3 support (no CGO required).
package database

// Note: This file intentionally has no imports to avoid CGO dependencies.
// When building without the 'sqlite' tag, SQLite3 functionality will not be available,
// but Pebble database backend will still work normally.
