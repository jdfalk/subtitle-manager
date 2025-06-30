// file: pkg/backups/compression.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174002

package backups

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
)

// GzipCompression implements the Compression interface using gzip.
type GzipCompression struct{}

// NewGzipCompression creates a new gzip compression instance.
func NewGzipCompression() *GzipCompression {
	return &GzipCompression{}
}

// Compress compresses the input data using gzip.
func (gc *GzipCompression) Compress(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	writer := gzip.NewWriter(&buf)

	if _, err := writer.Write(data); err != nil {
		writer.Close()
		return nil, fmt.Errorf("failed to compress data: %w", err)
	}

	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("failed to close gzip writer: %w", err)
	}

	return buf.Bytes(), nil
}

// Decompress decompresses the input data using gzip.
func (gc *GzipCompression) Decompress(data []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("failed to create gzip reader: %w", err)
	}
	defer reader.Close()

	decompressed, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to decompress data: %w", err)
	}

	return decompressed, nil
}