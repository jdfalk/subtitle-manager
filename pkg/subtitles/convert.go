package subtitles

import (
	"bytes"

	"github.com/asticode/go-astisub"
)

// ConvertToSRT reads a subtitle file and converts it to SRT format.
// It returns the resulting SRT bytes.
func ConvertToSRT(path string) ([]byte, error) {
	sub, err := astisub.OpenFile(path)
	if err != nil {
		return nil, err
	}
	buf := &bytes.Buffer{}
	if err := sub.WriteToSRT(buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
