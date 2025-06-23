// file: pkg/database/pebble.go
package database

import (
	"encoding/json"
	"errors"
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

func mediaPathKey(path string) []byte {
	return []byte("media_path:" + path)
}

func mediaKey(id string) []byte {
	return []byte("media:" + id)
}

func (p *PebbleStore) getMediaByPath(path string) (*MediaItem, string, error) {
	val, closer, err := p.db.Get(mediaPathKey(path))
	if err != nil {
		if errors.Is(err, pebble.ErrNotFound) {
			return nil, "", nil
		}
		return nil, "", err
	}
	closer.Close()
	id := string(val)
	data, closer, err := p.db.Get(mediaKey(id))
	if err != nil {
		return nil, "", err
	}
	defer closer.Close()
	var it MediaItem
	if err := json.Unmarshal(data, &it); err != nil {
		return nil, "", err
	}
	return &it, id, nil
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
	batch := p.db.NewBatch()
	if err := batch.Set(mediaKey(rec.ID), b, nil); err != nil {
		batch.Close()
		return err
	}
	if err := batch.Set(mediaPathKey(rec.Path), []byte(rec.ID), nil); err != nil {
		batch.Close()
		return err
	}
	return batch.Commit(pebble.Sync)
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
	item, id, err := p.getMediaByPath(path)
	if err != nil || item == nil {
		return err
	}
	batch := p.db.NewBatch()
	if err := batch.Delete(mediaKey(id), nil); err != nil {
		batch.Close()
		return err
	}
	if err := batch.Delete(mediaPathKey(path), nil); err != nil {
		batch.Close()
		return err
	}
	return batch.Commit(pebble.Sync)
}

// CountSubtitles returns the number of subtitle records.
func (p *PebbleStore) CountSubtitles() (int, error) {
	iter, err := p.db.NewIter(nil)
	if err != nil {
		return 0, err
	}
	defer iter.Close()
	count := 0
	for iter.First(); iter.Valid(); iter.Next() {
		if strings.HasPrefix(string(iter.Key()), "subtitle:") {
			count++
		}
	}
	if err := iter.Error(); err != nil {
		return 0, err
	}
	return count, nil
}

// CountDownloads returns the number of download records.
func (p *PebbleStore) CountDownloads() (int, error) {
	iter, err := p.db.NewIter(nil)
	if err != nil {
		return 0, err
	}
	defer iter.Close()
	count := 0
	for iter.First(); iter.Valid(); iter.Next() {
		if strings.HasPrefix(string(iter.Key()), "download:") {
			count++
		}
	}
	if err := iter.Error(); err != nil {
		return 0, err
	}
	return count, nil
}

// CountMediaItems returns the number of media item records.
func (p *PebbleStore) CountMediaItems() (int, error) {
	iter, err := p.db.NewIter(nil)
	if err != nil {
		return 0, err
	}
	defer iter.Close()
	count := 0
	for iter.First(); iter.Valid(); iter.Next() {
		if strings.HasPrefix(string(iter.Key()), "media:") {
			count++
		}
	}
	if err := iter.Error(); err != nil {
		return 0, err
	}
	return count, nil
}

// InsertTag is unsupported for PebbleStore and returns nil for compatibility.
func (p *PebbleStore) InsertTag(name string) error { return nil }

// ListTags returns no tags for PebbleStore.
func (p *PebbleStore) ListTags() ([]Tag, error) { return nil, nil }

// DeleteTag is a no-op for PebbleStore.
func (p *PebbleStore) DeleteTag(id int64) error { return nil }

// UpdateTag is a no-op for PebbleStore.
func (p *PebbleStore) UpdateTag(id int64, name string) error { return nil }

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

// SetMediaReleaseGroup stores the release group in the media item record.
func (p *PebbleStore) SetMediaReleaseGroup(path, group string) error {
	item, _, err := p.getMediaByPath(path)
	if err != nil || item == nil {
		return err
	}
	item.ReleaseGroup = group
	return p.InsertMediaItem(item)
}

// SetMediaAltTitles stores alternate titles in the media item record.
func (p *PebbleStore) SetMediaAltTitles(path string, titles []string) error {
	item, _, err := p.getMediaByPath(path)
	if err != nil || item == nil {
		return err
	}
	data, _ := json.Marshal(titles)
	item.AltTitles = string(data)
	return p.InsertMediaItem(item)
}

// SetMediaFieldLocks stores locked fields in the media item record.
func (p *PebbleStore) SetMediaFieldLocks(path, locks string) error {
	item, _, err := p.getMediaByPath(path)
	if err != nil || item == nil {
		return err
	}
	item.FieldLocks = locks
	return p.InsertMediaItem(item)
}

// SetMediaTitle updates the title in the media item record.
func (p *PebbleStore) SetMediaTitle(path, title string) error {
	item, _, err := p.getMediaByPath(path)
	if err != nil || item == nil {
		return err
	}
	item.Title = title
	return p.InsertMediaItem(item)
}
