<!-- file: docs/testing/E2E_QUICK_REFERENCE.md -->
<!-- version: 1.0.0 -->
<!-- guid: e2e-quick-ref-abcd-1234-5678-9012 -->

# E2E Testing Quick Reference

## Essential Commands

```bash
# Start E2E environment (complete setup)
make end2end-tests

# Check if environment is running
make status-e2e

# View logs in real-time
make logs-e2e

# Stop the environment
make stop-e2e

# Clean up everything
make clean-e2e
```

## Quick Setup

1. **Configure API Key:**

   ```bash
   cp .env.local.template .env.local
   # Edit .env.local and add your OpenAI API key
   ```

2. **Start Testing:**

   ```bash
   make end2end-tests
   ```

3. **Access Application:**
   - URL: <http://localhost:8080>
   - Username: `test`
   - Password: `test123`

## Test Categories

### Core Features

- **Folder Scanning**: Detect media directories and files
- **Subtitle Detection**: Find and parse subtitle files
- **Format Support**: SRT, ASS, SUB formats
- **Language Detection**: English, Spanish, French, Japanese

### Advanced Features

- **Subtitle Combining**: Merge multiple subtitle tracks
- **Translation**: OpenAI-powered subtitle translation
- **Format Conversion**: Convert between subtitle formats
- **Synchronization**: Adjust subtitle timing

## Test Data Locations

```
testdir/
├── movies/          # 4 test movies with various naming
├── tv/              # 4 TV shows with episodes
└── anime/           # 4 anime titles with Japanese subtitles
```

## Common Test Scenarios

### 1. Basic Functionality

- [ ] Navigate to media categories
- [ ] View individual item pages
- [ ] Check subtitle file detection
- [ ] Verify metadata display

### 2. Subtitle Operations

- [ ] Download original subtitles
- [ ] Combine multiple subtitle tracks
- [ ] Translate subtitles (requires API key)
- [ ] Convert between formats

### 3. Edge Cases

- [ ] Files with special characters
- [ ] Missing metadata
- [ ] Malformed subtitle files
- [ ] Network timeouts

## Troubleshooting

### Environment Not Starting

```bash
# Check for port conflicts
lsof -i :8080

# Verify API key
grep OPENAI_API_KEY .env.local

# Check logs
tail -f /tmp/subtitle-manager-e2e.log
```

### API Key Issues

```bash
# Test API key manually
curl -H "Authorization: Bearer $OPENAI_API_KEY" \
     https://api.openai.com/v1/models
```

### Build Problems

```bash
# Clean rebuild
make clean-all && make build

# Check Go environment
go version && go env
```

## File Structure

```
docs/testing/
├── E2E_SETUP_GUIDE.md          # Complete setup instructions
├── E2E_QA_TESTING_GUIDE.md     # Comprehensive QA scenarios
├── E2E_QUICK_REFERENCE.md      # This quick reference
└── e2e-config.yaml             # Environment configuration

scripts/
└── setup-e2e-environment.sh    # E2E startup script

testdir/                         # Test data directory
├── movies/                      # Movie test files
├── tv/                          # TV show test files
└── anime/                       # Anime test files

.env.local.template              # Environment configuration template
.env.local                       # Your local environment (create this)
```

## Key Make Targets

| Target              | Purpose                         |
| ------------------- | ------------------------------- |
| `end2end-tests`     | Complete E2E setup and start    |
| `setup-e2e-env`     | Environment configuration only  |
| `setup-e2e-testdir` | Create test data structure      |
| `status-e2e`        | Check if environment is running |
| `logs-e2e`          | View real-time logs             |
| `stop-e2e`          | Stop the E2E server             |
| `clean-e2e`         | Stop and clean up               |

## Environment Variables

```bash
# Required for translation testing
OPENAI_API_KEY=sk-your-api-key-here

# Optional customizations
SUBTITLE_MANAGER_PORT=8080
LOG_LEVEL=info
DATABASE_URL=sqlite:./subtitle_manager.db
TRANSLATION_TIMEOUT=60s
```

## Testing Checklist

### Pre-Testing

- [ ] API key configured in `.env.local`
- [ ] Port 8080 available
- [ ] Go and Node.js installed
- [ ] Test data directory exists

### During Testing

- [ ] All media categories load
- [ ] Subtitle files are detected
- [ ] Translation feature works
- [ ] No console errors
- [ ] Responsive design works

### Post-Testing

- [ ] Stop E2E environment
- [ ] Review logs for errors
- [ ] Document any issues found
- [ ] Clean up test artifacts

## Support Resources

- **Setup Guide**: `docs/testing/E2E_SETUP_GUIDE.md`
- **QA Guide**: `docs/testing/E2E_QA_TESTING_GUIDE.md`
- **Log File**: `/tmp/subtitle-manager-e2e.log`
- **Config**: `e2e-config.yaml`

## Common URLs

- **Application**: <http://localhost:8080>
- **Health Check**: <http://localhost:8080/health>
- **API Status**: <http://localhost:8080/api/status>
- **Movies**: <http://localhost:8080/movies>
- **TV Shows**: <http://localhost:8080/tv>
- **Anime**: <http://localhost:8080/anime>
