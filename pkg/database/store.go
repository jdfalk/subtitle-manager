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
	// Close releases any resources held by the store.
	Close() error
}
