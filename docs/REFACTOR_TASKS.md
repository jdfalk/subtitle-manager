<!-- file: docs/REFACTOR_TASKS.md -->
<!-- version: 1.0.0 -->
<!-- guid: f47ac10b-58cc-4372-a567-0e02b2c3d479 -->

# Subtitle Manager - gcommon Refactoring Task Breakdown

This document provides a comprehensive, detailed breakdown of all tasks required
to refactor subtitle-manager to use gcommon v1.8.0 packages. Each task is
designed to be independent and actionable by AI agents.

## Task Categories Overview

- **Phase 1**: Documentation & Setup (Tasks 1-5)
- **Phase 2**: Core Type Migration (Tasks 6-15)
- **Phase 3**: Package Replacements (Tasks 16-25)
- **Phase 4**: UI/UX Improvements (Tasks 26-40)
- **Phase 5**: Testing & QA (Tasks 41-55)
- **Phase 6**: Issue Management (Tasks 56-65)

## Required Resources

### Instruction Files (Located in docs/instructions/)

- `general-coding.instructions.md` - General coding standards
- `commit-messages.md` - Commit message format
- `pull-request-descriptions.md` - PR description format
- `test-generation.md` - Testing guidelines

### API Documentation (Located in docs/gcommon-api/)

- `common.md` - Core types (User, Session, etc.)
- `config.md` - Configuration management
- `database.md` - Database operations
- `health.md` - Health monitoring
- `media.md` - Media handling
- `metrics.md` - Metrics and monitoring
- `organization.md` - Organization management
- `queue.md` - Queue operations
- `web.md` - Web services

### UI/UX Reference Materials

- Bazarr Screenshots: https://wiki.bazarr.media/
- Modern UI Examples: https://github.com/morpheus65535/bazarr
- Current UI Issues: Located in TODO.md "High Priority UI/UX Improvements"

---

# PHASE 1: DOCUMENTATION & SETUP

## Task 1: Audit Current Local Protobuf Usage

**Objective**: Create comprehensive inventory of all local protobuf packages
that need replacement.

**Required Reading**:

- `docs/instructions/general-coding.instructions.md`
- `docs/gcommon-api/README.md`

**Steps**:

1. Scan entire codebase for imports of:
   - `github.com/jdfalk/subtitle-manager/pkg/configpb`
   - `github.com/jdfalk/subtitle-manager/pkg/databasepb`
   - `github.com/jdfalk/subtitle-manager/pkg/gcommonauth`
2. Create detailed inventory with:
   - File paths
   - Import aliases used
   - Specific types referenced
   - Method calls made
3. Generate mapping from local types to gcommon equivalents
4. Document opaque API requirements (setter/getter patterns)

**Output**: `docs/MIGRATION_INVENTORY.md`

**Acceptance Criteria**:

- Complete list of all affected files
- Mapping table: local type → gcommon type
- List of all method signatures that need updating
- Identified files that can be deleted post-migration

---

## Task 2: Create Migration Plan with Dependencies

**Objective**: Generate detailed migration sequence with dependency analysis.

**Required Reading**:

- `docs/MIGRATION_INVENTORY.md` (from Task 1)
- `docs/gcommon-api/common.md`
- `docs/gcommon-api/config.md`
- `docs/gcommon-api/database.md`

**Steps**:

1. Analyze dependencies between packages
2. Create migration phases based on dependency order
3. Identify packages that can be migrated in parallel
4. Plan testing strategy for each phase
5. Document rollback procedures

**Output**: `docs/MIGRATION_PLAN.md`

**Acceptance Criteria**:

- Clear phase breakdown with rationale
- Dependency graph visualization
- Risk assessment for each phase
- Estimated time per task
- Testing checkpoints defined

---

## Task 3: Set Up gcommon v1.8.0 Dependencies

**Objective**: Properly configure go.mod with gcommon dependencies and verify
imports.

**Required Reading**:

- `docs/instructions/general-coding.instructions.md`
- `docs/gcommon-api/README.md`

**Steps**:

1. Update `go.mod` to include `github.com/jdfalk/gcommon v1.8.0`
2. Verify all gcommon packages are accessible
3. Create test imports to validate package availability
4. Document import patterns for team
5. Set up IDE support for gcommon packages

**Output**: Updated `go.mod`, `docs/GCOMMON_SETUP.md`

**Acceptance Criteria**:

- `go mod tidy` runs without errors
- All 9 gcommon packages import successfully
- Test file demonstrates opaque API usage
- Documentation includes IDE setup instructions

---

## Task 4: Create Code Generation Scripts

**Objective**: Build automation scripts for repetitive migration tasks.

**Required Reading**:

- `docs/instructions/general-coding.instructions.md`
- `docs/MIGRATION_INVENTORY.md`
- `docs/MIGRATION_PLAN.md`

**Steps**:

1. Create script to replace import statements
2. Build type conversion generators
3. Generate setter/getter method replacements
4. Create validation scripts for migration accuracy
5. Add rollback automation

**Output**: `scripts/migration/` directory with automation tools

**Acceptance Criteria**:

- Import replacement script with dry-run mode
- Type conversion generator with templates
- Validation script that checks opaque API usage
- Comprehensive error handling and logging

---

## Task 5: Establish Testing Infrastructure

**Objective**: Set up comprehensive testing for migration validation.

**Required Reading**:

- `docs/instructions/test-generation.md`
- `docs/instructions/general-coding.instructions.md`

**Steps**:

1. Create migration test suite structure
2. Set up integration test environment
3. Build test data fixtures for all phases
4. Create performance benchmarks
5. Set up automated test execution

**Output**: `tests/migration/` directory structure

**Acceptance Criteria**:

- Test suite covers all migration phases
- Integration tests validate gcommon integration
- Performance benchmarks established
- CI/CD integration ready

---

# PHASE 2: CORE TYPE MIGRATION

## Task 6: Migrate User Type from gcommonauth to gcommon/common

**Objective**: Replace all local User types with gcommon User type using opaque
API.

**Required Reading**:

- `docs/gcommon-api/common.md` (User section)
- `docs/instructions/general-coding.instructions.md`
- `pkg/database/store.go` (current state)

**Files to Modify**:

- `pkg/database/store.go`
- `pkg/database/pebble.go`
- `pkg/database/database.go`
- `pkg/database/postgres.go`
- `pkg/webserver/auth.go`
- `pkg/webserver/users.go`
- `pkg/authserver/server.go`
- All test files using User type

**Steps**:

1. Replace `github.com/jdfalk/subtitle-manager/pkg/gcommonauth` imports with
   `github.com/jdfalk/gcommon/sdks/go/v1/common`
2. Update all User struct usages to use opaque API:
   - `user.Id` → `user.GetId()`
   - `user.Username` → `user.GetUsername()`
   - Direct assignment → `user.SetId("value")`
3. Fix return type mismatches: `[]common.User` → `[]*common.User`
4. Update method signatures to use `*common.User`
5. Implement missing authentication methods in SQLStore and PostgresStore

**Key Changes**:

```go
// OLD
import auth "github.com/jdfalk/subtitle-manager/pkg/gcommonauth"
user := auth.User{Id: "123", Username: "john"}

// NEW
import "github.com/jdfalk/gcommon/sdks/go/v1/common"
user := &common.User{}
user.SetId("123")
user.SetUsername("john")
```

**Acceptance Criteria**:

- All imports updated to gcommon/common
- All User field access uses opaque API (Set*/Get* methods)
- Return types corrected ([]\*common.User)
- All authentication methods implemented
- Tests pass with new User type
- No direct field access remains

---

## Task 7: Implement Missing Authentication Methods

**Objective**: Complete authentication method implementations in SQLStore and
PostgresStore.

**Required Reading**:

- `docs/gcommon-api/common.md` (User and Session sections)
- `pkg/database/pebble.go` (reference implementation)
- `docs/instructions/general-coding.instructions.md`

**Files to Modify**:

- `pkg/database/database.go` (SQLStore)
- `pkg/database/postgres.go` (PostgresStore)

**Missing Methods to Implement**:

1. `CreateOneTimeToken(ctx context.Context, userID string, token string, expiresAt time.Time) error`
2. `CleanupExpiredSessions(ctx context.Context) error`
3. `GetUserByID(ctx context.Context, id string) (*common.User, error)`
4. `GetUserByEmail(ctx context.Context, email string) (*common.User, error)`
5. `GetUserByUsername(ctx context.Context, username string) (*common.User, error)`
6. `ListUsers(ctx context.Context) ([]*common.User, error)`

**Steps**:

1. Study PebbleStore implementation as reference
2. Implement SQL queries for SQLStore
3. Implement PostgreSQL queries for PostgresStore
4. Use gcommon User type with opaque API
5. Add proper error handling and context support
6. Write unit tests for each method

**Acceptance Criteria**:

- All missing methods implemented for both backends
- SQL queries properly parameterized
- Context cancellation handled
- Unit tests cover all methods
- Error messages consistent with existing code

---

## Task 8: Create Session Type Migration

**Objective**: Migrate session management to use gcommon Session types.

**Required Reading**:

- `docs/gcommon-api/common.md` (Session section)
- `docs/instructions/general-coding.instructions.md`

**Files to Analyze and Modify**:

- `pkg/gcommonauth/session.go`
- `pkg/webserver/auth.go`
- `pkg/database/store.go`

**Steps**:

1. Map current session fields to gcommon Session type
2. Replace session creation/validation logic
3. Update database storage to use gcommon Session
4. Migrate session middleware
5. Update session cleanup logic

**Acceptance Criteria**:

