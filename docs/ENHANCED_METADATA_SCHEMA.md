# Enhanced Database Schema for Subtitle Metadata

## Overview

The subtitle manager database schema has been enhanced to provide richer tracking of subtitle sources, download history, and metadata. This enables better subtitle quality assessment, provider performance monitoring, and complete audit trails for subtitle operations.

## Enhanced Schema Features

### Subtitle Table Enhancements

The `subtitles` table now includes additional metadata fields:

- **`source_url`** (TEXT): Original download URL for the subtitle
- **`provider_metadata`** (TEXT): JSON-encoded metadata from the subtitle provider
- **`confidence_score`** (REAL): Quality/match confidence score (0.0 to 1.0)
- **`parent_id`** (INTEGER): References parent subtitle for tracking modifications
- **`modification_type`** (TEXT): Type of modification applied (sync, translate, etc.)

### Download Table Enhancements

The `downloads` table includes search and performance tracking:

- **`search_query`** (TEXT): Original search terms used
- **`match_score`** (REAL): How well the result matched the search (0.0 to 1.0)
- **`download_attempts`** (INTEGER): Number of download attempts made
- **`error_message`** (TEXT): Error message if download failed
- **`response_time_ms`** (INTEGER): Provider response time in milliseconds

### Subtitle Sources Table

New `subtitle_sources` table tracks provider performance:

- **`source_hash`** (TEXT): Unique hash identifying the subtitle content
- **`original_url`** (TEXT): Original source URL
- **`provider`** (TEXT): Provider name
- **`title`** (TEXT): Subtitle title from provider
- **`release_info`** (TEXT): Release information
- **`file_size`** (INTEGER): File size in bytes
- **`download_count`** (INTEGER): Total download attempts
- **`success_count`** (INTEGER): Successful downloads
- **`avg_rating`** (REAL): Average user rating (0-5)
- **`last_seen`** (TIMESTAMP): Last time this source was seen
- **`metadata`** (TEXT): JSON-encoded provider metadata

## Usage Examples

### Creating Enhanced Subtitle Records

```go
package main

import (
    "github.com/jdfalk/subtitle-manager/pkg/database"
)

func createEnhancedSubtitle() {
    // Create provider metadata
    metadata := &database.ProviderMetadata{
        Quality:    "excellent",
        Uploader:   "trusted_user",
        Rating:     4.8,
        Downloads:  1500,
        Format:     "srt",
        Encoding:   "utf-8",
        Language:   "en",
        Release:    "BluRay.x264-GROUP",
    }
    
    // Create subtitle record with enhanced metadata
    rec, err := database.CreateSubtitleRecord(
        "movie.en.srt",
        "movie.mkv", 
        "en", 
        "opensubtitles",
        metadata,
    )
    if err != nil {
        log.Fatal(err)
    }
    
    // Add source tracking information
    score := 0.95
    rec.SourceURL = "https://opensubtitles.org/download/123456"
    rec.ConfidenceScore = &score
    
    // Store in database
    store.InsertSubtitle(rec)
}
```

### Tracking Download Performance

```go
func trackDownloadPerformance() {
    // Create download record with search metadata
    downloadRec := database.CreateDownloadRecord(
        "movie.en.srt",
        "movie.mkv",
        "opensubtitles",
        "en",
        "The Matrix 1999 BluRay",
    )
    
    // Add performance metrics
    matchScore := 0.92
    responseTime := 1200
    downloadRec.MatchScore = &matchScore
    downloadRec.ResponseTimeMs = &responseTime
    
    store.InsertDownload(downloadRec)
}
```

### Managing Subtitle Sources

```go
func manageSubtitleSources() {
    // Create subtitle source tracking
    content := []byte("subtitle content here...")
    sourceHash := database.CalculateSubtitleHash(content)
    
    metadata := &database.ProviderMetadata{
        SourceName: "Movie Title",
        Release:    "BluRay.x264-GROUP",
        FileSize:   51200,
        Rating:     4.5,
    }
    
    source, err := database.CreateSubtitleSource(
        sourceHash,
        "https://provider.com/subtitle/789",
        "opensubtitles",
        metadata,
    )
    if err != nil {
        log.Fatal(err)
    }
    
    store.InsertSubtitleSource(source)
    
    // Update statistics after downloads
    newRating := 4.7
    store.UpdateSubtitleSourceStats(sourceHash, 25, 22, &newRating)
}
```

### Getting Provider Performance Statistics

```go
func getProviderStats() {
    stats, err := database.GetProviderPerformanceStats(store, "opensubtitles")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Provider: %s\n", stats.Provider)
    fmt.Printf("Total Sources: %d\n", stats.TotalSources)
    fmt.Printf("Success Rate: %.2f%%\n", stats.SuccessRate*100)
    if stats.AvgRating != nil {
        fmt.Printf("Average Rating: %.1f/5\n", *stats.AvgRating)
    }
}
```

## Modification Type Constants

The following constants are available for tracking subtitle modifications:

- `ModificationTypeOriginal`: Original downloaded subtitle
- `ModificationTypeSync`: Timing synchronization applied
- `ModificationTypeTranslate`: Translation applied
- `ModificationTypeManualEdit`: Manual user edits
- `ModificationTypeAutoCorrect`: Automatic corrections applied
- `ModificationTypeFormatConvert`: Format conversion (e.g., VTT to SRT)

## Subtitle Relationships

Track parent-child relationships between subtitles:

```go
func trackSubtitleModifications() {
    // Original subtitle record
    original := &database.SubtitleRecord{
        ID: "original-123",
        File: "original.srt",
        // ... other fields
    }
    
    // Create synced version that references the original
    synced := database.TrackSubtitleRelationship(
        original,
        "synced.srt",
        database.ModificationTypeSync,
    )
    
    store.InsertSubtitle(synced)
}
```

## Hash-Based Duplicate Detection

Use content hashing to detect duplicate subtitles:

```go
func detectDuplicates() {
    // Calculate hash from file content
    content, err := ioutil.ReadFile("subtitle.srt")
    if err != nil {
        log.Fatal(err)
    }
    
    hash := database.CalculateSubtitleHash(content)
    
    // Check if we've seen this subtitle before
    existing, err := store.GetSubtitleSource(hash)
    if err == nil {
        fmt.Printf("Duplicate detected! Original from: %s\n", existing.OriginalURL)
    }
}
```

## Validation

Helper functions are provided to validate metadata:

```go
func validateScores() {
    score := 0.85
    
    // Validate confidence score (0-1 range)
    if err := database.ValidateConfidenceScore(&score); err != nil {
        log.Printf("Invalid confidence score: %v", err)
    }
    
    // Validate match score (0-1 range)
    if err := database.ValidateMatchScore(&score); err != nil {
        log.Printf("Invalid match score: %v", err)
    }
}
```

## Backward Compatibility

All enhancements maintain full backward compatibility with existing subtitle and download records. Existing code will continue to work without modifications, while new features are available for enhanced tracking when needed.

Records created with the legacy schema will have default/empty values for new fields, and the system gracefully handles both old and new record formats.

## Database Migration

Schema changes are automatically applied using the `addColumnIfNotExists` helper function. When the application starts, it will detect missing columns and add them with appropriate default values, ensuring seamless upgrades from older database schemas.