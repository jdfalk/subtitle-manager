# TASK-03-001: GitHub Issue Management and Documentation

<!-- file: docs/tasks/03-github-management/TASK-03-001-issue-management.md -->
<!-- version: 1.0.0 -->
<!-- guid: k1l2m3n4-o5p6-7890-1234-f01234567890 -->

## 🎯 Objective

Comprehensively review all GitHub issues across the subtitle-manager repository, link them to appropriate gcommon migration tasks, resolve outdated issues, and create detailed resolution documentation.

## 📋 Acceptance Criteria

- [ ] Complete audit of all open and closed GitHub issues
- [ ] Link issues to specific gcommon migration tasks where applicable
- [ ] Create resolution documentation for each addressed issue
- [ ] Close resolved issues with detailed explanations
- [ ] Update issue labels and milestones for organization
- [ ] Create new issues for discovered gaps in functionality
- [ ] Generate comprehensive issue resolution report
- [ ] Set up automated issue management workflows

## 🔍 Current State Analysis

### Repository Issues Overview

Need to analyze:

1. **Open Issues**: Current problems and feature requests
2. **Closed Issues**: Historical context and patterns
3. **Issue Labels**: Organization and categorization
4. **Milestones**: Release planning and tracking
5. **Project Boards**: Task organization and workflow

### Common Issue Categories Expected

1. **gcommon Migration**: Protobuf package replacements
2. **UI/UX Issues**: Interface problems and improvements
3. **Authentication**: OAuth, session management, user roles
4. **Provider Issues**: Subtitle source configuration and reliability
5. **Performance**: Speed, memory usage, optimization
6. **Documentation**: Missing or outdated documentation
7. **Testing**: Test coverage and quality assurance

## 🔧 Implementation Steps

### Step 1: Set up GitHub API access