- All session operations use gcommon Session type
- Session serialization/deserialization works
- Session validation updated
- Database storage compatible

---

## Task 9: Migrate Authentication Flow

**Objective**: Update complete authentication flow to use gcommon types and
patterns.

**Required Reading**:

- `docs/gcommon-api/common.md`
- `pkg/webserver/oauth.go`
- `pkg/authserver/server.go`

**Files to Modify**:

- `pkg/webserver/oauth.go`
- `pkg/webserver/auth.go`
- `pkg/authserver/server.go`
- All authentication test files

**Steps**:

1. Update OAuth2 flow to create gcommon User types
2. Migrate login/logout handlers
3. Update middleware to work with gcommon types
4. Fix JWT token generation/validation
5. Update API key management

**Acceptance Criteria**:

- OAuth2 login creates proper gcommon User
- Session management uses gcommon Session
- All auth middleware updated
- API key system working
- Integration tests pass

---

## Task 10: Database Schema Compatibility

**Objective**: Ensure database schemas work with gcommon types and add any
missing fields.

**Required Reading**:

- `docs/gcommon-api/database.md`
- Current schema files in `pkg/database/`

**Files to Modify**:

- `pkg/database/migrations/`
- `pkg/database/schema.go`
- Database initialization files

**Steps**:

1. Compare current schema with gcommon requirements
2. Create migration scripts for schema updates
3. Add any missing fields for gcommon compatibility
4. Test migration on all database backends
5. Create rollback migrations

**Acceptance Criteria**:

- Schema supports all gcommon User fields
- Migration scripts tested on SQLite, PebbleDB, PostgreSQL
- Rollback procedures verified
- No data loss during migration

---

# PHASE 3: PACKAGE REPLACEMENTS

## Task 11: Replace configpb with gcommon/config

**Objective**: Completely replace local configpb package with gcommon config
types.

**Required Reading**:

- `docs/gcommon-api/config.md`
- `docs/MIGRATION_INVENTORY.md`
- `pkg/configpb/config.pb.go` (current implementation)

**Files to Modify**:

- `pkg/gcommon/config/config.go`
- `pkg/translatorpb/translator.pb.go`
- `pkg/subtitle/translator/v1/translator.pb.go`
- All files importing configpb

**Steps**:

1. Map SubtitleManagerConfig to appropriate gcommon config types
2. Replace all configpb imports with gcommon/config
3. Update configuration loading/saving logic
4. Migrate all config field access to opaque API
5. Update protobuf generation if needed
6. Remove pkg/configpb directory

**Key Mappings**:

```go
// OLD
import configpb "github.com/jdfalk/subtitle-manager/pkg/configpb"
config := configpb.SubtitleManagerConfig{...}

// NEW
import "github.com/jdfalk/gcommon/sdks/go/v1/config"
config := &config.ApplicationConfig{}
config.SetName("subtitle-manager")
```

**Acceptance Criteria**:

- All configpb imports removed
- Configuration loading uses gcommon types
- All config access uses opaque API (Set*/Get* methods)
- Tests updated and passing
- pkg/configpb directory deleted

---

## Task 12: Replace databasepb with gcommon/database

**Objective**: Replace local databasepb types with gcommon database types.

**Required Reading**:

- `docs/gcommon-api/database.md`
- `pkg/databasepb/databasepb.go` (current implementation)
- `pkg/database/pb_conversions.go`

**Files to Modify**:

- `pkg/database/pb_conversions.go`
- All files using SubtitleRecord, DownloadRecord types
- Database storage implementations

**Steps**:

1. Map SubtitleRecord to gcommon database types
2. Map DownloadRecord to gcommon database types
3. Update all database operations
4. Migrate serialization/deserialization
5. Update query builders
6. Remove pkg/databasepb directory

**Acceptance Criteria**:

- All databasepb types replaced
- Database operations use gcommon types
- Serialization works correctly
- Historical data compatibility maintained
- pkg/databasepb directory deleted

---

## Task 13: Health Monitoring Integration

**Objective**: Integrate gcommon health monitoring throughout the application.

**Required Reading**:

- `docs/gcommon-api/health.md`
- Current health checking implementations

**Files to Create/Modify**:

- `pkg/health/` (new package)
- `pkg/webserver/health.go`
- `cmd/health.go` (new command)

**Steps**:

1. Set up gcommon health monitoring
2. Create health check endpoints
3. Add database health checks
4. Add external service health checks
5. Integrate with existing monitoring

**Acceptance Criteria**:

- Health endpoints return gcommon HealthStatus
- Database connectivity monitored
- External services monitored
- Health metrics exposed

---

## Task 14: Media Package Integration

**Objective**: Use gcommon media types for video and subtitle file handling.

**Required Reading**:

- `docs/gcommon-api/media.md`
- Current media handling code

**Files to Analyze and Modify**:

