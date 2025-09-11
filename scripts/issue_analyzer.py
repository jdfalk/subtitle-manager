#!/usr/bin/env python3
# file: scripts/issue_analyzer.py
# version: 1.0.0
# guid: a1b2c3d4-e5f6-7890-abcd-ef1234567890

"""
GitHub Issue Analysis and Automation System

This module enhances the existing issue management infrastructure to provide:
1. Automated issue analysis and categorization
2. Priority scoring based on multiple factors
3. Label and milestone automation
4. Conflict detection and resolution

Works with existing:
- issue_manager.py: Core issue management operations
- enhanced_issue_manager.py: Enhanced timestamp tracking
- .github/issue-updates/: JSON-based update system
- VS Code tasks: Automated workflows
"""

import json
import logging
import os
import sys
from datetime import datetime
from pathlib import Path
from typing import Dict, List, Any
import uuid

# GitHub API
import requests

# Enhanced logging
logging.basicConfig(
    level=logging.INFO, format="%(asctime)s - %(name)s - %(levelname)s - %(message)s"
)
logger = logging.getLogger(__name__)


class IssueAnalyzer:
    """
    Analyzes GitHub issues for automated categorization and priority scoring.

    Integrates with existing issue management system to provide enhanced
    automation and intelligent processing.
    """

    def __init__(self, repo_owner: str, repo_name: str, github_token: str):
        self.repo_owner = repo_owner
        self.repo_name = repo_name
        self.github_token = github_token
        self.base_url = f"https://api.github.com/repos/{repo_owner}/{repo_name}"
        self.headers = {
            "Authorization": f"token {github_token}",
            "Accept": "application/vnd.github.v3+json",
        }

        # Load repository labels and configuration
        self.labels = self._load_labels()
        self.priority_weights = self._load_priority_config()

    def _load_labels(self) -> Dict[str, Dict]:
        """Load available repository labels from labels.json"""
        try:
            labels_file = Path("labels.json")
            if labels_file.exists():
                with open(labels_file, "r") as f:
                    labels_data = json.load(f)
                    return {label["name"]: label for label in labels_data}
            return {}
        except Exception as e:
            logger.warning(f"Could not load labels.json: {e}")
            return {}

    def _load_priority_config(self) -> Dict[str, int]:
        """Load priority scoring configuration"""
        return {
            # Bug severity multipliers
            "critical_bug": 100,
            "high_bug": 75,
            "medium_bug": 50,
            "low_bug": 25,
            # Feature importance
            "breaking_change": 90,
            "enhancement": 40,
            "documentation": 20,
            # Module criticality
            "module:auth": 80,
            "module:database": 70,
            "module:web": 60,
            "module:config": 50,
            "module:ui": 30,
            # Activity indicators
            "recent_activity": 20,
            "multiple_comments": 15,
            "long_open": -10,
            # Special flags
            "security": 95,
            "performance": 60,
            "refactor": 35,
        }

    def analyze_issue(self, issue_data: Dict) -> Dict[str, Any]:
        """
        Analyze a single issue and return categorization and priority data.

        Args:
            issue_data: GitHub issue data from API

        Returns:
            Analysis results including priority score, suggested labels, etc.
        """
        title = issue_data.get("title", "").lower()
        body = issue_data.get("body", "").lower()
        existing_labels = [label["name"] for label in issue_data.get("labels", [])]

        analysis = {
            "issue_number": issue_data["number"],
            "title": issue_data["title"],
            "current_labels": existing_labels,
            "priority_score": 0,
            "priority_factors": [],
            "suggested_labels": [],
            "suggested_module": None,
            "issue_type": None,
            "urgency_level": "medium",
            "automation_recommendations": [],
        }

        # Analyze issue type and content
        analysis.update(self._analyze_content(title, body))

        # Calculate priority score
        analysis["priority_score"] = self._calculate_priority_score(
            issue_data, analysis, existing_labels
        )

        # Determine urgency level
        analysis["urgency_level"] = self._determine_urgency(analysis["priority_score"])

        # Generate automation recommendations
        analysis["automation_recommendations"] = self._generate_recommendations(
            issue_data, analysis
        )

        return analysis

    def _analyze_content(self, title: str, body: str) -> Dict[str, Any]:
        """Analyze issue title and body for keywords and patterns"""
        content = f"{title} {body}"

        # Issue type detection
        issue_type = "unknown"
        if any(word in content for word in ["bug", "error", "fail", "broken", "crash"]):
            issue_type = "bug"
        elif any(
            word in content for word in ["feature", "enhancement", "add", "implement"]
        ):
            issue_type = "enhancement"
        elif any(word in content for word in ["test", "testing", "spec"]):
            issue_type = "test"
        elif any(word in content for word in ["doc", "documentation", "readme"]):
            issue_type = "documentation"
        elif any(word in content for word in ["refactor", "cleanup", "improve"]):
            issue_type = "refactor"

        # Module detection
        module = None
        module_keywords = {
            "auth": ["auth", "login", "session", "token", "oauth"],
            "database": ["database", "db", "sql", "migration", "schema"],
            "web": ["api", "http", "rest", "grpc", "server", "endpoint"],
            "ui": ["ui", "web", "frontend", "interface", "layout"],
            "config": ["config", "configuration", "settings", "yaml"],
            "cache": ["cache", "caching", "redis", "memory"],
            "queue": ["queue", "job", "worker", "task"],
            "metrics": ["metrics", "prometheus", "monitoring"],
        }

        for mod, keywords in module_keywords.items():
            if any(keyword in content for keyword in keywords):
                module = f"module:{mod}"
                break

        # Suggested labels based on content
        suggested_labels = []
        if issue_type != "unknown":
            suggested_labels.append(issue_type)
        if module:
            suggested_labels.append(module)

        # Security detection
        if any(
            word in content for word in ["security", "vulnerability", "cve", "exploit"]
        ):
            suggested_labels.append("security")

        # Performance detection
        if any(
            word in content for word in ["performance", "slow", "optimization", "speed"]
        ):
            suggested_labels.append("performance")

        return {
            "issue_type": issue_type,
            "suggested_module": module,
            "suggested_labels": suggested_labels,
        }

    def _calculate_priority_score(
        self, issue_data: Dict, analysis: Dict, existing_labels: List[str]
    ) -> int:
        """Calculate priority score based on multiple factors"""
        score = 0
        factors = []

        # Base score from issue type
        issue_type = analysis.get("issue_type", "unknown")
        if issue_type == "bug":
            score += self.priority_weights.get("medium_bug", 50)
            factors.append(f"Bug type: +{self.priority_weights.get('medium_bug', 50)}")
        elif issue_type == "enhancement":
            score += self.priority_weights.get("enhancement", 40)
            factors.append(
                f"Enhancement: +{self.priority_weights.get('enhancement', 40)}"
            )

        # Module importance
        module = analysis.get("suggested_module")
        if module and module in self.priority_weights:
            module_score = self.priority_weights[module]
            score += module_score
            factors.append(f"Module {module}: +{module_score}")

        # Existing label analysis
        for label in existing_labels:
            if label in self.priority_weights:
                label_score = self.priority_weights[label]
                score += label_score
                factors.append(f"Label {label}: +{label_score}")

        # Activity factors
        created_date = datetime.fromisoformat(
            issue_data["created_at"].replace("Z", "+00:00")
        )
        age_days = (datetime.now().astimezone() - created_date).days

        if age_days <= 7:
            score += self.priority_weights.get("recent_activity", 20)
            factors.append("Recent activity: +20")
        elif age_days > 90:
            score += self.priority_weights.get("long_open", -10)
            factors.append("Long open: -10")

        # Comment activity
        comments = issue_data.get("comments", 0)
        if comments > 5:
            score += self.priority_weights.get("multiple_comments", 15)
            factors.append("Multiple comments: +15")

        analysis["priority_factors"] = factors
        return max(0, score)  # Ensure non-negative score

    def _determine_urgency(self, priority_score: int) -> str:
        """Determine urgency level from priority score"""
        if priority_score >= 80:
            return "critical"
        elif priority_score >= 60:
            return "high"
        elif priority_score >= 40:
            return "medium"
        else:
            return "low"

    def _generate_recommendations(self, issue_data: Dict, analysis: Dict) -> List[str]:
        """Generate automation recommendations"""
        recommendations = []

        # Label recommendations
        current_labels = set(analysis["current_labels"])
        suggested_labels = set(analysis["suggested_labels"])
        missing_labels = suggested_labels - current_labels

        if missing_labels:
            recommendations.append(f"Add labels: {', '.join(missing_labels)}")

        # Priority label recommendation
        urgency = analysis["urgency_level"]
        priority_label = f"priority:{urgency}"
        if priority_label not in current_labels and priority_label in self.labels:
            recommendations.append(f"Add priority label: {priority_label}")

        # Module assignment
        if (
            analysis["suggested_module"]
            and analysis["suggested_module"] not in current_labels
        ):
            recommendations.append(f"Assign to module: {analysis['suggested_module']}")

        # Special handling recommendations
        if analysis["priority_score"] >= 80:
            recommendations.append("Consider immediate attention due to high priority")

        if "security" in analysis["suggested_labels"]:
            recommendations.append("Flag for security review")

        return recommendations


