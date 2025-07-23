// file: pkg/database/pb_conversions.go
// version: 1.0.0
// guid: c9cf2d1a-c284-46f4-90d1-70925cbe8b27

package database

import (
	"time"

	"github.com/jdfalk/subtitle-manager/pkg/databasepb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ToProto converts a SubtitleRecord to its protobuf representation.
func (r *SubtitleRecord) ToProto() *databasepb.SubtitleRecord {
	if r == nil {
		return nil
	}
	var score *float64
	if r.ConfidenceScore != nil {
		v := *r.ConfidenceScore
		score = &v
	}
	var parent *string
	if r.ParentID != nil {
		v := *r.ParentID
		parent = &v
	}
	return &databasepb.SubtitleRecord{
		Id:               r.ID,
		File:             r.File,
		VideoFile:        r.VideoFile,
		Release:          r.Release,
		Language:         r.Language,
		Service:          r.Service,
		Embedded:         r.Embedded,
		SourceUrl:        r.SourceURL,
		ProviderMetadata: r.ProviderMetadata,
		ConfidenceScore:  score,
		ParentId:         parent,
		ModificationType: r.ModificationType,
		CreatedAt:        timestamppb.New(r.CreatedAt.UTC()),
	}
}

// SubtitleRecordFromProto converts a protobuf record to internal struct.
func SubtitleRecordFromProto(pb *databasepb.SubtitleRecord) *SubtitleRecord {
	if pb == nil {
		return nil
	}
	rec := &SubtitleRecord{
		ID:               pb.Id,
		File:             pb.File,
		VideoFile:        pb.VideoFile,
		Release:          pb.Release,
		Language:         pb.Language,
		Service:          pb.Service,
		Embedded:         pb.Embedded,
		SourceURL:        pb.SourceUrl,
		ProviderMetadata: pb.ProviderMetadata,
		ModificationType: pb.ModificationType,
	}
	if pb.ConfidenceScore != nil {
		v := *pb.ConfidenceScore
		rec.ConfidenceScore = &v
	}
	if pb.ParentId != nil {
		v := *pb.ParentId
		rec.ParentID = &v
	}
	if pb.CreatedAt != nil {
		rec.CreatedAt = pb.CreatedAt.AsTime()
	} else {
		rec.CreatedAt = time.Time{}
	}
	return rec
}

// ToProto converts a DownloadRecord to its protobuf representation.
func (r *DownloadRecord) ToProto() *databasepb.DownloadRecord {
	if r == nil {
		return nil
	}
	var score *float64
	if r.MatchScore != nil {
		v := *r.MatchScore
		score = &v
	}
	var rt *int32
	if r.ResponseTimeMs != nil {
		v := int32(*r.ResponseTimeMs)
		rt = &v
	}
	return &databasepb.DownloadRecord{
		Id:               r.ID,
		File:             r.File,
		VideoFile:        r.VideoFile,
		Provider:         r.Provider,
		Language:         r.Language,
		SearchQuery:      r.SearchQuery,
		MatchScore:       score,
		DownloadAttempts: int32(r.DownloadAttempts),
		ErrorMessage:     r.ErrorMessage,
		ResponseTimeMs:   rt,
		CreatedAt:        timestamppb.New(r.CreatedAt.UTC()),
	}
}

// DownloadRecordFromProto converts protobuf to internal struct.
func DownloadRecordFromProto(pb *databasepb.DownloadRecord) *DownloadRecord {
	if pb == nil {
		return nil
	}
	rec := &DownloadRecord{
		ID:               pb.Id,
		File:             pb.File,
		VideoFile:        pb.VideoFile,
		Provider:         pb.Provider,
		Language:         pb.Language,
		SearchQuery:      pb.SearchQuery,
		DownloadAttempts: int(pb.DownloadAttempts),
		ErrorMessage:     pb.ErrorMessage,
	}
	if pb.MatchScore != nil {
		v := *pb.MatchScore
		rec.MatchScore = &v
	}
	if pb.ResponseTimeMs != nil {
		v := int(*pb.ResponseTimeMs)
		rec.ResponseTimeMs = &v
	}
	if pb.CreatedAt != nil {
		rec.CreatedAt = pb.CreatedAt.AsTime()
	}
	return rec
}
