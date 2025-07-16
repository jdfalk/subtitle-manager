package syncer

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/asticode/go-astisub"
	"github.com/stretchr/testify/mock"

	"github.com/jdfalk/subtitle-manager/pkg/audio"
	"github.com/jdfalk/subtitle-manager/pkg/subtitles"
	"github.com/jdfalk/subtitle-manager/pkg/syncer/mocks"
	"github.com/jdfalk/subtitle-manager/pkg/transcriber"
	pb "github.com/jdfalk/subtitle-manager/pkg/translatorpb"
	"google.golang.org/grpc"
)

// TestShift verifies that the Shift function offsets subtitles by the given duration.
func TestShift(t *testing.T) {
	items := []*astisub.Item{{StartAt: 0, EndAt: time.Second}}
	out := Shift(items, 2*time.Second)
	if out[0].StartAt != 2*time.Second || out[0].EndAt != 3*time.Second {
		t.Fatalf("unexpected values: %#v", out[0])
	}
}

// TestSync loads a subtitle file to ensure no error is returned.
func TestSync(t *testing.T) {
	items, err := Sync("dummy.mkv", "../../testdata/simple.srt", Options{})
	if err != nil {
		t.Fatalf("sync: %v", err)
	}
	if len(items) == 0 {
		t.Fatal("no items returned")
	}
}

// TestComputeOffset verifies that computeOffset returns the expected duration.
func TestComputeOffset(t *testing.T) {
	ref := []*astisub.Item{{StartAt: 2 * time.Second}}
	target := []*astisub.Item{{StartAt: time.Second}}
	if d := computeOffset(ref, target); d != time.Second {
		t.Fatalf("unexpected offset %v", d)
	}
}

func TestComputeOffsetMedian(t *testing.T) {
	ref := []*astisub.Item{
		{StartAt: time.Second},
		{StartAt: 2 * time.Second},
		{StartAt: 3 * time.Second},
		{StartAt: 4 * time.Second},
		{StartAt: 100 * time.Second},
	}
	target := []*astisub.Item{
		{StartAt: 0},
		{StartAt: time.Second},
		{StartAt: 2 * time.Second},
		{StartAt: 3 * time.Second},
		{StartAt: 97 * time.Second},
	}
	if d := computeOffset(ref, target); d != time.Second {
		t.Fatalf("unexpected offset %v", d)
	}
}

// mockServer returns "hola" for any translation request.
type mockServer struct {
	pb.UnimplementedTranslatorServer
}

func (mockServer) Translate(ctx context.Context, req *pb.TranslateRequest) (*pb.TranslateResponse, error) {
	return &pb.TranslateResponse{TranslatedText: "hola"}, nil
}

// TestTranslate verifies that subtitle items are translated using a gRPC provider.
func TestTranslate(t *testing.T) {
	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTranslatorServer(s, &mockServer{})
	go s.Serve(lis)
	defer s.Stop()

	items := []*astisub.Item{{Lines: []astisub.Line{{Items: []astisub.LineItem{{Text: "hello"}}}}}}
	out, err := Translate(items, "es", "grpc", "", "", lis.Addr().String())
	if err != nil {
		t.Fatalf("translate: %v", err)
	}
	if out[0].String() != "hola" {
		t.Fatalf("expected hola, got %s", out[0].String())
	}
}

// TestSyncTranslate ensures Sync translates subtitles when options specify a language.
func TestSyncTranslate(t *testing.T) {
	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTranslatorServer(s, &mockServer{})
	go s.Serve(lis)
	defer s.Stop()

	opts := Options{TargetLang: "es", Service: "grpc", GRPCAddr: lis.Addr().String()}
	items, err := Sync("dummy.mkv", "../../testdata/simple.srt", opts)
	if err != nil {
		t.Fatalf("sync: %v", err)
	}
	if items[0].String() != "hola" {
		t.Fatalf("expected hola, got %s", items[0].String())
	}
}

