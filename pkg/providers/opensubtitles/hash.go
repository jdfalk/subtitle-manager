// file: pkg/providers/opensubtitles/hash.go
package opensubtitles

import (
	"encoding/binary"
	"io"
	"os"

	"github.com/jdfalk/subtitle-manager/pkg/security"
)

// fileHash calculates the OpenSubtitles file hash.
// fileHashFunc is used internally to compute the movie hash.
// It is defined as a variable to allow tests to replace it.
var fileHashFunc = realFileHash

// realFileHash calculates the OpenSubtitles file hash.
// The provided path is validated to ensure it doesn't contain path traversal attempts.
func realFileHash(path string) (uint64, int64, error) {
	sanitizedPath, err := security.ValidateAndSanitizePath(path)
	if err != nil {
		return 0, 0, err
	}

	f, err := os.Open(sanitizedPath)
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