```bash
# Install required dependencies
pip install PyGithub requests tabulate markdown

# Create GitHub API client script
cat > scripts/github_issue_manager.py << 'EOF'
#!/usr/bin/env python3
# file: scripts/github_issue_manager.py
# version: 1.0.0
# guid: l2m3n4o5-p6q7-8901-2345-012345678901

import os
import json
import time
from datetime import datetime, timedelta
from github import Github
from tabulate import tabulate
import requests

class GitHubIssueManager:
    """Manage GitHub issues for subtitle-manager repository"""
    
    def __init__(self, token=None, repo_name="jdfalk/subtitle-manager"):
        self.token = token or os.getenv('GITHUB_TOKEN')
        if not self.token:
            raise ValueError("GitHub token required. Set GITHUB_TOKEN environment variable.")
        
        self.github = Github(self.token)
        self.repo = self.github.get_repo(repo_name)
        self.repo_name = repo_name
    
    def get_all_issues(self, state='all'):
        """Get all issues (including PRs) from repository"""
        issues = []
        
        for issue in self.repo.get_issues(state=state):
            issue_data = {
                'number': issue.number,
                'title': issue.title,
                'state': issue.state,
                'created_at': issue.created_at,
                'updated_at': issue.updated_at,
                'closed_at': issue.closed_at,
                'author': issue.user.login if issue.user else 'Unknown',
                'assignees': [a.login for a in issue.assignees],
                'labels': [l.name for l in issue.labels],
                'milestone': issue.milestone.title if issue.milestone else None,
                'body': issue.body or '',
                'comments': issue.comments,
                'is_pull_request': issue.pull_request is not None,
                'url': issue.html_url
            }
            issues.append(issue_data)
        
        return issues
    
    def categorize_issues(self, issues):
        """Categorize issues by type and relevance to gcommon migration"""
        categories = {
            'gcommon_migration': [],
            'ui_ux': [],
            'authentication': [],
            'providers': [],
            'performance': [],
            'documentation': [],
            'testing': [],
            'bugs': [],
            'features': [],
            'other': []
        }
        
        # Keywords for categorization
        keywords = {
            'gcommon_migration': ['protobuf', 'configpb', 'databasepb', 'gcommonauth', 'gcommon', 'proto'],
            'ui_ux': ['ui', 'ux', 'interface', 'frontend', 'design', 'layout', 'navigation'],
            'authentication': ['auth', 'login', 'oauth', 'session', 'user', 'permission'],
            'providers': ['provider', 'opensubtitles', 'source', 'api'],
            'performance': ['performance', 'slow', 'speed', 'memory', 'cpu', 'optimization'],
            'documentation': ['docs', 'documentation', 'readme', 'wiki'],
            'testing': ['test', 'testing', 'unittest', 'e2e', 'selenium'],
            'bugs': ['bug', 'error', 'crash', 'broken', 'fix'],
            'features': ['feature', 'enhancement', 'improvement', 'add']
        }
        
        for issue in issues:
            # Skip pull requests for now
            if issue['is_pull_request']:
                continue
            
            categorized = False
            text_to_search = f"{issue['title']} {issue['body']}".lower()
            
            # Check labels first
            for label in issue['labels']:
                label_lower = label.lower()
                for category, category_keywords in keywords.items():
                    if any(keyword in label_lower for keyword in category_keywords):
                        categories[category].append(issue)
                        categorized = True
                        break
                if categorized:
                    break
            
            # If not categorized by labels, check title and body
            if not categorized:
                for category, category_keywords in keywords.items():
                    if any(keyword in text_to_search for keyword in category_keywords):
                        categories[category].append(issue)
                        categorized = True
                        break
            
            # Default to 'other' if not categorized
            if not categorized:
                categories['other'].append(issue)
        
        return categories
    
    def generate_issue_report(self, categories):
        """Generate comprehensive issue report"""
        report = []
        report.append("# GitHub Issues Analysis Report")
        report.append(f"Generated: {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}")
        report.append(f"Repository: {self.repo_name}")
        report.append("")
        
        # Summary statistics
        total_issues = sum(len(issues) for issues in categories.values())
        open_issues = sum(1 for issues in categories.values() for issue in issues if issue['state'] == 'open')
        closed_issues = total_issues - open_issues
        
        report.append("## Summary Statistics")
        report.append(f"- Total Issues: {total_issues}")
        report.append(f"- Open Issues: {open_issues}")
        report.append(f"- Closed Issues: {closed_issues}")
        report.append("")
        
        # Category breakdown
        report.append("## Issues by Category")
        category_table = []
        for category, issues in categories.items():
            open_count = sum(1 for issue in issues if issue['state'] == 'open')
            closed_count = len(issues) - open_count
            category_table.append([
                category.replace('_', ' ').title(),
                len(issues),
                open_count,
                closed_count
            ])
        
        report.append(tabulate(
            category_table,
            headers=['Category', 'Total', 'Open', 'Closed'],
            tablefmt='markdown'
        ))
        report.append("")
        
        # Detailed category analysis
        for category, issues in categories.items():
            if not issues:
                continue
            
            report.append(f"## {category.replace('_', ' ').title()} Issues")
            
            issue_table = []
            for issue in issues:
                issue_table.append([
                    f"#{issue['number']}",
                    issue['title'][:50] + ('...' if len(issue['title']) > 50 else ''),
                    issue['state'].title(),
                    issue['created_at'].strftime('%Y-%m-%d'),
                    ', '.join(issue['labels'][:3]) + ('...' if len(issue['labels']) > 3 else '')
                ])
            
            report.append(tabulate(
                issue_table,
                headers=['Number', 'Title', 'State', 'Created', 'Labels'],
                tablefmt='markdown'
            ))
            report.append("")
        
        return '\n'.join(report)
    
    def link_issues_to_tasks(self, categories):
        """Link issues to gcommon migration tasks"""
        task_mappings = {
            'configpb': 'TASK-01-001-replace-configpb.md',
            'databasepb': 'TASK-01-002-replace-databasepb.md',
            'gcommonauth': 'TASK-01-003-replace-gcommonauth.md',
            'ui': 'TASK-02-001-fix-layout-navigation.md',
            'testing': 'TASK-02-002-selenium-testing.md'
        }
        
        linked_issues = {}
        
        for category, issues in categories.items():
            for issue in issues:
                text_to_search = f"{issue['title']} {issue['body']}".lower()
                
                for keyword, task_file in task_mappings.items():
                    if keyword in text_to_search:
                        if task_file not in linked_issues:
                            linked_issues[task_file] = []
                        linked_issues[task_file].append(issue)
        
        return linked_issues
    
    def create_issue_comment(self, issue_number, comment):
        """Add comment to an issue"""
        issue = self.repo.get_issue(issue_number)
        issue.create_comment(comment)
    
    def close_issue_with_resolution(self, issue_number, resolution_comment):
        """Close issue with resolution comment"""
        issue = self.repo.get_issue(issue_number)
        issue.create_comment(resolution_comment)
        issue.edit(state='closed')
    
    def update_issue_labels(self, issue_number, labels):
        """Update issue labels"""
        issue = self.repo.get_issue(issue_number)
        issue.edit(labels=labels)

# Example usage
if __name__ == "__main__":
    manager = GitHubIssueManager()
    
    print("Fetching all issues...")
    all_issues = manager.get_all_issues()
    
    print("Categorizing issues...")
    categories = manager.categorize_issues(all_issues)
    
    print("Generating report...")
    report = manager.generate_issue_report(categories)
    
    # Save report
    with open('github_issues_report.md', 'w') as f:
        f.write(report)
    
    print("Report saved to github_issues_report.md")
    
    # Link issues to tasks
    linked_issues = manager.link_issues_to_tasks(categories)
    
    print("Issues linked to tasks:")
    for task, issues in linked_issues.items():
        print(f"{task}: {len(issues)} issues")
EOF

chmod +x scripts/github_issue_manager.py
```

### Step 2: Analyze current repository issues

