<!-- file: docs/ISSUE_ANALYSIS_AND_AUTOMATION.md -->
<!-- version: 1.0.0 -->
<!-- guid: d4e5f6a7-b8c9-0123-def4-456789012345 -->

# GitHub Issue Analysis and Automation System

## Overview

This system enhances the existing GitHub issue management infrastructure with automated analysis, categorization, and priority scoring. It builds on the established issue update system and VS Code task integration to provide intelligent issue processing.

## System Components

### Core Components

1. **`scripts/issue_analyzer.py`** - Core analysis engine
   - Automated issue categorization
   - Priority scoring algorithm
   - Module detection and assignment
   - Recommendation generation

2. **`scripts/issue_automation_wrapper.py`** - VS Code task integration
   - Task-friendly interface
   - Logging integration with copilot-agent-util
   - Summary and detailed reporting
   - Status monitoring

3. **`scripts/issue_analysis_config.py`** - Configuration
   - Priority weight definitions
   - Module keyword mappings
   - Automation rules and thresholds

### Integration Points

- **Existing Issue Management**: Uses `.github/issue-updates/` JSON format
- **VS Code Tasks**: Integrated with copilot-agent-util logging
- **Enhanced Issue Manager**: Compatible with v2.0 timestamp tracking
- **GitHub API**: Direct integration for real-time analysis

## Features

### Automated Analysis

- **Priority Scoring**: Multi-factor algorithm considering:
  - Issue type (bug, enhancement, etc.)
  - Module criticality
  - Activity indicators
  - Existing labels and metadata

- **Module Detection**: Automatic module assignment based on:
  - Title and description content analysis
  - Keyword matching
  - Context clues

- **Label Automation**: Intelligent label suggestions for:
  - Priority levels (critical, high, medium, low)
  - Issue types (bug, enhancement, documentation, etc.)
  - Module assignments (auth, database, web, etc.)

### VS Code Integration

The system integrates with your existing VS Code task workflow:

#### Available Tasks

1. **Issue Analysis - High Priority**
   - Scans all open issues for high-priority items
   - Creates automation updates for processing
   - Provides summary report

2. **Issue Analysis - Specific Issue**
   - Analyzes a single issue by number
   - Provides detailed analysis and recommendations
   - Useful for manual review

3. **Issue Analysis - Repository Scan**
   - Custom priority threshold scanning
   - Comprehensive repository analysis
   - Configurable reporting

4. **Issue Status Check**
   - Shows pending and processed issue updates
   - Monitors automation queue status
   - System health check

## Usage

### Via VS Code Tasks

1. **Run High-Priority Analysis**:
   - Open Command Palette (Cmd+Shift+P)
   - Select "Tasks: Run Task"
   - Choose "Issue Analysis - High Priority"
   - Review results in logs/issue_analysis.log

2. **Analyze Specific Issue**:
   - Run "Issue Analysis - Specific Issue" task
   - Enter issue number when prompted
   - View detailed analysis results

3. **Check System Status**:
   - Run "Issue Status Check" task
   - View pending and processed updates

### Command Line Interface

```bash
# Analyze high-priority issues
python3 scripts/issue_automation_wrapper.py analyze

# Analyze specific issue
python3 scripts/issue_automation_wrapper.py analyze --issue-number 1789

# Custom priority threshold
python3 scripts/issue_automation_wrapper.py analyze --min-priority 80

# Check automation status
python3 scripts/issue_automation_wrapper.py status

# Direct analyzer usage
python3 scripts/issue_analyzer.py --repo-owner jdfalk --repo-name subtitle-manager --analyze-issue 1789
```

### Environment Setup

Set your GitHub token:
```bash
export GITHUB_TOKEN="your_github_token_here"
```

Or pass it directly:
```bash
python3 scripts/issue_automation_wrapper.py analyze --github-token "your_token"
```

## Priority Scoring Algorithm

The system uses a sophisticated multi-factor scoring algorithm:

### Base Scores by Type
- **Bug**: 50 points
- **Enhancement**: 40 points
- **Security**: 95 points
- **Performance**: 60 points
- **Documentation**: 20 points

### Module Criticality Multipliers
- **Auth**: +80 points
- **Database**: +70 points
- **Web/API**: +60 points
- **Config**: +50 points
- **UI**: +30 points

