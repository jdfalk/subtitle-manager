<!-- file: docs/testing/E2E_QA_TESTING_GUIDE.md -->
<!-- version: 1.0.0 -->
<!-- guid: qa-guide-12345-6789-abcd-ef01-234567890abc -->

# End-to-End QA Testing Guide for Subtitle Manager

## Overview

This guide provides step-by-step instructions for human QA testers to validate
the Subtitle Manager application using the e2e testing environment. The testing
environment includes sample media files with various naming conventions, special
characters, and subtitle formats to thoroughly test the application.

## Prerequisites

### System Requirements

- Access to the Subtitle Manager development environment
- Web browser (Chrome, Firefox, Safari, or Edge)
- Terminal access for running commands
- OpenAI API key (for translation features)

### Environment Setup

Before starting QA testing, ensure the e2e environment is properly configured:

1. **Check for API Key Configuration**
   - The system requires an OpenAI API key for translation features
   - Create `.env.local` file (not version controlled) with your API key
   - This will be automatically loaded during e2e setup

2. **Start the E2E Environment**

   ```bash
   make end2end-tests
   ```

3. **Verify System Access**
   - Web interface: http://localhost:8080
   - Default credentials: username=`test`, password=`test123`

## Test Data Structure

The e2e environment includes the following test data structure:

```
testdir/
├── movies/
│   ├── The Matrix (1999)/           # Spaces and parentheses
│   ├── Blade Runner 2049 (2017)/   # Spaces and parentheses
│   ├── Interstellar (2014)/        # Multiple subtitle languages
│   └── The_Dark_Knight/             # Underscores, no year
├── tv/
│   ├── Breaking Bad (2008)/         # Standard format
│   ├── The Office (2005)/           # Spaces and parentheses
│   ├── Game of Thrones (2011)/     # Multiple words, spaces
│   └── stranger_things_2016/        # Underscores, no parentheses
└── anime/
    ├── Attack on Titan (2013)/     # Standard format
    ├── Your Name (2016)/            # Spaces and parentheses
    ├── Spirited Away (2001)/        # Studio Ghibli test
    └── one_piece-1999/              # Underscores and dashes
```

### Subtitle Format Coverage

- **SRT**: SubRip (.srt) - Most common format
- **ASS**: Advanced SubStation Alpha (.ass) - Styled subtitles
- **SUB**: MicroDVD (.sub) - Frame-based timing
- **Multiple Languages**: English, Spanish, French, Japanese

## Testing Scenarios

### 1. Environment Startup and Access

#### Test Case 1.1: Application Launch

**Objective**: Verify the application starts correctly and is accessible

**Steps**:

1. Open terminal in the subtitle-manager project directory
2. Run: `make end2end-tests`
3. Wait for the "E2E Testing Environment Ready" message
4. Note the provided URL and credentials

**Expected Results**:

- Command completes without errors
- Server starts on http://localhost:8080
- Log files are created in `/tmp/subtitle-manager-e2e.log`
- Test credentials are displayed

**Validation**:

- [ ] Server starts successfully
- [ ] No error messages in startup log
- [ ] Health check endpoint responds
- [ ] Access credentials are displayed

#### Test Case 1.2: Web Interface Access

**Objective**: Verify web interface login and dashboard access

**Steps**:

1. Open web browser
2. Navigate to http://localhost:8080
3. Enter credentials: username=`test`, password=`test123`
4. Verify dashboard loads

**Expected Results**:

- Login page displays correctly
- Authentication succeeds with test credentials
- Dashboard/main interface loads
- Navigation elements are visible

**Validation**:

- [ ] Login page renders properly
- [ ] Test credentials work
- [ ] Dashboard loads without errors
- [ ] All UI elements are functional

### 2. Media Library Setup and Scanning

#### Test Case 2.1: Add Test Data Directory

**Objective**: Configure the application to scan the test data directory

**Steps**:

1. In the web interface, navigate to Settings or Library Management
2. Add library path: `./testdir`
3. Configure scanning options (if available)
4. Initiate library scan

**Expected Results**:

- Library path is accepted
- Scanning process starts
- Progress indication is shown
- Scan completes successfully

**Validation**:

- [ ] Path input accepts `./testdir`
- [ ] Scan process initiates
- [ ] Progress is visible
- [ ] Scan completes without errors

