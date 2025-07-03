// file: pkg/storage/errors.go
// version: 1.0.0
// guid: 23456789-01bc-def2-3456-789012345678

package storage

import "errors"

var (
	// ErrUnsupportedProvider is returned when an unsupported storage provider is requested.
	ErrUnsupportedProvider = errors.New("unsupported storage provider")

	// ErrNotFound is returned when a requested object is not found.
	ErrNotFound = errors.New("object not found")

	// ErrAlreadyExists is returned when trying to create an object that already exists.
	ErrAlreadyExists = errors.New("object already exists")

	// ErrInvalidKey is returned when an invalid key/path is provided.
	ErrInvalidKey = errors.New("invalid key or path")

	// ErrConfigurationMissing is returned when required configuration is missing.
	ErrConfigurationMissing = errors.New("required configuration missing")

	// ErrConnectionFailed is returned when connection to storage provider fails.
	ErrConnectionFailed = errors.New("connection to storage provider failed")
)
