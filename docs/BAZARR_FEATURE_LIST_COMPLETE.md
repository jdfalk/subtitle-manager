# file: docs/BAZARR_FEATURE_LIST_COMPLETE.md

# Comprehensive Bazarr Feature List

This document aggregates features described across the official Bazarr wiki and website.

## Installation & Platform Support

- Runs on **Windows**, **Linux**, **macOS**, **FreeBSD**, Raspberry Pi and in **Docker** containers.
- Provides Windows installer, source-based install, and Docker images.
- Supports autostart scripts for Windows, Linux, macOS and FreeBSD.

## Core Functionality

- **Automatic search** for missing subtitles and download when they become available.
- **Manual search** to choose specific subtitles from available matches.
- **Subtitle upgrade** process to periodically replace existing subtitles with better ones.
- Works alongside **Sonarr** and **Radarr** to manage episodes and movies.
- Mass-edit capability to apply language profiles to existing series and movies.
- Only media added after installation is automatically processed for subtitles unless manually triggered.

## Settings Overview

### General

- Configure bind address, port, and optional URL base for reverse proxies.
- Supports basic or form authentication, username/password and API key management.
- Proxy configuration with HTTP(S), Socks4 or Socks5 types and credential options.
- UI theme options and debug logging toggle.
- Optional anonymous analytics sharing.

### Sonarr and Radarr Integration

- Hostname, port, URL base, SSL and API key settings.
- Minimum score for downloads, monitored-only mode and excluded tags or series types.
- Path mappings when Bazarr and Sonarr/Radarr see files under different paths.

### Subtitles Options

- Select destination folder, upgrade previously downloaded subtitles and history depth.
- Upgrade manually downloaded subtitles toggle.
- Anti-captcha provider configuration.
- Performance options: adaptive searching, simultaneous provider search, use embedded subtitles, ignore PGS subtitles, show only desired languages.
- Post-processing features: UTF8 encoding, chmod permissions, automatic sync, score thresholds, custom scripts.

### Languages

- Enable/disable single language filenames and select enabled languages.
- Default language profiles for new series and movies including forced and hearing‑impaired options.

### Providers and Notifications

- Enable and configure multiple subtitle providers.
- Notification settings using Apprise configuration strings.

### Scheduler

- Adjustable intervals for Sonarr/Radarr sync, disk indexing and subtitle search/upgrade tasks.

## Additional Configuration

- **Performance Tuning** options such as disabling automatic sync, disabling disk scans, reducing search frequency and more.
- **PostgreSQL** database support with environment variables and migration from SQLite via PGLoader.
- **Reverse proxy** guidance for Nginx, Apache and Authelia setups including base URL support.
- **Webhooks** to POST JSON on events; example integration with Plex to trigger subtitle searches on playback.
- **Whisper provider** for generating subtitles using external whisper-asr-webservice with selectable model and backend, CPU/GPU deployment instructions and timeout configuration.

## Website Highlights

- Download packages for Windows, Python sources and Docker containers with automatic updates.
- Links to documentation, Discord, GitHub issues and subreddit for community support.

This list summarizes the capabilities referenced in the official Bazarr documentation to help ensure Subtitle Manager achieves or surpasses feature parity.
