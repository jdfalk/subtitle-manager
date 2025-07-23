<!-- file: .github/prompts/ai-rebase-context.md -->
<!-- version: 1.0.0 -->
<!-- guid: 7f8e9d6c-5b4a-3c2d-1e0f-6a5b4c3d2e1f -->

# Repository Context for AI Rebase

## Project Overview

This is the **Subtitle Manager** repository, a comprehensive subtitle file
processing and management application built with Go and featuring a modern web
UI. The project focuses on:

- Subtitle file format conversion and processing
- Web-based subtitle management interface
- Media file scanning and organization
- Advanced subtitle synchronization and editing
- Cloud storage integration for subtitle libraries
- Whisper ASR integration for automatic transcription

## Coding Standards

- Use Go conventions: CamelCase for exported functions, camelCase for private
- Follow standard file header format with path, version, and GUID
- Use conventional commit message format: `type(scope): description`
- All Go code should include comprehensive documentation
- Web UI follows React/JavaScript best practices with Material-UI components
- Database interactions use proper error handling and transactions
- API endpoints follow RESTful conventions

## Key Files to Reference

### README.md

Contains project overview, installation instructions, and usage examples for the
subtitle management system.

### .github/instructions/general-coding.instructions.md

Standard coding guidelines including file headers, version management, and
documentation requirements.

### go.mod

Go module definition showing dependencies on subtitle processing libraries, web
frameworks, and database drivers.

### main.go

Application entry point with command-line interface, server initialization, and
core subtitle processing logic.

## Common Conflict Patterns

### Go Code Conflicts

When resolving Go conflicts:

- Preserve both import statements when they're different packages
- Combine struct field additions from both branches
- Merge handler function improvements while maintaining API compatibility
- Keep both test cases when they test different scenarios
- Preserve error handling improvements from both sides

### Web UI Conflicts

For React/JavaScript conflicts:

- Combine component prop additions from both branches
- Merge CSS styling improvements while maintaining visual consistency
- Preserve both event handler additions when they serve different purposes
- Combine Material-UI component upgrades and customizations
- Keep both API endpoint additions when they serve different features

### Configuration Conflicts

For configuration file conflicts:

- Merge environment variable additions from both branches
- Combine Docker configuration improvements
- Preserve both database migration scripts
- Merge API route definitions when they don't overlap
- Keep both feature flag additions

## Dependencies and Imports

- **Core Go libraries**: `net/http`, `encoding/json`, `database/sql`
- **Subtitle processing**: Custom subtitle format parsers and converters
- **Web framework**: Gorilla Mux for routing, middleware for CORS/auth
- **Database**: SQLite/PostgreSQL drivers for metadata storage
- **Frontend**: React, Material-UI, modern JavaScript/TypeScript
- **Testing**: Go testing package, testify for assertions
- **Cloud storage**: AWS S3, DigitalOcean Spaces integration

## Project Structure

- `main.go` - Application entry point and CLI interface
- `pkg/` - Core subtitle processing libraries and utilities
- `webui/` - React-based web interface for subtitle management
- `cmd/` - Command-line tools and utilities
- `docs/` - Documentation including API specs and user guides
- `test/` - Integration tests and test data
- `scripts/` - Build scripts and development tools
- `sdks/` - Client SDKs for different programming languages

## Subtitle-Specific Context

### File Format Support

The application handles multiple subtitle formats with specific parsing logic:

- **SRT (SubRip)**: Most common format with simple time codes
- **VTT (WebVTT)**: Web-compatible format with styling support
- **ASS/SSA**: Advanced SubStation Alpha with complex styling
- **SUB**: MicroDVD format with frame-based timing

### Processing Pipeline

Subtitle processing follows a standard pipeline:

1. **Detection**: Identify subtitle format from file content
2. **Parsing**: Extract timing, text, and styling information
3. **Validation**: Check for timing overlaps and formatting errors
4. **Conversion**: Transform between different subtitle formats
5. **Optimization**: Adjust timing and improve readability

When resolving conflicts in subtitle processing code, preserve both format
improvements and ensure the pipeline remains functional for all supported
formats.