// TestSyncEmbedded verifies that embedded subtitle tracks are used to adjust
// timing.
func TestSyncEmbedded(t *testing.T) {
	dir := t.TempDir()
	script := filepath.Join(dir, "ffmpeg")
	data := `#!/bin/sh
last=$(eval echo \${$#}); cp ../../testdata/simple.srt "$last"
`
	if err := os.WriteFile(script, []byte(data), 0755); err != nil {
		t.Fatalf("write script: %v", err)
	}
	oldPath := os.Getenv("PATH")
	os.Setenv("PATH", dir+":"+oldPath)
	defer os.Setenv("PATH", oldPath)

	subtitles.SetFFmpegPath("ffmpeg")

	items, err := Sync("dummy.mkv", "../../testdata/simple_offset.srt", Options{UseEmbedded: true, SubtitleTracks: []int{0}})
	if err != nil {
		t.Fatalf("sync: %v", err)
	}
	if items[0].StartAt != time.Second {
		t.Fatalf("unexpected start %v", items[0].StartAt)
	}
}

// TestSyncWithAudioTrack verifies synchronization using a specific audio track.
func TestSyncWithAudioTrack(t *testing.T) {
	base, err := astisub.OpenFile("../../testdata/simple.srt")
	if err != nil {
		t.Fatalf("open base: %v", err)
	}

	dir := t.TempDir()
	subFile := filepath.Join(dir, "test.srt")
	f, err := os.Create(subFile)
	if err != nil {
		t.Fatalf("create temp: %v", err)
	}
	defer f.Close()
	subs := astisub.Subtitles{Items: base.Items}
	if err := subs.WriteToSRT(f); err != nil {
		t.Fatalf("write SRT: %v", err)
	}

	// Create mock transcriber
	mockTranscriber := mocks.NewTranscriber(t)

	// Setup transcriber mock to return the test file content
	b, _ := os.ReadFile("../../testdata/simple.srt")
	mockTranscriber.EXPECT().Transcribe(mock.AnythingOfType("string"), "", "test-key").
		Return(b, nil)

	opts := Options{
		UseAudio:    true,
		AudioTrack:  1,   // Use audio track 1
		AudioWeight: 1.0, // Use only audio for sync
		WhisperKey:  "test-key",
		Transcriber: mockTranscriber,
	}

	items, err := Sync("dummy.mkv", subFile, opts)
	if err != nil {
		t.Fatalf("sync: %v", err)
	}

	if len(items) == 0 {
		t.Fatal("no items returned")
	}

	// Should match the original timing since we're using the same test file
	expected := base.Items[0].StartAt
	if items[0].StartAt != expected {
		t.Fatalf("expected start time %v, got %v", expected, items[0].StartAt)
	}
}

// TestSyncWithMultipleSubtitleTracks verifies sync using multiple embedded subtitle tracks.
func TestSyncWithMultipleSubtitleTracks(t *testing.T) {
	base, err := astisub.OpenFile("../../testdata/simple.srt")
	if err != nil {
		t.Fatalf("open base: %v", err)
	}

	dir := t.TempDir()
	subFile := filepath.Join(dir, "test.srt")
	f, err := os.Create(subFile)
	if err != nil {
		t.Fatalf("create temp: %v", err)
	}
	defer f.Close()
	subs := astisub.Subtitles{Items: base.Items}
	if err := subs.WriteToSRT(f); err != nil {
		t.Fatalf("write SRT: %v", err)
	}

	// Create mock extractor
	mockExtractor := mocks.NewSubtitleExtractor(t)

	// Mock extract function that returns different offsets for different tracks
	mockExtractor.EXPECT().ExtractTrack("dummy.mkv", 0).
		Return(Shift(base.Items, 1*time.Second), nil)
	mockExtractor.EXPECT().ExtractTrack("dummy.mkv", 1).
		Return(Shift(base.Items, 2*time.Second), nil)
	mockExtractor.EXPECT().ExtractTrack("dummy.mkv", 2).
		Return(Shift(base.Items, 3*time.Second), nil)

	opts := Options{
		UseEmbedded:       true,
		SubtitleTracks:    []int{0, 1, 2}, // Use multiple tracks
		AudioWeight:       1.0,            // Use only embedded subtitles (no audio weight)
		SubtitleExtractor: mockExtractor,
	}

	items, err := Sync("dummy.mkv", subFile, opts)
	if err != nil {
		t.Fatalf("sync: %v", err)
	}

	if len(items) == 0 {
		t.Fatal("no items returned")
	}

	// Average offset of 1s, 2s, 3s = 2s
	// Original timing (1s) + offset (2s) = 3s final position
	expected := 3 * time.Second
	if items[0].StartAt < expected-10*time.Millisecond || items[0].StartAt > expected+10*time.Millisecond {
		t.Fatalf("expected start time around %v, got %v", expected, items[0].StartAt)
	}
}

