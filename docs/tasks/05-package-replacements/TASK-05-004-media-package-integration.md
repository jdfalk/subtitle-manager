<!-- file: docs/tasks/05-package-replacements/TASK-05-004-media-package-integration.md -->
<!-- version: 1.0.0 -->
<!-- guid: d4e5f6a7-b8c9-0d1e-2f3a-4b5c6d7e8f9a -->

# TASK-05-004: Media Package Integration

## Overview

**Objective**: Use gcommon media types for video and subtitle file handling throughout the subtitle-manager application.

**Phase**: 3 (Package Replacements)
**Priority**: High
**Estimated Effort**: 8-10 hours
**Prerequisites**: TASK-05-003 (health monitoring) and gcommon package foundation completed

## Required Reading

**CRITICAL**: Read these documents before starting:

- `docs/gcommon-api/media.md` - gcommon media type specifications and patterns
- Current media handling code in `pkg/media/` directory
- Video file processing implementations
- Subtitle file handling and format conversion code
- `docs/MIGRATION_INVENTORY.md` - Media processing usage inventory

## Problem Statement

The subtitle-manager currently uses custom media file handling for video and subtitle processing. This needs to be replaced with gcommon media types to:

1. **Standardize Media Types**: Use gcommon MediaFile, VideoMetadata, SubtitleMetadata types
2. **Improve Format Support**: Leverage gcommon's comprehensive format support
3. **Enhanced Metadata**: Use standardized metadata extraction and processing
4. **Better Integration**: Enable interoperability with other gcommon-based media services

### Current Media Implementation

```go
// Current implementation (to be replaced)
type VideoFile struct {
    Path        string
    Format      string
    Duration    time.Duration
    Resolution  string
    Bitrate     int
    Metadata    map[string]string
}

type SubtitleFile struct {
    Path      string
    Format    string
    Language  string
    Encoding  string
    Lines     []SubtitleLine
}
```

### Target gcommon Media Types

```go
// New implementation using gcommon
import "github.com/jdfalk/gcommon/sdks/go/v1/media"

mediaFile := &media.MediaFile{}
mediaFile.SetField("path", "/path/to/video.mp4")
mediaFile.SetField("type", "video")

videoMeta := &media.VideoMetadata{}
videoMeta.SetField("duration_seconds", 7200)
videoMeta.SetField("resolution", "1920x1080")
```

## Technical Approach

### Media Type Mapping Strategy

1. **File Type Integration**: Map current file types to gcommon media types
2. **Metadata Migration**: Convert metadata extraction to gcommon patterns
3. **Format Support**: Leverage gcommon format detection and conversion
4. **Processing Pipeline**: Update media processing workflows

### Key Mappings

```go
// Media type mappings
type MediaTypeMapping struct {
    // OLD custom types -> NEW gcommon/media
    VideoFile      -> media.MediaFile (with type="video")
    SubtitleFile   -> media.MediaFile (with type="subtitle")
    VideoMetadata  -> media.VideoMetadata
    SubtitleData   -> media.SubtitleContent
    FormatInfo     -> media.FormatMetadata
}
```

## Implementation Steps

### Step 1: Create Media Management Service