#### Test Case 2.2: Verify Media Detection

**Objective**: Confirm all test media items are detected with various naming
conventions

**Steps**:

1. Navigate to Movies section
2. Verify all movie entries are detected:
   - "The Matrix (1999)" - spaces and parentheses
   - "Blade Runner 2049 (2017)" - spaces and parentheses
   - "Interstellar (2014)" - multiple languages
   - "The_Dark_Knight" - underscores, no year
3. Navigate to TV Shows section
4. Verify all TV show entries:
   - "Breaking Bad (2008)" - standard format
   - "The Office (2005)" - spaces and parentheses
   - "Game of Thrones (2011)" - multiple words
   - "stranger_things_2016" - underscores format
5. Navigate to Anime section
6. Verify all anime entries:
   - "Attack on Titan (2013)" - standard format
   - "Your Name (2016)" - spaces and parentheses
   - "Spirited Away (2001)" - Studio Ghibli
   - "one_piece-1999" - underscores and dashes

**Expected Results**:

- All 12 media items are detected
- Names are parsed correctly despite different formats
- Categories (movies/tv/anime) are properly assigned
- Years are extracted where present

**Validation**:

- [ ] All movies detected (4 items)
- [ ] All TV shows detected (4 items)
- [ ] All anime detected (4 items)
- [ ] Special characters handled correctly
- [ ] Years parsed correctly

#### Test Case 2.3: Subtitle File Detection

**Objective**: Verify subtitle files are detected and associated correctly

**Steps**:

1. For each media item, click to view details
2. Verify subtitle files are detected:
   - Check for .srt files (SubRip)
   - Check for .ass files (Advanced SubStation Alpha)
   - Check for .sub files (MicroDVD)
3. Verify language detection:
   - English (en)
   - Spanish (es)
   - French (fr)
   - Japanese (ja)

**Expected Results**:

- All subtitle files are detected
- File formats are recognized
- Languages are identified correctly
- Files are associated with correct media items

**Validation**:

- [ ] SRT files detected
- [ ] ASS files detected
- [ ] SUB files detected
- [ ] Languages identified correctly
- [ ] File associations are correct

### 3. Individual Item Page Testing

#### Test Case 3.1: Movie Item Page Navigation

**Objective**: Test navigation to individual movie pages and subtitle management

**Steps**:

1. From the Movies section, click on "The Matrix (1999)"
2. Verify the item detail page loads
3. Check available subtitle files:
   - The.Matrix.1999.en.srt (English)
   - The.Matrix.1999.es.srt (Spanish)
   - The.Matrix.1999.en.ass (Advanced SubStation Alpha)
4. Test subtitle preview functionality
5. Check download options

**Expected Results**:

- Item page loads correctly
- Title displays with correct formatting
- All subtitle files are listed
- File details (format, language, size) are shown
- Preview and download functions work

**Validation**:

- [ ] Page loads without errors
- [ ] Title formatting is correct
- [ ] All subtitle files listed
- [ ] File metadata is accurate
- [ ] Preview functionality works
- [ ] Download links function

#### Test Case 3.2: TV Show Episode Management

**Objective**: Test TV show episode navigation and subtitle handling

**Steps**:

1. Navigate to "Breaking Bad (2008)"
2. Verify episode detection and listing
3. Click on episode "S01E01"
4. Check subtitle availability for the episode
5. Test episode-specific subtitle management

**Expected Results**:

- TV show page shows season/episode structure
- Episodes are listed with proper formatting
- Episode detail pages load correctly
- Episode-specific subtitles are managed separately

**Validation**:

- [ ] Season/episode structure detected
- [ ] Episode naming is correct
- [ ] Episode pages load properly
- [ ] Subtitle management per episode

#### Test Case 3.3: Anime with Japanese Content

**Objective**: Test handling of Japanese content and character encoding

**Steps**:

1. Navigate to "Attack on Titan (2013)"
2. Open the item detail page
3. Check Japanese subtitle file (Attack.on.Titan.S01E01.ja.srt)
4. Verify Japanese character display
5. Test preview of Japanese subtitles

**Expected Results**:

- Japanese characters display correctly
- No encoding issues
- Subtitle content renders properly
- Language detection works for Japanese

**Validation**:

