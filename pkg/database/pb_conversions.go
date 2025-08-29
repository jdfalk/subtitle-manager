// file: pkg/database/pb_conversions.go
// version: 2.0.0
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
	pb := &databasepb.SubtitleRecord{
		Id:        &r.ID,
		File:      &r.File,
		VideoFile: &r.VideoFile,
		Language:  &r.Language,
		CreatedAt: timestamppb.New(r.CreatedAt.UTC()),
	}
	return pb
}

// SubtitleRecordFromProto converts a protobuf record to an internal struct.
func SubtitleRecordFromProto(pb *databasepb.SubtitleRecord) *SubtitleRecord {
	if pb == nil {
		return nil
	}
	rec := &SubtitleRecord{
		ID:        pb.GetId(),
		File:      pb.GetFile(),
		VideoFile: pb.GetVideoFile(),
		Language:  pb.GetLanguage(),
	}
	if pb.GetCreatedAt() != nil {
		rec.CreatedAt = pb.GetCreatedAt().AsTime()
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
	pb := &databasepb.DownloadRecord{
		Id:        &r.ID,
		Url:       &r.File, // Using File field as URL for now
		Status:    &r.Provider, // Using Provider field as Status for now  
		CreatedAt: timestamppb.New(r.CreatedAt.UTC()),
	}
	return pb
}

// DownloadRecordFromProto converts a protobuf record to an internal struct.
func DownloadRecordFromProto(pb *databasepb.DownloadRecord) *DownloadRecord {
	if pb == nil {
		return nil
	}
	rec := &DownloadRecord{
		ID:       pb.GetId(),
		File:     pb.GetUrl(), // Using URL field as File for now
		Provider: pb.GetStatus(), // Using Status field as Provider for now
	}
	if pb.GetCreatedAt() != nil {
		rec.CreatedAt = pb.GetCreatedAt().AsTime()
	}
	return rec
}