- `pkg/media/`
- Video file processing code
- Subtitle file handling code

**Steps**:

1. Map current media types to gcommon media types
2. Update file processing logic
3. Migrate metadata extraction
4. Update media database storage
5. Test with various media formats

**Acceptance Criteria**:

- Media files processed using gcommon types
- Metadata extraction working
- Storage uses gcommon media types
- All media formats supported

---

## Task 15: Metrics and Monitoring Migration

**Objective**: Replace current metrics with gcommon metrics system.

**Required Reading**:

- `docs/gcommon-api/metrics.md`
- Current metrics implementation

**Files to Modify**:

- `pkg/metrics/`
- All files collecting metrics
- Monitoring dashboards

**Steps**:

1. Map current metrics to gcommon metrics types
2. Update metric collection points
3. Migrate dashboard configurations
4. Update alerting rules
5. Test metric accuracy

**Acceptance Criteria**:

- All metrics use gcommon types
- Dashboard displays correct data
- Alerting works properly
- Performance impact minimal

---

# PHASE 4: UI/UX IMPROVEMENTS

## Task 16: Fix Navigation and Layout Issues

**Objective**: Implement working navigation system with proper back button and
sidebar management.

**Required Reading**:

- `TODO.md` - "High Priority UI/UX Improvements" section
- Bazarr UI Reference: https://wiki.bazarr.media/
- `docs/instructions/general-coding.instructions.md`

**Files to Modify**:

- `webui/src/components/Navigation.jsx`
- `webui/src/components/Sidebar.jsx`
- `webui/src/App.jsx`
- `webui/src/Router.jsx`

**Requirements from TODO.md**:

1. **Fix user management display**: System/users shows blank usernames
2. **Move user management to settings**: Users interface should be part of
   settings
3. **Implement working back button**: Navigation history and proper back button
   functionality
4. **Add sidebar pinning**: Allow users to pin/unpin the sidebar
5. **Reorganize navigation order**: Dashboard → Media Library → Wanted → History
   → Settings → System

**UI Reference**: Study Bazarr's navigation at https://wiki.bazarr.media/ for
layout patterns

**Steps**:

1. Fix user display issues in system users page
2. Move user management component to settings section
3. Implement React Router history management for back button
4. Add sidebar pin/unpin state management
5. Reorder navigation menu items as specified
6. Add breadcrumb navigation
7. Implement responsive design improvements

**Acceptance Criteria**:

- User management displays usernames correctly
- User management accessible from Settings
- Back button works throughout application
- Sidebar can be pinned/unpinned and remembers state
- Navigation follows specified order
- Responsive design works on mobile

---

## Task 17: Settings Page Complete Redesign

**Objective**: Redesign settings page to match Bazarr's comprehensive settings
interface.

**Required Reading**:

- `TODO.md` - "Settings Page Enhancements" section
- Bazarr Settings Reference:
  https://wiki.bazarr.media/Additional-Configuration/Settings/
- `docs/instructions/general-coding.instructions.md`

**Files to Create/Modify**:

- `webui/src/pages/Settings/`
- `webui/src/components/Settings/`
- New settings components for each section

**Requirements from TODO.md**:

1. **Enhance General Settings**: Add Bazarr-compatible settings (Host:
   Address/Port/Base URL, Proxy, Updates, Logging with filters, Backups,
   Analytics)
2. **Improve Database Settings**: Add comprehensive database information and
   management options
3. **Redesign Authentication Page**: Card-based UI for each auth method with
   enable/disable toggles
4. **Add OAuth2 management**: Generate/regenerate client ID/secret, reset to
   defaults
5. **Enhance Notifications**: Card-based interface for each notification method
   with test buttons
6. **Create Languages Page**: Global language settings for subtitle downloads
7. **Add Scheduler Settings**: Integration into general settings

**Reference Implementation**: Use Bazarr's settings as gold standard -
https://wiki.bazarr.media/Additional-Configuration/Settings/

**Steps**:

1. Create settings page structure with navigation tabs
2. Implement General Settings with Bazarr-compatible options
3. Build Database Settings with management options
4. Create card-based Authentication settings
5. Add OAuth2 management interface
6. Build Notifications with test functionality
7. Create Languages configuration page
8. Add Scheduler settings integration

**Acceptance Criteria**:

- Settings organized in clear sections like Bazarr
- Card-based UI for authentication methods
- Database management tools working
- OAuth2 generation/regeneration working
- Notification test buttons functional
- Language settings comprehensive
- Scheduler integration complete

---

## Task 18: Provider Configuration Improvements

**Objective**: Fix provider configuration interface and implement proper
provider management.

**Required Reading**:

- `TODO.md` - "Provider System Improvements" section
- Current provider implementation in `pkg/providers/`
- Bazarr Provider Reference:
  https://wiki.bazarr.media/Additional-Configuration/Settings/#providers

**Files to Modify**:

