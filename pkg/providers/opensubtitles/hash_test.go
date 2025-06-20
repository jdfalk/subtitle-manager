package opensubtitles

import (
	"os"
	"testing"
)

func TestFileHash(t *testing.T) {
	f, err := os.CreateTemp("", "hash-test")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	data := make([]byte, 200000)
	for i := range data {
		data[i] = byte(i % 255)
	}
	if _, err := f.Write(data); err != nil {
		t.Fatal(err)
	}
	defer os.Remove(f.Name())

	h, size, err := realFileHash(f.Name())
	if err != nil {
		t.Fatal(err)
	}
	if size != int64(len(data)) {
		t.Fatalf("size mismatch: %d != %d", size, len(data))
	}
	if h == 0 {
		t.Fatal("hash should not be zero")
	}
}
