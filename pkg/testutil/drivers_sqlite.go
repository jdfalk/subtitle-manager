//go:build sqlite
// +build sqlite

// file: pkg/testutil/drivers_sqlite.go
// version: 1.0.0
// guid: 5c4d3e2f-1a0b-6c5d-8f7e-2b1c0d9e8a7f

// Package testutil provides test utilities and helpers.
// This file contains SQLite3 driver imports for testing with CGO-enabled builds.
package testutil

import (
	// Import SQLite3 driver for CGO-enabled test builds
	_ "github.com/mattn/go-sqlite3"
)
