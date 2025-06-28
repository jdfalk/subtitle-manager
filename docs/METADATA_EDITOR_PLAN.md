<!-- file: docs/METADATA_EDITOR_PLAN.md -->

# Media Metadata Editor Plan

This document outlines the approach for implementing a manual metadata editor.

## Requirements

- Allow manual search and selection of correct metadata when automatic lookup
  fails during import.
- Store alternate titles for anime and foreign releases so subtitle providers
  have more search options.
- Track release group information alongside each media item and maintain
  subtitle download history.
- Provide field-level locks to prevent updates from overwriting user edits.

## Proposed Design

1. **Database Updates**
   - Extend `media_items` table with new columns: `release_group`, `alt_titles`
     (JSON array) and `field_locks` (JSON object).
   - Provide helper `addColumnIfNotExists` in `initSchema` to allow automatic
     schema upgrades.
   - New CRUD functions for these fields in `pkg/database`.
2. **CLI Commands**
   - Add `metadata` command with `search` and `update` subcommands.
   - `search` queries TMDB for a title and prints the top results for manual
     selection.
   - `update` applies user supplied metadata and optional locks to an existing
     `media_item`.
3. **Subtitle History**
   - Reuse existing `DownloadRecord` and `SubtitleRecord` tables; add helper to
     list history per file.
4. **Locks**
   - Store locked fields as a JSON map: `{ "title": true, "year": true }`.
   - Update metadata import routines to check these locks before modifying a
     field.

## Implementation Steps

1. Update the database schema and structs.
2. Implement helper functions to set/retrieve release group, alternate titles
   and locks.
3. Create new CLI commands.
4. Write unit tests covering the new database logic.
5. Document usage in README.