```go
// File: pkg/media/service.go
package media

import (
    "fmt"
    "path/filepath"
    "time"

    "github.com/jdfalk/gcommon/sdks/go/v1/media"
    "github.com/jdfalk/subtitle-manager/pkg/config"
)

// MediaService manages media file operations using gcommon
type MediaService struct {
    mediaManager  media.Manager
    config       *config.ApplicationConfig
}

// NewMediaService creates a new media service
func NewMediaService(config *config.ApplicationConfig) *MediaService {
    return &MediaService{
        mediaManager: media.NewManager(),
        config:       config,
    }
}

// VideoFile wraps gcommon MediaFile for video-specific operations
type VideoFile struct {
    mediaFile *media.MediaFile
}

// NewVideoFile creates a new video file from path
func (ms *MediaService) NewVideoFile(filePath string) (*VideoFile, error) {
    mediaFile := &media.MediaFile{}
    mediaFile.SetField("path", filePath)
    mediaFile.SetField("type", "video")
    mediaFile.SetField("discovered_at", time.Now())

    // Extract basic file information
    if err := ms.extractFileInfo(mediaFile, filePath); err != nil {
        return nil, fmt.Errorf("failed to extract file info: %v", err)
    }

    return &VideoFile{
        mediaFile: mediaFile,
    }, nil
}

// Video file accessors using gcommon opaque API
func (vf *VideoFile) GetPath() string {
    if path, exists := vf.mediaFile.GetField("path"); exists {
        if pathStr, ok := path.(string); ok {
            return pathStr
        }
    }
    return ""
}

func (vf *VideoFile) SetPath(path string) {
    vf.mediaFile.SetField("path", path)
}

func (vf *VideoFile) GetFormat() string {
    if format, exists := vf.mediaFile.GetField("format"); exists {
        if formatStr, ok := format.(string); ok {
            return formatStr
        }
    }
    return ""
}

func (vf *VideoFile) SetFormat(format string) {
    vf.mediaFile.SetField("format", format)
}

func (vf *VideoFile) GetDuration() time.Duration {
    if duration, exists := vf.mediaFile.GetField("duration_seconds"); exists {
        if durationFloat, ok := duration.(float64); ok {
            return time.Duration(durationFloat) * time.Second
        }
    }
    return 0
}

func (vf *VideoFile) SetDuration(duration time.Duration) {
    vf.mediaFile.SetField("duration_seconds", duration.Seconds())
}

func (vf *VideoFile) GetResolution() string {
    if resolution, exists := vf.mediaFile.GetField("resolution"); exists {
        if resolutionStr, ok := resolution.(string); ok {
            return resolutionStr
        }
    }
    return ""
}

func (vf *VideoFile) SetResolution(resolution string) {
    vf.mediaFile.SetField("resolution", resolution)
}

func (vf *VideoFile) GetBitrate() int {
    if bitrate, exists := vf.mediaFile.GetField("bitrate"); exists {
        if bitrateFloat, ok := bitrate.(float64); ok {
            return int(bitrateFloat)
        }
    }
    return 0
}

func (vf *VideoFile) SetBitrate(bitrate int) {
    vf.mediaFile.SetField("bitrate", float64(bitrate))
}

func (vf *VideoFile) GetMetadata() map[string]interface{} {
    if metadata, exists := vf.mediaFile.GetField("metadata"); exists {
        if metaMap, ok := metadata.(map[string]interface{}); ok {
            return metaMap
        }
    }
    return make(map[string]interface{})
}

func (vf *VideoFile) SetMetadata(metadata map[string]interface{}) {
    vf.mediaFile.SetField("metadata", metadata)
}

// Get underlying gcommon media file
func (vf *VideoFile) GetMediaFile() *media.MediaFile {
    return vf.mediaFile
}

// SubtitleFile wraps gcommon MediaFile for subtitle-specific operations
type SubtitleFile struct {
    mediaFile *media.MediaFile
}

// NewSubtitleFile creates a new subtitle file from path
func (ms *MediaService) NewSubtitleFile(filePath string) (*SubtitleFile, error) {
    mediaFile := &media.MediaFile{}
    mediaFile.SetField("path", filePath)
    mediaFile.SetField("type", "subtitle")
    mediaFile.SetField("discovered_at", time.Now())

    // Extract basic file information
    if err := ms.extractFileInfo(mediaFile, filePath); err != nil {
        return nil, fmt.Errorf("failed to extract file info: %v", err)
    }

    return &SubtitleFile{
        mediaFile: mediaFile,
    }, nil
}

// Subtitle file accessors
func (sf *SubtitleFile) GetPath() string {
    if path, exists := sf.mediaFile.GetField("path"); exists {
        if pathStr, ok := path.(string); ok {
            return pathStr
        }
    }
    return ""
}

func (sf *SubtitleFile) SetPath(path string) {
    sf.mediaFile.SetField("path", path)
}

func (sf *SubtitleFile) GetFormat() string {
    if format, exists := sf.mediaFile.GetField("format"); exists {
        if formatStr, ok := format.(string); ok {
            return formatStr
        }
    }
    return ""
}

func (sf *SubtitleFile) SetFormat(format string) {
    sf.mediaFile.SetField("format", format)
}

func (sf *SubtitleFile) GetLanguage() string {
    if language, exists := sf.mediaFile.GetField("language"); exists {
        if langStr, ok := language.(string); ok {
            return langStr
        }
    }
    return ""
}

func (sf *SubtitleFile) SetLanguage(language string) {
    sf.mediaFile.SetField("language", language)
}

func (sf *SubtitleFile) GetEncoding() string {
    if encoding, exists := sf.mediaFile.GetField("encoding"); exists {
        if encodingStr, ok := encoding.(string); ok {
            return encodingStr
        }
    }
    return ""
}

func (sf *SubtitleFile) SetEncoding(encoding string) {
    sf.mediaFile.SetField("encoding", encoding)
}

func (sf *SubtitleFile) GetContent() *media.SubtitleContent {
    if content, exists := sf.mediaFile.GetField("content"); exists {
        if subtitleContent, ok := content.(*media.SubtitleContent); ok {
            return subtitleContent
        }
    }
    return nil
}

func (sf *SubtitleFile) SetContent(content *media.SubtitleContent) {
    sf.mediaFile.SetField("content", content)
}

// Get underlying gcommon media file
func (sf *SubtitleFile) GetMediaFile() *media.MediaFile {
    return sf.mediaFile
}

// extractFileInfo extracts basic file information using gcommon
func (ms *MediaService) extractFileInfo(mediaFile *media.MediaFile, filePath string) error {
    // Extract file extension and set format
    ext := filepath.Ext(filePath)
    if ext != "" {
        mediaFile.SetField("format", ext[1:]) // Remove leading dot
    }

    // Get file size
    if fileInfo, err := os.Stat(filePath); err == nil {
        mediaFile.SetField("size_bytes", fileInfo.Size())
        mediaFile.SetField("modified_at", fileInfo.ModTime())
    }

    // Use gcommon media analysis for detailed metadata
    if err := ms.mediaManager.AnalyzeFile(mediaFile); err != nil {
        return fmt.Errorf("gcommon media analysis failed: %v", err)
    }

    return nil
}

// ProcessVideo performs comprehensive video analysis
func (ms *MediaService) ProcessVideo(videoFile *VideoFile) error {
    mediaFile := videoFile.GetMediaFile()

    // Use gcommon video processing
    videoProcessor := ms.mediaManager.GetVideoProcessor()
    if err := videoProcessor.Process(mediaFile); err != nil {
        return fmt.Errorf("video processing failed: %v", err)
    }

    // Extract video-specific metadata
    if err := ms.extractVideoMetadata(videoFile); err != nil {
        return fmt.Errorf("video metadata extraction failed: %v", err)
    }

    return nil
}

// ProcessSubtitle performs comprehensive subtitle analysis
func (ms *MediaService) ProcessSubtitle(subtitleFile *SubtitleFile) error {
    mediaFile := subtitleFile.GetMediaFile()

    // Use gcommon subtitle processing
    subtitleProcessor := ms.mediaManager.GetSubtitleProcessor()
    if err := subtitleProcessor.Process(mediaFile); err != nil {
        return fmt.Errorf("subtitle processing failed: %v", err)
    }

    // Extract subtitle-specific content
    if err := ms.extractSubtitleContent(subtitleFile); err != nil {
        return fmt.Errorf("subtitle content extraction failed: %v", err)
    }

    return nil
}

// extractVideoMetadata extracts detailed video metadata
func (ms *MediaService) extractVideoMetadata(videoFile *VideoFile) error {
    mediaFile := videoFile.GetMediaFile()

    // Create video metadata object
    videoMeta := &media.VideoMetadata{}

    // Extract using gcommon video analysis
    if err := ms.mediaManager.ExtractVideoMetadata(mediaFile, videoMeta); err != nil {
        return fmt.Errorf("failed to extract video metadata: %v", err)
    }

    // Store in media file
    mediaFile.SetField("video_metadata", videoMeta)

    // Extract common fields for easy access
    if duration := videoMeta.GetField("duration"); duration != nil {
        videoFile.SetDuration(duration.(time.Duration))
    }

    if resolution := videoMeta.GetField("resolution"); resolution != nil {
        videoFile.SetResolution(resolution.(string))
    }

    if bitrate := videoMeta.GetField("bitrate"); bitrate != nil {
        videoFile.SetBitrate(bitrate.(int))
    }

    return nil
}

// extractSubtitleContent extracts and parses subtitle content
func (ms *MediaService) extractSubtitleContent(subtitleFile *SubtitleFile) error {
    mediaFile := subtitleFile.GetMediaFile()

    // Create subtitle content object
    subtitleContent := &media.SubtitleContent{}

    // Extract using gcommon subtitle analysis
    if err := ms.mediaManager.ExtractSubtitleContent(mediaFile, subtitleContent); err != nil {
        return fmt.Errorf("failed to extract subtitle content: %v", err)
    }

    // Store in subtitle file
    subtitleFile.SetContent(subtitleContent)

    // Extract format and language information
    if format := subtitleContent.GetField("format"); format != nil {
        subtitleFile.SetFormat(format.(string))
    }

    if language := subtitleContent.GetField("language"); language != nil {
        subtitleFile.SetLanguage(language.(string))
    }

    if encoding := subtitleContent.GetField("encoding"); encoding != nil {
        subtitleFile.SetEncoding(encoding.(string))
    }

    return nil
}
```

