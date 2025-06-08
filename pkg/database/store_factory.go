package database

// OpenStore selects a storage backend and returns a SubtitleStore.
// backend may be "sqlite" or "pebble". Any other value defaults to SQLite.
func OpenStore(path, backend string) (SubtitleStore, error) {
	switch backend {
	case "pebble":
		return OpenPebble(path)
	default:
		return OpenSQLStore(path)
	}
}