class IssueAutomation:
    """
    Automated issue processing using existing infrastructure.

    Creates issue updates in the .github/issue-updates/ format
    for processing by existing workflows.
    """

    def __init__(self, repo_path: str):
        self.repo_path = Path(repo_path)
        self.updates_dir = self.repo_path / ".github" / "issue-updates"
        self.updates_dir.mkdir(parents=True, exist_ok=True)

    def create_issue_update(self, issue_number: int, analysis: Dict[str, Any]) -> str:
        """
        Create an issue update JSON file based on analysis results.

        Uses the existing issue update format from enhanced_issue_manager.py
        """
        update_id = str(uuid.uuid4())
        timestamp = datetime.now().isoformat()

        # Determine actions based on analysis
        actions = []

        # Add suggested labels
        if analysis["automation_recommendations"]:
            for rec in analysis["automation_recommendations"]:
                if rec.startswith("Add labels:"):
                    labels_str = rec.replace("Add labels: ", "")
                    labels = [label.strip() for label in labels_str.split(",")]
                    actions.append(
                        {
                            "action": "update",
                            "labels": analysis["current_labels"] + labels,
                        }
                    )
                elif rec.startswith("Add priority label:"):
                    priority_label = rec.replace("Add priority label: ", "")
                    actions.append(
                        {
                            "action": "update",
                            "labels": analysis["current_labels"] + [priority_label],
                        }
                    )

        # Create update structure
        update_data = {
            "guid": update_id,
            "timestamp": timestamp,
            "issue_number": issue_number,
            "analysis": analysis,
            "recommended_actions": actions,
            "automation_applied": False,
            "version": "2.0",
        }

        # Write to updates directory
        update_file = self.updates_dir / f"{update_id}.json"
        with open(update_file, "w") as f:
            json.dump(update_data, f, indent=2)

        logger.info(f"Created issue update: {update_file}")
        return update_id

    def process_high_priority_issues(
        self, analyzer: IssueAnalyzer, min_priority: int = 60
    ) -> List[str]:
        """
        Process high-priority issues and create automation updates.

        Args:
            analyzer: IssueAnalyzer instance
            min_priority: Minimum priority score for processing

        Returns:
            List of created update IDs
        """
        # Get open issues
        issues_url = f"{analyzer.base_url}/issues"
        params = {"state": "open", "per_page": 100}

        response = requests.get(issues_url, headers=analyzer.headers, params=params)
        if response.status_code != 200:
            logger.error(f"Failed to fetch issues: {response.status_code}")
            return []

        issues = response.json()
        update_ids = []

        for issue in issues:
            if issue.get("pull_request"):  # Skip PRs
                continue

            analysis = analyzer.analyze_issue(issue)

            if analysis["priority_score"] >= min_priority:
                logger.info(
                    f"Processing high-priority issue #{issue['number']}: "
                    f"{issue['title']} (score: {analysis['priority_score']})"
                )

                update_id = self.create_issue_update(issue["number"], analysis)
                update_ids.append(update_id)

        return update_ids


