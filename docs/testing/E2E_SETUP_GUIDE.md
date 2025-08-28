<!-- file: docs/testing/E2E_SETUP_GUIDE.md -->
<!-- version: 1.0.0 -->
<!-- guid: e2e-setup-guide-12345-6789-abcd-ef01 -->

# E2E Testing Environment Setup Guide

## Quick Start

### 1. Configure API Key (Required for Translation Testing)

**Option A: Using .env.local file (Recommended)**

```bash
# Copy the template
cp .env.local.template .env.local

# Edit .env.local and add your OpenAI API key
# Replace 'your_openai_api_key_here' with your actual API key
```

**Option B: Using environment variable**

```bash
export OPENAI_API_KEY="your_actual_api_key_here"
```

### 2. Start E2E Environment

```bash
# This command will:
# - Check for API key configuration
# - Build the application if needed
# - Create test directory structure
# - Start the web server
# - Run basic health checks
make end2end-tests
```

### 3. Access the Application

- **URL**: http://localhost:8080
- **Username**: `test`
- **Password**: `test123`

## Detailed Setup Instructions

### Prerequisites

1. **Go Development Environment**
   - Go 1.21 or later
   - CGO enabled for SQLite support

2. **Node.js Environment** (for web UI)
   - Node.js 18+
   - npm package manager

3. **OpenAI API Key** (for translation features)
   - Sign up at https://platform.openai.com/
   - Generate an API key
   - Add billing information (required for API usage)

### Environment Configuration

#### Step 1: API Key Setup

The E2E testing environment requires an OpenAI API key for translation
functionality testing. You have several options:

**Method 1: .env.local File (Recommended)**

```bash
# Copy the template
cp .env.local.template .env.local

# Edit the file
nano .env.local  # or use your preferred editor

# Add your API key
OPENAI_API_KEY=sk-your-actual-api-key-here
```

**Method 2: Environment Variable**

```bash
# Add to your shell profile (.bashrc, .zshrc, etc.)
export OPENAI_API_KEY="sk-your-actual-api-key-here"

# Or set for current session
export OPENAI_API_KEY="sk-your-actual-api-key-here"
```

**Method 3: Runtime Configuration**

```bash
# Pass directly when running make command
OPENAI_API_KEY="sk-your-actual-api-key-here" make end2end-tests
```

#### Step 2: Verify Prerequisites

```bash
# Check Go installation
go version

# Check Node.js installation
node --version
npm --version

# Check if required tools are available
which curl jq
```

#### Step 3: Build and Test Setup

```bash
# Setup test directory structure
make setup-e2e-testdir

# Build the application
make build

# Verify binary works
make test-binary
```

### Starting the E2E Environment

#### Complete E2E Setup

```bash
# Run complete setup and start environment
make end2end-tests
```

This command performs the following actions:

1. **Environment Check**: Verifies API key configuration
2. **Dependencies**: Ensures all prerequisites are met
3. **Build**: Compiles the application if needed
4. **Test Data**: Creates/verifies test directory structure
5. **Server Start**: Launches the subtitle manager web server
6. **Health Check**: Runs basic connectivity tests
7. **Instructions**: Displays access information

#### Manual Step-by-Step Setup

```bash
# 1. Setup environment configuration
make setup-e2e-env

# 2. Create test directory structure
make setup-e2e-testdir

# 3. Build the application
make build

# 4. Start the environment manually
./scripts/setup-e2e-environment.sh
```

### Environment Management

#### Check Status

```bash
# Check if E2E environment is running
make status-e2e
```

#### View Logs

```bash
# Show real-time logs
make logs-e2e

# Or directly view log file
tail -f /tmp/subtitle-manager-e2e.log
```

#### Stop Environment

```bash
# Stop the E2E server
make stop-e2e
```

#### Clean Up

```bash
# Stop server and clean up artifacts
make clean-e2e
```

### Test Data Structure

The E2E environment creates a comprehensive test dataset:

```
testdir/
├── movies/
│   ├── The Matrix (1999)/
│   │   ├── The.Matrix.1999.en.srt
│   │   ├── The.Matrix.1999.es.srt
│   │   └── The.Matrix.1999.en.ass
│   ├── Blade Runner 2049 (2017)/
│   │   └── Blade.Runner.2049.2017.en.sub
│   ├── Interstellar (2014)/
│   │   ├── Interstellar.2014.en.srt
│   │   └── Interstellar.2014.fr.srt
│   └── The_Dark_Knight/
├── tv/
│   ├── Breaking Bad (2008)/
│   │   └── Breaking.Bad.S01E01.en.srt
│   ├── The Office (2005)/
│   ├── Game of Thrones (2011)/
│   │   └── GoT.S01E01.Winter.Is.Coming.en.srt
│   └── stranger_things_2016/
└── anime/
    ├── Attack on Titan (2013)/
    │   ├── Attack.on.Titan.S01E01.ja.srt
    │   └── Attack.on.Titan.S01E01.en.srt
    ├── Your Name (2016)/
    ├── Spirited Away (2001)/
    │   └── Spirited.Away.2001.ja.srt
    └── one_piece-1999/
```