```bash
# Set GitHub token (use your personal access token)
export GITHUB_TOKEN="your_github_token_here"

# Run issue analysis
cd scripts
python github_issue_manager.py

# Review the generated report
cat github_issues_report.md
```

### Step 3: Create task-issue mapping system

Create `scripts/task_issue_mapper.py`:

```python
#!/usr/bin/env python3
# file: scripts/task_issue_mapper.py
# version: 1.0.0
# guid: m3n4o5p6-q7r8-9012-3456-123456789012

import os
import re
import json
from pathlib import Path

class TaskIssueMapper:
    """Map GitHub issues to specific tasks and update task files"""
    
    def __init__(self, tasks_dir="docs/tasks"):
        self.tasks_dir = Path(tasks_dir)
        self.issue_mappings = {}
    
    def load_task_files(self):
        """Load all task files and extract metadata"""
        task_files = {}
        
        for task_file in self.tasks_dir.rglob("TASK-*.md"):
            with open(task_file, 'r') as f:
                content = f.read()
            
            # Extract task metadata
            task_id = task_file.stem
            title_match = re.search(r'^# (.+)$', content, re.MULTILINE)
            title = title_match.group(1) if title_match else "Unknown"
            
            # Extract objective
            objective_match = re.search(r'## 🎯 Objective\n\n(.+?)(?=\n##|\n\n##|$)', content, re.DOTALL)
            objective = objective_match.group(1).strip() if objective_match else ""
            
            task_files[task_id] = {
                'file_path': task_file,
                'title': title,
                'objective': objective,
                'content': content
            }
        
        return task_files
    
    def map_issues_to_tasks(self, issues, task_files):
        """Map issues to appropriate tasks based on content analysis"""
        mappings = {}
        
        # Define mapping rules
        mapping_rules = {
            'TASK-01-001-replace-configpb': [
                'configpb', 'config protobuf', 'configuration', 'settings protobuf'
            ],
            'TASK-01-002-replace-databasepb': [
                'databasepb', 'database protobuf', 'db schema', 'database types'
            ],
            'TASK-01-003-replace-gcommonauth': [
                'gcommonauth', 'authentication', 'auth types', 'user auth'
            ],
            'TASK-02-001-fix-layout-navigation': [
                'ui layout', 'navigation', 'sidebar', 'menu', 'interface', 'user interface'
            ],
            'TASK-02-002-selenium-testing': [
                'testing', 'e2e', 'selenium', 'browser testing', 'ui testing'
            ]
        }
        
        for issue in issues:
            issue_text = f"{issue['title']} {issue['body']}".lower()
            
            for task_id, keywords in mapping_rules.items():
                if any(keyword in issue_text for keyword in keywords):
                    if task_id not in mappings:
                        mappings[task_id] = []
                    mappings[task_id].append(issue)
        
        return mappings
    
    def update_task_files_with_issues(self, mappings, task_files):
        """Update task files to include related GitHub issues"""
        for task_id, issues in mappings.items():
            if task_id not in task_files:
                continue
            
            task_file = task_files[task_id]
            content = task_file['content']
            
            # Create GitHub issues section
            issues_section = self.create_issues_section(issues)
            
            # Insert or update the issues section
            if '## 🔗 Related GitHub Issues' in content:
                # Replace existing section
                pattern = r'## 🔗 Related GitHub Issues.*?(?=\n## |\n---|\Z)'
                content = re.sub(pattern, issues_section, content, flags=re.DOTALL)
            else:
                # Add new section before "Notes for AI Agent"
                if '## 📝 Notes for AI Agent' in content:
                    content = content.replace(
                        '## 📝 Notes for AI Agent',
                        f'{issues_section}\n\n## 📝 Notes for AI Agent'
                    )
                else:
                    # Add at the end
                    content += f'\n\n{issues_section}'
            
            # Write updated content back to file
            with open(task_file['file_path'], 'w') as f:
                f.write(content)
    
    def create_issues_section(self, issues):
        """Create GitHub issues section for task file"""
        if not issues:
            return ""
        
        section = ["## 🔗 Related GitHub Issues"]
        section.append("")
        section.append("This task addresses the following GitHub issues:")
        section.append("")
        
        for issue in issues:
            status_emoji = "🟢" if issue['state'] == 'closed' else "🔴"
            section.append(f"- {status_emoji} [#{issue['number']}]({issue['url']}) - {issue['title']}")
            
            if issue['labels']:
                labels_str = ', '.join(f"`{label}`" for label in issue['labels'])
                section.append(f"  - Labels: {labels_str}")
            
            if issue['body']:
                # Add first line of body as description
                first_line = issue['body'].split('\n')[0][:100]
                if len(issue['body']) > 100:
                    first_line += "..."
                section.append(f"  - Description: {first_line}")
        
        section.append("")
        section.append("### Issue Resolution Strategy")
        section.append("")
        section.append("When completing this task:")
        section.append("")
        section.append("1. **Reference Issues**: Include issue numbers in commit messages")
        section.append("2. **Update Issues**: Comment on issues with resolution details")
        section.append("3. **Close Issues**: Close resolved issues with explanation")
        section.append("4. **Document Changes**: Update issue descriptions if needed")
        section.append("")
        
        return '\n'.join(section)
    
    def generate_resolution_comments(self, issues):
        """Generate resolution comments for issues"""
        comments = {}
        
        for issue in issues:
            comment = f"""## ✅ Issue Resolution

This issue has been addressed as part of the comprehensive subtitle-manager refactoring to use gcommon packages.

### Resolution Details:

**Changes Made:**
- Replaced local protobuf packages with gcommon equivalents
- Updated import statements and type references
- Implemented opaque API pattern for protobuf access
- Added comprehensive testing and validation

**Related Tasks:**
- Configuration migration (configpb → gcommon/config)
- Database types migration (databasepb → gcommon/database)  
- Authentication types migration (gcommonauth → gcommon/common)

**Testing:**
- Unit tests updated for new package structure
- Integration tests verify functionality
- End-to-end tests validate user workflows

**Documentation:**
- Updated API documentation
- Added migration guides
- Updated development setup instructions

### Verification Steps:

To verify this issue is resolved:

1. Check that imports use gcommon packages:
   ```bash
   grep -r "gcommon/v1" --include="*.go" .
   ```

2. Verify no old package references remain:
   ```bash
   grep -r "configpb\\|databasepb\\|gcommonauth" --include="*.go" . || echo "All cleaned up!"
   ```

3. Run tests to ensure functionality:
   ```bash
   go test ./...
   ```

### Additional Notes:

This refactoring improves:
- Code maintainability through shared package usage
- Type consistency across applications
- Reduced duplication of protobuf definitions
- Better alignment with gcommon package standards

If you encounter any issues with this resolution, please reopen this issue with specific details about the problem.

---
*This issue was automatically resolved as part of task completion.*"""
            
            comments[issue['number']] = comment
        
        return comments

# Example usage
if __name__ == "__main__":
    from github_issue_manager import GitHubIssueManager
    
    # Load issues and tasks
    issue_manager = GitHubIssueManager()
    all_issues = issue_manager.get_all_issues()
    
    mapper = TaskIssueMapper()
    task_files = mapper.load_task_files()
    
    print(f"Found {len(task_files)} task files")
    print(f"Found {len(all_issues)} GitHub issues")
    
    # Create mappings
    mappings = mapper.map_issues_to_tasks(all_issues, task_files)
    
    print("\nMappings created:")
    for task_id, issues in mappings.items():
        print(f"{task_id}: {len(issues)} issues")
    
    # Update task files
    mapper.update_task_files_with_issues(mappings, task_files)
    print("\nTask files updated with GitHub issue references")
    
    # Generate resolution comments
    all_mapped_issues = [issue for issues in mappings.values() for issue in issues]
    resolution_comments = mapper.generate_resolution_comments(all_mapped_issues)
    
    # Save resolution comments for review
    with open('issue_resolution_comments.json', 'w') as f:
        json.dump(resolution_comments, f, indent=2)
    
    print(f"Generated resolution comments for {len(resolution_comments)} issues")
    print("Comments saved to issue_resolution_comments.json")
```

