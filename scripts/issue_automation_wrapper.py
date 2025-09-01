#!/usr/bin/env python3
# file: scripts/issue_automation_wrapper.py
# version: 1.0.0
# guid: b2c3d4e5-f6a7-8901-bcde-f23456789012

"""
Issue Automation Wrapper for VS Code Tasks

Provides a unified interface for issue analysis and automation that works
with VS Code tasks and the copilot-agent-util system.

Integrates with:
- issue_analyzer.py: Core analysis engine
- existing issue management infrastructure
- VS Code task system
- copilot-agent-util logging
"""

import json
import logging
import os
import sys
from pathlib import Path
from typing import Dict, List, Any

# Import our issue analyzer (conditionally to avoid import errors for status command)
try:
    sys.path.append(str(Path(__file__).parent))
    from issue_analyzer import IssueAnalyzer, IssueAutomation
    ANALYZER_AVAILABLE = True
except ImportError as e:
    ANALYZER_AVAILABLE = False
    IMPORT_ERROR = str(e)

# Setup logging for copilot-agent-util compatibility
logging.basicConfig(
    level=logging.INFO,
    format='[%(asctime)s] %(levelname)s: %(message)s',
    datefmt='%Y-%m-%d %H:%M:%S'
)
logger = logging.getLogger(__name__)

class TaskIntegration:
    """Integration layer for VS Code tasks and automation"""
    
    def __init__(self, repo_path: str = "."):
        self.repo_path = Path(repo_path)
        self.logs_dir = self.repo_path / "logs"
        self.logs_dir.mkdir(exist_ok=True)
        
    def run_issue_analysis(self, repo_owner: str, repo_name: str, 
                          github_token: str, **kwargs) -> Dict[str, Any]:
        """
        Run comprehensive issue analysis with task-friendly output
        
        Args:
            repo_owner: GitHub repository owner
            repo_name: GitHub repository name  
            github_token: GitHub API token
            **kwargs: Additional analysis parameters
            
        Returns:
            Analysis results and summary
        """
        log_file = self.logs_dir / "issue_analysis.log"
        
        try:
            logger.info("Starting issue analysis...")
            
            # Initialize analyzer
            analyzer = IssueAnalyzer(repo_owner, repo_name, github_token)
            automation = IssueAutomation(str(self.repo_path))
            
            # Get analysis parameters
            min_priority = kwargs.get('min_priority', 60)
            specific_issue = kwargs.get('issue_number')
            
            results = {
                'success': True,
                'timestamp': str(Path().cwd()),
                'analysis_summary': {},
                'high_priority_issues': [],
                'automation_updates': [],
                'recommendations': []
            }
            
            if specific_issue:
                # Analyze specific issue
                results.update(self._analyze_single_issue(
                    analyzer, specific_issue
                ))
            else:
                # Analyze all high-priority issues
                results.update(self._analyze_high_priority_issues(
                    analyzer, automation, min_priority
                ))
            
            # Write results to log
            with open(log_file, 'w') as f:
                json.dump(results, f, indent=2)
            
            logger.info(f"Analysis complete. Results logged to: {log_file}")
            return results
            
        except Exception as e:
            error_msg = f"Issue analysis failed: {str(e)}"
            logger.error(error_msg)
            
            error_result = {
                'success': False,
                'error': error_msg,
                'timestamp': str(Path().cwd())
            }
            
            with open(log_file, 'w') as f:
                json.dump(error_result, f, indent=2)
                
            return error_result
    
    def _analyze_single_issue(self, analyzer: IssueAnalyzer, 
                            issue_number: int) -> Dict[str, Any]:
        """Analyze a single issue"""
        import requests
        
        issue_url = f"{analyzer.base_url}/issues/{issue_number}"
        response = requests.get(issue_url, headers=analyzer.headers)
        
        if response.status_code != 200:
            raise Exception(f"Failed to fetch issue #{issue_number}")
        
        issue = response.json()
        analysis = analyzer.analyze_issue(issue)
        
        return {
            'analysis_type': 'single_issue',
            'issue_number': issue_number,
            'analysis_summary': analysis,
            'recommendations': analysis['automation_recommendations']
        }
    
    def _analyze_high_priority_issues(self, analyzer: IssueAnalyzer,
                                    automation: IssueAutomation,
                                    min_priority: int) -> Dict[str, Any]:
        """Analyze all high-priority issues"""
        import requests
        
        # Get open issues
        issues_url = f"{analyzer.base_url}/issues"
        params = {"state": "open", "per_page": 100}
        
        response = requests.get(issues_url, headers=analyzer.headers, params=params)
        if response.status_code != 200:
            raise Exception(f"Failed to fetch issues: {response.status_code}")
        
        issues = response.json()
        high_priority_issues = []
        total_analyzed = 0
        
        for issue in issues:
            if issue.get('pull_request'):  # Skip PRs
                continue
                
            total_analyzed += 1
            analysis = analyzer.analyze_issue(issue)
            
            if analysis['priority_score'] >= min_priority:
                high_priority_issues.append({
                    'number': issue['number'],
                    'title': issue['title'],
                    'priority_score': analysis['priority_score'],
                    'urgency_level': analysis['urgency_level'],
                    'recommendations': analysis['automation_recommendations']
                })
        
        # Create automation updates for high-priority issues
        update_ids = automation.process_high_priority_issues(analyzer, min_priority)
        
        return {
            'analysis_type': 'repository_scan',
            'total_issues_analyzed': total_analyzed,
            'high_priority_count': len(high_priority_issues),
            'min_priority_threshold': min_priority,
            'high_priority_issues': high_priority_issues,
            'automation_updates': update_ids,
            'summary_stats': self._generate_summary_stats(high_priority_issues)
        }
    
    def _generate_summary_stats(self, high_priority_issues: List[Dict]) -> Dict[str, Any]:
        """Generate summary statistics for analysis results"""
        if not high_priority_issues:
            return {'message': 'No high-priority issues found'}
        
        urgency_counts = {}
        total_score = 0
        
        for issue in high_priority_issues:
            urgency = issue['urgency_level']
            urgency_counts[urgency] = urgency_counts.get(urgency, 0) + 1
            total_score += issue['priority_score']
        
        avg_score = total_score / len(high_priority_issues)
        
        return {
            'total_high_priority': len(high_priority_issues),
            'average_priority_score': round(avg_score, 2),
            'urgency_breakdown': urgency_counts,
            'highest_priority': max(high_priority_issues, key=lambda x: x['priority_score']),
            'needs_immediate_attention': len([
                issue for issue in high_priority_issues 
                if issue['urgency_level'] == 'critical'
            ])
        }

