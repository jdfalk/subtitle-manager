package subtitles

import "testing"

func TestConvertToSRT(t *testing.T) {
	out, err := ConvertToSRT("../../testdata/simple.srt")
	if err != nil {
		t.Fatalf("convert: %v", err)
	}
	if len(out) == 0 {
		t.Fatal("no output")
	}
	if string(out) == "" {
		t.Fatal("empty output")
	}
}
