<!-- file: docs/tasks/01-gcommon-migration/config-migration-map.md -->
<!-- version: 1.0.0 -->
<!-- guid: a1b2c3d4-e5f6-7890-1234-567890abcdef -->

# Config Migration Mapping

This file maps local configpb types to their gcommon equivalents for the
migration process.

| Local Type                     | gcommon Type           | Import Path                                 |
| ------------------------------ | ---------------------- | ------------------------------------------- |
| configpb.LogLevel              | common.LogLevel        | github.com/jdfalk/gcommon/sdks/go/v1/common |
| configpb.SubtitleManagerConfig | Local Config + gcommon | Use local struct with gcommon types         |

## Migration Strategy

1. **LogLevel Migration**: Replace all `configpb.LogLevel` with
   `common.LogLevel`
2. **SubtitleManagerConfig**: Keep as local struct but use gcommon types for
   fields
3. **Opaque API**: Convert all direct field access to getter/setter methods

## Files to Update

1. `pkg/config/types.go` - Replace LogLevel enum
2. `pkg/gcommonlog/logrus_provider.go` - Update LogLevel usage
3. `pkg/types/types.go` - Update any config references
4. Generated protobuf files - Regenerate after proto updates

## Enum Value Mapping

| Local LogLevel       | gcommon LogLevel               |
| -------------------- | ------------------------------ |
| LogLevel_UNSPECIFIED | LogLevel_LOG_LEVEL_UNSPECIFIED |
| LogLevel_TRACE       | LogLevel_LOG_LEVEL_TRACE       |
| LogLevel_DEBUG       | LogLevel_LOG_LEVEL_DEBUG       |
| LogLevel_INFO        | LogLevel_LOG_LEVEL_INFO        |
| LogLevel_WARN        | LogLevel_LOG_LEVEL_WARN        |
| LogLevel_ERROR       | LogLevel_LOG_LEVEL_ERROR       |
| LogLevel_FATAL       | LogLevel_LOG_LEVEL_FATAL       |