def main():
    """Main CLI interface for task integration"""
    import argparse
    
    parser = argparse.ArgumentParser(
        description="Issue Analysis and Automation for VS Code Tasks"
    )
    parser.add_argument("command", choices=["analyze", "status", "help"],
                       help="Command to execute")
    parser.add_argument("--repo-owner", default="jdfalk", 
                       help="Repository owner (default: jdfalk)")
    parser.add_argument("--repo-name", default="subtitle-manager",
                       help="Repository name (default: subtitle-manager)")
    parser.add_argument("--github-token", 
                       help="GitHub token (or set GITHUB_TOKEN env var)")
    parser.add_argument("--issue-number", type=int,
                       help="Analyze specific issue number")
    parser.add_argument("--min-priority", type=int, default=60,
                       help="Minimum priority score for high-priority processing")
    parser.add_argument("--output", choices=["summary", "detailed", "json"], 
                       default="summary", help="Output format")
    
    args = parser.parse_args()
    
    if args.command == "help":
        parser.print_help()
        return
    
    if args.command == "status":
        # Status command doesn't need GitHub token
        updates_dir = Path(".github/issue-updates")
        if updates_dir.exists():
            pending_updates = list(updates_dir.glob("*.json"))
            processed_dir = updates_dir / "processed"
            processed_updates = list(processed_dir.glob("*.json")) if processed_dir.exists() else []
            
            print("Issue Update Status:")
            print(f"  Pending updates: {len(pending_updates)}")
            print(f"  Processed updates: {len(processed_updates)}")
            
            if pending_updates:
                print("\nPending updates:")
                for update_file in pending_updates[:5]:  # Show first 5
                    try:
                        with open(update_file) as f:
                            data = json.load(f)
                        print(f"  - {update_file.name}: Issue #{data.get('issue_number', 'unknown')}")
                    except Exception:
                        print(f"  - {update_file.name}: (invalid format)")
                
                if len(pending_updates) > 5:
                    print(f"  ... and {len(pending_updates) - 5} more")
        else:
            print("No issue updates directory found")
        return
    
    # For analyze command, we need GitHub token and analyzer
    if not ANALYZER_AVAILABLE:
        logger.error(f"Issue analyzer not available: {IMPORT_ERROR}")
        logger.error("Please install required dependencies: pip install -r requirements.txt")
        sys.exit(1)
        
    github_token = args.github_token or os.getenv("GITHUB_TOKEN")
    if not github_token:
        logger.error("GitHub token required. Set GITHUB_TOKEN env var or use --github-token")
        sys.exit(1)
    
    # Initialize task integration
    task_integration = TaskIntegration()
    
    if args.command == "analyze":
        # Run analysis
        kwargs = {}
        if args.issue_number:
            kwargs['issue_number'] = args.issue_number
        if args.min_priority != 60:
            kwargs['min_priority'] = args.min_priority
        
        results = task_integration.run_issue_analysis(
            args.repo_owner, args.repo_name, github_token, **kwargs
        )
        
        if not results['success']:
            logger.error(f"Analysis failed: {results.get('error', 'Unknown error')}")
            sys.exit(1)
        
        # Output results
        if args.output == "json":
            print(json.dumps(results, indent=2))
        elif args.output == "detailed":
            _print_detailed_results(results)
        else:
            _print_summary_results(results)
    
    elif args.command == "status":
        # Show current status
        updates_dir = Path(".github/issue-updates")
        if updates_dir.exists():
            pending_updates = list(updates_dir.glob("*.json"))
            processed_dir = updates_dir / "processed"
            processed_updates = list(processed_dir.glob("*.json")) if processed_dir.exists() else []
            
            print("Issue Update Status:")
            print(f"  Pending updates: {len(pending_updates)}")
            print(f"  Processed updates: {len(processed_updates)}")
            
            if pending_updates:
                print("\nPending updates:")
                for update_file in pending_updates[:5]:  # Show first 5
                    try:
                        with open(update_file) as f:
                            data = json.load(f)
                        print(f"  - {update_file.name}: Issue #{data.get('issue_number', 'unknown')}")
                    except Exception:
                        print(f"  - {update_file.name}: (invalid format)")
                
                if len(pending_updates) > 5:
                    print(f"  ... and {len(pending_updates) - 5} more")
        else:
            print("No issue updates directory found")