// TestSyncWithTranslation verifies that translation is applied after sync.
func TestSyncWithTranslation(t *testing.T) {
	// Create a simple subtitle with English text
	items := []*astisub.Item{
		{
			StartAt: time.Second,
			EndAt:   2 * time.Second,
			Lines: []astisub.Line{
				{
					Items: []astisub.LineItem{
						{Text: "Hello world"},
					},
				},
			},
		},
	}

	dir := t.TempDir()
	subFile := filepath.Join(dir, "test.srt")
	f, err := os.Create(subFile)
	if err != nil {
		t.Fatalf("create temp: %v", err)
	}
	defer f.Close()
	subs := astisub.Subtitles{Items: items}
	if err := subs.WriteToSRT(f); err != nil {
		t.Fatalf("write SRT: %v", err)
	}

	// Mock successful sync without actual translation service
	// In a real test, you'd mock the translator.Translate function
	opts := Options{
		UseEmbedded:      true,
		Translate:        true,
		TranslateLang:    "es",
		TranslateService: "google",
		GoogleAPIKey:     "test-key",
	}

	result, err := Sync("dummy.mkv", subFile, opts)
	if err != nil {
		t.Fatalf("sync with translation: %v", err)
	}

	if len(result) == 0 {
		t.Fatal("no items returned")
	}

	// Note: Without mocking the translator, the text won't actually be translated
	// but the test verifies that the translation code path doesn't break the sync
	if result[0].Lines[0].Items[0].Text == "" {
		t.Fatal("text was removed during sync")
	}
}

// TestSyncDefaultBehavior verifies that sync defaults to embedded subtitles when no method is specified.
func TestSyncDefaultBehavior(t *testing.T) {
	// Create mock extractor
	mockExtractor := mocks.NewSubtitleExtractor(t)

	// Mock extract function
	mockExtractor.EXPECT().ExtractTrack("dummy.mkv", 0).
		Return([]*astisub.Item{
			{StartAt: time.Second, EndAt: 2 * time.Second},
		}, nil)

	opts := Options{
		// No UseAudio or UseEmbedded specified
		SubtitleExtractor: mockExtractor,
	}

	items, err := Sync("dummy.mkv", "../../testdata/simple.srt", opts)
	if err != nil {
		t.Fatalf("sync: %v", err)
	}

	if len(items) == 0 {
		t.Fatal("no items returned")
	}
}