- `webui/src/components/Providers/`
- `webui/src/pages/Settings/Providers.jsx`
- `pkg/webserver/providers.go`

**Requirements from TODO.md**:

1. **Fix provider configuration modals**: Proper provider selection dropdowns
   and configuration options
2. **Improve embedded provider config**: Working dropdown and proper
   configuration display
3. **Implement global language settings**: Move language settings from
   provider-level to global
4. **Add language profiles**: Bazarr-style language profiles for different
   content types

**Steps**:

1. Fix provider selection dropdown functionality
2. Implement working provider configuration modals
3. Add embedded provider configuration interface
4. Create global language settings page
5. Implement language profiles system
6. Add provider priority management
7. Create provider testing interface

**Acceptance Criteria**:

- Provider dropdowns work correctly
- Configuration modals save settings properly
- Embedded provider config functional
- Global language settings implemented
- Language profiles working
- Provider testing interface operational

---

## Task 19: Dashboard and Media Library Enhancement

**Objective**: Implement comprehensive dashboard and improve media library
interface.

**Required Reading**:

- Bazarr Dashboard Reference: https://wiki.bazarr.media/
- Current dashboard implementation
- `docs/instructions/general-coding.instructions.md`

**Files to Modify**:

- `webui/src/pages/Dashboard.jsx`
- `webui/src/pages/MediaLibrary.jsx`
- `webui/src/components/Stats/`
- `webui/src/components/MediaLibrary/`

**Requirements**:

1. **Statistics Dashboard**: Show download stats, provider performance, library
   status
2. **Recent Activity Feed**: Latest downloads, errors, system events
3. **System Health**: Database status, provider availability, disk space
4. **Media Library**: Improved file browser, metadata display, bulk operations

**Steps**:

1. Create statistics collection system
2. Build activity feed component
3. Add system health monitoring display
4. Enhance media library with better file management
5. Add bulk operations interface
6. Implement search and filtering

**Acceptance Criteria**:

- Dashboard shows comprehensive statistics
- Activity feed displays recent events
- System health monitoring working
- Media library has improved navigation
- Bulk operations functional
- Search and filtering working

---

## Task 20: History and Wanted Pages

**Objective**: Implement comprehensive history tracking and wanted subtitle
management.

**Required Reading**:

- Current history implementation
- Bazarr History Reference: https://wiki.bazarr.media/
- `docs/instructions/general-coding.instructions.md`

**Files to Modify**:

- `webui/src/pages/History.jsx`
- `webui/src/pages/Wanted.jsx`
- `webui/src/components/History/`
- `webui/src/components/Wanted/`

**Steps**:

1. Enhance history page with detailed tracking
2. Add filtering and search to history
3. Implement wanted subtitles tracking
4. Add retry functionality for failed downloads
5. Create bulk operations for wanted items
6. Add export/import functionality

**Acceptance Criteria**:

- History shows detailed download information
- Filtering and search working
- Wanted subtitles properly tracked
- Retry functionality working
- Bulk operations available
- Export/import functional

---

# PHASE 5: TESTING & QA

## Task 21: Unit Test Suite for gcommon Integration

**Objective**: Create comprehensive unit tests for all gcommon package
integrations.

**Required Reading**:

- `docs/instructions/test-generation.md`
- `docs/gcommon-api/` (all packages)
- `docs/instructions/general-coding.instructions.md`

**Files to Create**:

- `tests/gcommon/common_test.go`
- `tests/gcommon/config_test.go`
- `tests/gcommon/database_test.go`
- `tests/gcommon/health_test.go`
- `tests/gcommon/media_test.go`
- `tests/gcommon/metrics_test.go`
- `tests/gcommon/organization_test.go`
- `tests/gcommon/queue_test.go`
- `tests/gcommon/web_test.go`

**Test Categories**:

1. **Opaque API Tests**: Verify all setter/getter methods work correctly
2. **Type Conversion Tests**: Test conversion between old and new types
3. **Database Integration Tests**: Verify gcommon types work with all database
   backends
4. **Serialization Tests**: Test protobuf serialization/deserialization
5. **Error Handling Tests**: Verify proper error handling with gcommon types

**Steps**:

1. Create test structure following test-generation.md guidelines
2. Implement Arrange-Act-Assert pattern for all tests
3. Test opaque API usage patterns
4. Verify database compatibility
5. Test serialization edge cases
6. Add performance benchmarks
7. Create integration test scenarios

**Acceptance Criteria**:

- 95%+ code coverage for gcommon integration
- All opaque API methods tested
- Database backends tested with gcommon types
- Serialization/deserialization verified
- Performance benchmarks established
- Tests follow established patterns from test-generation.md

---

## Task 22: Integration Test Suite

**Objective**: Create comprehensive integration tests for complete workflows
using gcommon.

**Required Reading**:

- `docs/instructions/test-generation.md`
- `docs/MIGRATION_PLAN.md`
- Current integration test structure

