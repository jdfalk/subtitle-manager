# Proof of Concept Tasks

This document enumerates the small steps required to demonstrate a working prototype of Subtitle Manager with a web interface.

## 1. Media Library Scanning
- [x] Create `pkg/metadata` package to parse file names and query TheMovieDB for movie/episode details.
- [x] Extend database schema with a `media_items` table storing video path, title, season and episode numbers.
- [x] Add CLI command `scanlib` that uses the new metadata package to populate the table.
- [ ] Implement REST endpoint `/api/scan` to trigger a library scan from the web UI.
- [x] Display scan progress and results in React dashboard.

## 2. Subtitle Extraction
- [x] Expose `subtitles.ExtractFromMedia` via new REST endpoint `/api/extract`.
- [x] Allow the UI to request extraction for a selected media item.
- [x] Store extracted subtitle paths in the database for later reference.

## 3. Subtitle Downloading
- [ ] Implement REST endpoint `/api/download` calling `providers.Get` and `scanner.ProcessFile`.
- [ ] Provide UI controls to choose language and provider when requesting a download.
- [ ] Record download events using `database.InsertDownload`.

## 4. Conversion and Translation
- [ ] Add REST endpoint `/api/convert` wrapping the `convert` CLI logic.
- [ ] Add REST endpoint `/api/translate` wrapping the `translate` CLI logic.
- [ ] UI forms for converting and translating uploaded subtitle files.

## 5. Saving Subtitles
- [ ] Ensure downloaded or translated subtitles are written next to the video file using `<name>.<lang>.srt`.
- [ ] Verify Plex and similar tools detect the saved subtitles.
- [ ] Provide a delete function to remove generated subtitles.

## 6. Additional POC Requirements
- [ ] Authentication flow for the web UI using existing session system.
- [ ] Basic settings page allowing configuration of API keys and preferred languages.
- [ ] Unit tests for new REST endpoints and metadata parsing.
- [ ] Update README and CHANGELOG once POC is functional.