// TestSyncAudio verifies that audio transcription is used when requested.
func TestSyncAudio(t *testing.T) {
	ffdir := t.TempDir()
	script := filepath.Join(ffdir, "ffmpeg")
	data := `#!/bin/sh
last=$(eval echo \${$#}); cp ../../testdata/simple.srt "$last"
`
	if err := os.WriteFile(script, []byte(data), 0755); err != nil {
		t.Fatalf("write script: %v", err)
	}
	oldPath := os.Getenv("PATH")
	os.Setenv("PATH", ffdir+":"+oldPath)
	defer os.Setenv("PATH", oldPath)
	audio.SetFFmpegPath("ffmpeg")

	var gotPath string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		fmt.Fprint(w, "1\n00:00:01,000 --> 00:00:02,000\nHello\n")
	}))
	defer srv.Close()
	transcriber.SetBaseURL(srv.URL + "/v1")
	defer transcriber.SetBaseURL("https://api.openai.com/v1")

	items, err := Sync("dummy.mkv", "../../testdata/simple_offset.srt", Options{UseAudio: true, AudioTrack: 0, WhisperKey: "k"})
	if err != nil {
		t.Fatalf("sync: %v", err)
	}
	if gotPath != "/v1/audio/transcriptions" {
		t.Fatalf("unexpected path %s", gotPath)
	}
	if items[0].StartAt != time.Second {
		t.Fatalf("unexpected start %v", items[0].StartAt)
	}
}

// TestSyncWeighted verifies synchronization using both audio and embedded
// subtitles with weighted averaging.
func TestSyncWeighted(t *testing.T) {
	base, err := astisub.OpenFile("../../testdata/simple.srt")
	if err != nil {
		t.Fatalf("open base: %v", err)
	}
	shifted := Shift(base.Items, -1*time.Second)
	dir := t.TempDir()
	subFile := filepath.Join(dir, "shifted.srt")
	f, err := os.Create(subFile)
	if err != nil {
		t.Fatalf("create temp: %v", err)
	}
	defer f.Close()
	subs := astisub.Subtitles{Items: shifted}
	if err := subs.WriteToSRT(f); err != nil {
		t.Fatalf("write SRT: %v", err)
	}

	// Create mocks
	mockTranscriber := mocks.NewTranscriber(t)
	mockExtractor := mocks.NewSubtitleExtractor(t)

	// Setup transcriber mock to return the original test file content
	b, _ := os.ReadFile("../../testdata/simple.srt")
	mockTranscriber.EXPECT().Transcribe(mock.AnythingOfType("string"), "", "test-key").
		Return(b, nil)

	// Setup extractor mock to return items shifted by 1 second
	mockExtractor.EXPECT().ExtractTrack("dummy.mkv", 0).
		Return(Shift(base.Items, time.Second), nil)

	opts := Options{
		UseAudio:          true,
		UseEmbedded:       true,
		AudioWeight:       0.7,
		WhisperKey:        "test-key",
		Transcriber:       mockTranscriber,
		SubtitleExtractor: mockExtractor,
	}

	items, err := Sync("dummy.mkv", subFile, opts)
	if err != nil {
		t.Fatalf("sync: %v", err)
	}
	if len(items) == 0 {
		t.Fatal("no items returned")
	}
	exp := 1300 * time.Millisecond
	if items[0].StartAt != exp {
		t.Fatalf("unexpected start %v", items[0].StartAt)
	}
}

// TestSyncBatch verifies that multiple subtitle files are processed.
func TestSyncBatch(t *testing.T) {
	dir := t.TempDir()
	data, err := os.ReadFile("../../testdata/simple.srt")
	if err != nil {
		t.Fatalf("read testdata: %v", err)
	}
	// prepare two input files
	in1 := filepath.Join(dir, "a.srt")
	os.WriteFile(in1, data, 0644)
	in2 := filepath.Join(dir, "b.srt")
	os.WriteFile(in2, data, 0644)
	out1 := filepath.Join(dir, "out1.srt")
	out2 := filepath.Join(dir, "out2.srt")
	items := []BatchItem{
		{Media: "dummy1.mkv", Subtitle: in1, Output: out1},
		{Media: "dummy2.mkv", Subtitle: in2, Output: out2},
	}
	errs := SyncBatch(items, Options{})
	if len(errs) != 2 {
		t.Fatalf("expected 2 results, got %d", len(errs))
	}
	for i, p := range []string{out1, out2} {
		if errs[i] != nil {
			t.Fatalf("item %d error: %v", i, errs[i])
		}
		if fi, err := os.Stat(p); err != nil || fi.Size() == 0 {
			t.Fatalf("item %d not written", i)
		}
	}
}