### Step 4: Create issue resolution workflow

Create `scripts/resolve_issues.py`:

```python
#!/usr/bin/env python3
# file: scripts/resolve_issues.py
# version: 1.0.0
# guid: n4o5p6q7-r8s9-0123-4567-234567890123

import json
import time
from github_issue_manager import GitHubIssueManager

class IssueResolver:
    """Resolve GitHub issues with detailed comments and closures"""
    
    def __init__(self, dry_run=True):
        self.issue_manager = GitHubIssueManager()
        self.dry_run = dry_run
    
    def resolve_issues_batch(self, issue_numbers, resolution_template=None):
        """Resolve multiple issues with consistent messaging"""
        if not resolution_template:
            resolution_template = self.get_default_resolution_template()
        
        resolved_count = 0
        
        for issue_number in issue_numbers:
            try:
                if self.dry_run:
                    print(f"[DRY RUN] Would resolve issue #{issue_number}")
                else:
                    self.issue_manager.close_issue_with_resolution(
                        issue_number, 
                        resolution_template
                    )
                    print(f"✅ Resolved issue #{issue_number}")
                    resolved_count += 1
                    
                    # Add delay to respect rate limits
                    time.sleep(1)
                    
            except Exception as e:
                print(f"❌ Failed to resolve issue #{issue_number}: {str(e)}")
        
        return resolved_count
    
    def get_default_resolution_template(self):
        """Get default resolution template"""
        return """## ✅ Issue Resolved - gcommon Migration Complete

This issue has been resolved as part of the comprehensive subtitle-manager refactoring to use gcommon packages.

### 🔄 Changes Made:

**Package Migration:**
- ✅ Replaced `configpb` with `gcommon/v1/config`
- ✅ Replaced `databasepb` with `gcommon/v1/database`  
- ✅ Replaced `gcommonauth` with `gcommon/v1/common`
- ✅ Updated all import statements and type references
- ✅ Implemented opaque API pattern for protobuf access

**Code Quality:**
- ✅ Added comprehensive unit tests
- ✅ Implemented integration testing
- ✅ Added end-to-end Selenium tests
- ✅ Updated documentation and examples

**UI/UX Improvements:**
- ✅ Fixed navigation and layout issues
- ✅ Improved settings interface
- ✅ Enhanced user management display
- ✅ Added responsive design elements

### 🧪 Verification:

The resolution has been verified through:

1. **Automated Testing:**
   ```bash
   go test ./...  # All tests passing
   npm test      # Frontend tests passing
   ```

2. **Package Usage Verification:**
   ```bash
   grep -r "gcommon/v1" --include="*.go" .  # New packages in use
   grep -r "configpb\|databasepb" --include="*.go" . || echo "Old packages removed"
   ```

3. **End-to-End Testing:**
   - User authentication flows
   - Media library management
   - Subtitle operations
   - Provider configuration
   - Settings management

### 📚 Documentation:

- Updated API documentation with new package usage
- Added migration guides for developers
- Created comprehensive task documentation
- Updated development setup instructions

### 🔗 Related Tasks:

This issue was resolved through the following detailed tasks:
- `TASK-01-001`: Replace configpb package
- `TASK-01-002`: Replace databasepb package  
- `TASK-01-003`: Replace gcommonauth package
- `TASK-02-001`: Fix UI layout and navigation
- `TASK-02-002`: Implement comprehensive testing

### 💡 Future Improvements:

The migration also enables:
- Better code maintainability through shared packages
- Improved type consistency across applications
- Reduced duplication of protobuf definitions
- Enhanced development workflow standardization

---

**Issue Status:** ✅ **RESOLVED**  
**Resolution Date:** {timestamp}  
**Migration Phase:** Complete  

If you experience any issues related to this change, please open a new issue with specific details about the problem and steps to reproduce."""

    def update_issue_with_task_links(self, issue_number, task_links):
        """Update issue with links to related tasks"""
        comment = f"""## 🔗 Related Implementation Tasks

This issue is being addressed through the following detailed implementation tasks:

{chr(10).join(f'- 📋 [{task}](../docs/tasks/{task})' for task in task_links)}

Each task includes:
- ✅ Detailed acceptance criteria
- 📝 Step-by-step implementation guide  
- 🧪 Comprehensive testing requirements
- 📚 Complete documentation references
- 🎯 Success metrics and validation

### 📊 Progress Tracking:

You can track progress on this issue by monitoring the completion of the linked tasks above. Each task is designed to be completed independently by automated agents or developers.

### ⏱️ Expected Timeline:

The implementation tasks are prioritized as follows:
1. **Phase 1**: gcommon package migration (TASK-01-*)
2. **Phase 2**: UI/UX improvements (TASK-02-*)  
3. **Phase 3**: Testing and validation (ongoing)

This issue will be automatically updated as tasks are completed."""
        
        if self.dry_run:
            print(f"[DRY RUN] Would update issue #{issue_number} with task links")
        else:
            self.issue_manager.create_issue_comment(issue_number, comment)

# Example usage and resolution workflow
if __name__ == "__main__":
    # Load previously mapped issues
    with open('issue_resolution_comments.json', 'r') as f:
        resolution_data = json.load(f)
    
    resolver = IssueResolver(dry_run=True)  # Set to False to actually resolve
    
    # Example resolution workflow
    gcommon_issues = [1, 2, 3]  # Replace with actual issue numbers
    ui_issues = [4, 5, 6]       # Replace with actual issue numbers
    
    print("Resolving gcommon migration issues...")
    resolver.resolve_issues_batch(gcommon_issues)
    
    print("Resolving UI/UX issues...")
    resolver.resolve_issues_batch(ui_issues)
    
    print("Resolution workflow completed!")
```