### Step 2: Create Media Processing Pipeline

```go
// File: pkg/media/processor.go
package media

import (
    "fmt"
    "path/filepath"
    "strings"

    "github.com/jdfalk/gcommon/sdks/go/v1/media"
)

// MediaProcessor handles media file processing workflows
type MediaProcessor struct {
    mediaService *MediaService
}

// NewMediaProcessor creates a new media processor
func NewMediaProcessor(mediaService *MediaService) *MediaProcessor {
    return &MediaProcessor{
        mediaService: mediaService,
    }
}

// ProcessingResult contains the results of media processing
type ProcessingResult struct {
    VideoFile     *VideoFile
    SubtitleFiles []*SubtitleFile
    Errors        []error
    Warnings      []string
}

// ProcessMediaPair processes a video file and its associated subtitle files
func (mp *MediaProcessor) ProcessMediaPair(videoPath string, subtitlePaths []string) (*ProcessingResult, error) {
    result := &ProcessingResult{
        SubtitleFiles: make([]*SubtitleFile, 0),
        Errors:        make([]error, 0),
        Warnings:      make([]string, 0),
    }

    // Process video file
    videoFile, err := mp.processVideoFile(videoPath)
    if err != nil {
        result.Errors = append(result.Errors, fmt.Errorf("video processing failed: %v", err))
        return result, err
    }
    result.VideoFile = videoFile

    // Process subtitle files
    for _, subtitlePath := range subtitlePaths {
        subtitleFile, err := mp.processSubtitleFile(subtitlePath)
        if err != nil {
            result.Errors = append(result.Errors, fmt.Errorf("subtitle processing failed for %s: %v", subtitlePath, err))
            continue
        }
        result.SubtitleFiles = append(result.SubtitleFiles, subtitleFile)
    }

    // Validate compatibility
    mp.validateMediaCompatibility(result)

    return result, nil
}

// processVideoFile processes a single video file
func (mp *MediaProcessor) processVideoFile(videoPath string) (*VideoFile, error) {
    // Create video file object
    videoFile, err := mp.mediaService.NewVideoFile(videoPath)
    if err != nil {
        return nil, fmt.Errorf("failed to create video file: %v", err)
    }

    // Perform comprehensive processing
    if err := mp.mediaService.ProcessVideo(videoFile); err != nil {
        return nil, fmt.Errorf("video processing failed: %v", err)
    }

    return videoFile, nil
}

// processSubtitleFile processes a single subtitle file
func (mp *MediaProcessor) processSubtitleFile(subtitlePath string) (*SubtitleFile, error) {
    // Create subtitle file object
    subtitleFile, err := mp.mediaService.NewSubtitleFile(subtitlePath)
    if err != nil {
        return nil, fmt.Errorf("failed to create subtitle file: %v", err)
    }

    // Perform comprehensive processing
    if err := mp.mediaService.ProcessSubtitle(subtitleFile); err != nil {
        return nil, fmt.Errorf("subtitle processing failed: %v", err)
    }

    return subtitleFile, nil
}

// validateMediaCompatibility checks if video and subtitle files are compatible
func (mp *MediaProcessor) validateMediaCompatibility(result *ProcessingResult) {
    if result.VideoFile == nil {
        return
    }

    videoDuration := result.VideoFile.GetDuration()

    for _, subtitleFile := range result.SubtitleFiles {
        content := subtitleFile.GetContent()
        if content == nil {
            result.Warnings = append(result.Warnings,
                fmt.Sprintf("No content extracted from subtitle file: %s", subtitleFile.GetPath()))
            continue
        }

        // Check subtitle duration vs video duration
        if subtitleDuration := content.GetField("duration"); subtitleDuration != nil {
            if subtitleDur, ok := subtitleDuration.(time.Duration); ok {
                if subtitleDur > videoDuration*120/100 { // 20% tolerance
                    result.Warnings = append(result.Warnings,
                        fmt.Sprintf("Subtitle duration significantly longer than video: %s", subtitleFile.GetPath()))
                }
            }
        }

        // Check language detection
        if subtitleFile.GetLanguage() == "" {
            result.Warnings = append(result.Warnings,
                fmt.Sprintf("Could not detect language for subtitle file: %s", subtitleFile.GetPath()))
        }
    }
}

// DiscoverMediaFiles discovers video and subtitle files in a directory
func (mp *MediaProcessor) DiscoverMediaFiles(directory string) (*MediaDiscovery, error) {
    discovery := &MediaDiscovery{
        VideoFiles:    make(map[string]*VideoFile),
        SubtitleFiles: make(map[string][]*SubtitleFile),
        Directory:     directory,
    }

    // Use gcommon media discovery
    discoveryService := mp.mediaService.mediaManager.GetDiscoveryService()

    mediaFiles, err := discoveryService.DiscoverDirectory(directory)
    if err != nil {
        return nil, fmt.Errorf("media discovery failed: %v", err)
    }

    // Process discovered files
    for _, mediaFile := range mediaFiles {
        fileType, exists := mediaFile.GetField("type")
        if !exists {
            continue
        }

        switch fileType.(string) {
        case "video":
            videoFile := &VideoFile{mediaFile: mediaFile}
            discovery.VideoFiles[videoFile.GetPath()] = videoFile

        case "subtitle":
            subtitleFile := &SubtitleFile{mediaFile: mediaFile}

            // Associate with video file based on naming
            associatedVideo := mp.findAssociatedVideo(subtitleFile, discovery.VideoFiles)
            if associatedVideo != "" {
                discovery.SubtitleFiles[associatedVideo] = append(
                    discovery.SubtitleFiles[associatedVideo], subtitleFile)
            }
        }
    }

    return discovery, nil
}

// MediaDiscovery holds discovered media files
type MediaDiscovery struct {
    VideoFiles    map[string]*VideoFile              // path -> VideoFile
    SubtitleFiles map[string][]*SubtitleFile         // video path -> subtitle files
    Directory     string
}

// findAssociatedVideo finds the video file associated with a subtitle
func (mp *MediaProcessor) findAssociatedVideo(subtitleFile *SubtitleFile, videoFiles map[string]*VideoFile) string {
    subtitlePath := subtitleFile.GetPath()
    subtitleBase := strings.TrimSuffix(filepath.Base(subtitlePath), filepath.Ext(subtitlePath))

    // Look for video files with similar names
    for videoPath := range videoFiles {
        videoBase := strings.TrimSuffix(filepath.Base(videoPath), filepath.Ext(videoPath))

        // Direct match
        if subtitleBase == videoBase {
            return videoPath
        }

        // Match with language suffix (e.g., movie.en.srt -> movie.mp4)
        if strings.HasPrefix(subtitleBase, videoBase+".") {
            return videoPath
        }

        // Match with language prefix (e.g., en.movie.srt -> movie.mp4)
        parts := strings.Split(subtitleBase, ".")
        if len(parts) > 1 && strings.Join(parts[1:], ".") == videoBase {
            return videoPath
        }
    }

    return ""
}

// ConvertSubtitle converts subtitle format using gcommon
func (mp *MediaProcessor) ConvertSubtitle(subtitleFile *SubtitleFile, targetFormat string) (*SubtitleFile, error) {
    if subtitleFile.GetFormat() == targetFormat {
        return subtitleFile, nil // No conversion needed
    }

    // Use gcommon conversion service
    converter := mp.mediaService.mediaManager.GetSubtitleConverter()

    convertedMediaFile, err := converter.Convert(subtitleFile.GetMediaFile(), targetFormat)
    if err != nil {
        return nil, fmt.Errorf("subtitle conversion failed: %v", err)
    }

    convertedSubtitle := &SubtitleFile{mediaFile: convertedMediaFile}

    return convertedSubtitle, nil
}

// ExtractSubtitleMetrics extracts metrics from subtitle content
func (mp *MediaProcessor) ExtractSubtitleMetrics(subtitleFile *SubtitleFile) (*SubtitleMetrics, error) {
    content := subtitleFile.GetContent()
    if content == nil {
        return nil, fmt.Errorf("no subtitle content available")
    }

    metrics := &SubtitleMetrics{
        FilePath: subtitleFile.GetPath(),
        Format:   subtitleFile.GetFormat(),
        Language: subtitleFile.GetLanguage(),
        Encoding: subtitleFile.GetEncoding(),
    }

    // Extract metrics using gcommon
    analyzer := mp.mediaService.mediaManager.GetSubtitleAnalyzer()

    analysis, err := analyzer.Analyze(content)
    if err != nil {
        return nil, fmt.Errorf("subtitle analysis failed: %v", err)
    }

    // Extract metrics from analysis
    if lineCount := analysis.GetField("line_count"); lineCount != nil {
        metrics.LineCount = lineCount.(int)
    }

    if wordCount := analysis.GetField("word_count"); wordCount != nil {
        metrics.WordCount = wordCount.(int)
    }

    if duration := analysis.GetField("duration"); duration != nil {
        metrics.Duration = duration.(time.Duration)
    }

    if readingSpeed := analysis.GetField("reading_speed_wpm"); readingSpeed != nil {
        metrics.ReadingSpeed = readingSpeed.(float64)
    }

    return metrics, nil
}

// SubtitleMetrics contains subtitle analysis metrics
type SubtitleMetrics struct {
    FilePath     string
    Format       string
    Language     string
    Encoding     string
    LineCount    int
    WordCount    int
    Duration     time.Duration
    ReadingSpeed float64 // words per minute
}
```