- [ ] Japanese characters render correctly
- [ ] No encoding errors
- [ ] Subtitle preview works
- [ ] Language detection accurate

### 4. Subtitle Combining and Merging

#### Test Case 4.1: Combine Two Subtitle Files

**Objective**: Test the ability to combine multiple subtitle files into one

**Steps**:

1. Navigate to "Interstellar (2014)"
2. Select both available subtitle files:
   - Interstellar.2014.en.srt (English)
   - Interstellar.2014.fr.srt (French)
3. Use the "Combine Subtitles" feature
4. Configure combination settings:
   - Choose layout (side-by-side, top-bottom)
   - Set language priorities
   - Adjust timing if needed
5. Generate combined subtitle file
6. Preview the result
7. Download the combined file

**Expected Results**:

- Both files are selectable for combination
- Combination interface is intuitive
- Layout options are available
- Combined file contains both languages
- Timing synchronization is maintained
- Download produces valid subtitle file

**Validation**:

- [ ] File selection works
- [ ] Combination options available
- [ ] Layout choices function
- [ ] Combined output is correct
- [ ] Timing is synchronized
- [ ] Download works properly

#### Test Case 4.2: Merge Overlapping Subtitles

**Objective**: Test merging of subtitles with overlapping time codes

**Steps**:

1. Select subtitle files with timing conflicts
2. Use merge functionality
3. Configure merge behavior:
   - Priority settings for overlaps
   - Gap handling
   - Duplicate removal
4. Process the merge
5. Validate the output

**Expected Results**:

- Merge options handle overlaps intelligently
- User has control over conflict resolution
- Output maintains readability
- No timing corruption occurs

**Validation**:

- [ ] Overlap detection works
- [ ] Merge options available
- [ ] Conflict resolution functions
- [ ] Output quality is good

### 5. Translation Testing

#### Test Case 5.1: Translate English to Spanish

**Objective**: Test automatic translation functionality using OpenAI API

**Prerequisites**: Ensure OpenAI API key is configured in `.env.local`

**Steps**:

1. Select an English subtitle file (e.g., "The Matrix" English SRT)
2. Navigate to Translation options
3. Select target language: Spanish
4. Configure translation settings:
   - Translation model preferences
   - Subtitle formatting retention
   - Timing preservation
5. Initiate translation process
6. Monitor progress
7. Review translated output
8. Compare timing with original
9. Download translated file

**Expected Results**:

- Translation interface is accessible
- Language options are available
- API integration works correctly
- Translation preserves timing
- Output quality is reasonable
- Subtitle formatting is maintained

**Validation**:

- [ ] Translation UI loads
- [ ] Language selection works
- [ ] API connection successful
- [ ] Translation completes
- [ ] Timing preserved
- [ ] Quality is acceptable
- [ ] Download functions

#### Test Case 5.2: Translate Japanese to English

**Objective**: Test translation from Japanese to English for anime content

**Steps**:

1. Select Japanese subtitle file from anime content
2. Initiate translation to English
3. Verify Japanese character handling
4. Check translation quality for anime-specific terms
5. Compare with existing English subtitle if available

**Expected Results**:

- Japanese characters are processed correctly
- Translation handles anime terminology
- Cultural references are translated appropriately
- Technical terms are preserved or explained

**Validation**:

- [ ] Japanese input processed
- [ ] Anime terms handled well
- [ ] Cultural context preserved
- [ ] Technical accuracy maintained

#### Test Case 5.3: Batch Translation

**Objective**: Test translation of multiple subtitle files in batch

**Steps**:

1. Select multiple subtitle files for batch translation
2. Configure batch settings:
   - Target language
   - Processing order
   - Error handling
3. Start batch process
4. Monitor progress for multiple files
5. Review results for consistency

**Expected Results**:

- Batch selection interface works
- Progress tracking is clear
- Error handling is robust
- Results are consistent across files

**Validation**:

- [ ] Batch selection works
- [ ] Progress tracking clear
- [ ] Error handling robust
- [ ] Consistent results

### 6. Subtitle Synchronization and Timing

#### Test Case 6.1: Adjust Subtitle Timing

**Objective**: Test manual timing adjustment capabilities

**Steps**:

1. Select a subtitle file with timing issues
2. Open timing adjustment tools
3. Test adjustment options:
   - Global offset (shift all subtitles)
   - Speed adjustment (stretch/compress timing)
   - Individual subtitle timing
