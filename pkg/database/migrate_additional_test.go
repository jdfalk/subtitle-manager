// file: pkg/database/migrate_additional_test.go
// version: 1.0.0
// guid: b0bf3918-666d-47a5-990c-fb8f6eecc709

package database_test

import (
	"errors"
	"testing"

	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/database/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestMigrate_Success_CopiesRecords(t *testing.T) {
	// Arrange
	src := mocks.NewMockSubtitleStore(t)
	dest := mocks.NewMockSubtitleStore(t)

	subtitles := []database.SubtitleRecord{
		{ID: "sub-1", File: "first.srt", VideoFile: "first.mkv", Language: "en", Service: "service"},
		{ID: "sub-2", File: "second.srt", VideoFile: "second.mkv", Language: "fr", Service: "service"},
	}
	downloads := []database.DownloadRecord{
		{ID: "down-1", File: "download.srt", VideoFile: "download.mkv", Provider: "provider", Language: "en"},
	}
	media := []database.MediaItem{
		{ID: "media-1", Path: "/media/one.mkv", Title: "One"},
	}

	src.EXPECT().ListSubtitles().Return(subtitles, nil).Once()
	src.EXPECT().ListDownloads().Return(downloads, nil).Once()
	src.EXPECT().ListMediaItems().Return(media, nil).Once()

	dest.EXPECT().InsertSubtitle(mock.MatchedBy(func(rec *database.SubtitleRecord) bool {
		return rec.File == subtitles[0].File && rec.VideoFile == subtitles[0].VideoFile
	})).Return(nil).Once()
	dest.EXPECT().InsertSubtitle(mock.MatchedBy(func(rec *database.SubtitleRecord) bool {
		return rec.File == subtitles[1].File && rec.VideoFile == subtitles[1].VideoFile
	})).Return(nil).Once()
	dest.EXPECT().InsertDownload(mock.MatchedBy(func(rec *database.DownloadRecord) bool {
		return rec.File == downloads[0].File && rec.VideoFile == downloads[0].VideoFile
	})).Return(nil).Once()
	dest.EXPECT().InsertMediaItem(mock.MatchedBy(func(rec *database.MediaItem) bool {
		return rec.Path == media[0].Path && rec.Title == media[0].Title
	})).Return(nil).Once()

	// Act
	err := database.Migrate(src, dest)

	// Assert
	require.NoError(t, err)
}

func TestMigrate_ListSubtitlesError_ReturnsError(t *testing.T) {
	// Arrange
	src := mocks.NewMockSubtitleStore(t)
	dest := mocks.NewMockSubtitleStore(t)

	listErr := errors.New("list subtitles failed")
	src.EXPECT().ListSubtitles().Return(nil, listErr).Once()

	// Act
	err := database.Migrate(src, dest)

	// Assert
	require.ErrorIs(t, err, listErr)
}

func TestMigrate_ListDownloadsError_ReturnsError(t *testing.T) {
	// Arrange
	src := mocks.NewMockSubtitleStore(t)
	dest := mocks.NewMockSubtitleStore(t)

	listErr := errors.New("list downloads failed")
	src.EXPECT().ListSubtitles().Return([]database.SubtitleRecord{}, nil).Once()
	src.EXPECT().ListDownloads().Return(nil, listErr).Once()

	// Act
	err := database.Migrate(src, dest)

	// Assert
	require.ErrorIs(t, err, listErr)
}

func TestMigrate_ListMediaItemsError_ReturnsError(t *testing.T) {
	// Arrange
	src := mocks.NewMockSubtitleStore(t)
	dest := mocks.NewMockSubtitleStore(t)

	listErr := errors.New("list media failed")
	src.EXPECT().ListSubtitles().Return([]database.SubtitleRecord{}, nil).Once()
	src.EXPECT().ListDownloads().Return([]database.DownloadRecord{}, nil).Once()
	src.EXPECT().ListMediaItems().Return(nil, listErr).Once()

	// Act
	err := database.Migrate(src, dest)

	// Assert
	require.ErrorIs(t, err, listErr)
}

func TestMigrate_InsertSubtitleError_ReturnsError(t *testing.T) {
	// Arrange
	src := mocks.NewMockSubtitleStore(t)
	dest := mocks.NewMockSubtitleStore(t)

	insertErr := errors.New("insert subtitle failed")
	subtitles := []database.SubtitleRecord{
		{ID: "sub-1", File: "first.srt", VideoFile: "first.mkv"},
		{ID: "sub-2", File: "second.srt", VideoFile: "second.mkv"},
	}

	src.EXPECT().ListSubtitles().Return(subtitles, nil).Once()
	src.EXPECT().ListDownloads().Return([]database.DownloadRecord{}, nil).Once()
	src.EXPECT().ListMediaItems().Return([]database.MediaItem{}, nil).Once()

	dest.EXPECT().InsertSubtitle(mock.MatchedBy(func(rec *database.SubtitleRecord) bool {
		return rec.File == subtitles[0].File && rec.VideoFile == subtitles[0].VideoFile
	})).Return(insertErr).Once()

	// Act
	err := database.Migrate(src, dest)

	// Assert
	require.ErrorIs(t, err, insertErr)
}

func TestMigrate_InsertDownloadError_ReturnsError(t *testing.T) {
	// Arrange
	src := mocks.NewMockSubtitleStore(t)
	dest := mocks.NewMockSubtitleStore(t)

	insertErr := errors.New("insert download failed")
	downloads := []database.DownloadRecord{
		{ID: "down-1", File: "download.srt", VideoFile: "download.mkv"},
	}

	src.EXPECT().ListSubtitles().Return([]database.SubtitleRecord{}, nil).Once()
	src.EXPECT().ListDownloads().Return(downloads, nil).Once()
	src.EXPECT().ListMediaItems().Return([]database.MediaItem{}, nil).Once()

	dest.EXPECT().InsertDownload(mock.MatchedBy(func(rec *database.DownloadRecord) bool {
		return rec.File == downloads[0].File && rec.VideoFile == downloads[0].VideoFile
	})).Return(insertErr).Once()

	// Act
	err := database.Migrate(src, dest)

	// Assert
	require.ErrorIs(t, err, insertErr)
}

func TestMigrate_InsertMediaItemError_ReturnsError(t *testing.T) {
	// Arrange
	src := mocks.NewMockSubtitleStore(t)
	dest := mocks.NewMockSubtitleStore(t)

	insertErr := errors.New("insert media failed")
	media := []database.MediaItem{
		{ID: "media-1", Path: "/media/one.mkv", Title: "One"},
		{ID: "media-2", Path: "/media/two.mkv", Title: "Two"},
	}

	src.EXPECT().ListSubtitles().Return([]database.SubtitleRecord{}, nil).Once()
	src.EXPECT().ListDownloads().Return([]database.DownloadRecord{}, nil).Once()
	src.EXPECT().ListMediaItems().Return(media, nil).Once()

	dest.EXPECT().InsertMediaItem(mock.MatchedBy(func(rec *database.MediaItem) bool {
		return rec.Path == media[0].Path && rec.Title == media[0].Title
	})).Return(insertErr).Once()

	// Act
	err := database.Migrate(src, dest)

	// Assert
	require.ErrorIs(t, err, insertErr)
}
