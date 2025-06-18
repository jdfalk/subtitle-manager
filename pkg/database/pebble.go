package database

import (
	"encoding/json"
	"sort"
	"strings"
	"time"

	"github.com/cockroachdb/pebble"
	"github.com/google/uuid"
)

// PebbleStore wraps a Pebble database and implements basic CRUD operations
// for SubtitleRecord documents.
// Keys are generated using UUIDs with the "subtitle:" prefix.
// Values are stored as JSON encoded SubtitleRecord structures.
type PebbleStore struct {
	db *pebble.DB
}

// OpenPebble opens a Pebble database at path and returns a PebbleStore.
func OpenPebble(path string) (*PebbleStore, error) {
	db, err := pebble.Open(path, &pebble.Options{})
	if err != nil {
		return nil, err
	}
	return &PebbleStore{db: db}, nil
}

// Close closes the underlying Pebble database.
func (p *PebbleStore) Close() error { return p.db.Close() }

// InsertSubtitle stores a subtitle translation record.
// The ID field of rec will be filled with a generated UUID if empty.
func (p *PebbleStore) InsertSubtitle(rec *SubtitleRecord) error {
	if rec.ID == "" {
		rec.ID = uuid.NewString()
	}
	if rec.CreatedAt.IsZero() {
		rec.CreatedAt = time.Now()
	}
	b, err := json.Marshal(rec)
	if err != nil {
		return err
	}
	key := []byte("subtitle:" + rec.ID)
	return p.db.Set(key, b, pebble.Sync)
}

// ListSubtitles returns all stored subtitle records sorted by creation time
// in descending order.
func (p *PebbleStore) ListSubtitles() ([]SubtitleRecord, error) {
	iter, err := p.db.NewIter(nil)
	if err != nil {
		return nil, err
	}
	defer iter.Close()
	var recs []SubtitleRecord
	for iter.First(); iter.Valid(); iter.Next() {
		if !strings.HasPrefix(string(iter.Key()), "subtitle:") {
			continue
		}
		var r SubtitleRecord
		if err := json.Unmarshal(iter.Value(), &r); err != nil {
			return nil, err
		}
		recs = append(recs, r)
	}
	if err := iter.Error(); err != nil {
		return nil, err
	}
	sort.Slice(recs, func(i, j int) bool {
		return recs[i].CreatedAt.After(recs[j].CreatedAt)
	})
	return recs, nil
}

// DeleteSubtitle removes all records matching file from the store.
func (p *PebbleStore) DeleteSubtitle(file string) error {
	iter, err := p.db.NewIter(nil)
	if err != nil {
		return err
	}
	defer iter.Close()
	for iter.First(); iter.Valid(); iter.Next() {
		if !strings.HasPrefix(string(iter.Key()), "subtitle:") {
			continue
		}
		var r SubtitleRecord
		if err := json.Unmarshal(iter.Value(), &r); err != nil {
			return err
		}
		if r.File == file {
			if err := p.db.Delete(iter.Key(), pebble.Sync); err != nil {
				return err
			}
		}
	}
	return iter.Error()
}

// InsertDownload stores a download record in PebbleDB.
// The ID field is generated when empty.
func (p *PebbleStore) InsertDownload(rec *DownloadRecord) error {
	if rec.ID == "" {
		rec.ID = uuid.NewString()
	}
	if rec.CreatedAt.IsZero() {
		rec.CreatedAt = time.Now()
	}
	b, err := json.Marshal(rec)
	if err != nil {
		return err
	}
	key := []byte("download:" + rec.ID)
	return p.db.Set(key, b, pebble.Sync)
}

// ListDownloads returns all download records sorted by creation time descending.
func (p *PebbleStore) ListDownloads() ([]DownloadRecord, error) {
	iter, err := p.db.NewIter(nil)
	if err != nil {
		return nil, err
	}
	defer iter.Close()
	var recs []DownloadRecord
	for iter.First(); iter.Valid(); iter.Next() {
		if !strings.HasPrefix(string(iter.Key()), "download:") {
			continue
		}
		var r DownloadRecord
		if err := json.Unmarshal(iter.Value(), &r); err != nil {
			return nil, err
		}
		recs = append(recs, r)
	}
	if err := iter.Error(); err != nil {
		return nil, err
	}
	sort.Slice(recs, func(i, j int) bool {
		return recs[i].CreatedAt.After(recs[j].CreatedAt)
	})
	return recs, nil
}