def main():
    """Main CLI entry point for issue analysis"""
    import argparse

    parser = argparse.ArgumentParser(description="GitHub Issue Analysis and Automation")
    parser.add_argument("--repo-owner", required=True, help="Repository owner")
    parser.add_argument("--repo-name", required=True, help="Repository name")
    parser.add_argument(
        "--github-token", help="GitHub token (or set GITHUB_TOKEN env var)"
    )
    parser.add_argument(
        "--analyze-issue", type=int, help="Analyze specific issue number"
    )
    parser.add_argument(
        "--process-high-priority",
        action="store_true",
        help="Process all high-priority issues",
    )
    parser.add_argument(
        "--min-priority",
        type=int,
        default=60,
        help="Minimum priority score for processing",
    )
    parser.add_argument(
        "--output-format",
        choices=["json", "table"],
        default="table",
        help="Output format",
    )

    args = parser.parse_args()

    # Get GitHub token
    github_token = args.github_token or os.getenv("GITHUB_TOKEN")
    if not github_token:
        logger.error(
            "GitHub token required. Set GITHUB_TOKEN env var or use --github-token"
        )
        sys.exit(1)

    # Initialize analyzer
    analyzer = IssueAnalyzer(args.repo_owner, args.repo_name, github_token)
    automation = IssueAutomation(".")

    if args.analyze_issue:
        # Analyze specific issue
        issue_url = f"{analyzer.base_url}/issues/{args.analyze_issue}"
        response = requests.get(issue_url, headers=analyzer.headers)

        if response.status_code != 200:
            logger.error(f"Failed to fetch issue #{args.analyze_issue}")
            sys.exit(1)

        issue = response.json()
        analysis = analyzer.analyze_issue(issue)

        if args.output_format == "json":
            print(json.dumps(analysis, indent=2))
        else:
            print(f"\n=== Issue #{analysis['issue_number']}: {analysis['title']} ===")
            print(
                f"Priority Score: {analysis['priority_score']} ({analysis['urgency_level']})"
            )
            print(f"Issue Type: {analysis['issue_type']}")
            print(f"Current Labels: {', '.join(analysis['current_labels'])}")
            print(f"Suggested Labels: {', '.join(analysis['suggested_labels'])}")
            print("\nPriority Factors:")
            for factor in analysis["priority_factors"]:
                print(f"  - {factor}")
            print("\nRecommendations:")
            for rec in analysis["automation_recommendations"]:
                print(f"  - {rec}")

    elif args.process_high_priority:
        # Process high-priority issues
        update_ids = automation.process_high_priority_issues(
            analyzer, args.min_priority
        )
        print(f"Created {len(update_ids)} issue updates for high-priority issues")
        for update_id in update_ids:
            print(f"  - {update_id}")

    else:
        parser.print_help()


if __name__ == "__main__":
    main()
