# file: docs/SCHEDULING.md

# Automated Task Scheduling

Subtitle Manager includes several maintenance tasks that keep the system running
smoothly. Each task can be scheduled using a simple frequency string such as
`hourly`, `daily`, `weekly` or any valid duration like `12h`.

## Tasks

### Database Cleanup

Removes expired sessions and optimizes the SQLite database when applicable.

Configuration key: `db_cleanup_frequency`

### Metadata Refresh

Updates stored media items by querying TMDB and OMDb. Requires valid
`tmdb_api_key` and `omdb_api_key` values.

Configuration key: `metadata_refresh_frequency`

### Disk Scan

Calculates disk usage under the configured `db_path`.

Configuration key: `disk_scan_frequency`

### Self Test

Periodically verifies core components are operational. Currently checks database
connectivity and exits on failure so containers can restart.

Configuration key: `selftest_frequency`

Set the desired frequency in `config.yaml` or via the `/api/config` endpoint.
The web UI scheduling page provides a convenient interface to adjust these
values.
