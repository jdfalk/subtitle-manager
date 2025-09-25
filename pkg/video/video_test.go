// file: pkg/video/video_test.go
// version: 1.0.0
// guid: 8a9b0c1d-2e3f-4a5b-6c7d-8e9f0a1b2c3e

package video

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestSetFFprobePath tests the SetFFprobePath function
func TestSetFFprobePath(t *testing.T) {
	originalPath := ffprobePath
	defer func() {
		ffprobePath = originalPath // Restore original path
	}()

	testPath := "/custom/path/to/ffprobe"
	SetFFprobePath(testPath)

	assert.Equal(t, testPath, ffprobePath)
}

// TestVideoInfoMethods tests all methods on the VideoInfo struct
func TestVideoInfoMethods(t *testing.T) {
	tests := []struct {
		name                   string
		videoInfo              *VideoInfo
		wantResolution         string
		wantResolutionCategory string
		wantIsHD               bool
		wantAspectRatio        float64
	}{
		{
			name: "4K video",
			videoInfo: &VideoInfo{
				Width:  3840,
				Height: 2160,
			},
			wantResolution:         "3840x2160",
			wantResolutionCategory: "2160p",
			wantIsHD:               true,
			wantAspectRatio:        1.7777777777777777, // 16:9
		},
		{
			name: "1440p video",
			videoInfo: &VideoInfo{
				Width:  2560,
				Height: 1440,
			},
			wantResolution:         "2560x1440",
			wantResolutionCategory: "1440p",
			wantIsHD:               true,
			wantAspectRatio:        1.7777777777777777, // 16:9
		},
		{
			name: "1080p video",
			videoInfo: &VideoInfo{
				Width:  1920,
				Height: 1080,
			},
			wantResolution:         "1920x1080",
			wantResolutionCategory: "1080p",
			wantIsHD:               true,
			wantAspectRatio:        1.7777777777777777, // 16:9
		},
		{
			name: "720p video",
			videoInfo: &VideoInfo{
				Width:  1280,
				Height: 720,
			},
			wantResolution:         "1280x720",
			wantResolutionCategory: "720p",
			wantIsHD:               true,
			wantAspectRatio:        1.7777777777777777, // 16:9
		},
		{
			name: "480p video",
			videoInfo: &VideoInfo{
				Width:  854,
				Height: 480,
			},
			wantResolution:         "854x480",
			wantResolutionCategory: "480p",
			wantIsHD:               false,
			wantAspectRatio:        1.7791666666666666,
		},
		{
			name: "SD video",
			videoInfo: &VideoInfo{
				Width:  640,
				Height: 360,
			},
			wantResolution:         "640x360",
			wantResolutionCategory: "SD",
			wantIsHD:               false,
			wantAspectRatio:        1.7777777777777777, // 16:9
		},
		{
			name: "4:3 aspect ratio",
			videoInfo: &VideoInfo{
				Width:  1024,
				Height: 768,
			},
			wantResolution:         "1024x768",
			wantResolutionCategory: "720p", // Height >= 720
			wantIsHD:               true,
			wantAspectRatio:        1.3333333333333333, // 4:3
		},
		{
			name: "zero dimensions",
			videoInfo: &VideoInfo{
				Width:  0,
				Height: 0,
			},
			wantResolution:         "",
			wantResolutionCategory: "SD",
			wantIsHD:               false,
			wantAspectRatio:        0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantResolution, tt.videoInfo.GetResolution())
			assert.Equal(t, tt.wantResolutionCategory, tt.videoInfo.GetResolutionCategory())
			assert.Equal(t, tt.wantIsHD, tt.videoInfo.IsHD())
			assert.InDelta(t, tt.wantAspectRatio, tt.videoInfo.GetAspectRatio(), 0.0001)
		})
	}
}

