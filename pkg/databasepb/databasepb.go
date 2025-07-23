// file: pkg/databasepb/databasepb.go
// version: 1.0.0
// guid: b2ed7ce4-1a47-4da1-bbde-63d964be4896

package databasepb

import "google.golang.org/protobuf/types/known/timestamppb"

// SubtitleRecord mirrors proto message for storing subtitle history.
type SubtitleRecord struct {
	Id               string
	File             string
	VideoFile        string
	Release          string
	Language         string
	Service          string
	Embedded         bool
	SourceUrl        string
	ProviderMetadata string
	ConfidenceScore  *float64
	ParentId         *string
	ModificationType string
	CreatedAt        *timestamppb.Timestamp
}

// DownloadRecord mirrors proto message for subtitle downloads.
type DownloadRecord struct {
	Id               string
	File             string
	VideoFile        string
	Provider         string
	Language         string
	SearchQuery      string
	MatchScore       *float64
	DownloadAttempts int32
	ErrorMessage     string
	ResponseTimeMs   *int32
	CreatedAt        *timestamppb.Timestamp
}