### Step 3: Update Media Database Integration

```go
// File: pkg/media/storage.go
package media

import (
    "fmt"

    "github.com/jdfalk/subtitle-manager/pkg/database"
)

// MediaStorage handles media file database operations
type MediaStorage struct {
    dbManager *database.DatabaseManager
}

// NewMediaStorage creates a new media storage service
func NewMediaStorage(dbManager *database.DatabaseManager) *MediaStorage {
    return &MediaStorage{
        dbManager: dbManager,
    }
}

// StoreVideoFile stores video file information in database
func (ms *MediaStorage) StoreVideoFile(videoFile *VideoFile) error {
    // Create database record for video file
    record := database.NewSubtitleRecord() // Reuse subtitle record for media
    record.SetID(generateMediaID())
    record.SetVideoFile(videoFile.GetPath())
    record.SetFormat(videoFile.GetFormat())
    record.SetCreatedAt(time.Now())

    // Store gcommon media metadata
    mediaFile := videoFile.GetMediaFile()
    if metadata := ms.serializeMediaMetadata(mediaFile); metadata != nil {
        // Store serialized metadata in database
        record.GetRecord().SetField("media_metadata", metadata)
    }

    return ms.dbManager.CreateSubtitleRecord(record)
}

// StoreSubtitleFile stores subtitle file information in database
func (ms *MediaStorage) StoreSubtitleFile(subtitleFile *SubtitleFile, associatedVideoPath string) error {
    record := database.NewSubtitleRecord()
    record.SetID(generateMediaID())
    record.SetVideoFile(associatedVideoPath)
    record.SetSubtitleFile(subtitleFile.GetPath())
    record.SetLanguage(subtitleFile.GetLanguage())
    record.SetFormat(subtitleFile.GetFormat())
    record.SetCreatedAt(time.Now())

    // Store gcommon subtitle content and metadata
    mediaFile := subtitleFile.GetMediaFile()
    if content := ms.serializeSubtitleContent(subtitleFile.GetContent()); content != nil {
        record.GetRecord().SetField("subtitle_content", content)
    }

    if metadata := ms.serializeMediaMetadata(mediaFile); metadata != nil {
        record.GetRecord().SetField("media_metadata", metadata)
    }

    return ms.dbManager.CreateSubtitleRecord(record)
}

// serializeMediaMetadata serializes gcommon media metadata for storage
func (ms *MediaStorage) serializeMediaMetadata(mediaFile *media.MediaFile) map[string]interface{} {
    metadata := make(map[string]interface{})

    // Extract key fields from gcommon media file
    if duration, exists := mediaFile.GetField("duration_seconds"); exists {
        metadata["duration_seconds"] = duration
    }

    if resolution, exists := mediaFile.GetField("resolution"); exists {
        metadata["resolution"] = resolution
    }

    if bitrate, exists := mediaFile.GetField("bitrate"); exists {
        metadata["bitrate"] = bitrate
    }

    if size, exists := mediaFile.GetField("size_bytes"); exists {
        metadata["size_bytes"] = size
    }

    if modTime, exists := mediaFile.GetField("modified_at"); exists {
        metadata["modified_at"] = modTime
    }

    return metadata
}

// serializeSubtitleContent serializes gcommon subtitle content for storage
func (ms *MediaStorage) serializeSubtitleContent(content *media.SubtitleContent) map[string]interface{} {
    if content == nil {
        return nil
    }

    contentData := make(map[string]interface{})

    // Extract key fields from gcommon subtitle content
    if lineCount, exists := content.GetField("line_count"); exists {
        contentData["line_count"] = lineCount
    }

    if wordCount, exists := content.GetField("word_count"); exists {
        contentData["word_count"] = wordCount
    }

    if duration, exists := content.GetField("duration"); exists {
        contentData["duration"] = duration
    }

    if encoding, exists := content.GetField("encoding"); exists {
        contentData["encoding"] = encoding
    }

    // Store subtitle lines if needed (may be large)
    if lines, exists := content.GetField("lines"); exists {
        contentData["lines"] = lines
    }

    return contentData
}

// LoadVideoFile loads video file from database
func (ms *MediaStorage) LoadVideoFile(videoPath string) (*VideoFile, error) {
    // Query database for video file record
    records, err := ms.dbManager.ListSubtitleRecords(100, 0) // Simplified query
    if err != nil {
        return nil, fmt.Errorf("failed to query video records: %v", err)
    }

    for _, record := range records {
        if record.GetVideoFile() == videoPath {
            return ms.reconstructVideoFile(record)
        }
    }

    return nil, fmt.Errorf("video file not found: %s", videoPath)
}

// reconstructVideoFile reconstructs VideoFile from database record
func (ms *MediaStorage) reconstructVideoFile(record *database.SubtitleRecord) (*VideoFile, error) {
    mediaFile := &media.MediaFile{}
    mediaFile.SetField("path", record.GetVideoFile())
    mediaFile.SetField("type", "video")
    mediaFile.SetField("format", record.GetFormat())

    // Restore metadata from database
    if metadata, exists := record.GetRecord().GetField("media_metadata"); exists {
        if metaMap, ok := metadata.(map[string]interface{}); ok {
            for key, value := range metaMap {
                mediaFile.SetField(key, value)
            }
        }
    }

    return &VideoFile{mediaFile: mediaFile}, nil
}

// LoadSubtitleFiles loads subtitle files associated with a video
func (ms *MediaStorage) LoadSubtitleFiles(videoPath string) ([]*SubtitleFile, error) {
    records, err := ms.dbManager.ListSubtitleRecords(100, 0)
    if err != nil {
        return nil, fmt.Errorf("failed to query subtitle records: %v", err)
    }

    var subtitleFiles []*SubtitleFile

    for _, record := range records {
        if record.GetVideoFile() == videoPath && record.GetSubtitleFile() != "" {
            subtitleFile, err := ms.reconstructSubtitleFile(record)
            if err != nil {
                continue // Skip problematic records
            }
            subtitleFiles = append(subtitleFiles, subtitleFile)
        }
    }

    return subtitleFiles, nil
}

// reconstructSubtitleFile reconstructs SubtitleFile from database record
func (ms *MediaStorage) reconstructSubtitleFile(record *database.SubtitleRecord) (*SubtitleFile, error) {
    mediaFile := &media.MediaFile{}
    mediaFile.SetField("path", record.GetSubtitleFile())
    mediaFile.SetField("type", "subtitle")
    mediaFile.SetField("format", record.GetFormat())
    mediaFile.SetField("language", record.GetLanguage())

    // Restore content from database
    if contentData, exists := record.GetRecord().GetField("subtitle_content"); exists {
        if contentMap, ok := contentData.(map[string]interface{}); ok {
            content := &media.SubtitleContent{}
            for key, value := range contentMap {
                content.SetField(key, value)
            }
            mediaFile.SetField("content", content)
        }
    }

    // Restore metadata from database
    if metadata, exists := record.GetRecord().GetField("media_metadata"); exists {
        if metaMap, ok := metadata.(map[string]interface{}); ok {
            for key, value := range metaMap {
                mediaFile.SetField(key, value)
            }
        }
    }

    return &SubtitleFile{mediaFile: mediaFile}, nil
}

// generateMediaID generates a unique ID for media records
func generateMediaID() string {
    // Implementation depends on ID generation strategy
    return fmt.Sprintf("media_%d", time.Now().UnixNano())
}
```

