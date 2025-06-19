package database

// SubtitleStore abstracts subtitle history storage backends.
type SubtitleStore interface {
	// InsertSubtitle stores a subtitle record.
	InsertSubtitle(rec *SubtitleRecord) error
	// ListSubtitles retrieves all subtitle records sorted by creation time.
	ListSubtitles() ([]SubtitleRecord, error)
	// DeleteSubtitle removes all records for the specified file.
	DeleteSubtitle(file string) error
	// InsertDownload stores a download record.
	InsertDownload(rec *DownloadRecord) error
	// ListDownloads retrieves all download records sorted by creation time.
	ListDownloads() ([]DownloadRecord, error)
	// DeleteDownload removes download records for the specified subtitle file.
	DeleteDownload(file string) error
	// InsertMediaItem stores a media library record.
	InsertMediaItem(rec *MediaItem) error
	// ListMediaItems retrieves all media items sorted by creation time.
	ListMediaItems() ([]MediaItem, error)
	// DeleteMediaItem removes a record for the specified media path.
	DeleteMediaItem(path string) error
	// InsertTag stores a new tag value.
	InsertTag(name string) error
	// ListTags retrieves all tags.
	ListTags() ([]Tag, error)
	// DeleteTag removes a tag by ID.
	DeleteTag(id int64) error
	// UpdateTag renames a tag by ID.
	UpdateTag(id int64, name string) error
	// AssignTagToUser associates a tag with a user.
	AssignTagToUser(userID, tagID int64) error
	// RemoveTagFromUser disassociates a tag from a user.
	RemoveTagFromUser(userID, tagID int64) error
	// ListTagsForUser returns tags associated with a user.
	ListTagsForUser(userID int64) ([]Tag, error)
	// AssignTagToMedia associates a tag with a media item.
	AssignTagToMedia(mediaID, tagID int64) error
	// RemoveTagFromMedia disassociates a tag from a media item.
	RemoveTagFromMedia(mediaID, tagID int64) error
	// ListTagsForMedia returns tags associated with a media item.
	ListTagsForMedia(mediaID int64) ([]Tag, error)
	// Close releases any resources held by the store.
	Close() error
}
