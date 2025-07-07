package database

import (
	"time"

	"github.com/jdfalk/subtitle-manager/pkg/profiles"
)

type LanguageProfile = profiles.LanguageProfile

type LanguageProfileAssignment = profiles.MediaProfileAssignment

// SubtitleStore abstracts subtitle history storage backends.
type SubtitleStore interface {
	// InsertSubtitle stores a subtitle record.
	InsertSubtitle(rec *SubtitleRecord) error
	// ListSubtitles retrieves all subtitle records sorted by creation time.
	ListSubtitles() ([]SubtitleRecord, error)
	// ListSubtitlesByVideo returns subtitle records for a video file.
	ListSubtitlesByVideo(video string) ([]SubtitleRecord, error)
	// CountSubtitles returns the total number of subtitle records.
	CountSubtitles() (int, error)
	// DeleteSubtitle removes all records for the specified file.
	DeleteSubtitle(file string) error
	// InsertDownload stores a download record.
	InsertDownload(rec *DownloadRecord) error
	// ListDownloads retrieves all download records sorted by creation time.
	ListDownloads() ([]DownloadRecord, error)
	// ListDownloadsByVideo returns download records for a video file.
	ListDownloadsByVideo(video string) ([]DownloadRecord, error)
	// CountDownloads returns the total number of download records.
	CountDownloads() (int, error)
	// DeleteDownload removes download records for the specified subtitle file.
	DeleteDownload(file string) error
	// InsertMediaItem stores a media library record.
	InsertMediaItem(rec *MediaItem) error
	// ListMediaItems retrieves all media items sorted by creation time.
	ListMediaItems() ([]MediaItem, error)
	// CountMediaItems returns the total number of media items.
	CountMediaItems() (int, error)
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
	// SetMediaReleaseGroup sets the release group for a media item.
	SetMediaReleaseGroup(path, group string) error
	// SetMediaAltTitles stores alternate titles for a media item.
	SetMediaAltTitles(path string, titles []string) error
	// SetMediaFieldLocks updates locked fields for a media item.
	SetMediaFieldLocks(path, locks string) error
	// GetMediaReleaseGroup retrieves the release group for a media item.
	GetMediaReleaseGroup(path string) (string, error)
	// GetMediaAltTitles retrieves alternate titles for a media item.
	GetMediaAltTitles(path string) ([]string, error)
	// GetMediaFieldLocks retrieves locked fields for a media item.
	GetMediaFieldLocks(path string) (string, error)
	// SetMediaTitle updates the title for a media item.
	SetMediaTitle(path, title string) error
	// CreateLanguageProfile stores a new language profile.
	CreateLanguageProfile(profile *LanguageProfile) error
	// GetLanguageProfile retrieves a language profile by ID.
	GetLanguageProfile(id string) (*LanguageProfile, error)
	// ListLanguageProfiles retrieves all language profiles.
	ListLanguageProfiles() ([]LanguageProfile, error)
	// UpdateLanguageProfile updates an existing language profile.
	UpdateLanguageProfile(profile *LanguageProfile) error
	// DeleteLanguageProfile removes a language profile by ID.
	DeleteLanguageProfile(id string) error
	// SetDefaultLanguageProfile marks a profile as the default.
	SetDefaultLanguageProfile(id string) error
	// GetDefaultLanguageProfile retrieves the default language profile.
	GetDefaultLanguageProfile() (*LanguageProfile, error)
	// AssignProfileToMedia assigns a language profile to a media item.
	AssignProfileToMedia(mediaID, profileID string) error
	// RemoveProfileFromMedia removes language profile assignment from a media item.
	RemoveProfileFromMedia(mediaID string) error
	// GetMediaProfile retrieves the language profile assigned to a media item.
	GetMediaProfile(mediaID string) (*LanguageProfile, error)
	// Subtitle source tracking operations
	InsertSubtitleSource(src *SubtitleSource) error
	GetSubtitleSource(sourceHash string) (*SubtitleSource, error)
	UpdateSubtitleSourceStats(sourceHash string, downloadCount, successCount int, avgRating *float64) error
	ListSubtitleSources(provider string, limit int) ([]SubtitleSource, error)
	DeleteSubtitleSource(sourceHash string) error
	// Monitoring methods
	// InsertMonitoredItem stores a monitored item record.
	InsertMonitoredItem(rec *MonitoredItem) error
	// ListMonitoredItems retrieves all monitored items.
	ListMonitoredItems() ([]MonitoredItem, error)
	// UpdateMonitoredItem updates an existing monitored item.
	UpdateMonitoredItem(rec *MonitoredItem) error
	// DeleteMonitoredItem removes a monitored item by ID.
	DeleteMonitoredItem(id string) error
	// GetMonitoredItemsToCheck returns items that need monitoring.
	GetMonitoredItemsToCheck(interval time.Duration) ([]MonitoredItem, error)
	// Close releases any resources held by the store.
	Close() error
}