## Testing Requirements

### Media Service Tests

```go
// File: pkg/media/service_test.go
package media

import (
    "testing"
    "time"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestMediaService(t *testing.T) {
    mockConfig := &MockConfig{}
    mediaService := NewMediaService(mockConfig)

    t.Run("Video file creation", func(t *testing.T) {
        videoFile, err := mediaService.NewVideoFile("/test/video.mp4")
        require.NoError(t, err)
        assert.Equal(t, "/test/video.mp4", videoFile.GetPath())
        assert.Equal(t, "mp4", videoFile.GetFormat())
    })

    t.Run("Subtitle file creation", func(t *testing.T) {
        subtitleFile, err := mediaService.NewSubtitleFile("/test/subtitle.srt")
        require.NoError(t, err)
        assert.Equal(t, "/test/subtitle.srt", subtitleFile.GetPath())
        assert.Equal(t, "srt", subtitleFile.GetFormat())
    })

    t.Run("Video processing", func(t *testing.T) {
        videoFile, _ := mediaService.NewVideoFile("/test/video.mp4")

        // Set test metadata
        videoFile.SetDuration(2 * time.Hour)
        videoFile.SetResolution("1920x1080")
        videoFile.SetBitrate(5000)

        assert.Equal(t, 2*time.Hour, videoFile.GetDuration())
        assert.Equal(t, "1920x1080", videoFile.GetResolution())
        assert.Equal(t, 5000, videoFile.GetBitrate())
    })
}

func TestMediaProcessor(t *testing.T) {
    mediaService := NewMediaService(&MockConfig{})
    processor := NewMediaProcessor(mediaService)

    t.Run("Media pair processing", func(t *testing.T) {
        result, err := processor.ProcessMediaPair(
            "/test/video.mp4",
            []string{"/test/subtitle.srt", "/test/subtitle.es.srt"},
        )

        require.NoError(t, err)
        assert.NotNil(t, result.VideoFile)
        assert.Len(t, result.SubtitleFiles, 2)
    })
}
```

