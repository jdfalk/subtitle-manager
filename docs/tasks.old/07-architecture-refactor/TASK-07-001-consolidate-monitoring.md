# TASK-07-001: Consolidate File Monitoring Services

<!-- file: docs/tasks/07-architecture-refactor/TASK-07-001-consolidate-monitoring.md -->
<!-- version: 1.0.0 -->
<!-- guid: a7b8c9d0-e1f2-3456-7890-abcdef123456 -->

## 🎯 Objective

Consolidate the three redundant file monitoring services (`autoscan`, `watch`,
`monitor`) into a single unified monitoring service with configurable behavior.

## 📋 Current Problem

The application currently has three overlapping monitoring services:

- `autoscan` - Periodic directory scanning
- `watch` - Real-time file system monitoring
- `monitor` - Sonarr/Radarr integration monitoring

This creates confusion, code duplication, and maintenance overhead.

## ✅ Acceptance Criteria

- [ ] Single `monitor` command that combines all three functionalities
- [ ] Default configuration enables all three monitoring modes
- [ ] Clear configuration options to enable/disable each mode independently
- [ ] Backwards compatibility maintained for existing configurations
- [ ] Performance optimized to avoid duplicate scanning

## 🔧 Implementation Steps

### Phase 1: Analysis and Design

1. Analyze current implementations of all three services
2. Identify common functionality and interfaces
3. Design unified monitoring service architecture
4. Create configuration schema for combined service

### Phase 2: Core Implementation

1. Create new unified monitoring service in `pkg/monitoring/unified.go`
2. Implement configuration-driven monitoring modes:
   - Periodic scanning (autoscan)
   - Real-time file watching (watch)
   - Media server integration (Sonarr/Radarr)
3. Add intelligent deduplication to prevent overlap
4. Implement proper resource cleanup and shutdown

### Phase 3: Command Line Interface

1. Update `cmd/monitor.go` to handle unified functionality
2. Add subcommands for specific modes if needed
3. Remove old `cmd/autoscan.go` and `cmd/watch.go`
4. Update help text and documentation

### Phase 4: Migration and Testing

1. Create migration guide for existing users
2. Add comprehensive tests for unified service
3. Test backwards compatibility
4. Performance testing to ensure no regression

## 📚 Required Configuration

```yaml
monitor:
  # Enable/disable specific monitoring modes
  periodic_scan:
    enabled: true
    interval: '1h'
    cron_schedule: '' # Optional cron override

  real_time_watch:
    enabled: true
    recursive: true
    debounce_interval: '5s'

  media_server:
    enabled: true
    sonarr:
      enabled: true
      sync_interval: '24h'
    radarr:
      enabled: true
      sync_interval: '24h'

  # Shared settings
  scan_workers: 4
  upgrade_existing: false
```

## 🧪 Testing Requirements

- [ ] Unit tests for unified monitoring service
- [ ] Integration tests for all three monitoring modes
- [ ] Performance tests comparing old vs new implementation
- [ ] Migration tests for existing configurations
- [ ] Memory leak tests for long-running monitors

## 📖 Documentation Updates

- [ ] Update README with new monitoring command structure
- [ ] Create migration guide from old commands
- [ ] Update API documentation
- [ ] Update Docker compose examples

## 🎯 Success Metrics

- Single monitoring command handles all use cases
- No increase in memory or CPU usage vs previous implementation
- All existing functionality preserved
- Configuration migration works for 100% of existing setups
- User documentation is clear and comprehensive

## 🔗 Dependencies

- Must complete after any pending monitoring service changes
- Should coordinate with TASK-07-002 (All-in-one mode)

## 📝 Notes

This task eliminates redundancy while preserving all functionality. The unified
service should be more efficient and easier to maintain while providing better
user experience.
