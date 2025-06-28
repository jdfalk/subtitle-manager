//go:build sqlite
// +build sqlite

// file: pkg/database/drivers_sqlite.go
// version: 1.0.0
// guid: 3a2b1c9d-8e7f-4c5d-6b4a-9c8b7a6d5e4f

// Package database provides database abstraction layer for subtitle-manager.
// This file contains SQLite3 driver imports that require CGO compilation.
package database

import (
	// Import SQLite3 driver for CGO-enabled builds
	_ "github.com/mattn/go-sqlite3"
)