4. Preview changes in real-time
5. Apply adjustments
6. Verify synchronization

**Expected Results**:

- Timing tools are intuitive and responsive
- Real-time preview works correctly
- Adjustments are applied accurately
- Synchronization improves

**Validation**:

- [ ] Timing tools accessible
- [ ] Real-time preview works
- [ ] Adjustments apply correctly
- [ ] Synchronization improved

#### Test Case 6.2: Automatic Synchronization

**Objective**: Test automatic subtitle synchronization features

**Steps**:

1. Use subtitles with known timing offsets
2. Apply automatic synchronization
3. Configure sync parameters
4. Process the synchronization
5. Compare before and after timing

**Expected Results**:

- Auto-sync detects timing issues
- Synchronization improves accuracy
- User retains control over process

**Validation**:

- [ ] Auto-sync detection works
- [ ] Accuracy improves
- [ ] User control maintained

### 7. File Format Conversion

#### Test Case 7.1: Convert SRT to ASS

**Objective**: Test conversion between subtitle formats

**Steps**:

1. Select an SRT file
2. Choose conversion to ASS format
3. Configure conversion options:
   - Style settings
   - Font preferences
   - Color options
4. Process conversion
5. Verify ASS output
6. Test advanced features in ASS format

**Expected Results**:

- Conversion interface is clear
- Format options are available
- Conversion produces valid ASS file
- Advanced features are accessible

**Validation**:

- [ ] Conversion UI clear
- [ ] Format options available
- [ ] Valid output produced
- [ ] Advanced features work

#### Test Case 7.2: Convert ASS to SRT

**Objective**: Test conversion from advanced format to simple format

**Steps**:

1. Select an ASS file with styling
2. Convert to SRT format
3. Verify how styling is handled
4. Check text preservation
5. Validate timing preservation

**Expected Results**:

- Conversion strips styling appropriately
- Text content is preserved
- Timing remains accurate

**Validation**:

- [ ] Styling handled correctly
- [ ] Text preserved
- [ ] Timing accurate

### 8. Error Handling and Edge Cases

#### Test Case 8.1: Malformed Subtitle Files

**Objective**: Test handling of corrupted or malformed subtitle files

**Steps**:

1. Attempt to process files with known issues:
   - Invalid time codes
   - Missing timestamps
   - Encoding problems
   - Truncated files
2. Verify error handling
3. Check recovery options

**Expected Results**:

- Errors are detected and reported clearly
- Recovery options are provided where possible
- System remains stable

**Validation**:

- [ ] Errors detected properly
- [ ] Clear error messages
- [ ] Recovery options available
- [ ] System stability maintained

#### Test Case 8.2: Large File Handling

**Objective**: Test performance with large subtitle files

**Steps**:

1. Create or use large subtitle files (>1000 entries)
2. Test processing performance
3. Check memory usage
4. Verify UI responsiveness

**Expected Results**:

- Large files are handled efficiently
- Performance remains acceptable
- Memory usage is reasonable

**Validation**:

- [ ] Large files process efficiently
- [ ] Performance acceptable
- [ ] Memory usage reasonable

#### Test Case 8.3: Concurrent Operations

**Objective**: Test system behavior with multiple simultaneous operations

**Steps**:

1. Start multiple operations simultaneously:
   - Translation in progress
   - File conversion
   - Timing adjustment
2. Monitor system performance
3. Verify operation completion
4. Check for conflicts or errors

**Expected Results**:

- System handles concurrent operations
- No data corruption occurs
- Operations complete successfully

**Validation**:

- [ ] Concurrent operations handled
- [ ] No data corruption
- [ ] All operations complete

### 9. API and Integration Testing

#### Test Case 9.1: API Endpoint Validation

**Objective**: Test REST API endpoints if available

**Steps**:

1. Test health check endpoint: `GET /health`
2. Test subtitle listing: `GET /api/v1/subtitles`
3. Test file upload: `POST /api/v1/subtitles/upload`
4. Test translation: `POST /api/v1/subtitles/translate`
5. Verify API responses and error codes

**Expected Results**:

- API endpoints respond correctly
- JSON responses are well-formed
- Error codes are appropriate
- Authentication works if required

**Validation**:

- [ ] Health check responds
- [ ] Subtitle API works
- [ ] Upload functionality
- [ ] Translation API
- [ ] Proper error codes

#### Test Case 9.2: External Service Integration

**Objective**: Test integration with external services (OpenAI API)

**Steps**:

1. Verify OpenAI API key configuration
2. Test API connectivity
3. Validate translation service response
4. Check error handling for API failures

**Expected Results**:

- API key is properly configured
- External services are accessible
- Error handling is robust

**Validation**:

- [ ] API key configured
- [ ] External services accessible
- [ ] Error handling robust

### 10. Performance and Load Testing

#### Test Case 10.1: Response Time Testing

**Objective**: Measure system response times for key operations

**Steps**:

1. Measure page load times
2. Time subtitle processing operations
3. Monitor translation response times
4. Document performance baselines

**Expected Results**:

- Page loads are under 3 seconds
- Processing operations complete reasonably
- Translation times are acceptable

**Validation**:

- [ ] Page load times acceptable
- [ ] Processing times reasonable
- [ ] Translation performance good

#### Test Case 10.2: Memory and Resource Usage

**Objective**: Monitor system resource consumption

**Steps**:

1. Monitor memory usage during operations
2. Check CPU utilization
3. Observe disk usage patterns
4. Test with multiple concurrent users

**Expected Results**:

- Memory usage remains stable
- CPU utilization is reasonable
- Disk usage is efficient

**Validation**:

- [ ] Memory usage stable
- [ ] CPU utilization reasonable
- [ ] Disk usage efficient

## Environment Management Commands

### Starting the Environment

```bash
# Start complete e2e environment
make end2end-tests

# Check if environment is running
curl http://localhost:8080/health
```

### Monitoring and Logs

```bash
# View application logs
tail -f /tmp/subtitle-manager-e2e.log

# Check process status
ps aux | grep subtitle-manager
```

### Stopping the Environment

```bash
# Stop the e2e environment
make stop-e2e

# Clean up test artifacts
make clean-e2e
```

### Environment Variables

The following environment variables are automatically set during e2e testing:

- `SUBTITLE_MANAGER_PORT=8080`
- `SUBTITLE_MANAGER_USERNAME=test`
- `SUBTITLE_MANAGER_PASSWORD=test123`
- `SUBTITLE_MANAGER_MEDIA_PATH=./testdir`
- `OPENAI_API_KEY` (from .env.local if present)

## Troubleshooting Common Issues

### API Key Issues

- **Problem**: Translation not working
- **Solution**: Verify `.env.local` contains valid `OPENAI_API_KEY`
- **Check**: Test API connectivity manually

### Port Conflicts

- **Problem**: Server won't start on port 8080
- **Solution**: Kill existing processes or change port
- **Check**: `lsof -i :8080`

### File Permission Issues

- **Problem**: Cannot access test files
- **Solution**: Check file permissions in testdir/
- **Check**: `ls -la testdir/`

### Memory Issues

- **Problem**: Application crashes with large files
- **Solution**: Increase available memory or reduce file size
- **Check**: Monitor memory usage

## Test Results Documentation

### Test Report Template

For each test session, document:

1. **Environment Details**
   - Date and time of testing
   - Version information
   - Configuration used

2. **Test Results Summary**
   - Total test cases executed
   - Passed/Failed/Skipped counts
   - Critical issues identified

3. **Issue Details**
   - Bug descriptions
   - Steps to reproduce
   - Expected vs actual behavior
   - Severity and impact

4. **Performance Metrics**
   - Response times
   - Resource usage
   - Throughput measurements

5. **Recommendations**
   - Priority fixes needed
   - Feature improvements
   - Performance optimizations

### Quality Assurance Checklist

Before marking QA complete, verify:

- [ ] All test scenarios executed
- [ ] Critical functionality works
- [ ] Error handling is robust
- [ ] Performance is acceptable
- [ ] User experience is positive
- [ ] Documentation is accurate
- [ ] Security considerations addressed
- [ ] Edge cases handled properly

## Conclusion

This comprehensive QA testing guide ensures thorough validation of the Subtitle
Manager application across all major features and edge cases. Regular execution
of these test scenarios helps maintain quality and identify issues early in the
development cycle.

For questions or issues with this testing guide, refer to the development team
or update this document with new test scenarios as features are added.
