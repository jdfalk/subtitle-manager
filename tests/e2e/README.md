# file: tests/e2e/README.md
# version: 1.0.0
# guid: b3c4d5e6-f7a8-b9c0-d1e2-f3a4b5c6d7e8

# E2E Testing for Subtitle Manager

This directory contains end-to-end tests using Selenium WebDriver to validate UI functionality and user workflows.

## 🎯 Purpose

Our E2E testing framework provides:

- **Automated UI Testing**: Verify that our UI changes actually work
- **Regression Prevention**: Catch when changes break existing functionality
- **Cross-browser Validation**: Test on Chrome and Firefox
- **Video Recording**: Capture test execution for debugging
- **Performance Monitoring**: Track page load times and responsiveness

## 🚀 Quick Start

### Prerequisites

1. **Python 3.8+** with pip
2. **Chrome or Firefox** browser installed
3. **Node.js** for frontend development server

### Installation

```bash
# Install Python dependencies
cd tests/e2e
pip install -r requirements.txt

# Install Chrome driver (handled automatically by webdriver-manager)
# Or install manually from https://chromedriver.chromium.org/
```

### Running Tests

#### Quick Smoke Test
```bash
# Start frontend server first
cd webui && npm run dev

# In another terminal, run smoke tests
cd tests/e2e
python test_simple.py
```

#### Full Test Suite
```bash
# Using the test runner (starts services automatically)
cd tests/e2e
python run_tests.py --start-services --smoke

# Or manually with frontend running:
python run_tests.py --frontend-only --critical
```

#### Specific Test Categories
```bash
# Settings navigation tests (our recent fix)
python run_tests.py --settings

# Critical functionality only
python run_tests.py --critical

# All smoke tests
python run_tests.py --smoke

# Headless mode (no browser window)
python run_tests.py --headless --smoke
```

## 📁 Test Structure

```
tests/e2e/
├── conftest.py              # Pytest configuration and fixtures
├── requirements.txt         # Python dependencies
├── pytest.ini             # Pytest settings
├── run_tests.py            # Test runner script
├── test_simple.py          # Simple validation script
├── test_smoke.py           # Smoke tests
├── test_settings_navigation.py  # Settings routing tests
├── test_ui_critical_functionality.py  # Core UI tests
├── reports/                # HTML test reports
├── screenshots/            # Test failure screenshots
├── recordings/             # Video recordings
└── logs/                   # Test execution logs
```

## 🧪 Test Categories

### Smoke Tests (`@pytest.mark.smoke`)
- Basic application loading
- Navigation menu presence
- No obvious errors
- **Purpose**: Quick validation that app is functional

### Critical Tests (`@pytest.mark.critical`)
- Settings navigation (our recent fix)
- Main user workflows
- Core functionality
- **Purpose**: Must-pass tests for essential features

### Regression Tests (`@pytest.mark.regression`)
- Known bug fixes
- Previously failing scenarios
- **Purpose**: Ensure fixed issues don't reoccur

### Slow Tests (`@pytest.mark.slow`)
- Performance testing
- Load time validation
- **Purpose**: Performance regression detection

## 🔧 Configuration

### Environment Variables

Set these in your shell or `.env` file:

```bash
# Application URLs
TEST_BASE_URL=http://localhost:5173     # Frontend URL
TEST_BACKEND_URL=http://localhost:8080  # Backend API URL

# Browser Configuration
TEST_BROWSER=chrome                     # chrome or firefox
TEST_HEADLESS=false                     # true for headless mode
TEST_WINDOW_SIZE=1920,1080             # Browser window size

# Test Behavior
TEST_TIMEOUT=30                         # Element wait timeout (seconds)
TEST_VIDEO_RECORDING=true               # Record test videos
TEST_SCREENSHOT_ON_FAILURE=true         # Screenshot failed tests
TEST_PARALLEL_WORKERS=2                 # Parallel test execution
```

## 📊 Test Reports

After running tests, check:

- **HTML Report**: `reports/test_report.html` - Detailed test results
- **Screenshots**: `screenshots/` - Failure debugging images
- **Videos**: `recordings/` - Full test execution videos
- **Logs**: `logs/` - Detailed execution logs

## 🎥 Video Recording

Tests automatically record video for debugging:

1. **Automatic Recording**: Each test session is recorded
2. **Failure Focus**: Failed tests get special attention
3. **Performance Review**: See exactly how the UI behaves

Videos are saved as MP4 files in `recordings/` directory.

## 🔍 Debugging Failed Tests

When a test fails:

1. **Check Screenshot**: Look at `screenshots/test_name_failure.png`
2. **Watch Video**: Review `recordings/test_name_timestamp.mp4`
3. **Read Logs**: Check `logs/test_name_date.log`
4. **Console Errors**: Check browser console output in logs

## 🚨 Common Issues

### Tests Can't Find Elements
- Elements may load slowly - increase timeout
- Check selectors in browser dev tools
- Verify CSS classes haven't changed

### Browser Doesn't Start
- Install Chrome/Firefox
- Check webdriver installation
- Try headless mode: `--headless`

### Frontend Not Running
- Start with: `cd webui && npm run dev`
- Check port availability (default 5173)
- Verify in browser first

### Backend API Errors
- Some tests need backend running
- Start with: `go run . web`
- Check backend logs for errors

## 🎯 Testing Our Recent Changes

### Settings Navigation Fix

Our recent fix changed `/settings` routing to show the tabbed interface instead of the 3-card overview. Test this with:

```bash
# Specific test for settings fix
python run_tests.py --settings

# Or run the simple validator
python test_simple.py
```

**What we're testing**:
- ✓ Clicking "Settings" goes to tabbed interface
- ✓ URL shows `/settings`
- ✓ Multiple tabs are visible (Providers, General, etc.)
- ✗ NOT the old 3-card overview page

## 🔄 Continuous Integration

To integrate with CI/CD:

```bash
# In CI pipeline
cd tests/e2e
pip install -r requirements.txt

# Start services and run tests
python run_tests.py --start-services --critical --headless

# Check exit code
if [ $? -eq 0 ]; then
    echo "Tests passed"
else
    echo "Tests failed"
    exit 1
fi
```

## 📈 Performance Testing

Monitor performance with:

```bash
# Run performance-focused tests
python run_tests.py --slow

# Check load times in logs
grep "load time" logs/*.log
```

## 🤝 Contributing

When adding new tests:

1. **Use Appropriate Markers**: `@pytest.mark.smoke`, `@pytest.mark.critical`
2. **Take Screenshots**: Use `TestUtils.take_screenshot()`
3. **Add to Documentation**: Update this README
4. **Test Locally First**: Verify tests pass before committing

## 📚 Additional Resources

- [Selenium Documentation](https://selenium-python.readthedocs.io/)
- [Pytest Documentation](https://docs.pytest.org/)
- [WebDriver Manager](https://github.com/SergeyPirogov/webdriver_manager)
- [Material-UI Testing](https://mui.com/guides/testing/)