// mockFFprobeOutput creates a temporary script that outputs mock ffprobe JSON
func mockFFprobeOutput(t *testing.T, mockJSON string) string {
	tempDir := t.TempDir()
	scriptPath := filepath.Join(tempDir, "mock-ffprobe")

	content := fmt.Sprintf(`#!/bin/bash
echo '%s'
`, mockJSON)

	err := os.WriteFile(scriptPath, []byte(content), 0755)
	require.NoError(t, err)

	return scriptPath
}

// TestAnalyzeVideo tests the AnalyzeVideo function with mocked ffprobe output
func TestAnalyzeVideo(t *testing.T) {
	tests := []struct {
		name          string
		mockJSON      string
		wantVideoInfo *VideoInfo
		wantError     bool
		wantErrorMsg  string
	}{
		{
			name: "complete video analysis",
			mockJSON: `{
				"format": {
					"filename": "test.mp4",
					"format_name": "mp4,m4a,3gp,3g2,mj2",
					"duration": "120.500000",
					"size": "10485760",
					"bit_rate": "699050"
				},
				"streams": [
					{
						"index": 0,
						"codec_type": "video",
						"codec_name": "h264",
						"width": 1920,
						"height": 1080,
						"avg_frame_rate": "30000/1001"
					},
					{
						"index": 1,
						"codec_type": "audio",
						"codec_name": "aac",
						"sample_rate": "48000",
						"channels": 2
					}
				]
			}`,
			wantVideoInfo: &VideoInfo{
				Duration:   time.Duration(120.5 * float64(time.Second)),
				Width:      1920,
				Height:     1080,
				Bitrate:    699050,
				FrameRate:  29.97002997002997, // 30000/1001
				Codec:      "h264",
				Format:     "mp4,m4a,3gp,3g2,mj2",
				FileSize:   10485760,
				AudioCodec: "aac",
				AudioRate:  48000,
			},
		},
		{
			name: "video with zero frame rate",
			mockJSON: `{
				"format": {
					"filename": "test.mp4",
					"format_name": "mp4",
					"duration": "60.0",
					"size": "5242880",
					"bit_rate": "699050"
				},
				"streams": [
					{
						"index": 0,
						"codec_type": "video",
						"codec_name": "h264",
						"width": 1280,
						"height": 720,
						"avg_frame_rate": "0/0"
					}
				]
			}`,
			wantVideoInfo: &VideoInfo{
				Duration:  60 * time.Second,
				Width:     1280,
				Height:    720,
				Bitrate:   699050,
				FrameRate: 0, // Should remain 0 for 0/0 frame rate
				Codec:     "h264",
				Format:    "mp4",
				FileSize:  5242880,
			},
		},
		{
			name: "audio only file",
			mockJSON: `{
				"format": {
					"filename": "test.mp3",
					"format_name": "mp3",
					"duration": "180.0",
					"size": "2097152"
				},
				"streams": [
					{
						"index": 0,
						"codec_type": "audio",
						"codec_name": "mp3",
						"sample_rate": "44100",
						"channels": 2
					}
				]
			}`,
			wantVideoInfo: &VideoInfo{
				Duration:   180 * time.Second,
				Format:     "mp3",
				FileSize:   2097152,
				AudioCodec: "mp3",
				AudioRate:  44100,
			},
		},
		{
			name: "minimal format info",
			mockJSON: `{
				"format": {
					"filename": "test.mkv",
					"format_name": "matroska"
				},
				"streams": []
			}`,
			wantVideoInfo: &VideoInfo{
				Format: "matroska",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock ffprobe script
			mockPath := mockFFprobeOutput(t, tt.mockJSON)
			originalPath := ffprobePath
			SetFFprobePath(mockPath)
			defer SetFFprobePath(originalPath)

			// Test the analysis
			info, err := AnalyzeVideo("dummy-path")

			if tt.wantError {
				require.Error(t, err)
				if tt.wantErrorMsg != "" {
					assert.Contains(t, err.Error(), tt.wantErrorMsg)
				}
				return
			}

			require.NoError(t, err)
			require.NotNil(t, info)

			// Compare all fields
			assert.Equal(t, tt.wantVideoInfo.Duration, info.Duration)
			assert.Equal(t, tt.wantVideoInfo.Width, info.Width)
			assert.Equal(t, tt.wantVideoInfo.Height, info.Height)
			assert.Equal(t, tt.wantVideoInfo.Bitrate, info.Bitrate)
			assert.InDelta(t, tt.wantVideoInfo.FrameRate, info.FrameRate, 0.001)
			assert.Equal(t, tt.wantVideoInfo.Codec, info.Codec)
			assert.Equal(t, tt.wantVideoInfo.Format, info.Format)
			assert.Equal(t, tt.wantVideoInfo.FileSize, info.FileSize)
			assert.Equal(t, tt.wantVideoInfo.AudioCodec, info.AudioCodec)
			assert.Equal(t, tt.wantVideoInfo.AudioRate, info.AudioRate)
		})
	}
}

