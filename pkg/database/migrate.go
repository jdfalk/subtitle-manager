package database

// MigrateToPebble copies subtitle history from a SQLite database into a Pebble store.
func MigrateToPebble(sqlitePath, pebblePath string) error {
	sqlStore, err := OpenSQLStore(sqlitePath)
	if err != nil {
		return err
	}
	defer sqlStore.Close()

	recs, err := sqlStore.ListSubtitles()
	if err != nil {
		return err
	}

	pebbleStore, err := OpenPebble(pebblePath)
	if err != nil {
		return err
	}
	defer pebbleStore.Close()

	for _, r := range recs {
		rec := r
		if err := pebbleStore.InsertSubtitle(&rec); err != nil {
			return err
		}
	}
	return nil
}