### Step 5: Create automated issue management

Create `scripts/automated_issue_manager.py`:

```python
#!/usr/bin/env python3
# file: scripts/automated_issue_manager.py
# version: 1.0.0
# guid: o5p6q7r8-s9t0-1234-5678-345678901234

import schedule
import time
from datetime import datetime, timedelta
from github_issue_manager import GitHubIssueManager

class AutomatedIssueManager:
    """Automated issue management and monitoring"""
    
    def __init__(self):
        self.issue_manager = GitHubIssueManager()
    
    def daily_issue_review(self):
        """Daily automated issue review"""
        print(f"[{datetime.now()}] Starting daily issue review...")
        
        # Get recent issues
        recent_issues = []
        cutoff_date = datetime.now() - timedelta(days=1)
        
        for issue in self.issue_manager.repo.get_issues(state='open'):
            if issue.updated_at > cutoff_date:
                recent_issues.append(issue)
        
        print(f"Found {len(recent_issues)} recently updated issues")
        
        # Analyze and categorize
        for issue in recent_issues:
            self.analyze_and_label_issue(issue)
        
        # Generate daily report
        self.generate_daily_report(recent_issues)
    
    def analyze_and_label_issue(self, issue):
        """Analyze issue and suggest appropriate labels"""
        title_body = f"{issue.title} {issue.body or ''}".lower()
        
        # Suggested labels based on content
        suggested_labels = []
        
        label_keywords = {
            'gcommon-migration': ['configpb', 'databasepb', 'gcommonauth', 'protobuf'],
            'ui/ux': ['ui', 'interface', 'design', 'layout', 'navigation'],
            'bug': ['bug', 'error', 'crash', 'broken', 'not working'],
            'enhancement': ['feature', 'enhancement', 'improvement', 'add'],
            'performance': ['slow', 'performance', 'speed', 'memory'],
            'documentation': ['docs', 'documentation', 'readme'],
            'testing': ['test', 'testing', 'e2e', 'selenium']
        }
        
        for label, keywords in label_keywords.items():
            if any(keyword in title_body for keyword in keywords):
                suggested_labels.append(label)
        
        # Priority assessment
        priority = self.assess_priority(issue)
        if priority:
            suggested_labels.append(f'priority-{priority}')
        
        # Add labels if they don't exist
        current_labels = [l.name for l in issue.labels]
        new_labels = [label for label in suggested_labels if label not in current_labels]
        
        if new_labels:
            print(f"Issue #{issue.number}: Suggested labels: {', '.join(new_labels)}")
            # Note: In production, you might want to actually apply these labels
            # self.issue_manager.update_issue_labels(issue.number, current_labels + new_labels)
    
    def assess_priority(self, issue):
        """Assess issue priority based on content and context"""
        title_body = f"{issue.title} {issue.body or ''}".lower()
        
        # High priority indicators
        high_priority_keywords = ['crash', 'broken', 'critical', 'urgent', 'security']
        if any(keyword in title_body for keyword in high_priority_keywords):
            return 'high'
        
        # Medium priority indicators
        medium_priority_keywords = ['bug', 'error', 'performance', 'slow']
        if any(keyword in title_body for keyword in medium_priority_keywords):
            return 'medium'
        
        # Low priority for enhancements
        low_priority_keywords = ['enhancement', 'feature', 'improvement']
        if any(keyword in title_body for keyword in low_priority_keywords):
            return 'low'
        
        return None
    
    def generate_daily_report(self, issues):
        """Generate daily issue management report"""
        report = [f"# Daily Issue Report - {datetime.now().strftime('%Y-%m-%d')}"]
        report.append("")
        
        report.append(f"## Summary")
        report.append(f"- Issues reviewed: {len(issues)}")
        report.append(f"- Open issues: {len([i for i in issues if i.state == 'open'])}")
        report.append(f"- Closed issues: {len([i for i in issues if i.state == 'closed'])}")
        report.append("")
        
        if issues:
            report.append("## Recent Activity")
            for issue in issues[:10]:  # Top 10 most recent
                report.append(f"- [#{issue.number}]({issue.html_url}) - {issue.title}")
                report.append(f"  - State: {issue.state}")
                report.append(f"  - Updated: {issue.updated_at.strftime('%Y-%m-%d %H:%M')}")
                report.append(f"  - Labels: {', '.join([l.name for l in issue.labels])}")
                report.append("")
        
        # Save report
        with open(f'reports/daily_issue_report_{datetime.now().strftime("%Y%m%d")}.md', 'w') as f:
            f.write('\n'.join(report))
        
        print("Daily report generated")
    
    def weekly_cleanup(self):
        """Weekly issue cleanup and maintenance"""
        print(f"[{datetime.now()}] Starting weekly cleanup...")
        
        # Find stale issues (no activity for 30+ days)
        cutoff_date = datetime.now() - timedelta(days=30)
        stale_issues = []
        
        for issue in self.issue_manager.repo.get_issues(state='open'):
            if issue.updated_at < cutoff_date:
                stale_issues.append(issue)
        
        print(f"Found {len(stale_issues)} stale issues")
        
        # Comment on stale issues
        for issue in stale_issues:
            self.handle_stale_issue(issue)
    
    def handle_stale_issue(self, issue):
        """Handle stale issue with automated comment"""
        stale_comment = """## 🤖 Automated Stale Issue Notice

This issue has been automatically flagged as stale because it has not had recent activity.

### Next Steps:
- If this issue is still relevant, please comment to keep it open
- If this issue is resolved, please close it
- If this issue needs more information, please provide details

This issue will be automatically closed in 7 days if no activity occurs.

### Related Resources:
- [Current Project Status](../README.md)
- [Known Issues](../docs/known-issues.md)
- [How to Contribute](../CONTRIBUTING.md)"""
        
        print(f"Commenting on stale issue #{issue.number}")
        # Note: In production, uncomment to actually create comments
        # self.issue_manager.create_issue_comment(issue.number, stale_comment)
    
    def start_automated_monitoring(self):
        """Start automated issue monitoring"""
        print("Starting automated issue monitoring...")
        
        # Schedule daily reviews
        schedule.every().day.at("09:00").do(self.daily_issue_review)
        
        # Schedule weekly cleanup
        schedule.every().monday.at("08:00").do(self.weekly_cleanup)
        
        print("Automated monitoring scheduled:")
        print("- Daily reviews at 09:00")
        print("- Weekly cleanup on Mondays at 08:00")
        
        # Keep running
        while True:
            schedule.run_pending()
            time.sleep(60)  # Check every minute

if __name__ == "__main__":
    manager = AutomatedIssueManager()
    
    # Run immediate review
    manager.daily_issue_review()
    
    # Optionally start automated monitoring
    # manager.start_automated_monitoring()
```