// TestAnalyzeVideoErrors tests error conditions in AnalyzeVideo
func TestAnalyzeVideoErrors(t *testing.T) {
	tests := []struct {
		name         string
		setupMock    func(t *testing.T) string // Returns mock ffprobe path
		wantErrorMsg string
	}{
		{
			name: "ffprobe command fails",
			setupMock: func(t *testing.T) string {
				tempDir := t.TempDir()
				scriptPath := filepath.Join(tempDir, "failing-ffprobe")
				content := `#!/bin/bash
exit 1
`
				err := os.WriteFile(scriptPath, []byte(content), 0755)
				require.NoError(t, err)
				return scriptPath
			},
			wantErrorMsg: "ffprobe analysis failed",
		},
		{
			name: "invalid JSON output",
			setupMock: func(t *testing.T) string {
				return mockFFprobeOutput(t, `invalid json{`)
			},
			wantErrorMsg: "failed to parse ffprobe output",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockPath := tt.setupMock(t)
			originalPath := ffprobePath
			SetFFprobePath(mockPath)
			defer SetFFprobePath(originalPath)

			info, err := AnalyzeVideo("dummy-path")

			require.Error(t, err)
			assert.Nil(t, info)
			assert.Contains(t, err.Error(), tt.wantErrorMsg)
		})
	}
}

// TestGetSubtitleTracks tests the GetSubtitleTracks function
func TestGetSubtitleTracks(t *testing.T) {
	tests := []struct {
		name         string
		mockJSON     string
		wantTracks   []map[string]string
		wantError    bool
		wantErrorMsg string
	}{
		{
			name: "video with subtitle tracks",
			mockJSON: `{
				"streams": [
					{
						"index": 2,
						"codec_type": "subtitle",
						"codec_name": "subrip"
					},
					{
						"index": 3,
						"codec_type": "subtitle",
						"codec_name": "ass"
					}
				]
			}`,
			wantTracks: []map[string]string{
				{"index": "2", "codec": "subrip"},
				{"index": "3", "codec": "ass"},
			},
		},
		{
			name: "video with no subtitle tracks",
			mockJSON: `{
				"streams": [
					{
						"index": 0,
						"codec_type": "video",
						"codec_name": "h264"
					},
					{
						"index": 1,
						"codec_type": "audio",
						"codec_name": "aac"
					}
				]
			}`,
			wantTracks: nil, // Returns nil when no subtitle tracks found
		},
		{
			name:       "empty streams",
			mockJSON:   `{"streams": []}`,
			wantTracks: nil, // Returns nil when no streams at all
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock ffprobe script
			mockPath := mockFFprobeOutput(t, tt.mockJSON)
			originalPath := ffprobePath
			SetFFprobePath(mockPath)
			defer SetFFprobePath(originalPath)

			// Test the subtitle track analysis
			tracks, err := GetSubtitleTracks("dummy-path")

			if tt.wantError {
				require.Error(t, err)
				if tt.wantErrorMsg != "" {
					assert.Contains(t, err.Error(), tt.wantErrorMsg)
				}
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.wantTracks, tracks)
		})
	}
}

