package database

// OpenStore selects a storage backend and returns a SubtitleStore.
// Backend may be "sqlite", "pebble" or "postgres". Any other value defaults to SQLite.
func OpenStore(path, backend string) (SubtitleStore, error) {
	switch backend {
	case "pebble":
		return OpenPebble(path)
	case "postgres":
		return OpenPostgresStore(path)
	default:
		return OpenSQLStore(path)
	}
}

// OpenStoreWithConfig opens a store using the current configuration.
func OpenStoreWithConfig() (SubtitleStore, error) {
	path := GetDatabasePath()
	backend := GetDatabaseBackend()
	return OpenStore(path, backend)
}
