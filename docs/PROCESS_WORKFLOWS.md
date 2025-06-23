<!-- file: docs/PROCESS_WORKFLOWS.md -->
# Process Workflows

This document illustrates key workflows in Subtitle Manager using ASCII diagrams.

## Convert Workflow

```
User -> convert command -> subtitles.ConvertToSRT -> write output
```

1. User runs `subtitle-manager convert input.ssa output.srt`.
2. Command loads configuration and logger.
3. `ConvertToSRT` parses the input and produces SRT bytes.
4. Result is saved to `output.srt`.

## Translate Workflow

```
User -> translate command -> translator.Translate -> database.InsertSubtitle
```

1. User runs `subtitle-manager translate movie.srt translated.srt fr`.
2. Translation backend is selected based on configuration.
3. Translated subtitles are written to disk and recorded in the database.

## Sync Workflow

```
User -> sync command -> syncer.Sync -> subtitles.Merge -> write output
```

1. Sync reads media and subtitle files.
2. If `--use-audio` is set, an audio track is extracted for alignment.
3. Merged subtitles are written to the target file.

## Fetch Workflow

```
Cron -> scheduler -> providers.Fetch -> database.InsertDownload
```

1. Scheduled scan triggers `providers.Fetch` for missing subtitles.
2. Downloads are stored in the database and written to disk.

