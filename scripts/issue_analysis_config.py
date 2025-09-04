# file: scripts/issue_analysis_config.py
# version: 1.0.0
# guid: c3d4e5f6-a7b8-9012-cdef-345678901234

"""
Configuration for GitHub Issue Analysis System

Defines priority weights, module mappings, and automation rules
for the issue analysis and automation system.
"""

# Priority scoring weights
PRIORITY_WEIGHTS = {
    # Bug severity multipliers
    "critical_bug": 100,
    "high_bug": 75,
    "medium_bug": 50,
    "low_bug": 25,
    # Issue type base scores
    "bug": 50,
    "enhancement": 40,
    "documentation": 20,
    "test": 30,
    "refactor": 35,
    "security": 95,
    "performance": 60,
    # Module criticality scores
    "module:auth": 80,
    "module:database": 70,
    "module:web": 60,
    "module:config": 50,
    "module:ui": 30,
    "module:cache": 45,
    "module:queue": 55,
    "module:metrics": 40,
    # Priority labels
    "priority:critical": 90,
    "priority:high": 70,
    "priority:medium": 50,
    "priority:low": 25,
    # Activity indicators
    "recent_activity": 20,
    "multiple_comments": 15,
    "long_open": -10,
    "stale": -20,
    # Special categories
    "breaking_change": 90,
    "codex": 5,  # AI-generated issues get slight boost
    "needs-triage": 10,
}

# Module detection keywords
MODULE_KEYWORDS = {
    "auth": [
        "auth",
        "login",
        "session",
        "token",
        "oauth",
        "authentication",
        "authorization",
    ],
    "database": ["database", "db", "sql", "migration", "schema", "sqlite", "postgres"],
    "web": ["api", "http", "rest", "grpc", "server", "endpoint", "handler"],
    "ui": ["ui", "web", "frontend", "interface", "layout", "webui", "react"],
    "config": ["config", "configuration", "settings", "yaml", "json"],
    "cache": ["cache", "caching", "redis", "memory"],
    "queue": ["queue", "job", "worker", "task", "background"],
    "metrics": ["metrics", "prometheus", "monitoring", "telemetry"],
}

# Issue type detection patterns
ISSUE_TYPE_PATTERNS = {
    "bug": ["bug", "error", "fail", "broken", "crash", "exception", "panic"],
    "enhancement": ["feature", "enhancement", "add", "implement", "improve"],
    "test": ["test", "testing", "spec", "coverage"],
    "documentation": ["doc", "documentation", "readme", "guide"],
    "refactor": ["refactor", "cleanup", "improve", "reorganize"],
    "security": ["security", "vulnerability", "cve", "exploit", "auth"],
    "performance": ["performance", "slow", "optimization", "speed", "memory"],
}

# Automation rules
AUTOMATION_RULES = {
    # Auto-assign priority labels
    "auto_priority_labels": True,
    # Auto-assign module labels
    "auto_module_labels": True,
    # Auto-assign type labels
    "auto_type_labels": True,
    # Minimum score for automatic processing
    "auto_process_threshold": 60,
    # Critical threshold for immediate attention
    "critical_threshold": 80,
    # Days before considering an issue stale
    "stale_days": 90,
    # Days to consider an issue "recent"
    "recent_days": 7,
    # Minimum comments for "active discussion"
    "active_comments_threshold": 5,
}

# Label mappings for standardization
LABEL_MAPPINGS = {
    # Standardize bug severity
    "bug-critical": "priority:critical",
    "bug-high": "priority:high",
    "bug-medium": "priority:medium",
    "bug-low": "priority:low",
    # Standardize feature priority
    "feature-high": "priority:high",
    "feature-medium": "priority:medium",
    "feature-low": "priority:low",
    # Module standardization
    "backend": "module:web",
    "frontend": "module:ui",
    "api": "module:web",
    "cli": "module:config",
}

# Repository-specific configuration
REPO_CONFIG = {
    "default_priority": "medium",
    "default_labels": ["needs-triage"],
    "critical_modules": ["auth", "database", "web"],
    "auto_close_stale": False,
    "auto_assign_reviewers": True,
}
