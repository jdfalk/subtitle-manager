//go:build !sqlite
// +build !sqlite

// file: pkg/testutil/drivers_nosqlite.go
// version: 1.0.0
// guid: 6d5e4f3a-2b1c-7d6e-9a8f-3c2b1d0e9f8a

// Package testutil provides test utilities and helpers.
// This file is used for pure Go test builds without SQLite3 support (no CGO required).
package testutil

// Note: This file intentionally has no imports to avoid CGO dependencies.
// When building tests without the 'sqlite' tag, SQLite3 testing functionality
// will not be available, but Pebble database testing will still work normally.
