# TODO

This file tracks remaining work and implementation status for Subtitle Manager. **Note: The project is ~99% complete with full production readiness achieved.**

---

## 🔄 Low Priority Improvements

### Testing Enhancements

- [ ] **Mock PostgreSQL tests**: Mock out PostgreSQL logic as much as possible to reduce dependency on local PostgreSQL installation for test coverage
  - Current: Tests skip gracefully if PostgreSQL unavailable
  - Goal: Better test coverage without requiring PostgreSQL setup
  - Priority: Low (current approach works well)

---

## 📊 Complete Bazarr Feature Parity Analysis

This section provides a comprehensive comparison between Bazarr and Subtitle Manager, cataloguing every feature from Bazarr's [official documentation](https://wiki.bazarr.media/) and [repository](https://github.com/morpheus65535/bazarr).

### Executive Summary: Feature Parity Status

| **Feature Category**   | **Bazarr Status**         | **Our Implementation**                                                               | **Gold Standard**             |
| ---------------------- | ------------------------- | ------------------------------------------------------------------------------------ | ----------------------------- |
| **Subtitle Providers** | 40+ providers supported   | ✅ 40+ providers ([registry.go](pkg/providers/registry.go))                          | ✅ Full parity achieved       |
| **Web Interface**      | Modern React UI           | ✅ Complete React app ([webui/src/](webui/src/))                                     | ✅ Production ready           |
| **Authentication**     | Basic auth + API keys     | ✅ Password, OAuth2, API keys, RBAC ([pkg/auth/](pkg/auth/))                         | ✅ Enterprise grade           |
| **Library Management** | Sonarr/Radarr integration | ✅ Full integration ([cmd/sonarr.go](cmd/sonarr.go), [cmd/radarr.go](cmd/radarr.go)) | ✅ Complete                   |
| **Database Support**   | SQLite + PostgreSQL       | ✅ SQLite, PebbleDB, PostgreSQL ([pkg/database/](pkg/database/))                     | ✅ Full enterprise DB support |
| **REST API**           | Comprehensive API         | ✅ Full coverage ([pkg/webserver/](pkg/webserver/))                                  | ✅ Complete                   |
| **Configuration**      | YAML + Web UI             | ✅ Viper + Web settings ([cmd/root.go](cmd/root.go))                                 | ✅ Full parity                |
| **Webhooks**           | Basic webhook support     | ✅ Sonarr/Radarr/custom endpoints ([pkg/webhooks/](pkg/webhooks/))                   | ✅ Enhanced system            |
| **Notifications**      | Apprise integration       | ✅ Discord/Telegram/SMTP ([pkg/notifications/](pkg/notifications/))                  | ✅ Multi-provider support     |
| **Anti-Captcha**       | Anti-captcha.com          | ✅ Anti-Captcha.com + 2captcha.com ([pkg/captcha/](pkg/captcha/))                    | ✅ Multi-service support      |
| **Scheduling**         | Basic cron support        | ✅ Advanced cron expressions ([pkg/scheduler/](pkg/scheduler/))                      | ✅ Enhanced scheduling        |

### Detailed Feature Analysis

#### 1. Core Subtitle Operations ✅ 100% Complete

| Bazarr Feature               | Implementation Status | Code Reference                       |
| ---------------------------- | --------------------- | ------------------------------------ |
| Format conversion            | ✅ Complete           | [cmd/convert.go](cmd/convert.go)     |
| Subtitle merging             | ✅ Complete           | [cmd/merge.go](cmd/merge.go)         |
| Media extraction             | ✅ Complete           | [cmd/extract.go](cmd/extract.go)     |
| Translation (Google/ChatGPT) | ✅ Complete           | [cmd/translate.go](cmd/translate.go) |
| Batch processing             | ✅ Complete           | [cmd/batch.go](cmd/batch.go)         |
| History tracking             | ✅ Complete           | [cmd/history.go](cmd/history.go)     |

#### 2. Authentication & Authorization ✅ 100% Complete

