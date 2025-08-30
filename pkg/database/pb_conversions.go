// file: pkg/database/pb_conversions.go
// version: 3.0.0
// guid: c9cf2d1a-c284-46f4-90d1-70925cbe8b27

package database

import (
	"time"

	"github.com/jdfalk/gcommon/sdks/go/v1/database"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ToProto converts a SubtitleRecord to its gcommon Row representation.
func (r *SubtitleRecord) ToProto() *database.Row {
	if r == nil {
		return nil
	}
	
	row := &database.Row{}
	
	// Convert fields to Any values for the generic Row
	values := make([]*anypb.Any, 0)
	
	// Add ID
	if idAny, err := anypb.New(&timestamppb.Timestamp{}); err == nil {
		idAny.TypeUrl = "subtitle_manager/id"
		idAny.Value = []byte(r.ID)
		values = append(values, idAny)
	}
	
	// Add File
	if fileAny, err := anypb.New(&timestamppb.Timestamp{}); err == nil {
		fileAny.TypeUrl = "subtitle_manager/file"
		fileAny.Value = []byte(r.File)
		values = append(values, fileAny)
	}
	
	// Add VideoFile
	if videoAny, err := anypb.New(&timestamppb.Timestamp{}); err == nil {
		videoAny.TypeUrl = "subtitle_manager/video_file"
		videoAny.Value = []byte(r.VideoFile)
		values = append(values, videoAny)
	}
	
	// Add Language
	if langAny, err := anypb.New(&timestamppb.Timestamp{}); err == nil {
		langAny.TypeUrl = "subtitle_manager/language"
		langAny.Value = []byte(r.Language)
		values = append(values, langAny)
	}
	
	// Add CreatedAt
	if timeAny, err := anypb.New(timestamppb.New(r.CreatedAt.UTC())); err == nil {
		values = append(values, timeAny)
	}
	
	row.SetValues(values)
	return row
}

// SubtitleRecordFromProto converts a gcommon Row to an internal SubtitleRecord struct.
func SubtitleRecordFromProto(row *database.Row) *SubtitleRecord {
	if row == nil || len(row.GetValues()) < 4 {
		return nil
	}
	
	values := row.GetValues()
	rec := &SubtitleRecord{}
	
	// Extract values from Any array - simplified approach for migration
	if len(values) > 0 {
		rec.ID = string(values[0].Value) // Simplified extraction
	}
	if len(values) > 1 {
		rec.File = string(values[1].Value)
	}
	if len(values) > 2 {
		rec.VideoFile = string(values[2].Value)
	}
	if len(values) > 3 {
		rec.Language = string(values[3].Value)
	}
	if len(values) > 4 {
		// Extract timestamp from the last value
		var ts timestamppb.Timestamp
		if err := values[4].UnmarshalTo(&ts); err == nil {
			rec.CreatedAt = ts.AsTime()
		} else {
			rec.CreatedAt = time.Now()
		}
	} else {
		rec.CreatedAt = time.Now()
	}
	
	return rec
}

// ToProto converts a DownloadRecord to its gcommon Row representation.
func (r *DownloadRecord) ToProto() *database.Row {
	if r == nil {
		return nil
	}
	
	row := &database.Row{}
	values := make([]*anypb.Any, 0)
	
	// Add basic fields as Any values
	if idAny, err := anypb.New(&timestamppb.Timestamp{}); err == nil {
		idAny.TypeUrl = "subtitle_manager/download_id"
		idAny.Value = []byte(r.ID)
		values = append(values, idAny)
	}
	
	if fileAny, err := anypb.New(&timestamppb.Timestamp{}); err == nil {
		fileAny.TypeUrl = "subtitle_manager/download_file"
		fileAny.Value = []byte(r.File)
		values = append(values, fileAny)
	}
	
	if providerAny, err := anypb.New(&timestamppb.Timestamp{}); err == nil {
		providerAny.TypeUrl = "subtitle_manager/provider"
		providerAny.Value = []byte(r.Provider)
		values = append(values, providerAny)
	}
	
	// Add CreatedAt
	if timeAny, err := anypb.New(timestamppb.New(r.CreatedAt.UTC())); err == nil {
		values = append(values, timeAny)
	}
	
	row.SetValues(values)
	return row
}

// DownloadRecordFromProto converts a gcommon Row to an internal DownloadRecord struct.
func DownloadRecordFromProto(row *database.Row) *DownloadRecord {
	if row == nil || len(row.GetValues()) < 3 {
		return nil
	}
	
	values := row.GetValues()
	rec := &DownloadRecord{}
	
	// Extract values from Any array
	if len(values) > 0 {
		rec.ID = string(values[0].Value)
	}
	if len(values) > 1 {
		rec.File = string(values[1].Value)
	}
	if len(values) > 2 {
		rec.Provider = string(values[2].Value)
	}
	if len(values) > 3 {
		// Extract timestamp
		var ts timestamppb.Timestamp
		if err := values[3].UnmarshalTo(&ts); err == nil {
			rec.CreatedAt = ts.AsTime()
		} else {
			rec.CreatedAt = time.Now()
		}
	} else {
		rec.CreatedAt = time.Now()
	}
	
	return rec
}