### Step 6: Execute comprehensive issue management

```bash
# Run the complete issue management workflow

# Step 1: Analyze all issues
python scripts/github_issue_manager.py

# Step 2: Map issues to tasks and update task files
python scripts/task_issue_mapper.py

# Step 3: Review generated mappings and resolution comments
cat issue_resolution_comments.json

# Step 4: Execute issue resolutions (review first with dry_run=True)
python scripts/resolve_issues.py

# Step 5: Start automated monitoring (optional)
# python scripts/automated_issue_manager.py
```

### Step 7: Create issue resolution documentation

Create `docs/issue-resolution-report.md`:

```markdown
# GitHub Issue Resolution Report

<!-- file: docs/issue-resolution-report.md -->
<!-- version: 1.0.0 -->
<!-- guid: p6q7r8s9-t0u1-2345-6789-456789012345 -->

## Executive Summary

This report documents the comprehensive resolution of GitHub issues as part of the subtitle-manager refactoring to use gcommon packages.

### Resolution Statistics

- **Total Issues Reviewed**: [To be filled by script]
- **Issues Resolved**: [To be filled by script]
- **Issues Linked to Tasks**: [To be filled by script]
- **New Issues Created**: [To be filled by script]

## Issue Categories and Resolutions

### gcommon Migration Issues

**Issues Addressed:**
- Configuration package migration (configpb → gcommon/config)
- Database package migration (databasepb → gcommon/database)
- Authentication package migration (gcommonauth → gcommon/common)

**Resolution Method:**
- Systematic package replacement
- Import statement updates
- Opaque API implementation
- Comprehensive testing

### UI/UX Issues

**Issues Addressed:**
- Navigation layout problems
- User management display issues
- Provider configuration modals
- Settings interface improvements

**Resolution Method:**
- Complete UI redesign following Bazarr patterns
- Responsive design implementation
- Component refactoring
- User experience testing

### Testing and Quality Assurance

**Issues Addressed:**
- Lack of comprehensive testing
- Missing end-to-end test coverage
- Performance testing gaps
- Visual regression testing needs

**Resolution Method:**
- Selenium-based E2E testing framework
- Video recording capabilities
- Performance benchmarking
- Visual regression detection

## Automated Issue Management

### Daily Monitoring

- Automated issue categorization
- Priority assessment
- Label management
- Activity tracking

### Weekly Maintenance

- Stale issue detection
- Cleanup workflows
- Progress reporting
- Milestone tracking

## Future Issue Prevention

### Process Improvements

1. **Automated Testing**: Comprehensive test suites prevent regressions
2. **Code Review**: Standardized review process catches issues early
3. **Documentation**: Complete documentation reduces user confusion
4. **Monitoring**: Automated monitoring detects issues quickly

### Development Workflow

1. **Issue Templates**: Standardized issue creation
2. **Task Linking**: Direct connection between issues and tasks
3. **Progress Tracking**: Transparent progress visibility
4. **Resolution Documentation**: Detailed resolution records

## Recommendations

### Short-term Actions

1. Continue automated issue monitoring
2. Maintain comprehensive test coverage
3. Update documentation regularly
4. Monitor user feedback channels

### Long-term Strategy

1. Implement preventive measures
2. Enhance automated workflows
3. Improve user onboarding
4. Expand testing coverage

---

*This report is automatically updated as issues are resolved and new issues are created.*
```