| Bazarr Feature            | Implementation Status | Code Reference                                   |
| ------------------------- | --------------------- | ------------------------------------------------ |
| Password authentication   | ✅ Complete           | [pkg/auth/auth.go](pkg/auth/auth.go)             |
| API key management        | ✅ Complete           | [cmd/user.go](cmd/user.go)                       |
| Session management        | ✅ Complete           | [pkg/webserver/auth.go](pkg/webserver/auth.go)   |
| Role-based access control | ✅ Complete           | [pkg/auth/rbac.go](pkg/auth/rbac.go)             |
| OAuth2 (GitHub)           | ✅ Complete           | [pkg/webserver/oauth.go](pkg/webserver/oauth.go) |
| One-time tokens           | ✅ Complete           | [cmd/user.go](cmd/user.go)                       |

#### 3. Subtitle Providers ✅ 100% Complete - Full Bazarr Parity

| Provider Category      | Bazarr Count | Our Implementation                                           | Status             |
| ---------------------- | ------------ | ------------------------------------------------------------ | ------------------ |
| Major providers        | ~40          | ✅ 40+ providers                                             | ✅ Parity achieved |
| OpenSubtitles variants | 3            | ✅ Complete ([opensubtitles/](pkg/providers/opensubtitles/)) | ✅                 |
| Regional providers     | ~25          | ✅ Complete (Greek, Turkish, etc.)                           | ✅                 |
| Torrent-based          | ~8           | ✅ Complete (YIFY, etc.)                                     | ✅                 |
| Embedded extraction    | 1            | ✅ Complete ([embedded/](pkg/providers/embedded/))           | ✅                 |
| Whisper transcription  | 1            | ✅ Complete ([transcribe.go](cmd/transcribe.go))             | ✅                 |

#### 4. Web Interface Pages ✅ 100% Complete

| Bazarr Page         | Implementation Status     | Code Reference                                     |
| ------------------- | ------------------------- | -------------------------------------------------- |
| Dashboard           | ✅ Complete               | [webui/src/Dashboard.jsx](webui/src/Dashboard.jsx) |
| Settings            | ✅ Complete               | [webui/src/Settings.jsx](webui/src/Settings.jsx)   |
| History             | ✅ Complete               | [webui/src/History.jsx](webui/src/History.jsx)     |
| Wanted              | ✅ Complete               | [webui/src/Wanted.jsx](webui/src/Wanted.jsx)       |
| System/Logs         | ✅ Complete               | [webui/src/System.jsx](webui/src/System.jsx)       |
| Providers           | ✅ Integrated in Settings | Settings page                                      |
| Subtitle extraction | ✅ Complete               | [webui/src/Extract.jsx](webui/src/Extract.jsx)     |

#### 5. Integration Features ✅ 90% Complete

| Bazarr Feature     | Implementation Status | Code Reference                                     |
| ------------------ | --------------------- | -------------------------------------------------- |
| Sonarr integration | ✅ Complete           | [cmd/sonarr.go](cmd/sonarr.go)                     |
| Radarr integration | ✅ Complete           | [cmd/radarr.go](cmd/radarr.go)                     |
| Plex integration   | ✅ Complete           | [cmd/plex.go](cmd/plex.go), [pkg/plex/](pkg/plex/) |
| Library scanning   | ✅ Complete           | [cmd/scan.go](cmd/scan.go)                         |
| Directory watching | ✅ Complete           | [cmd/watch.go](cmd/watch.go)                       |
| Webhooks           | ✅ Complete           | [pkg/webhooks](pkg/webhooks/)                      |
| Notifications      | 🔶 Planned            | [TODO] Discord/Telegram/Email                      |

#### 6. Advanced Features ✅ 100% Complete

| Bazarr Feature        | Implementation Status | Code Reference                            |
| --------------------- | --------------------- | ----------------------------------------- |
| PostgreSQL support    | ✅ Complete           | SQLite, PebbleDB and PostgreSQL available |
| Reverse proxy support | 🔶 Partial            | Basic configuration available             |
| Anti-captcha service  | ✅ Complete           | [pkg/captcha/](pkg/captcha/)              |
| Performance tuning    | ✅ Complete           | Concurrent workers, pools                 |
| Custom scheduling     | ✅ Complete           | [pkg/scheduler/](pkg/scheduler/)          |
| Bazarr config import  | ✅ Complete           | [cmd/import.go](cmd/import.go)            |
| Webhook system        | ✅ Complete           | [pkg/webhooks/](pkg/webhooks/)            |
| Notifications         | ✅ Complete           | [pkg/notifications/](pkg/notifications/)  |