**Files to Create**:

- `tests/integration/auth_flow_test.go`
- `tests/integration/subtitle_workflow_test.go`
- `tests/integration/provider_integration_test.go`
- `tests/integration/database_migration_test.go`

**Test Scenarios**:

1. **Complete Authentication Flow**: Login → Session → API calls → Logout
2. **Subtitle Processing Workflow**: Upload → Process → Store → Retrieve
3. **Provider Integration**: Configure → Search → Download → Validate
4. **Database Migration**: Old data → Migration → New format → Validation

**Steps**:

1. Set up integration test environment
2. Create test data fixtures using gcommon types
3. Implement complete workflow tests
4. Add database migration testing
5. Test cross-backend compatibility
6. Add performance and stress testing

**Acceptance Criteria**:

- All major workflows tested end-to-end
- Database migration tested thoroughly
- Cross-backend compatibility verified
- Performance requirements met
- Tests can run in CI/CD pipeline

---

## Task 23: Selenium Web UI Test Suite

**Objective**: Create comprehensive Selenium-based tests for all UI workflows
with video recording.

**Required Reading**:

- `docs/instructions/test-generation.md`
- `docs/instructions/general-coding.instructions.md`
- UI workflow requirements from TODO.md

**Tools Required**:

- Selenium WebDriver
- Video recording capability (consider using Selenoid or Selenium Grid with
  video)
- Browser automation (Chrome, Firefox)

**Files to Create**:

- `tests/selenium/`
- `tests/selenium/auth_test.py`
- `tests/selenium/navigation_test.py`
- `tests/selenium/settings_test.py`
- `tests/selenium/provider_test.py`
- `tests/selenium/media_management_test.py`
- `tests/selenium/dashboard_test.py`

**Test Scenarios with Video Recording**:

1. **Authentication Workflow**:
   - Login with password
   - OAuth2 login flow
   - Session management
   - Logout process

2. **Navigation Testing**:
   - Sidebar navigation
   - Breadcrumb navigation
   - Back button functionality
   - Mobile responsive navigation

3. **Settings Management**:
   - General settings configuration
   - Provider setup and testing
   - Authentication method configuration
   - Database settings management

4. **Provider Configuration**:
   - Add new provider
   - Configure provider settings
   - Test provider connection
   - Enable/disable providers

5. **Media Management**:
   - File upload process
   - Library scanning
   - Subtitle search
   - Download management

6. **Dashboard Interaction**:
   - Statistics display
   - Activity monitoring
   - System health checks
   - Quick actions

**Video Recording Requirements**:

- Record full screen interactions
- Capture audio for debugging
- Generate video artifacts for each test
- Store videos with test results
- Provide playback capability

**Steps**:

1. Set up Selenium testing environment with video recording
2. Create base test classes with video recording setup
3. Implement authentication workflow tests
4. Create navigation and UI interaction tests
5. Build settings configuration tests
6. Add provider management tests
7. Implement media workflow tests
8. Create dashboard interaction tests
9. Set up video artifact storage and playback
10. Integrate with CI/CD for automated execution

**Technical Implementation**:

```python
# Example Selenium test with video recording
import pytest
from selenium import webdriver
from selenium.webdriver.chrome.options import Options
from selenium_video import VideoRecorder

class TestAuthentication:
    def setup_method(self):
        chrome_options = Options()
        chrome_options.add_argument("--enable-video-recording")
        self.driver = webdriver.Chrome(options=chrome_options)
        self.video_recorder = VideoRecorder(self.driver)
        self.video_recorder.start_recording("auth_test")

    def teardown_method(self):
        self.video_recorder.stop_recording()
        self.driver.quit()

    def test_login_workflow(self):
        # Test implementation with video recording
        pass
```

**Acceptance Criteria**:

- All major UI workflows have Selenium tests
- Video recording works for each test
- Videos stored as CI artifacts
- Tests run reliably in headless mode
- Cross-browser compatibility tested
- Mobile responsive tests included
- Test execution time < 30 minutes total
- Video artifacts accessible for debugging

---

## Task 24: Performance and Load Testing

**Objective**: Ensure gcommon integration doesn't impact performance and system
can handle expected load.

**Required Reading**:

- `docs/instructions/test-generation.md`
- `docs/gcommon-api/` (performance considerations)

**Files to Create**:

- `tests/performance/`
- `tests/performance/database_benchmark_test.go`
- `tests/performance/api_load_test.go`
- `tests/performance/memory_usage_test.go`
- `tests/performance/concurrency_test.go`

**Test Categories**:

1. **Database Performance**: Compare old vs new type performance
2. **API Response Times**: Measure response times with gcommon types
3. **Memory Usage**: Monitor memory consumption with new types
4. **Concurrency**: Test concurrent operations with gcommon types
5. **Provider Performance**: Measure provider operation speeds