def _print_summary_results(results: Dict[str, Any]):
    """Print summary format results"""
    print("\n=== Issue Analysis Summary ===")
    print(f"Status: {'✅ Success' if results['success'] else '❌ Failed'}")
    
    if results.get('analysis_type') == 'single_issue':
        analysis = results['analysis_summary']
        print(f"Issue #{analysis['issue_number']}: {analysis['title']}")
        print(f"Priority Score: {analysis['priority_score']} ({analysis['urgency_level']})")
        print(f"Recommendations: {len(analysis['automation_recommendations'])}")
        
    elif results.get('analysis_type') == 'repository_scan':
        stats = results.get('summary_stats', {})
        print(f"Total Issues Analyzed: {results['total_issues_analyzed']}")
        print(f"High Priority Issues: {results['high_priority_count']}")
        print(f"Automation Updates Created: {len(results['automation_updates'])}")
        
        if stats and stats.get('needs_immediate_attention', 0) > 0:
            print(f"⚠️  Critical Issues Requiring Immediate Attention: {stats['needs_immediate_attention']}")
        
        if results['high_priority_issues']:
            print("\nTop Priority Issue:")
            top_issue = max(results['high_priority_issues'], key=lambda x: x['priority_score'])
            print(f"  #{top_issue['number']}: {top_issue['title']}")
            print(f"  Priority: {top_issue['priority_score']} ({top_issue['urgency_level']})")

def _print_detailed_results(results: Dict[str, Any]):
    """Print detailed format results"""
    print(json.dumps(results, indent=2))

if __name__ == "__main__":
    main()
