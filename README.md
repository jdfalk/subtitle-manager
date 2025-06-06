# Subtitle Manager

A simple command-line tool for converting, merging and translating subtitles.

## Commands

- `convert [input] [output]` - convert a subtitle file to SRT
- `merge [sub1] [sub2] [output]` - merge two subtitles sorted by start time
- `translate [input] [output] [lang]` - translate a subtitle using Google Translate or ChatGPT

Use `--log-levels` to set per-component log levels, e.g. `--log-levels translate=debug`.

API keys for Google Translate and OpenAI can be provided via flags `--google-key` and `--openai-key` or through configuration.

Configuration can be stored in `$HOME/.subtitle-manager.yaml`.
