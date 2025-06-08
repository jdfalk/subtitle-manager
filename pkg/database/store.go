package database

// SubtitleStore abstracts subtitle history storage backends.
type SubtitleStore interface {
	// InsertSubtitle stores a subtitle record.
	InsertSubtitle(rec *SubtitleRecord) error
	// ListSubtitles retrieves all subtitle records sorted by creation time.
	ListSubtitles() ([]SubtitleRecord, error)
	// DeleteSubtitle removes all records for the specified file.
	DeleteSubtitle(file string) error
	// Close releases any resources held by the store.
	Close() error
}
