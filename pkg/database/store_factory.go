package database

// OpenStore selects a storage backend and returns a SubtitleStore with performance optimizations.
// Backend may be "sqlite", "pebble" or "postgres". Any other value defaults to SQLite.
//
// This function applies connection pooling, database indexes, and performance settings
// automatically for SQL-based backends to improve response times and reduce resource usage.
func OpenStore(path, backend string) (SubtitleStore, error) {
	switch backend {
	case "pebble":
		return OpenPebble(path)
	case "postgres":
		store, err := OpenPostgresStore(path)
		if err != nil {
			return nil, err
		}
		// Apply performance optimizations for PostgreSQL
		if store.db != nil {
			OptimizeConnectionPool(store.db, "postgres")
			_ = CreatePerformanceIndexes(store.db, "postgres")
			_ = OptimizeDatabaseSettings(store.db, "postgres")
		}
		return store, nil
	default:
		store, err := OpenSQLStore(path)
		if err != nil {
			return nil, err
		}
		// Apply performance optimizations for SQLite
		if store.db != nil {
			OptimizeConnectionPool(store.db, "sqlite")
			_ = CreatePerformanceIndexes(store.db, "sqlite")
			_ = OptimizeDatabaseSettings(store.db, "sqlite")
		}
		return store, nil
	}
}