### Step 8: Create issue templates for future use

Create `.github/ISSUE_TEMPLATE/bug_report.yml`:

```yaml
name: Bug Report
description: Report a bug to help us improve
title: "[Bug]: "
labels: ["bug"]
body:
  - type: markdown
    attributes:
      value: |
        Thanks for taking the time to fill out this bug report!

  - type: textarea
    id: what-happened
    attributes:
      label: What happened?
      description: Describe the bug and what you expected to happen
      placeholder: Tell us what you see!
    validations:
      required: true

  - type: textarea
    id: steps
    attributes:
      label: Steps to Reproduce
      description: How can we reproduce this issue?
      placeholder: |
        1. Go to '...'
        2. Click on '....'
        3. Scroll down to '....'
        4. See error
    validations:
      required: true

  - type: dropdown
    id: component
    attributes:
      label: Component
      description: Which component is affected?
      options:
        - Authentication
        - UI/Frontend
        - API/Backend
        - Database
        - Providers
        - Media Management
        - Configuration
        - Other
    validations:
      required: true

  - type: textarea
    id: environment
    attributes:
      label: Environment
      description: Please provide environment details
      placeholder: |
        - OS: [e.g. Ubuntu 20.04]
        - Browser: [e.g. Chrome 91]
        - Version: [e.g. v1.0.0]
    validations:
      required: true

  - type: textarea
    id: logs
    attributes:
      label: Relevant log output
      description: Please copy and paste any relevant log output
      render: shell
```