### Activity Factors
- **Recent Activity** (< 7 days): +20 points
- **Multiple Comments** (> 5): +15 points
- **Long Open** (> 90 days): -10 points
- **Stale Issues**: -20 points

### Priority Thresholds
- **Critical**: 80+ points
- **High**: 60-79 points
- **Medium**: 40-59 points
- **Low**: < 40 points

## Automation Workflow

1. **Analysis Phase**:
   - Fetch open issues from GitHub API
   - Apply priority scoring algorithm
   - Generate categorization and recommendations

2. **Update Creation Phase**:
   - Create JSON update files in `.github/issue-updates/`
   - Use existing v2.0 enhanced format
   - Include analysis metadata and recommendations

3. **Processing Phase**:
   - Existing workflows process update files
   - Apply label changes and assignments
   - Move processed files to processed/ directory

4. **Monitoring Phase**:
   - Track automation success/failure
   - Log all operations to logs/ directory
   - Provide status reporting

## Configuration

### Priority Weights

Modify `scripts/issue_analysis_config.py` to adjust:
- Module importance scores
- Issue type base values
- Activity factor weights
- Automation thresholds

### Module Keywords

Update keyword mappings for better module detection:
```python
MODULE_KEYWORDS = {
    'auth': ['auth', 'login', 'session', 'token'],
    'database': ['database', 'db', 'sql', 'migration'],
    # Add custom keywords
}
```

### Automation Rules

Configure automation behavior:
```python
AUTOMATION_RULES = {
    'auto_priority_labels': True,
    'auto_module_labels': True,
    'auto_process_threshold': 60,
    'critical_threshold': 80
}
```

## Integration with Existing System

### Compatibility

- **Issue Manager**: Fully compatible with existing issue_manager.py
- **Enhanced Manager**: Uses v2.0 timestamp format
- **Workflows**: Integrates with existing GitHub Actions
- **VS Code Tasks**: Extends current task system

### Data Flow

```
GitHub Issues → Analysis Engine → Update JSON Files → Existing Workflows → Applied Changes
```

### Logging

All operations logged to:
- `logs/issue_analysis.log` - Analysis results
- VS Code task output via copilot-agent-util
- GitHub workflow logs for update processing

## Best Practices

### Regular Analysis

Run high-priority analysis regularly:
- Daily for active repositories
- Weekly for maintenance mode
- On-demand for issue triage sessions

### Custom Thresholds

Adjust priority thresholds based on:
- Repository activity level
- Team capacity
- Project phase (development vs maintenance)

### Review Recommendations

Always review automation recommendations:
- Verify module assignments
- Confirm priority levels
- Check label suggestions

### Monitor Status

Use status checks to:
- Ensure automation is working
- Track processing delays
- Identify configuration issues

## Troubleshooting

### Common Issues

1. **GitHub Token Issues**:
   - Ensure GITHUB_TOKEN is set
   - Verify token has repo access
   - Check rate limiting

2. **Analysis Errors**:
   - Review logs/issue_analysis.log
   - Check network connectivity
   - Verify repository permissions

3. **Task Failures**:
   - Ensure Python dependencies installed
   - Check copilot-agent-util installation
   - Verify script permissions

### Debug Mode

Enable detailed logging:
```bash
python3 scripts/issue_analyzer.py --repo-owner jdfalk --repo-name subtitle-manager --process-high-priority --output-format json
```

## Future Enhancements

### Planned Features

- **Machine Learning Integration**: Learn from manual label assignments
- **Cross-Repository Analysis**: Analyze patterns across multiple repos
- **Slack/Discord Integration**: Automated notifications for critical issues
- **Advanced Filtering**: More sophisticated query capabilities

### Extension Points

- **Custom Analyzers**: Plugin system for domain-specific analysis
- **External Data Sources**: Integration with project management tools
- **Advanced Metrics**: Detailed analytics and reporting

## Contributing

To enhance the issue analysis system:

1. **Modify Configuration**: Update `issue_analysis_config.py`
2. **Extend Analysis**: Add new factors to priority scoring
3. **Improve Detection**: Enhance keyword matching algorithms
4. **Add Integrations**: Connect with additional tools and services

All changes should maintain compatibility with the existing issue management infrastructure.