### Complete Provider Implementation Analysis

**Reference**: [Bazarr Providers](https://wiki.bazarr.media/Additional-Configuration/Settings/#providers) vs [Our Registry](pkg/providers/registry.go)

#### Implemented Providers (40+ with Full Parity)

| Provider                | Bazarr | Our Implementation | Documentation                                                               |
| ----------------------- | ------ | ------------------ | --------------------------------------------------------------------------- |
| Addic7ed                | ✅     | ✅                 | [addic7ed/](pkg/providers/addic7ed/)                                        |
| AnimeKalesi             | ✅     | ✅                 | [animekalesi/](pkg/providers/animekalesi/)                                  |
| Animetosho              | ✅     | ✅                 | [animetosho/](pkg/providers/animetosho/)                                    |
| Assrt                   | ✅     | ✅                 | [assrt/](pkg/providers/assrt/)                                              |
| AvistaZ/CinemaZ         | ✅     | ✅                 | [avistaz/](pkg/providers/avistaz/)                                          |
| BetaSeries              | ✅     | ✅                 | [betaseries/](pkg/providers/betaseries/)                                    |
| BSplayer                | ✅     | ✅                 | [bsplayer/](pkg/providers/bsplayer/)                                        |
| Embedded Subtitles      | ✅     | ✅                 | [embedded/](pkg/providers/embedded/)                                        |
| Gestdown.info           | ✅     | ✅                 | [gestdown/](pkg/providers/gestdown/)                                        |
| GreekSubs               | ✅     | ✅                 | [greeksubs/](pkg/providers/greeksubs/)                                      |
| GreekSubtitles          | ✅     | ✅                 | [greeksubtitles/](pkg/providers/greeksubtitles/)                            |
| HDBits.org              | ✅     | ✅                 | [hdbits/](pkg/providers/hdbits/)                                            |
| Hosszupuska             | ✅     | ✅                 | [hosszupuska/](pkg/providers/hosszupuska/)                                  |
| Karagarga.in            | ✅     | ✅                 | [karagarga/](pkg/providers/karagarga/)                                      |
| Ktuvit                  | ✅     | ✅                 | [ktuvit/](pkg/providers/ktuvit/)                                            |
| LegendasDivx            | ✅     | ✅                 | [legendasdivx/](pkg/providers/legendasdivx/)                                |
| Legendas.net            | ✅     | ✅                 | [legendasnet/](pkg/providers/legendasnet/)                                  |
| Napiprojekt             | ✅     | ✅                 | [napiprojekt/](pkg/providers/napiprojekt/)                                  |
| Napisy24                | ✅     | ✅                 | [napisy24/](pkg/providers/napisy24/)                                        |
| Nekur                   | ✅     | ✅                 | [nekur/](pkg/providers/nekur/)                                              |
| OpenSubtitles.com       | ✅     | ✅                 | [opensubtitlescom/](pkg/providers/opensubtitlescom/)                        |
| OpenSubtitles.org (VIP) | ✅     | ✅                 | [opensubtitlesvip/](pkg/providers/opensubtitlesvip/)                        |
| Podnapisi               | ✅     | ✅                 | [podnapisi/](pkg/providers/podnapisi/)                                      |
| RegieLive               | ✅     | ✅                 | [regielive/](pkg/providers/regielive/)                                      |
| Sous-Titres.eu          | ✅     | ✅                 | [soustitres/](pkg/providers/soustitres/)                                    |
| Subdivx                 | ✅     | ✅                 | [subdivx/](pkg/providers/subdivx/)                                          |
| subf2m.co               | ✅     | ✅                 | [subf2m/](pkg/providers/subf2m/)                                            |
| Subs.sab.bz             | ✅     | ✅                 | [subssabbz/](pkg/providers/subssabbz/)                                      |
| Subs4Free               | ✅     | ✅                 | [subs4free/](pkg/providers/subs4free/)                                      |
| Subs4Series             | ✅     | ✅                 | [subs4series/](pkg/providers/subs4series/)                                  |
| Subscene                | ✅     | ✅                 | [subscene/](pkg/providers/subscene/)                                        |
| Subscenter              | ✅     | ✅                 | [subscenter/](pkg/providers/subscenter/)                                    |
| Subsunacs.net           | ✅     | ✅                 | [subsunacs/](pkg/providers/subsunacs/)                                      |
| SubSynchro              | ✅     | ✅                 | [subsynchro/](pkg/providers/subsynchro/)                                    |
| Subtitrari-noi.ro       | ✅     | ✅                 | [subtitrarinoi/](pkg/providers/subtitrarinoi/)                              |
| subtitri.id.lv          | ✅     | ✅                 | [subtitriidlv/](pkg/providers/subtitriidlv/)                                |
| Subtitulamos.tv         | ✅     | ✅                 | [subtitulamos/](pkg/providers/subtitulamos/)                                |
| Supersubtitles          | ✅     | ✅                 | [supersubtitles/](pkg/providers/supersubtitles/)                            |
| Titlovi                 | ✅     | ✅                 | [titlovi/](pkg/providers/titlovi/)                                          |
| Titrari.ro              | ✅     | ✅                 | [titrariro/](pkg/providers/titrariro/)                                      |
| Titulky.com             | ✅     | ✅                 | [titulky/](pkg/providers/titulky/)                                          |
| Turkcealtyazi.org       | ✅     | ✅                 | [turkcealtyazi/](pkg/providers/turkcealtyazi/)                              |
| TuSubtitulo             | ✅     | ✅                 | [tusubtitulo/](pkg/providers/tusubtitulo/)                                  |
| TVSubtitles             | ✅     | ✅                 | [tvsubtitles/](pkg/providers/tvsubtitles/)                                  |
| Whisper                 | ✅     | ✅                 | [whisper/](pkg/providers/whisper/) + [cmd/transcribe.go](cmd/transcribe.go) |
| Wizdom                  | ✅     | ✅                 | [wizdom/](pkg/providers/wizdom/)                                            |
| XSubs                   | ✅     | ✅                 | [xsubs/](pkg/providers/xsubs/)                                              |
| Yavka.net               | ✅     | ✅                 | [yavka/](pkg/providers/yavka/)                                              |
| YIFY Subtitles          | ✅     | ✅                 | [yifysubtitles/](pkg/providers/yifysubtitles/)                              |
| Zimuku                  | ✅     | ✅                 | [zimuku/](pkg/providers/zimuku/)                                            |

#### Bazarr Settings Comparison Analysis

**Reference**: [Bazarr Settings](https://wiki.bazarr.media/Additional-Configuration/Settings/)

| Bazarr Setting Category      | Implementation Status | Our Location                     | Bazarr Reference                                                                                                                      |
| ---------------------------- | --------------------- | -------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| **Host Settings**            | ✅ Complete           | [cmd/root.go](cmd/root.go)       | [Host](https://wiki.bazarr.media/Additional-Configuration/Settings/#host)                                                             |
| - Bind Address               | ✅ Complete           | Viper config                     | -                                                                                                                                     |
| - Port Number                | ✅ Complete           | Viper config                     | -                                                                                                                                     |
| - URL Base                   | ✅ Complete           | Reverse proxy support            | [URL Base](https://wiki.bazarr.media/Additional-Configuration/Settings/#url-base)                                                     |
| **Security Settings**        | ✅ Complete           | [pkg/auth/](pkg/auth/)           | [Security](https://wiki.bazarr.media/Additional-Configuration/Settings/#security)                                                     |
| - Authentication             | ✅ Enhanced           | Multi-mode auth                  | [Authentication](https://wiki.bazarr.media/Additional-Configuration/Settings/#authentication)                                         |
| - Username/Password          | ✅ Complete           | Hashed storage                   | -                                                                                                                                     |
| - API Key                    | ✅ Enhanced           | Multiple keys                    | [API Key](https://wiki.bazarr.media/Additional-Configuration/Settings/#api-key)                                                       |
| **Proxy Settings**           | ✅ Complete           | HTTP client config               | [Proxy](https://wiki.bazarr.media/Additional-Configuration/Settings/#proxy)                                                           |
| **Sonarr Integration**       | ✅ Complete           | [cmd/sonarr.go](cmd/sonarr.go)   | [Sonarr](https://wiki.bazarr.media/Additional-Configuration/Settings/#sonarr)                                                         |
| - Host Configuration         | ✅ Complete           | Viper config                     | -                                                                                                                                     |
| - API Key                    | ✅ Complete           | Secure storage                   | -                                                                                                                                     |
| - Path Mappings              | ✅ Complete           | Config mappings                  | [Path Mappings](https://wiki.bazarr.media/Additional-Configuration/Settings/#path-mappings)                                           |
| **Radarr Integration**       | ✅ Complete           | [cmd/radarr.go](cmd/radarr.go)   | [Radarr](https://wiki.bazarr.media/Additional-Configuration/Settings/#radarr)                                                         |
| **Subtitle Options**         | ✅ Complete           | [pkg/subtitles/](pkg/subtitles/) | [Subtitles](https://wiki.bazarr.media/Additional-Configuration/Settings/#subtitles)                                                   |
| - Subtitle Folder            | ✅ Complete           | Config option                    | -                                                                                                                                     |
| - Upgrade Logic              | ✅ Complete           | Auto-upgrade                     | [Upgrade Previously Downloaded](https://wiki.bazarr.media/Additional-Configuration/Settings/#upgrade-previously-downloaded-subtitles) |
| **Anti-Captcha**             | ✅ Basic              | [pkg/captcha/](pkg/captcha/)     | [Anti-Captcha Options](https://wiki.bazarr.media/Additional-Configuration/Settings/#anti-captcha-options)                             |
| **Performance/Optimization** | ✅ Complete           | Worker pools                     | [Performance](https://wiki.bazarr.media/Additional-Configuration/Settings/#performance-optimization)                                  |
| - Adaptive Searching         | 🔶 Basic              | Simple scheduling                | [Adaptive Searching](https://wiki.bazarr.media/Additional-Configuration/Settings/#adaptive-searching)                                 |
| - Simultaneous Search        | ✅ Complete           | Concurrent workers               | -                                                                                                                                     |
| - Embedded Subtitles         | ✅ Complete           | Full support                     | [Use Embedded Subtitles](https://wiki.bazarr.media/Additional-Configuration/Settings/#use-embedded-subtitles)                         |
| **Post-Processing**          | ✅ Complete           | UTF-8 encoding                   | [Post-Processing](https://wiki.bazarr.media/Additional-Configuration/Settings/#post-processing)                                       |
| **Languages**                | ✅ Complete           | 180+ languages                   | [Languages](https://wiki.bazarr.media/Additional-Configuration/Settings/#languages)                                                   |
| **Providers**                | ✅ Complete           | Full registry                    | [Providers](https://wiki.bazarr.media/Additional-Configuration/Settings/#providers)                                                   |
| **Notifications**            | ✅ Basic              | Infrastructure ready             | [Notifications](https://wiki.bazarr.media/Additional-Configuration/Settings/#notifications)                                           |
| **Scheduler**                | ✅ Basic              | Auto-scan available              | [Scheduler](https://wiki.bazarr.media/Additional-Configuration/Settings/#scheduler)                                                   |

### Missing Features Analysis

#### High Priority Missing (1% of project)

1. **PostgreSQL Backend** - Enterprise database support

   - Status: ✅ Complete for large deployments
   - Current: SQLite, PebbleDB and PostgreSQL fully functional
   - Reference: [PostgreSQL Database](https://wiki.bazarr.media/Additional-Configuration/PostgreSQL-Database/)

2. **Advanced Webhook System** - Enhanced event notifications

   - Status: ✅ Complete with Sonarr/Radarr/custom webhook endpoints implemented
   - Reference: [Webhooks](https://wiki.bazarr.media/Additional-Configuration/Webhooks/)

3. **Notification Services** - Discord, Telegram, Email alerts

   - Status: ✅ Complete implementation with Discord, Telegram and SMTP notifiers
   - Current: Full multi-provider notification system available
   - Reference: [Notifications](https://wiki.bazarr.media/Additional-Configuration/Settings/#notifications)

4. **Anti-Captcha Integration** - For challenging providers

   - Status: ✅ Complete with Anti-Captcha.com and 2captcha.com support
   - Current: Multi-service captcha solving available
   - Reference: [Anti-Captcha Options](https://wiki.bazarr.media/Additional-Configuration/Settings/#anti-captcha-options)

5. **Advanced Scheduling** - Granular scan controls
   - Status: ✅ Complete with cron-based scheduler implemented
   - Current: Supports interval or cron expression with full granular control
   - Reference: [Scheduler](https://wiki.bazarr.media/Additional-Configuration/Settings/#scheduler)

#### Medium Priority Missing (Convenience Features)

1. **Bazarr Settings Import** - Automated migration

   - Status: 🔶 Partial implementation
   - Current: Manual configuration transfer works
   - Reference: [docs/BAZARR_SETTINGS_SYNC.md](docs/BAZARR_SETTINGS_SYNC.md)

2. **Advanced Scheduling** - Granular scan controls

   - Status: ✅ Cron-based scheduler implemented
   - Current: Supports interval or cron expression
   - Reference: [Scheduler](https://wiki.bazarr.media/Additional-Configuration/Settings/#scheduler)

3. **Reverse Proxy Enhancement** - Base URL configuration
   - Status: 🔶 Basic support exists
   - Current: Works behind reverse proxies
   - Reference: [Reverse Proxy Help](https://wiki.bazarr.media/Additional-Configuration/Reverse-Proxy-Help/)

---

## 🎯 Remaining Optional Features (1%)

### 1. Advanced Database Support

- [x] **PostgreSQL backend**: Alternative to SQLite/PebbleDB for large deployments
  - Location: `pkg/database/postgres.go`
  - Reference: [PostgreSQL Database](https://wiki.bazarr.media/Additional-Configuration/PostgreSQL-Database/)
  - Note: PostgreSQL tests require a local PostgreSQL installation and will skip gracefully if unavailable
  - TODO: Mock out PostgreSQL logic as much as possible for better test coverage (low priority)
- [x] **Database migration tools**: Enhanced migration between database types
  - Location: `cmd/migrate.go`

### 2. Advanced Integration Features

- [x] **Webhook support**: Enhanced Plex event integration
  - Location: `pkg/webhooks/` (complete implementation)
  - Reference: [Webhooks](https://wiki.bazarr.media/Additional-Configuration/Webhooks/)
- [x] **Anti-captcha service**: For providers requiring captcha solving
  - Location: `pkg/captcha/` (complete with Anti-Captcha.com and 2captcha.com)
- [x] **Notification services**: Discord, Telegram, Email alerts
  - Location: `pkg/notifications/` (complete implementation)
- [ ] **Reverse proxy support**: Base URL configuration for proxy deployments
  - Location: `pkg/webserver/server.go` (enhance existing)
  - Reference: [Reverse Proxy Help](https://wiki.bazarr.media/Additional-Configuration/Reverse-Proxy-Help/)
- [x] **Advanced scheduler**: Enhanced periodic scanning with more granular controls
  - Location: `pkg/scheduler/` (cron support complete)

### 3. Bazarr Configuration Import (Complete)

- [x] Implement `import-bazarr` command that fetches settings from `/api/system/settings`
      using the user's API key.
  - Location: `cmd/import.go` (complete implementation)
  - Reference: `pkg/bazarr/client.go` (complete implementation)
- [x] Map Bazarr preferences for languages, providers and network options into
      the Viper configuration.
  - Location: `pkg/bazarr/mapper.go` (complete implementation)
- [x] Document the synchronization process in `docs/BAZARR_SETTINGS_SYNC.md` and
      expose it through the welcome workflow.
  - Reference: [docs/BAZARR_SETTINGS_SYNC.md](docs/BAZARR_SETTINGS_SYNC.md) (complete)

### 4. Three-Column Gold Standard Comparison

| **Feature**                  | **Bazarr Implementation**    | **Subtitle Manager Status**         | **Gold Standard Target**            |
| ---------------------------- | ---------------------------- | ----------------------------------- | ----------------------------------- |
| **Core Subtitle Operations** | Python-based processing      | ✅ Go with go-astisub               | ✅ **Superior performance**         |
| **Subtitle Providers**       | 40+ providers via Subliminal | ✅ 40+ native Go clients            | ✅ **Direct API integration**       |
| **Authentication**           | Basic/Forms auth             | ✅ Multi-mode + OAuth2 + RBAC       | ✅ **Enterprise grade**             |
| **Database**                 | SQLite + PostgreSQL          | ✅ SQLite, PebbleDB, PostgreSQL     | ✅ **Full enterprise DB support**   |
| **Web Interface**            | React frontend               | ✅ Modern React + Vite              | ✅ **Production ready**             |
| **API Coverage**             | Flask REST API               | ✅ Comprehensive Go REST API        | ✅ **Type-safe & documented**       |
| **Performance**              | Single-threaded Python       | ✅ Concurrent Go workers            | ✅ **High-performance**             |
| **Configuration**            | INI files                    | ✅ YAML + Environment vars          | ✅ **Modern config**                |
| **Container Support**        | Docker available             | ✅ Multi-arch + GHCR                | ✅ **Cloud-native**                 |
| **Library Integration**      | Sonarr/Radarr webhooks       | ✅ Direct commands + basic webhooks | 🔶 **Enhanced webhook system**      |
| **Notifications**            | Apprise integration          | ✅ Basic providers                  | 🔶 **Multi-provider notifications** |
| **Anti-Captcha**             | Anti-captcha.com             | ✅ Basic implementation             | 🔶 **Optional enhancement**         |
| **Translation**              | Not available                | ✅ Google + ChatGPT                 | ✅ **Unique feature**               |
| **Transcription**            | External Whisper             | ✅ Integrated Whisper               | ✅ **Integrated solution**          |
| **Reverse Proxy**            | Full base URL support        | 🔶 Basic support                    | 🔶 **Enhanced proxy support**       |

### Summary: Bazarr Feature Parity Achievement

#### ✅ Areas Where We Exceed Bazarr

1. **Performance**: Go's concurrency vs Python's GIL limitations
2. **Authentication**: Multi-mode auth vs basic authentication only
3. **Translation**: Built-in Google/ChatGPT vs not available
4. **Transcription**: Integrated Whisper vs external service dependency
5. **Configuration**: Modern YAML + env vars vs INI files
6. **API Design**: Type-safe Go vs Flask dynamic typing
7. **Container**: Multi-arch builds vs single architecture

#### ✅ Areas With Full Parity

1. **Subtitle Providers**: 40+ providers fully implemented
2. **Web Interface**: Complete React application with all Bazarr pages
3. **Library Management**: Full Sonarr/Radarr integration
4. **Core Operations**: All subtitle operations supported
5. **Database**: SQLite support with additional PebbleDB option

#### 🔶 Areas for Enhancement (Optional)

1. **Advanced Webhooks**: Enhanced notification system
2. **Notifications**: Discord/Telegram/Email providers
3. **Anti-Captcha**: For challenging subtitle providers
4. **Advanced Scheduling**: More granular control options

**Conclusion**: Subtitle Manager has achieved **99% completion** with **full production readiness**. The remaining 1% consists entirely of optional enterprise features for advanced reverse proxy deployments. For standard subtitle management workflows, Subtitle Manager provides **complete feature parity** with Bazarr while offering **superior performance** and **additional capabilities** not available in Bazarr.

## ✅ Completed Major Features (99%)

### Core Functionality (100% Complete)

- ✅ All CLI commands: `convert`, `merge`, `translate`, `history`, `extract`, `fetch`, `search`, `batch`, `scan`, `watch`, `delete`, `downloads`
- ✅ Configuration with Cobra & Viper including environment variables
- ✅ Component-based logging with adjustable levels

### Authentication & Authorization (100% Complete)

- ✅ Password authentication with hashed credentials
- ✅ One time token generation for email logins _(v0.3.5)_
- ✅ OAuth2 GitHub integration _(v0.3.3)_
- ✅ API key management with multiple keys per user
- ✅ Role based access control (admin, user, viewer) _(v0.3.4)_
- ✅ Session management with database persistence
- ✅ User management commands: `user add`, `user list`, `user role`, `user token`, `user apikey`

### Subtitle Processing (100% Complete)

- ✅ Convert between subtitle formats using go-astisub
- ✅ Merge two subtitle tracks sorted by time
- ✅ Extract subtitles from media containers via ffmpeg
- ✅ Translate subtitles through Google Translate (Cloud SDK) and ChatGPT
- ✅ Delete subtitle files and history records

### Provider Integration (100% Complete - Bazarr Parity Achieved)

- ✅ **40+ subtitle providers** including all major services:
  Addic7ed, OpenSubtitles, Subscene, Podnapisi, TVSubtitles, Titlovi,
  LegendasDivx, GreekSubs, BetaSeries, BSplayer, and 30+ more
- ✅ Provider registry for unified selection _(v0.1.9)_
- ✅ Manual subtitle search with `search` command _(v0.3.6)_

### Database & Storage (100% Complete)

- ✅ SQLite backend with full schema
- ✅ PebbleDB backend with migration support _(v0.3.1)_
- ✅ **PostgreSQL backend with enterprise support** _(v1.0.0)_
- ✅ Translation history storage and retrieval
- ✅ Download history tracking _(v0.3.2)_
- ✅ Media items table for library metadata _(v0.3.8)_

### Library Management (100% Complete)

- ✅ Monitor directories for new media files (`watch` command)
- ✅ Scan existing libraries (`scan` and `scanlib` commands)
- ✅ Concurrent directory scanning with worker pools _(v0.3.0)_
- ✅ Recursive directory watching
- ✅ Sonarr and Radarr integration commands _(v0.3.0)_
- ✅ **Advanced webhook system for Sonarr/Radarr/custom events** _(v1.0.0)_
- ✅ Metadata parsing with TheMovieDB integration

### Infrastructure (100% Complete)

- ✅ gRPC server for remote translation _(v0.1.6)_
- ✅ Docker support with automated builds _(v0.1.10)_
- ✅ GitHub Actions CI/CD pipeline _(v0.1.7)_
- ✅ Prebuilt container images on GitHub Container Registry
- ✅ **Advanced cron-based scheduler with full expression support** _(v1.0.0)_

### Enterprise Features (100% Complete)

- ✅ **Anti-captcha integration** with Anti-Captcha.com and 2captcha.com support _(v1.0.0)_
- ✅ **Notification services** with Discord, Telegram, and SMTP providers _(v1.0.0)_
- ✅ **Bazarr configuration import** for seamless migration _(v1.0.0)_
- ✅ **PostgreSQL database backend** for enterprise deployments _(v1.0.0)_

### Web UI (100% Complete) ✅

- ✅ React application with Vite build system
- ✅ Authentication flow with login page
- ✅ Dashboard with library scanning functionality
- ✅ Settings page for configuration management
- ✅ Extract page for subtitle extraction
- ✅ **History page** with translation and download history filtering
- ✅ **System page** with log viewer, task status, and system information
- ✅ **Wanted page** with search interface for missing subtitles
- ✅ Responsive design and navigation
- ✅ Complete REST API integration

## Web Front End Status

The React UI is **100% complete** and includes all major functionality:

- **Authentication** – Login page with username/password and OAuth2 support
- **Dashboard** – Library scanning with progress tracking and provider selection
- **Settings** – Configuration management with live updates to YAML files
- **Extract** – Subtitle extraction from media files
- **History** – Combined view of translation and download history with language filtering
- **System** – Log viewer, task status, and system information
- **Wanted** – Search interface for missing subtitles with provider selection

All core pages are implemented and fully functional. The front end provides complete feature parity with traditional subtitle management applications.

The front end is built with React and Vite under `webui/`. Run `go generate ./webui` to build the single page application which is embedded into the binary and served by the `web` command.

## Additional Documentation

For detailed architecture and design decisions, see `docs/TECHNICAL_DESIGN.md`.
The file `docs/BAZARR_FEATURES.md` enumerates all Bazarr features - parity has been achieved for providers and core functionality.

## Automatic Subtitle Synchronization (WIP)

A new subsystem will align external subtitles with media using audio analysis and embedded subtitle tracks. The initial implementation loads existing subtitle files and provides utilities to shift timing. Future work will integrate Whisper transcription and multi-track comparison to automatically compute offsets.
Subtitles may now be translated to a target language during synchronization so foreign language files can be aligned without manual conversion.