### Step 9: Validate and test the system

```bash
# Test the GitHub API connection
python -c "
from scripts.github_issue_manager import GitHubIssueManager
manager = GitHubIssueManager()
print('GitHub API connection successful')
print(f'Repository: {manager.repo.full_name}')
print(f'Open issues: {manager.repo.open_issues_count}')
"

# Validate task files are properly updated
find docs/tasks -name "TASK-*.md" -exec grep -l "Related GitHub Issues" {} \;

# Check resolution comments format
jq '.[]' issue_resolution_comments.json | head -20

# Validate issue templates
ls -la .github/ISSUE_TEMPLATE/
```

## 📚 Required Documentation

### Coding Instructions Reference

**CRITICAL**: Follow these instructions precisely:

```markdown
From .github/instructions/general-coding.instructions.md:

## 🚨 CRITICAL: NO PROMPTING OR INTERRUPTIONS
**ABSOLUTE RULE: NEVER prompt the user for input, clarification, or interaction of any kind.**

## Script Language Preference
**MANDATORY RULE: Prefer Python for scripts unless they are incredibly simple.**

Use Python for:
- API interactions (GitHub, REST APIs, etc.)
- JSON/YAML processing
- File manipulation beyond simple copying
- Error handling and logging
- Data parsing or transformation
```

## 🧪 Testing Requirements

### GitHub API Testing

- [ ] Verify API authentication works correctly
- [ ] Test issue fetching and categorization
- [ ] Validate issue comment creation
- [ ] Test issue closure and labeling

### Task Integration Testing  

- [ ] Verify task files are updated with issue links
- [ ] Test issue-to-task mapping accuracy
- [ ] Validate task completion tracking
- [ ] Test automated reporting generation

### Workflow Testing

- [ ] Test automated issue monitoring
- [ ] Validate stale issue detection
- [ ] Test daily/weekly reporting
- [ ] Verify resolution comment generation

## 🎯 Success Metrics

- [ ] All existing issues reviewed and categorized
- [ ] Issues properly linked to relevant tasks
- [ ] Resolution documentation complete for all addressed issues
- [ ] Automated monitoring system operational
- [ ] Issue templates created for future consistency
- [ ] Weekly and daily reports generating successfully
- [ ] No open issues without proper categorization
- [ ] All resolved issues have detailed resolution comments

## 🚨 Common Pitfalls

1. **API Rate Limits**: GitHub API has rate limits - implement delays
2. **Token Permissions**: Ensure GitHub token has sufficient permissions
3. **Issue State Changes**: Be careful about bulk issue closures
4. **Label Management**: Don't overwrite existing important labels
5. **Comment Formatting**: Ensure markdown formatting is correct
6. **Task File Updates**: Don't break existing task file format
7. **Dry Run Testing**: Always test with dry_run=True first

## 📖 Additional Resources

- [GitHub REST API Documentation](https://docs.github.com/en/rest)
- [PyGithub Documentation](https://pygithub.readthedocs.io/)
- [GitHub Issue Templates](https://docs.github.com/en/communities/using-templates-to-encourage-useful-issues-and-pull-requests)
- [Markdown Formatting Guide](https://guides.github.com/features/mastering-markdown/)

## 🔄 Related Tasks

- **TASK-01-001 to TASK-01-003**: gcommon migration tasks (issues will be linked to these)
- **TASK-02-001**: UI fixes (UI-related issues will reference this)
- **TASK-02-002**: Testing framework (testing issues will reference this)

## 📝 Notes for AI Agent

- Set up GitHub token as environment variable before running scripts
- Always use dry_run=True for initial testing
- Review generated resolution comments before applying them
- Be conservative with automatic issue closures
- Focus on linking issues to tasks rather than immediate closure
- Standard Python tooling - no special libraries beyond PyGithub required
- Generate comprehensive reports for human review before bulk operations
- If any GitHub API calls fail, continue with remaining operations and report errors
