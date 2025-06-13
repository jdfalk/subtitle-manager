package database

// Migrate copies subtitle history and download records from src to dest.
func Migrate(src, dest SubtitleStore) error {
	recs, err := src.ListSubtitles()
	if err != nil {
		return err
	}
	downloads, err := src.ListDownloads()
	if err != nil {
		return err
	}
	media, err := src.ListMediaItems()
	if err != nil {
		return err
	}
	for _, r := range recs {
		rec := r
		if err := dest.InsertSubtitle(&rec); err != nil {
			return err
		}
	}
	for _, d := range downloads {
		dr := d
		if err := dest.InsertDownload(&dr); err != nil {
			return err
		}
	}
	for _, m := range media {
		mr := m
		if err := dest.InsertMediaItem(&mr); err != nil {
			return err
		}
	}
	return nil
}

// MigrateToPebble copies subtitle history from a SQLite database into a Pebble store.
func MigrateToPebble(sqlitePath, pebblePath string) error {
	sqlStore, err := OpenSQLStore(sqlitePath)
	if err != nil {
		return err
	}
	defer sqlStore.Close()

	pebbleStore, err := OpenPebble(pebblePath)
	if err != nil {
		return err
	}
	defer pebbleStore.Close()

	return Migrate(sqlStore, pebbleStore)
}
