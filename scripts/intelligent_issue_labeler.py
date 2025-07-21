#!/usr/bin/env python3
# file: scripts/intelligent_issue_labeler.py
# version: 1.0.0
# guid: 4e5f6a7b-8c9d-0e1f-2a3b-4c5d6e7f8a9b

"""
Intelligent Issue Labeler for GitHub repositories.
Analyzes GitHub issues and automatically applies appropriate labels using pattern matching and AI fallback.

This script fixes the issues with the inline GitHub Actions version:
1. Properly handles NLTK data requirements
2. Simplified pattern matching without heavy NLP dependencies
3. Better error handling and logging
4. More efficient processing
"""

import argparse
import json
import os
import re
import sys
from collections import defaultdict
from typing import Dict, List

import requests
import yaml

try:
    import openai

    OPENAI_AVAILABLE = True
except ImportError:
    OPENAI_AVAILABLE = False
    print("Warning: OpenAI not available. AI fallback will be disabled.")


class IntelligentIssueLabeler:
    def __init__(self, github_token: str, repo: str, config: Dict):
        self.token = github_token
        self.repo = repo
        self.config = config
        self.session = requests.Session()
        self.session.headers.update(
            {
                "Authorization": f"token {github_token}",
                "Accept": "application/vnd.github.v3+json",
            }
        )

        # Statistics tracking
        self.stats = {
            "issues_processed": 0,
            "labels_applied": 0,
            "ai_fallback_used": 0,
            "patterns_matched": 0,
            "errors": 0,
        }

        # Initialize OpenAI if available and configured
        self.openai_client = None
        if OPENAI_AVAILABLE and os.getenv("OPENAI_API_KEY"):
            try:
                self.openai_client = openai.OpenAI(api_key=os.getenv("OPENAI_API_KEY"))
            except Exception as e:
                print(f"Warning: Failed to initialize OpenAI client: {e}")

    def get_repository_labels(self) -> Dict[str, Dict]:
        """Fetch all available labels from the repository."""
        try:
            response = self.session.get(
                f"https://api.github.com/repos/{self.repo}/labels?per_page=100"
            )
            response.raise_for_status()

            labels = {}
            for label in response.json():
                labels[label["name"]] = {
                    "color": label["color"],
                    "description": label.get("description", ""),
                }
            return labels
        except Exception as e:
            print(f"Error fetching repository labels: {e}")
            return {}

    def get_issues_to_process(self, batch_size: int) -> List[Dict]:
        """Get issues that need intelligent labeling."""
        try:
            # Get all issues with pagination
            all_issues = []
            page = 1
            per_page = min(100, batch_size)  # GitHub API max is 100 per page

            while len(all_issues) < batch_size:
                params = {
                    "state": "all",
                    "per_page": per_page,
                    "page": page,
                    "sort": "updated",
                    "direction": "desc",
                }

                print(f"  Fetching page {page} (per_page={per_page})...")
                response = self.session.get(
                    f"https://api.github.com/repos/{self.repo}/issues", params=params
                )
                response.raise_for_status()

                page_data = response.json()
                print(f"  Got {len(page_data)} items from API...")

                if not page_data:  # No more issues
                    break

                # Filter out pull requests
                page_issues = []
                for issue in page_data:
                    if "pull_request" not in issue:
                        page_issues.append(issue)

                print(f"  Found {len(page_issues)} actual issues (non-PRs)")
                all_issues.extend(page_issues)

                # If we got less than per_page, we're at the end
                if len(page_data) < per_page:
                    break

                page += 1

            print(f"Total issues collected: {len(all_issues)}")
            return all_issues[:batch_size]  # Limit to requested batch size

        except Exception as e:
            print(f"Error fetching issues: {e}")
            return []

    def simple_tokenize(self, text: str) -> List[str]:
        """Simple tokenization without NLTK dependencies."""
        if not text:
            return []

        # Convert to lowercase and split on non-alphanumeric characters
        tokens = re.findall(r"\b[a-zA-Z]+\b", text.lower())

        # Simple stop words list
        stop_words = {
            "the",
            "a",
            "an",
            "and",
            "or",
            "but",
            "in",
            "on",
            "at",
            "to",
            "for",
            "of",
            "with",
            "by",
            "this",
            "that",
            "these",
            "those",
            "i",
            "you",
            "he",
            "she",
            "it",
            "we",
            "they",
            "is",
            "are",
            "was",
            "were",
            "be",
            "been",
            "being",
            "have",
            "has",
            "had",
            "do",
            "does",
            "did",
            "will",
            "would",
            "could",
            "should",
            "may",
            "might",
            "can",
            "must",
            "shall",
        }

        # Filter out stop words and short tokens
        return [token for token in tokens if len(token) > 2 and token not in stop_words]

    def extract_features(self, issue: Dict) -> Dict[str, any]:
        """Extract features from issue for intelligent labeling."""
        title = issue.get("title", "").lower()
        body = issue.get("body", "").lower() if issue.get("body") else ""

        # Combine title and body for analysis
        text_content = f"{title} {body}"

        # Simple tokenization
        tokens = self.simple_tokenize(text_content)

        features = {
            "title": title,
            "body": body,
            "text_content": text_content,
            "tokens": tokens,
            "labels": [label["name"] for label in issue.get("labels", [])],
            "author": issue.get("user", {}).get("login", ""),
            "state": issue.get("state", ""),
            "created_at": issue.get("created_at", ""),
            "updated_at": issue.get("updated_at", ""),
        }

        return features

    def analyze_content_patterns(self, features: Dict) -> Dict[str, float]:
        """Analyze content patterns to suggest labels with confidence scores."""
        suggestions = defaultdict(float)
        text = features["text_content"]
        title = features["title"]

        # Issue type detection
        if any(
            word in text
            for word in [
                "bug",
                "error",
                "fail",
                "broken",
                "issue",
                "problem",
                "crash",
                "exception",
            ]
        ):
            suggestions["bug"] += 0.8
        if any(
            word in text
            for word in ["feature", "enhancement", "improve", "add", "new", "implement"]
        ):
            suggestions["enhancement"] += 0.8
        if any(
            word in text
            for word in ["doc", "documentation", "readme", "guide", "manual"]
        ):
            suggestions["documentation"] += 0.9
        if any(
            word in text
            for word in ["question", "help", "how", "why", "what", "clarification"]
        ):
            suggestions["question"] += 0.7

        # Priority detection
        if any(
            word in text
            for word in ["urgent", "critical", "blocker", "asap", "emergency"]
        ):
            suggestions["priority-high"] += 0.9
        elif any(word in text for word in ["low priority", "minor", "nice to have"]):
            suggestions["priority-low"] += 0.8
        else:
            suggestions["priority-medium"] += 0.6

        # Technology detection
        tech_patterns = {
            "tech:go": ["go", "golang", ".go", "gofmt"],
            "tech:python": ["python", ".py", "pip", "pytest"],
            "tech:javascript": ["javascript", "js", ".js", "npm", "node"],
            "tech:typescript": ["typescript", "ts", ".ts"],
            "tech:protobuf": ["protobuf", "proto", ".proto", "grpc"],
            "tech:docker": ["docker", "dockerfile", "container"],
            "tech:kubernetes": ["kubernetes", "k8s", "kubectl", "helm"],
            "tech:shell": ["bash", "shell", ".sh", "script"],
        }

        for label, keywords in tech_patterns.items():
            if any(keyword in text for keyword in keywords):
                suggestions[label] += 0.85
                self.stats["patterns_matched"] += 1

        # Module detection
        module_patterns = {
            "module:auth": ["auth", "authentication", "login", "password", "oauth"],
            "module:cache": ["cache", "redis", "memcache", "caching"],
            "module:config": ["config", "configuration", "settings", "env"],
            "module:database": ["database", "db", "sql", "mysql", "postgres"],
            "module:metrics": ["metrics", "monitoring", "prometheus", "grafana"],
            "module:queue": ["queue", "job", "worker", "task", "background"],
            "module:web": ["web", "http", "server", "api", "rest"],
            "module:ui": ["ui", "interface", "frontend", "html", "css"],
        }

        for label, keywords in module_patterns.items():
            if any(keyword in text for keyword in keywords):
                suggestions[label] += 0.8
                self.stats["patterns_matched"] += 1

        # Workflow detection
        workflow_patterns = {
            "workflow:automation": ["automation", "script", "workflow", "ci/cd"],
            "workflow:github-actions": ["github actions", "workflow", ".yml", "action"],
            "workflow:deployment": ["deploy", "deployment", "release", "production"],
            "github-actions": ["github actions", "workflow", "action", ".github"],
        }

        for label, keywords in workflow_patterns.items():
            if any(keyword in text for keyword in keywords):
                suggestions[label] += 0.8
                self.stats["patterns_matched"] += 1

        # Security and performance
        if any(
            word in text for word in ["security", "vulnerability", "cve", "exploit"]
        ):
            suggestions["security"] += 0.9
        if any(word in text for word in ["performance", "slow", "optimize", "speed"]):
            suggestions["performance"] += 0.8
        if any(word in text for word in ["breaking", "break", "compatibility"]):
            suggestions["breaking-change"] += 0.8

        # Special labels
        if any(word in title for word in ["good first issue", "beginner", "starter"]):
            suggestions["good first issue"] += 0.9
        if any(word in text for word in ["help wanted", "need help", "assistance"]):
            suggestions["help wanted"] += 0.8

        return dict(suggestions)

    def use_ai_fallback(
        self, issue: Dict, initial_suggestions: Dict[str, float]
    ) -> Dict[str, float]:
        """Use AI (OpenAI) as fallback for complex labeling decisions."""
        if not self.openai_client:
            print("AI fallback not available (OpenAI not configured)")
            return initial_suggestions

        try:
            # Get comprehensive label list from repository
            repo_labels = self.get_repository_labels()
            label_descriptions = []
            for name, info in repo_labels.items():
                desc = info.get("description", "")
                label_descriptions.append(f"- {name}: {desc}")

            # Limit content to avoid token limits
            title = issue.get("title", "")
            body = issue.get("body", "")[:1000] if issue.get("body") else ""

            prompt = f"""
Analyze this GitHub issue and suggest appropriate labels from the available list.

Issue Title: {title}
Issue Body: {body}

Available Labels:
{chr(10).join(label_descriptions[:50])}

Current Suggestions: {", ".join(initial_suggestions.keys())}

Please suggest 3-5 most relevant labels with confidence scores (0.0-1.0).
Respond in JSON format: {{"label_name": confidence_score}}
"""

            response = self.openai_client.chat.completions.create(
                model="gpt-4o-mini",
                messages=[
                    {
                        "role": "system",
                        "content": "You are an expert at analyzing GitHub issues and applying appropriate labels. Return only valid JSON.",
                    },
                    {"role": "user", "content": prompt},
                ],
                max_tokens=500,
                temperature=0.3,
            )

            ai_suggestions = json.loads(response.choices[0].message.content)

            # Merge AI suggestions with initial suggestions
            combined_suggestions = initial_suggestions.copy()
            for label, confidence in ai_suggestions.items():
                if label in repo_labels:  # Only use valid labels
                    combined_suggestions[label] = max(
                        combined_suggestions.get(label, 0), confidence
                    )

            self.stats["ai_fallback_used"] += 1
            return combined_suggestions

        except Exception as e:
            print(f"AI fallback failed: {e}")
            self.stats["errors"] += 1
            return initial_suggestions

    def apply_labels_to_issue(
        self, issue_number: int, suggested_labels: List[str], current_labels: List[str]
    ) -> bool:
        """Apply suggested labels to an issue."""
        try:
            if self.config.get("preserve_existing_labels", True):
                # Combine existing and new labels, removing duplicates
                all_labels = list(set(current_labels + suggested_labels))
            else:
                all_labels = suggested_labels

            if self.config.get("dry_run", False):
                print(
                    f"DRY RUN: Would apply labels {suggested_labels} to issue #{issue_number}"
                )
                return True

            data = {"labels": all_labels}
            response = self.session.patch(
                f"https://api.github.com/repos/{self.repo}/issues/{issue_number}",
                json=data,
            )
            response.raise_for_status()

            self.stats["labels_applied"] += len(suggested_labels)
            return True

        except Exception as e:
            print(f"Error applying labels to issue #{issue_number}: {e}")
            self.stats["errors"] += 1
            return False

    def process_issues(self):
        """Main processing function."""
        issues = self.get_issues_to_process(self.config.get("batch_size", 10))
        print(f"Processing {len(issues)} issues for intelligent labeling...")

        repo_labels = self.get_repository_labels()
        confidence_threshold = self.config.get("confidence_threshold", 0.7)
        max_labels = self.config.get("max_labels_per_issue", 8)
        use_ai = self.config.get("use_ai_fallback", True)

        for issue in issues:
            issue_number = issue["number"]
            print(f"\nProcessing issue #{issue_number}: {issue['title'][:60]}...")

            try:
                # Extract features
                features = self.extract_features(issue)
                current_labels = features["labels"]

                # Get initial suggestions from pattern analysis
                suggestions = self.analyze_content_patterns(features)

                # Use AI fallback if enabled and we have few suggestions
                if use_ai and len(suggestions) < 3:
                    suggestions = self.use_ai_fallback(issue, suggestions)

                # Filter suggestions by confidence threshold and max labels
                filtered_suggestions = {
                    label: confidence
                    for label, confidence in suggestions.items()
                    if confidence >= confidence_threshold and label in repo_labels
                }

                # Sort by confidence and limit
                sorted_suggestions = sorted(
                    filtered_suggestions.items(), key=lambda x: x[1], reverse=True
                )[:max_labels]

                suggested_labels = [label for label, _ in sorted_suggestions]

                # Only apply labels that aren't already present
                new_labels = [
                    label for label in suggested_labels if label not in current_labels
                ]

                if new_labels:
                    print(f"  Suggested labels: {new_labels}")
                    if self.apply_labels_to_issue(
                        issue_number, new_labels, current_labels
                    ):
                        print(f"  ✅ Applied {len(new_labels)} new labels")
                    else:
                        print("  ❌ Failed to apply labels")
                else:
                    print("  ⏭️ No new labels needed")

                self.stats["issues_processed"] += 1

            except Exception as e:
                print(f"  ❌ Error processing issue #{issue_number}: {e}")
                self.stats["errors"] += 1

        return self.stats