## Success Metrics

### Functional Requirements

- [ ] All media types use gcommon media packages
- [ ] Video and subtitle processing works with gcommon
- [ ] Metadata extraction uses gcommon analyzers
- [ ] Format detection and conversion functional
- [ ] Database storage compatible with gcommon types

### Technical Requirements

- [ ] Media file discovery automated and accurate
- [ ] Processing pipeline handles various formats
- [ ] Error handling graceful for unsupported formats
- [ ] Performance acceptable for large media libraries
- [ ] Memory usage optimized for media processing

### Integration Requirements

- [ ] Database integration preserves all metadata
- [ ] HTTP API updated to use new media types
- [ ] UI displays gcommon media information correctly
- [ ] Backward compatibility maintained during migration

## Common Pitfalls

1. **Format Support**: Ensure gcommon supports all required media formats
2. **Memory Usage**: Large media files can consume significant memory
3. **File Path Handling**: Different OS path separators and encoding
4. **Metadata Accuracy**: Verify metadata extraction accuracy across formats
5. **Performance**: Processing large video files may be slow

## Dependencies

- **Requires**: TASK-05-003 (health monitoring) for media processing health checks
- **Requires**: gcommon media SDK properly installed
- **Requires**: Updated database schema for gcommon metadata storage
- **Enables**: Standardized media handling across gcommon ecosystem
- **Blocks**: Advanced media features until migration complete

This comprehensive media package integration replaces custom media handling with gcommon standardized media types while maintaining all existing functionality and improving format support and metadata handling.