#### Test Data Features

- **Naming Conventions**: Various formats with/without spaces, parentheses,
  underscores
- **Special Characters**: International characters, punctuation
- **File Formats**: SRT, ASS, SUB subtitle formats
- **Languages**: English, Spanish, French, Japanese
- **Categories**: Movies, TV Shows, Anime
- **Edge Cases**: Missing years, different naming patterns

### Troubleshooting

#### Common Issues and Solutions

**Problem: API Key Not Working**

```bash
# Check API key configuration
grep OPENAI_API_KEY .env.local

# Test API key manually
curl -H "Authorization: Bearer $OPENAI_API_KEY" \
     -H "Content-Type: application/json" \
     https://api.openai.com/v1/models
```

**Problem: Port 8080 Already in Use**

```bash
# Find what's using the port
lsof -i :8080

# Kill existing processes
pkill -f subtitle-manager

# Or use a different port
SUBTITLE_MANAGER_PORT=8081 make end2end-tests
```

**Problem: Build Failures**

```bash
# Clean and rebuild everything
make clean-all
make build

# Check Go environment
go env
```

**Problem: Web UI Not Loading**

```bash
# Rebuild web UI
make webui-rebuild

# Check web UI dependencies
cd webui && npm ci
```

**Problem: Permission Errors**

```bash
# Fix script permissions
chmod +x scripts/setup-e2e-environment.sh

# Check directory permissions
ls -la testdir/
```

### Advanced Configuration

#### Custom Configuration

Create additional configuration in `.env.local`:

```bash
# Custom port
SUBTITLE_MANAGER_PORT=8081

# Custom database
DATABASE_URL=sqlite:./e2e_test.db

# Debug logging
LOG_LEVEL=debug

# Custom timeouts
TRANSLATION_TIMEOUT=60s
MAX_CONCURRENT_TRANSLATIONS=3
```

#### Using Docker for E2E Testing

```bash
# Build Docker image
make docker

# Run E2E tests in Docker
docker run -p 8080:8080 -v $(pwd)/testdir:/app/testdir \
           -e OPENAI_API_KEY="$OPENAI_API_KEY" \
           subtitle-manager:latest
```

#### CI/CD Integration

```bash
# For automated testing environments
export OPENAI_API_KEY="$CI_OPENAI_API_KEY"
export SUBTITLE_MANAGER_PORT=8080
export LOG_LEVEL=info

# Run headless E2E tests
make end2end-tests

# Check status programmatically
make status-e2e && echo "E2E environment ready"
```

### Security Considerations

1. **API Key Protection**
   - Never commit `.env.local` to version control
   - Use environment variables in CI/CD
   - Rotate API keys regularly

2. **Test Environment Isolation**
   - Use separate API keys for testing
   - Isolate test data from production
   - Clean up after testing

3. **Network Security**
   - E2E server binds to localhost only
   - No external network access required
   - Test credentials are for local use only

### Performance Testing

#### Load Testing Setup

```bash
# Start E2E environment
make end2end-tests

# Run concurrent requests
for i in {1..10}; do
  curl -s http://localhost:8080/health &
done
wait

# Monitor resource usage
top -p $(cat /tmp/subtitle-manager-e2e.pid)
```

#### Memory Testing

```bash
# Monitor memory usage
watch -n 1 'ps -p $(cat /tmp/subtitle-manager-e2e.pid) -o pid,vsz,rss,pmem,comm'

# Test with large files
# (Create larger test subtitle files if needed)
```

### Next Steps

After setting up the E2E environment:

1. **Read the QA Testing Guide**: `docs/testing/E2E_QA_TESTING_GUIDE.md`
2. **Run Test Scenarios**: Follow the step-by-step testing instructions
3. **Report Issues**: Document any problems found during testing
4. **Contribute**: Add new test cases or improve existing ones

### Support

If you encounter issues with the E2E setup:

1. Check the troubleshooting section above
2. Review log files: `/tmp/subtitle-manager-e2e.log`
3. Verify prerequisites are properly installed
4. Ensure API key is valid and has sufficient credits
5. Check GitHub issues for known problems

### Contributing to E2E Testing

To improve the E2E testing environment:

1. Add new test scenarios to the QA guide
2. Create additional test data files
3. Enhance the setup scripts
4. Document edge cases and solutions
5. Improve error handling and user experience