**Steps**:

1. Establish performance baselines
2. Create database benchmarks
3. Implement API load testing
4. Monitor memory usage patterns
5. Test concurrent operations
6. Measure provider performance
7. Generate performance reports

**Acceptance Criteria**:

- Performance within 5% of baseline
- Memory usage not increased significantly
- Concurrent operations stable
- Load testing passes requirements
- Performance reports generated

---

## Task 25: End-to-End Workflow Testing

**Objective**: Test complete user workflows from start to finish using gcommon
types.

**Required Reading**:

- `docs/instructions/test-generation.md`
- Complete user workflow documentation

**Files to Create**:

- `tests/e2e/`
- `tests/e2e/complete_setup_test.go`
- `tests/e2e/subtitle_processing_test.go`
- `tests/e2e/user_management_test.go`
- `tests/e2e/backup_restore_test.go`

**Workflow Tests**:

1. **Complete Setup**: Fresh install → Configuration → First use
2. **Subtitle Processing**: Search → Download → Process → Store
3. **User Management**: Create → Manage → Delete users
4. **Backup/Restore**: Backup → Restore → Verify data integrity

**Steps**:

1. Create realistic test scenarios
2. Implement complete workflow tests
3. Add data validation at each step
4. Test error recovery scenarios
5. Verify data consistency
6. Test backup/restore procedures

**Acceptance Criteria**:

- All workflows complete successfully
- Data integrity maintained
- Error recovery works
- Backup/restore functional
- Tests run reliably

---

# PHASE 6: ISSUE MANAGEMENT

## Task 26: GitHub Issues Audit and Categorization

**Objective**: Review all GitHub issues, categorize them, and identify which
ones are resolved by gcommon refactoring.

**Required Reading**:

- `docs/instructions/general-coding.instructions.md`
- `docs/instructions/commit-messages.md`
- Current GitHub issues

**Steps**:

1. Export all open GitHub issues
2. Categorize issues by type (bug, feature, enhancement)
3. Identify issues resolved by gcommon migration
4. Create mapping of issues to specific tasks
5. Prioritize remaining open issues
6. Create resolution plan for each category

**Output**: `docs/ISSUE_AUDIT.md`

**Acceptance Criteria**:

- Complete issue inventory
- Clear categorization
- Resolution mapping created
- Priority assessment complete

---

## Task 27: Close Resolved Authentication Issues

**Objective**: Close GitHub issues that are resolved by the gcommon
authentication migration.

**Required Reading**:

- `docs/instructions/commit-messages.md` - for proper closing documentation
- GitHub issues related to authentication
- Task 6-9 results (authentication migration)

**Issues to Review**:

- Search for issues with tags: authentication, auth, login, user, session
- Look for OAuth2 related issues
- Find API key management issues
- Identify RBAC or permission issues

**Steps**:

1. Identify all authentication-related issues
2. Test each issue against new gcommon implementation
3. Create detailed resolution documentation
4. Close issues with proper PR references
5. Update issue labels and milestones

**Closing Template**:

```markdown
## Issue Resolution

This issue has been resolved by the gcommon authentication migration in PR #XXX.

### What was fixed:

- [Detailed explanation of resolution]

### Changes made:

- Migrated to gcommon v1.8.0 User types with opaque API
- Updated authentication flow to use gcommon Session management
- Implemented proper RBAC with gcommon types

### Testing completed:

- Unit tests pass for authentication flow
- Integration tests verify OAuth2 functionality
- Selenium tests confirm UI authentication works

### Verification:

To verify this fix:

1. [Step-by-step verification instructions]

Closes via commit: [commit hash] Related PR: #XXX
```

**Acceptance Criteria**:

- All resolved auth issues identified
- Each issue closed with detailed resolution
- PR references included
- Verification steps provided
- Labels updated appropriately

---

## Task 28: Close Resolved Configuration Issues

**Objective**: Close GitHub issues resolved by configuration system migration to
gcommon.

**Required Reading**:

- Task 11 results (configpb replacement)
- Issues related to configuration management
- `docs/instructions/commit-messages.md`

**Steps**:

1. Find all configuration-related issues
2. Verify resolution with gcommon config types
3. Document how gcommon resolved each issue
4. Close with detailed explanations
5. Update documentation references

**Acceptance Criteria**:

- Configuration issues identified and verified
- Detailed resolution documentation
- Proper issue closing with PR references

---

## Task 29: Close Resolved Database Issues

**Objective**: Close GitHub issues resolved by database migration to gcommon
types.

**Required Reading**:

- Task 12 results (databasepb replacement)
- Database-related GitHub issues
- Database backend compatibility issues

**Steps**:

1. Identify database-related issues
2. Test issues against new gcommon database types
3. Verify cross-backend compatibility fixes
4. Document performance improvements
5. Close with comprehensive resolution details

**Acceptance Criteria**:

- Database issues properly resolved
- Cross-backend compatibility verified
- Performance improvements documented

---

## Task 30: Close Resolved UI/UX Issues

**Objective**: Close GitHub issues resolved by UI/UX improvements.

**Required Reading**:

- Tasks 16-20 results (UI improvements)
- UI/UX related GitHub issues
- Selenium test results

**Steps**:

1. Map UI issues to specific improvements
2. Verify fixes with Selenium test recordings
3. Document user experience improvements
4. Include video evidence where possible
5. Close with before/after comparisons

**Acceptance Criteria**:

- UI issues mapped to specific fixes
- Video evidence of improvements
- User experience improvements documented

---

## Task 31: Create New Issues for Outstanding Items

**Objective**: Create GitHub issues for any remaining work identified during
refactoring.

**Required Reading**:

- `docs/instructions/general-coding.instructions.md`
- Findings from all migration tasks
- `TODO.md` outstanding items

**Steps**:

1. Review findings from all tasks
2. Identify new issues discovered
3. Create well-documented GitHub issues
4. Assign appropriate labels and milestones
5. Link to related work and PRs

**Issue Template**:

```markdown
## Description

[Clear description of the issue]

## Context

Discovered during gcommon migration in [task/PR reference]

## Expected Behavior

[What should happen]

## Current Behavior

[What actually happens]

## Steps to Reproduce

1. [Step 1]
2. [Step 2]
3. [Step 3]

## Additional Context

- Related to gcommon migration: [yes/no]
- Affects: [components affected]
- Priority: [high/medium/low]

## Definition of Done

- [ ] [Specific completion criteria]
- [ ] Tests written and passing
- [ ] Documentation updated
```

**Acceptance Criteria**:

- All new issues properly documented
- Clear acceptance criteria defined
- Appropriate labels and milestones assigned
- Links to related work included

---

# TASK EXECUTION GUIDELINES

## General Requirements for All Tasks

### Before Starting Any Task:

1. **Read ALL specified documentation** in "Required Reading" section
2. **Review relevant gcommon API documentation** in docs/gcommon-api/
3. **Follow coding standards** from
   docs/instructions/general-coding.instructions.md
4. **Check current state** of files before making changes

### During Task Execution:

1. **Follow opaque API patterns** - use Set*/Get* methods, never direct field
   access
2. **Maintain backward compatibility** where possible
3. **Add comprehensive error handling** and logging
4. **Write tests** following test-generation.md guidelines
5. **Update documentation** as needed

### Commit and PR Requirements:

1. **Follow commit message format** from docs/instructions/commit-messages.md
2. **Create PR descriptions** following
   docs/instructions/pull-request-descriptions.md
3. **Reference related tasks** and issues
4. **Include test results** and verification steps

### Testing Requirements:

1. **Run all existing tests** and ensure they pass
2. **Add new tests** for all new functionality
3. **Verify cross-platform compatibility** (SQLite, PebbleDB, PostgreSQL)
4. **Test UI changes** in multiple browsers
5. **Include performance verification**

### Quality Assurance:

1. **Code review** against gcommon API usage
2. **Integration testing** with real data
3. **User experience validation** for UI changes
4. **Performance impact assessment**
5. **Security review** for authentication changes

## Task Dependencies

### Sequential Dependencies:

- Tasks 1-5 must be completed before Phase 2
- Tasks 6-10 must be completed before Phase 3
- Tasks 11-15 must be completed before Phase 4
- Tasks 21-25 should run parallel with implementation
- Tasks 26-31 can be done after major phases

### Parallel Execution Possible:

- Tasks within same phase can often be done in parallel
- Testing tasks (21-25) can be developed alongside implementation
- Issue management (26-31) can be done independently

## Success Criteria

### Technical Success:

- All tests passing with gcommon types
- Performance within acceptable limits
- No regressions in functionality
- Clean code following standards

### User Experience Success:

- UI improvements match requirements
- Navigation works intuitively
- All workflows complete successfully
- Video tests show proper operation

### Project Management Success:

- All GitHub issues properly categorized
- Resolved issues closed with documentation
- New issues created for remaining work
- Clear progress tracking

## Emergency Procedures

### If Task Blocks Progress:

1. Document the blocking issue
2. Create GitHub issue with full context
3. Implement workaround if possible
4. Escalate to project maintainer
5. Continue with non-dependent tasks

### If Tests Fail:

1. Isolate the failing component
2. Check gcommon API usage patterns
3. Verify opaque API implementation
4. Test with minimal reproduction case
5. Document findings and seek help

### If Performance Degrades:

1. Run performance benchmarks
2. Profile the problematic code
3. Compare with baseline implementation
4. Optimize gcommon usage patterns
5. Consider caching or batching

This comprehensive task breakdown provides independent, actionable work items
that can be executed by AI agents with minimal risk of errors or conflicts.