// TestGetSubtitleTracksErrors tests error conditions in GetSubtitleTracks
func TestGetSubtitleTracksErrors(t *testing.T) {
	tests := []struct {
		name         string
		setupMock    func(t *testing.T) string
		wantErrorMsg string
	}{
		{
			name: "ffprobe command fails",
			setupMock: func(t *testing.T) string {
				tempDir := t.TempDir()
				scriptPath := filepath.Join(tempDir, "failing-ffprobe")
				content := `#!/bin/bash
exit 1
`
				err := os.WriteFile(scriptPath, []byte(content), 0755)
				require.NoError(t, err)
				return scriptPath
			},
			wantErrorMsg: "ffprobe subtitle analysis failed",
		},
		{
			name: "invalid JSON output",
			setupMock: func(t *testing.T) string {
				return mockFFprobeOutput(t, `invalid json{`)
			},
			wantErrorMsg: "failed to parse ffprobe output",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockPath := tt.setupMock(t)
			originalPath := ffprobePath
			SetFFprobePath(mockPath)
			defer SetFFprobePath(originalPath)

			tracks, err := GetSubtitleTracks("dummy-path")

			require.Error(t, err)
			assert.Nil(t, tracks)
			assert.Contains(t, err.Error(), tt.wantErrorMsg)
		})
	}
}

// Integration tests - these require ffprobe to be available
func TestVideoIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode")
	}

	// Check if ffprobe is available
	if _, err := exec.LookPath("ffprobe"); err != nil {
		t.Skip("ffprobe not available, skipping integration tests")
	}

	t.Run("real ffprobe availability test", func(t *testing.T) {
		// Test with a non-existent file to ensure ffprobe is working
		// but the file doesn't exist (we expect an error, but not a "command not found" error)
		_, err := AnalyzeVideo("/path/that/does/not/exist")
		require.Error(t, err)
		// Should contain ffprobe analysis failed, not command not found
		assert.Contains(t, err.Error(), "ffprobe analysis failed")
	})
}

// Benchmark tests
func BenchmarkVideoInfoMethods(b *testing.B) {
	info := &VideoInfo{
		Width:     1920,
		Height:    1080,
		Duration:  120 * time.Second,
		Bitrate:   1000000,
		FrameRate: 30.0,
	}

	b.Run("GetResolution", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = info.GetResolution()
		}
	})

	b.Run("GetResolutionCategory", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = info.GetResolutionCategory()
		}
	})

	b.Run("GetAspectRatio", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = info.GetAspectRatio()
		}
	})

	b.Run("IsHD", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = info.IsHD()
		}
	})
}

// TestCompleteVideoInfoStructure tests that VideoInfo struct can be properly serialized
func TestCompleteVideoInfoStructure(t *testing.T) {
	info := VideoInfo{
		Duration:   120 * time.Second,
		Width:      1920,
		Height:     1080,
		Bitrate:    1000000,
		FrameRate:  29.97,
		Codec:      "h264",
		Format:     "mp4",
		FileSize:   50 * 1024 * 1024, // 50MB
		AudioCodec: "aac",
		AudioRate:  48000,
	}

	// Test JSON serialization
	data, err := json.Marshal(info)
	require.NoError(t, err)

	var unmarshaled VideoInfo
	err = json.Unmarshal(data, &unmarshaled)
	require.NoError(t, err)

	// Verify all fields are preserved
	assert.Equal(t, info.Duration, unmarshaled.Duration)
	assert.Equal(t, info.Width, unmarshaled.Width)
	assert.Equal(t, info.Height, unmarshaled.Height)
	assert.Equal(t, info.Bitrate, unmarshaled.Bitrate)
	assert.Equal(t, info.FrameRate, unmarshaled.FrameRate)
	assert.Equal(t, info.Codec, unmarshaled.Codec)
	assert.Equal(t, info.Format, unmarshaled.Format)
	assert.Equal(t, info.FileSize, unmarshaled.FileSize)
	assert.Equal(t, info.AudioCodec, unmarshaled.AudioCodec)
	assert.Equal(t, info.AudioRate, unmarshaled.AudioRate)
}