def load_config_file(config_path: str) -> Dict:
    """Load configuration from YAML file."""
    try:
        if os.path.exists(config_path):
            with open(config_path, "r") as f:
                return yaml.safe_load(f)
        else:
            print(f"Configuration file {config_path} not found, using defaults")
            return {}
    except Exception as e:
        print(f"Error loading configuration file: {e}")
        return {}


def main():
    parser = argparse.ArgumentParser(description="Intelligent Issue Labeler")
    parser.add_argument("--repo", required=True, help="GitHub repository (owner/name)")
    parser.add_argument("--token", help="GitHub token (or set GITHUB_TOKEN env var)")
    parser.add_argument("--dry-run", action="store_true", help="Run in dry-run mode")
    parser.add_argument(
        "--batch-size", type=int, default=10, help="Number of issues to process"
    )
    parser.add_argument(
        "--confidence-threshold", type=float, default=0.7, help="Confidence threshold"
    )
    parser.add_argument(
        "--max-labels", type=int, default=8, help="Max labels per issue"
    )
    parser.add_argument(
        "--config", default=".github/intelligent-labeling.yml", help="Config file path"
    )
    parser.add_argument(
        "--use-ai", action="store_true", default=True, help="Use AI fallback"
    )
    parser.add_argument(
        "--preserve-existing",
        action="store_true",
        default=True,
        help="Preserve existing labels",
    )

    args = parser.parse_args()

    # Get GitHub token
    github_token = args.token or os.getenv("GITHUB_TOKEN")
    if not github_token:
        print("Error: GitHub token is required (--token or GITHUB_TOKEN env var)")
        sys.exit(1)

    # Load configuration
    file_config = load_config_file(args.config)

    # Merge command line args with file config (CLI takes precedence)
    config = {
        "dry_run": args.dry_run,
        "batch_size": args.batch_size,
        "use_ai_fallback": args.use_ai,
        "confidence_threshold": args.confidence_threshold,
        "max_labels_per_issue": args.max_labels,
        "preserve_existing_labels": args.preserve_existing,
    }

    # Override with file config where not specified on command line
    for key, value in file_config.get("global", {}).items():
        if key not in config:
            config[key] = value

    print("🤖 Starting Intelligent Issue Labeling...")
    print(f"Repository: {args.repo}")
    print(f"Configuration: {config}")

    labeler = IntelligentIssueLabeler(github_token, args.repo, config)
    stats = labeler.process_issues()

    print("\n📊 Labeling Statistics:")
    print(f"  Issues processed: {stats['issues_processed']}")
    print(f"  Labels applied: {stats['labels_applied']}")
    print(f"  Patterns matched: {stats['patterns_matched']}")
    print(f"  AI fallback used: {stats['ai_fallback_used']} times")
    print(f"  Errors: {stats['errors']}")

    if stats["errors"] > 0:
        sys.exit(1)


if __name__ == "__main__":
    main()
