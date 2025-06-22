// file: pkg/providers/opensubtitles/hash.go
package opensubtitles

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// fileHash calculates the OpenSubtitles file hash.
// fileHashFunc is used internally to compute the movie hash.
// It is defined as a variable to allow tests to replace it.
var fileHashFunc = realFileHash

// realFileHash calculates the OpenSubtitles file hash.
// The provided path is validated to ensure it doesn't contain path traversal attempts.
func realFileHash(path string) (uint64, int64, error) {
	// Clean the path to resolve .. and . elements
	cleanPath := filepath.Clean(path)

	// Security check: ensure no path traversal components remain
	if strings.Contains(cleanPath, "..") {
		return 0, 0, fmt.Errorf("path traversal detected: %s", path)
	}

	// Convert to absolute path for consistent checking
	absPath, err := filepath.Abs(cleanPath)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid file path: %s", path)
	}

	f, err := os.Open(absPath)
	if err != nil {
		return 0, 0, err
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return 0, 0, err
	}
	size := fi.Size()
	const chunk = 64 * 1024
	buf := make([]byte, chunk)
	var h uint64
	h += uint64(size)
	if _, err := io.ReadFull(f, buf); err != nil {
		return 0, 0, err
	}
	for i := 0; i < chunk/8; i++ {
		h += binary.LittleEndian.Uint64(buf[i*8:])
	}
	if size > chunk {
		if _, err := f.Seek(-chunk, io.SeekEnd); err != nil {
			return 0, 0, err
		}
		if _, err := io.ReadFull(f, buf); err != nil {
			return 0, 0, err
		}
		for i := 0; i < chunk/8; i++ {
			h += binary.LittleEndian.Uint64(buf[i*8:])
		}
	}
	return h, size, nil
}