// DeleteDownload removes download records for the given subtitle file.
func (p *PebbleStore) DeleteDownload(file string) error {
	iter, err := p.db.NewIter(nil)
	if err != nil {
		return err
	}
	defer iter.Close()
	for iter.First(); iter.Valid(); iter.Next() {
		if !strings.HasPrefix(string(iter.Key()), "download:") {
			continue
		}
		var r DownloadRecord
		if err := json.Unmarshal(iter.Value(), &r); err != nil {
			return err
		}
		if r.File == file {
			if err := p.db.Delete(iter.Key(), pebble.Sync); err != nil {
				return err
			}
		}
	}
	return iter.Error()
}

// InsertMediaItem stores a media item.
func (p *PebbleStore) InsertMediaItem(rec *MediaItem) error {
	if rec.ID == "" {
		rec.ID = uuid.NewString()
	}
	if rec.CreatedAt.IsZero() {
		rec.CreatedAt = time.Now()
	}
	b, err := json.Marshal(rec)
	if err != nil {
		return err
	}
	key := []byte("media:" + rec.ID)
	return p.db.Set(key, b, pebble.Sync)
}

// ListMediaItems returns stored media items sorted by creation time.
func (p *PebbleStore) ListMediaItems() ([]MediaItem, error) {
	iter, err := p.db.NewIter(nil)
	if err != nil {
		return nil, err
	}
	defer iter.Close()
	var recs []MediaItem
	for iter.First(); iter.Valid(); iter.Next() {
		if !strings.HasPrefix(string(iter.Key()), "media:") {
			continue
		}
		var r MediaItem
		if err := json.Unmarshal(iter.Value(), &r); err != nil {
			return nil, err
		}
		recs = append(recs, r)
	}
	if err := iter.Error(); err != nil {
		return nil, err
	}
	sort.Slice(recs, func(i, j int) bool { return recs[i].CreatedAt.After(recs[j].CreatedAt) })
	return recs, nil
}

// DeleteMediaItem removes records with matching path.
func (p *PebbleStore) DeleteMediaItem(path string) error {
	iter, err := p.db.NewIter(nil)
	if err != nil {
		return err
	}
	defer iter.Close()
	for iter.First(); iter.Valid(); iter.Next() {
		if !strings.HasPrefix(string(iter.Key()), "media:") {
			continue
		}
		var r MediaItem
		if err := json.Unmarshal(iter.Value(), &r); err != nil {
			return err
		}
		if r.Path == path {
			if err := p.db.Delete(iter.Key(), pebble.Sync); err != nil {
				return err
			}
		}
	}
	return iter.Error()
}

// InsertTag is unsupported for PebbleStore and returns nil for compatibility.
func (p *PebbleStore) InsertTag(name string) error { return nil }

// ListTags returns no tags for PebbleStore.
func (p *PebbleStore) ListTags() ([]Tag, error) { return nil, nil }

// DeleteTag is a no-op for PebbleStore.
func (p *PebbleStore) DeleteTag(id int64) error { return nil }

// AssignTagToUser is a no-op for PebbleStore.
func (p *PebbleStore) AssignTagToUser(userID, tagID int64) error { return nil }

// RemoveTagFromUser is a no-op for PebbleStore.
func (p *PebbleStore) RemoveTagFromUser(userID, tagID int64) error { return nil }

// ListTagsForUser returns no tags for PebbleStore.
func (p *PebbleStore) ListTagsForUser(userID int64) ([]Tag, error) { return nil, nil }

// AssignTagToMedia is a no-op for PebbleStore.
func (p *PebbleStore) AssignTagToMedia(mediaID, tagID int64) error { return nil }

// RemoveTagFromMedia is a no-op for PebbleStore.
func (p *PebbleStore) RemoveTagFromMedia(mediaID, tagID int64) error { return nil }

// ListTagsForMedia returns no tags for PebbleStore.
func (p *PebbleStore) ListTagsForMedia(mediaID int64) ([]Tag, error) { return nil, nil }
