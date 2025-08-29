// file: pkg/database/pb_conversions.go
// version: 1.1.0
// guid: c9cf2d1a-c284-46f4-90d1-70925cbe8b27

package database

import (
	"time"

	"github.com/jdfalk/gcommon/sdks/go/v1/database"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ToProto converts a SubtitleRecord to its protobuf representation.
func (r *SubtitleRecord) ToProto() *database.SubtitleRecord {
	if r == nil {
		return nil
	}
	pb := &database.SubtitleRecord{}
	pb.SetId(r.ID)
	pb.SetFile(r.File)
	pb.SetVideoFile(r.VideoFile)
	pb.SetRelease(r.Release)
	pb.SetLanguage(r.Language)
	pb.SetService(r.Service)
	pb.SetEmbedded(r.Embedded)
	pb.SetSourceUrl(r.SourceURL)
	pb.SetProviderMetadata(r.ProviderMetadata)
	if r.ConfidenceScore != nil {
		pb.SetConfidenceScore(*r.ConfidenceScore)
	}
	if r.ParentID != nil {
		pb.SetParentId(*r.ParentID)
	}
	pb.SetModificationType(r.ModificationType)
	pb.SetCreatedAt(timestamppb.New(r.CreatedAt.UTC()))
	return pb
}

// SubtitleRecordFromProto converts a protobuf record to an internal struct.
func SubtitleRecordFromProto(pb *database.SubtitleRecord) *SubtitleRecord {
	if pb == nil {
		return nil
	}
	rec := &SubtitleRecord{
		ID:               pb.GetId(),
		File:             pb.GetFile(),
		VideoFile:        pb.GetVideoFile(),
		Release:          pb.GetRelease(),
		Language:         pb.GetLanguage(),
		Service:          pb.GetService(),
		Embedded:         pb.GetEmbedded(),
		SourceURL:        pb.GetSourceUrl(),
		ProviderMetadata: pb.GetProviderMetadata(),
		ModificationType: pb.GetModificationType(),
	}
	if pb.GetConfidenceScore() != 0 {
		v := pb.GetConfidenceScore()
		rec.ConfidenceScore = &v
	}
	if pb.GetParentId() != "" {
		v := pb.GetParentId()
		rec.ParentID = &v
	}
	if pb.GetCreatedAt() != nil {
		rec.CreatedAt = pb.GetCreatedAt().AsTime()
	} else {
		rec.CreatedAt = time.Time{}
	}
	return rec
}

// ToProto converts a DownloadRecord to its protobuf representation.
func (r *DownloadRecord) ToProto() *database.DownloadRecord {
	if r == nil {
		return nil
	}
	pb := &database.DownloadRecord{}
	pb.SetId(r.ID)
	pb.SetFile(r.File)
	pb.SetVideoFile(r.VideoFile)
	pb.SetProvider(r.Provider)
	pb.SetLanguage(r.Language)
	pb.SetSearchQuery(r.SearchQuery)
	if r.MatchScore != nil {
		pb.SetMatchScore(*r.MatchScore)
	}
	pb.SetDownloadAttempts(int32(r.DownloadAttempts))
	pb.SetErrorMessage(r.ErrorMessage)
	if r.ResponseTimeMs != nil {
		pb.SetResponseTimeMs(int32(*r.ResponseTimeMs))
	}
	pb.SetCreatedAt(timestamppb.New(r.CreatedAt.UTC()))
	return pb
}

// DownloadRecordFromProto converts a protobuf record to an internal struct.
func DownloadRecordFromProto(pb *database.DownloadRecord) *DownloadRecord {
	if pb == nil {
		return nil
	}
	rec := &DownloadRecord{
		ID:               pb.GetId(),
		File:             pb.GetFile(),
		VideoFile:        pb.GetVideoFile(),
		Provider:         pb.GetProvider(),
		Language:         pb.GetLanguage(),
		SearchQuery:      pb.GetSearchQuery(),
		DownloadAttempts: int(pb.GetDownloadAttempts()),
		ErrorMessage:     pb.GetErrorMessage(),
	}
	if pb.GetMatchScore() != 0 {
		v := pb.GetMatchScore()
		rec.MatchScore = &v
	}
	if pb.GetResponseTimeMs() != 0 {
		v := int(pb.GetResponseTimeMs())
		rec.ResponseTimeMs = &v
	}
	if pb.GetCreatedAt() != nil {
		rec.CreatedAt = pb.GetCreatedAt().AsTime()
	}
	return rec
}
