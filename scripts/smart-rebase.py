#!/usr/bin/env python3
# file: scripts/smart-rebase.py
# version: 1.0.0
# guid: 8a9b0c1d-2e3f-4a5b-6c7d-8e9f0a1b2c3d

"""
Smart Git Rebase Tool - Comprehensive Python Implementation

A comprehensive Git rebase automation tool with intelligent conflict resolution,
persistent state management, backup management, and detailed logging. This tool
combines the best features from multiple rebase implementations into a single
robust solution.

Features:
- Persistent state management with resume capability
- Intelligent conflict resolution based on file types and patterns
- Automatic backup branch creation and recovery
- Comprehensive logging and progress tracking
- Multiple operation modes (interactive, automated, smart)
- Recovery instructions and rollback capabilities
- Support for dry-run operations
- Real-time progress updates and conflict tracking
- Smart conflict resolution strategies with file pattern matching
- Automatic cleanup and summary generation
"""

import argparse
import json
import logging
import os
import re
import subprocess
import sys
import time
from dataclasses import dataclass, field
from datetime import datetime
from enum import Enum
from pathlib import Path
from typing import List


class RebaseMode(Enum):
    """Rebase operation modes"""

    INTERACTIVE = "interactive"
    AUTOMATED = "automated"
    SMART = "smart"


class ConflictStrategy(Enum):
    """Conflict resolution strategies"""

    PREFER_INCOMING = "prefer_incoming"
    PREFER_CURRENT = "prefer_current"
    SMART_MERGE = "smart_merge"
    MANUAL_REVIEW = "manual_review"
    SAVE_BOTH = "save_both"
    AUTO_RESOLVE = "auto_resolve"


class RebaseResult(Enum):
    """Rebase operation results"""

    SUCCESS = "success"
    CONFLICTS = "conflicts"
    FAILED = "failed"
    ABORTED = "aborted"
    RESUMED = "resumed"
    IN_PROGRESS = "in_progress"


class RebasePhase(Enum):
    """Rebase operation phases"""

    INIT = "init"
    PREREQUISITES = "prerequisites"
    BACKUP = "backup"
    REBASE_START = "rebase_start"
    REBASE_IN_PROGRESS = "rebase_in_progress"
    CONFLICT_RESOLUTION = "conflict_resolution"
    REBASE_CONTINUE = "rebase_continue"
    PUSH = "push"
    CLEANUP = "cleanup"
    COMPLETE = "complete"


@dataclass
class ConflictFile:
    """Represents a file with merge conflicts"""

    path: str
    strategy: ConflictStrategy
    resolved: bool = False
    resolution_method: str = ""
    backup_created: bool = False
    error_message: str = ""


@dataclass
class RebaseState:
    """Represents the current state of a rebase operation"""

    phase: RebasePhase = RebasePhase.INIT
    step: str = ""
    current_branch: str = ""
    target_branch: str = ""
    source_branch: str = ""
    backup_branch: str = ""
    timestamp: str = ""
    pid: int = 0
    total_commits: int = 0
    processed_commits: int = 0
    progress_percent: int = 0
    conflicts: List[ConflictFile] = field(default_factory=list)
    resolved_conflicts: int = 0
    force_push: bool = False
    dry_run: bool = False
    verbose: bool = False
    session_id: str = ""
    git_status: str = ""
    error_messages: List[str] = field(default_factory=list)
    recovery_instructions: List[str] = field(default_factory=list)


class GitRebaseError(Exception):
    """Custom exception for rebase operations"""

    pass


class SmartRebase:
    """
    Smart Git Rebase implementation with intelligent conflict resolution
    and persistent state management.
    """

    def __init__(self, verbose: bool = False, dry_run: bool = False):
        self.verbose = verbose
        self.dry_run = dry_run
        self.start_time = datetime.now()
        self.session_id = f"rebase_{int(time.time())}"

        # State management
        self.state_dir = Path.cwd() / ".rebase-state"
        self.backup_dir = Path.cwd() / ".rebase-backup"
        self.state_file = self.state_dir / "rebase.state"
        self.progress_file = self.state_dir / "progress.json"
        self.log_file = self.state_dir / "rebase.log"
        self.conflict_log = self.state_dir / "conflicts.log"
        self.summary_file = self.state_dir / "summary.md"

        # Initialize state
        self.state = RebaseState()
        self.state.session_id = self.session_id
        self.state.verbose = verbose
        self.state.dry_run = dry_run

        # Initialize directories and logging
        self._init_state_management()
        self._setup_logging()

        # File pattern mappings for conflict resolution
        self.conflict_patterns = {
            # Documentation files - prefer incoming
            r".*\.md$": ConflictStrategy.PREFER_INCOMING,
            r".*\.rst$": ConflictStrategy.PREFER_INCOMING,
            r".*\.txt$": ConflictStrategy.PREFER_INCOMING,
            r"docs/.*": ConflictStrategy.PREFER_INCOMING,
            r"README.*": ConflictStrategy.PREFER_INCOMING,
            r"CHANGELOG.*": ConflictStrategy.SMART_MERGE,
            r"TODO.*": ConflictStrategy.SMART_MERGE,
            # Build and CI files - prefer incoming
            r"\.github/.*": ConflictStrategy.PREFER_INCOMING,
            r"Dockerfile.*": ConflictStrategy.PREFER_INCOMING,
            r"docker-compose.*": ConflictStrategy.PREFER_INCOMING,
            r".*\.yml$": ConflictStrategy.PREFER_INCOMING,
            r".*\.yaml$": ConflictStrategy.PREFER_INCOMING,
            r"Makefile.*": ConflictStrategy.PREFER_INCOMING,
            # Package management files - smart merge
            r"go\.mod$": ConflictStrategy.SMART_MERGE,
            r"go\.sum$": ConflictStrategy.PREFER_INCOMING,
            r"package\.json$": ConflictStrategy.SMART_MERGE,
            r"package-lock\.json$": ConflictStrategy.PREFER_INCOMING,
            r"requirements\.txt$": ConflictStrategy.SMART_MERGE,
            r"Pipfile.*": ConflictStrategy.SMART_MERGE,
            # Source code files - manual review by default
            r".*\.go$": ConflictStrategy.AUTO_RESOLVE,
            r".*\.py$": ConflictStrategy.AUTO_RESOLVE,
            r".*\.js$": ConflictStrategy.AUTO_RESOLVE,
            r".*\.ts$": ConflictStrategy.AUTO_RESOLVE,
            r".*\.java$": ConflictStrategy.AUTO_RESOLVE,
            r".*\.cpp$": ConflictStrategy.AUTO_RESOLVE,
            r".*\.c$": ConflictStrategy.AUTO_RESOLVE,
            r".*\.h$": ConflictStrategy.AUTO_RESOLVE,
            # Test files - prefer current (keep local tests)
            r".*_test\.go$": ConflictStrategy.PREFER_CURRENT,
            r".*_test\.py$": ConflictStrategy.PREFER_CURRENT,
            r".*\.test\.js$": ConflictStrategy.PREFER_CURRENT,
            r"test/.*": ConflictStrategy.PREFER_CURRENT,
            r"tests/.*": ConflictStrategy.PREFER_CURRENT,
            # Configuration files - manual review
            r".*\.conf$": ConflictStrategy.MANUAL_REVIEW,
            r".*\.config$": ConflictStrategy.MANUAL_REVIEW,
            r".*\.ini$": ConflictStrategy.MANUAL_REVIEW,
            r".*\.toml$": ConflictStrategy.MANUAL_REVIEW,
        }

    def _init_state_management(self):
        """Initialize state management directories"""
        self.state_dir.mkdir(exist_ok=True)
        self.backup_dir.mkdir(exist_ok=True)

        # Initialize session info
        self.state.pid = os.getpid()
        self.state.timestamp = datetime.now().isoformat()

    def _setup_logging(self):
        """Setup logging configuration"""
        logging.basicConfig(
            level=logging.DEBUG if self.verbose else logging.INFO,
            format="%(asctime)s - %(levelname)s - %(message)s",
            handlers=[
                logging.FileHandler(self.log_file),
                logging.StreamHandler(sys.stdout),
            ],
        )
        self.logger = logging.getLogger(__name__)

    def log_info(self, message: str) -> None:
        """Log info message"""
        self.logger.info(message)

    def log_success(self, message: str) -> None:
        """Log success message"""
        self.logger.info(f"✓ {message}")

    def log_warning(self, message: str) -> None:
        """Log warning message"""
        self.logger.warning(f"⚠ {message}")

    def log_error(self, message: str) -> None:
        """Log error message"""
        self.logger.error(f"✗ {message}")
        self.state.error_messages.append(message)

    def log_verbose(self, message: str) -> None:
        """Log verbose message"""
        if self.verbose:
            self.logger.debug(f"🔍 {message}")

    def log_step(self, message: str) -> None:
        """Log step message"""
        self.logger.info(f"📋 {message}")

    def save_state(self, phase: RebasePhase, step: str = ""):
        """Save current state to persistent storage"""
        self.state.phase = phase
        self.state.step = step
        self.state.timestamp = datetime.now().isoformat()

        # Update progress information
        self._update_progress()

        # Save state to JSON file
        state_data = {
            "phase": phase.value,
            "step": step,
            "current_branch": self.state.current_branch,
            "target_branch": self.state.target_branch,
            "source_branch": self.state.source_branch,
            "backup_branch": self.state.backup_branch,
            "timestamp": self.state.timestamp,
            "pid": self.state.pid,
            "session_id": self.state.session_id,
            "total_commits": self.state.total_commits,
            "processed_commits": self.state.processed_commits,
            "progress_percent": self.state.progress_percent,
            "conflicts": [
                {
                    "path": cf.path,
                    "strategy": cf.strategy.value,
                    "resolved": cf.resolved,
                    "resolution_method": cf.resolution_method,
                    "backup_created": cf.backup_created,
                    "error_message": cf.error_message,
                }
                for cf in self.state.conflicts
            ],
            "resolved_conflicts": self.state.resolved_conflicts,
            "force_push": self.state.force_push,
            "dry_run": self.state.dry_run,
            "verbose": self.state.verbose,
            "git_status": self.state.git_status,
            "error_messages": self.state.error_messages,
            "recovery_instructions": self.state.recovery_instructions,
        }

        with open(self.state_file, "w") as f:
            json.dump(state_data, f, indent=2)

        # Also save progress file for quick reference
        progress_data = {
            "phase": phase.value,
            "step": step,
            "progress_percent": self.state.progress_percent,
            "conflicts_total": len(self.state.conflicts),
            "conflicts_resolved": self.state.resolved_conflicts,
            "timestamp": self.state.timestamp,
        }

        with open(self.progress_file, "w") as f:
            json.dump(progress_data, f, indent=2)

        self.log_verbose(f"State saved: {phase.value}/{step}")

    def load_state(self) -> bool:
        """Load state from persistent storage"""
        if not self.state_file.exists():
            return False

        try:
            with open(self.state_file, "r") as f:
                state_data = json.load(f)

            # Restore state
            self.state.phase = RebasePhase(state_data.get("phase", "init"))
            self.state.step = state_data.get("step", "")
            self.state.current_branch = state_data.get("current_branch", "")
            self.state.target_branch = state_data.get("target_branch", "")
            self.state.source_branch = state_data.get("source_branch", "")
            self.state.backup_branch = state_data.get("backup_branch", "")
            self.state.total_commits = state_data.get("total_commits", 0)
            self.state.processed_commits = state_data.get("processed_commits", 0)
            self.state.progress_percent = state_data.get("progress_percent", 0)
            self.state.resolved_conflicts = state_data.get("resolved_conflicts", 0)
            self.state.force_push = state_data.get("force_push", False)
            self.state.error_messages = state_data.get("error_messages", [])
            self.state.recovery_instructions = state_data.get(
                "recovery_instructions", []
            )

            # Restore conflicts
            self.state.conflicts = []
            for cf_data in state_data.get("conflicts", []):
                cf = ConflictFile(
                    path=cf_data["path"],
                    strategy=ConflictStrategy(cf_data["strategy"]),
                    resolved=cf_data.get("resolved", False),
                    resolution_method=cf_data.get("resolution_method", ""),
                    backup_created=cf_data.get("backup_created", False),
                    error_message=cf_data.get("error_message", ""),
                )
                self.state.conflicts.append(cf)

            return True

        except Exception as e:
            self.log_error(f"Failed to load state: {e}")
            return False

    def _update_progress(self):
        """Update progress information"""
        try:
            # Get current Git status
            result = self.run_command(
                ["git", "status", "--porcelain"], capture_output=True
            )
            self.state.git_status = result.stdout.strip()

            # Update commit progress if we have branches
            if self.state.current_branch and self.state.target_branch:
                try:
                    # Get total commits to rebase
                    result = self.run_command(
                        [
                            "git",
                            "rev-list",
                            "--count",
                            f"{self.state.target_branch}..{self.state.current_branch}",
                        ],
                        capture_output=True,
                    )
                    self.state.total_commits = int(result.stdout.strip())

                    # Try to determine processed commits
                    if self._is_rebase_in_progress():
                        rebase_todo = Path(".git/rebase-merge/git-rebase-todo")
                        if rebase_todo.exists():
                            with open(rebase_todo, "r") as f:
                                remaining = len(
                                    [
                                        line
                                        for line in f
                                        if line.strip().startswith(("pick", "p "))
                                    ]
                                )
                            self.state.processed_commits = (
                                self.state.total_commits - remaining
                            )

                    # Calculate progress
                    if self.state.total_commits > 0:
                        self.state.progress_percent = int(
                            (self.state.processed_commits / self.state.total_commits)
                            * 100
                        )

                except subprocess.CalledProcessError:
                    pass  # Ignore errors in progress calculation

        except Exception as e:
            self.log_verbose(f"Progress update failed: {e}")

    def _is_rebase_in_progress(self) -> bool:
        """Check if a rebase is currently in progress"""
        return Path(".git/rebase-merge").exists() or Path(".git/rebase-apply").exists()

    def show_progress(self):
        """Display current progress"""
        self.log_info("Rebase Progress:")
        self.log_info(f"  Phase: {self.state.phase.value} ({self.state.step})")
        self.log_info(
            f"  Progress: {self.state.progress_percent}% "
            f"({self.state.processed_commits}/{self.state.total_commits} commits)"
        )
        self.log_info(
            f"  Conflicts resolved: {self.state.resolved_conflicts}/{len(self.state.conflicts)}"
        )

        if self.state.current_branch and self.state.target_branch:
            self.log_info(
                f"  Branch: {self.state.current_branch} -> {self.state.target_branch}"
            )

    def check_resume(self) -> bool:
        """Check if we can resume from a previous session"""
        if not self.load_state():
            return False

        if self.state.phase == RebasePhase.COMPLETE:
            self.log_info("Previous rebase session completed successfully")
            return False

        self.log_info("Found previous rebase session:")
        self.show_progress()

        response = input("\nResume previous rebase session? [Y/n]: ").strip().lower()
        if response in ["", "y", "yes"]:
            self.log_info("Resuming previous session...")
            return True
        else:
            self.log_info("Starting fresh rebase session")
            self.cleanup_state()
            return False

    def cleanup_state(self):
        """Clean up state files"""
        try:
            if self.state_file.exists():
                self.state_file.unlink()
            if self.progress_file.exists():
                self.progress_file.unlink()
            self.log_verbose("State files cleaned up")
        except Exception as e:
            self.log_warning(f"Failed to clean up state files: {e}")

    def run_command(
        self,
        cmd: List[str],
        capture_output: bool = True,
        check: bool = True,
        input_text: str = None,
    ) -> subprocess.CompletedProcess:
        """Run a command with proper logging and error handling"""
        cmd_str = " ".join(cmd)
        self.log_verbose(f"Running command: {cmd_str}")

        if self.dry_run and not cmd[0] == "git" or "status" in cmd or "log" in cmd:
            self.log_info(f"[DRY RUN] Would run: {cmd_str}")
            return subprocess.CompletedProcess(cmd, 0, stdout="", stderr="")

        try:
            result = subprocess.run(
                cmd,
                capture_output=capture_output,
                text=True,
                check=check,
                input=input_text,
            )

            if result.stdout and self.verbose:
                self.log_verbose(f"Command output: {result.stdout.strip()}")
            if result.stderr and self.verbose:
                self.log_verbose(f"Command error: {result.stderr.strip()}")

            return result

        except subprocess.CalledProcessError as e:
            error_msg = f"Command failed: {cmd_str}"
            if e.stderr:
                error_msg += f"\nError: {e.stderr.strip()}"
            self.log_error(error_msg)
            raise GitRebaseError(error_msg) from e

    def get_current_branch(self) -> str:
        """Get current Git branch"""
        try:
            result = self.run_command(
                ["git", "branch", "--show-current"], capture_output=True
            )
            return result.stdout.strip()
        except GitRebaseError:
            return ""

    def branch_exists(self, branch: str) -> bool:
        """Check if a branch exists"""
        try:
            self.run_command(
                ["git", "show-ref", "--verify", f"refs/heads/{branch}"],
                capture_output=True,
            )
            return True
        except GitRebaseError:
            return False

    def get_conflicted_files(self) -> List[str]:
        """Get list of files with merge conflicts"""
        try:
            result = self.run_command(
                ["git", "diff", "--name-only", "--diff-filter=U"], capture_output=True
            )
            return [f.strip() for f in result.stdout.split("\n") if f.strip()]
        except GitRebaseError:
            return []

    def determine_conflict_strategy(self, file_path: str) -> ConflictStrategy:
        """Determine conflict resolution strategy for a file"""
        for pattern, strategy in self.conflict_patterns.items():
            if re.match(pattern, file_path):
                return strategy

        # Default strategy
        return ConflictStrategy.MANUAL_REVIEW

    def create_backup_branch(self, branch: str) -> str:
        """Create a backup branch"""
        backup_name = f"{branch}_backup_{int(time.time())}"

        try:
            self.run_command(
                ["git", "branch", backup_name, branch], capture_output=True
            )
            self.log_success(f"Created backup branch: {backup_name}")
            return backup_name
        except GitRebaseError as e:
            self.log_error(f"Failed to create backup branch: {e}")
            raise

    def resolve_conflict_auto(self, file_path: str) -> bool:
        """Automatically resolve conflicts using smart strategies"""
        try:
            # Read the conflicted file
            with open(file_path, "r", encoding="utf-8") as f:
                content = f.read()

            # Check if it has conflict markers
            if not (
                "<<<<<<< " in content and "=======" in content and ">>>>>>> " in content
            ):
                return True  # No conflicts to resolve

            # Split content into sections
            lines = content.split("\n")
            resolved_lines = []
            in_conflict = False
            current_section = []
            incoming_section = []

            for line in lines:
                if line.startswith("<<<<<<< "):
                    in_conflict = True
                    current_section = []
                    incoming_section = []
                elif line.startswith("======= "):
                    # Switch to incoming section
                    pass
                elif line.startswith(">>>>>>> "):
                    in_conflict = False

                    # Apply resolution strategy
                    strategy = self.determine_conflict_strategy(file_path)

                    if strategy == ConflictStrategy.PREFER_CURRENT:
                        resolved_lines.extend(current_section)
                    elif strategy == ConflictStrategy.PREFER_INCOMING:
                        resolved_lines.extend(incoming_section)
                    elif strategy == ConflictStrategy.SMART_MERGE:
                        # Try to intelligently merge
                        merged = self._smart_merge_sections(
                            current_section, incoming_section, file_path
                        )
                        resolved_lines.extend(merged)
                    else:
                        # For manual review, prefer incoming as safe default
                        resolved_lines.extend(incoming_section)

                elif in_conflict:
                    if "=======" not in line:
                        if not any(
                            line.startswith("=======") for line in resolved_lines[-10:]
                        ):
                            current_section.append(line)
                        else:
                            incoming_section.append(line)
                else:
                    resolved_lines.append(line)

            # Write resolved content
            with open(file_path, "w", encoding="utf-8") as f:
                f.write("\n".join(resolved_lines))

            # Stage the resolved file
            self.run_command(["git", "add", file_path], capture_output=True)

            return True

        except Exception as e:
            self.log_error(f"Auto-resolution failed for {file_path}: {e}")
            return False

    def _smart_merge_sections(
        self, current: List[str], incoming: List[str], file_path: str
    ) -> List[str]:
        """Intelligently merge conflicting sections"""
        # For certain file types, apply specific merge strategies
        if file_path.endswith(".md"):
            # For markdown files, combine content intelligently
            return self._merge_markdown_sections(current, incoming)
        elif file_path.endswith((".go", ".py", ".js", ".ts")):
            # For code files, prefer incoming but preserve important local changes
            return self._merge_code_sections(current, incoming)
        else:
            # Default: prefer incoming
            return incoming

    def _merge_markdown_sections(
        self, current: List[str], incoming: List[str]
    ) -> List[str]:
        """Merge markdown sections intelligently"""
        # Combine unique lines from both sections
        incoming_set = set(line.strip() for line in incoming if line.strip())

        # Start with incoming content
        merged = incoming[:]

        # Add unique lines from current
        for line in current:
            if line.strip() and line.strip() not in incoming_set:
                merged.append(line)

        return merged

    def _merge_code_sections(
        self, current: List[str], incoming: List[str]
    ) -> List[str]:
        """Merge code sections intelligently"""
        # For code, prefer incoming but look for important local additions
        merged = incoming[:]

        # Look for comments, imports, or other safe additions in current
        for line in current:
            line_stripped = line.strip()
            if line_stripped.startswith(
                ("import ", "from ", "#", "//", "/*")
            ) and line_stripped not in [
                incoming_line.strip() for incoming_line in incoming
            ]:
                merged.append(line)

        return merged

    def resolve_conflicts(self) -> bool:
        """Resolve all conflicts in the repository"""
        conflicted_files = self.get_conflicted_files()

        if not conflicted_files:
            self.log_info("No conflicts to resolve")
            return True

        self.log_info(f"Found {len(conflicted_files)} conflicted files")

        # Initialize conflict tracking
        self.state.conflicts = []
        for file_path in conflicted_files:
            strategy = self.determine_conflict_strategy(file_path)
            conflict = ConflictFile(path=file_path, strategy=strategy)
            self.state.conflicts.append(conflict)

        self.save_state(RebasePhase.CONFLICT_RESOLUTION, "analyzing_conflicts")

        # Resolve each conflict
        for conflict in self.state.conflicts:
            self.log_step(
                f"Resolving conflict in {conflict.path} using {conflict.strategy.value}"
            )

            # Create backup if needed
            if not conflict.backup_created:
                backup_path = (
                    self.backup_dir
                    / f"{conflict.path.replace('/', '_')}_backup_{int(time.time())}"
                )
                backup_path.parent.mkdir(parents=True, exist_ok=True)

                try:
                    import shutil

                    shutil.copy2(conflict.path, backup_path)
                    conflict.backup_created = True
                except Exception as e:
                    self.log_warning(
                        f"Failed to create backup for {conflict.path}: {e}"
                    )

            # Attempt resolution
            if conflict.strategy == ConflictStrategy.MANUAL_REVIEW:
                self.log_warning(f"Manual review required for {conflict.path}")
                conflict.resolution_method = "manual_review_required"
                continue

            try:
                if self.resolve_conflict_auto(conflict.path):
                    conflict.resolved = True
                    conflict.resolution_method = (
                        f"auto_resolved_{conflict.strategy.value}"
                    )
                    self.state.resolved_conflicts += 1
                    self.log_success(f"Resolved conflict in {conflict.path}")
                else:
                    conflict.error_message = "Auto-resolution failed"
                    self.log_error(f"Failed to resolve conflict in {conflict.path}")

            except Exception as e:
                conflict.error_message = str(e)
                self.log_error(f"Error resolving conflict in {conflict.path}: {e}")

        # Check if all conflicts are resolved
        unresolved = [cf for cf in self.state.conflicts if not cf.resolved]

        if unresolved:
            self.log_warning(f"{len(unresolved)} conflicts require manual resolution:")
            for conflict in unresolved:
                self.log_warning(f"  - {conflict.path} ({conflict.strategy.value})")

            self.save_state(
                RebasePhase.CONFLICT_RESOLUTION, "manual_resolution_required"
            )
            return False

        self.log_success("All conflicts resolved successfully")
        return True

    def perform_rebase(
        self, target_branch: str, force_push: bool = False
    ) -> RebaseResult:
        """Perform the complete rebase operation"""
        self.state.target_branch = target_branch
        self.state.force_push = force_push

        try:
            # Check if we can resume
            if self.check_resume():
                return self._resume_rebase()

            # Initialize fresh rebase
            self.state.current_branch = self.get_current_branch()
            self.state.source_branch = self.state.current_branch

            if not self.state.current_branch:
                raise GitRebaseError("Not on a Git branch")

            if self.state.current_branch == target_branch:
                self.log_info("Already on target branch")
                return RebaseResult.SUCCESS

            # Prerequisites check
            self.save_state(RebasePhase.PREREQUISITES, "checking_prerequisites")
            self._check_prerequisites()

            # Create backup
            self.save_state(RebasePhase.BACKUP, "creating_backup")
            self.state.backup_branch = self.create_backup_branch(
                self.state.current_branch
            )

            # Start rebase
            self.save_state(RebasePhase.REBASE_START, "starting_rebase")
            self.log_step(
                f"Starting rebase of {self.state.current_branch} onto {target_branch}"
            )

            try:
                self.run_command(["git", "rebase", target_branch], capture_output=True)

                # Rebase completed successfully
                self.save_state(RebasePhase.PUSH, "preparing_push")
                return self._complete_rebase()

            except GitRebaseError:
                # Rebase has conflicts
                self.save_state(RebasePhase.REBASE_IN_PROGRESS, "conflicts_detected")

                if self.resolve_conflicts():
                    # Continue rebase
                    self.save_state(RebasePhase.REBASE_CONTINUE, "continuing_rebase")
                    return self._continue_rebase()
                else:
                    # Manual intervention required
                    self._generate_recovery_instructions()
                    return RebaseResult.CONFLICTS

        except Exception as e:
            self.log_error(f"Rebase failed: {e}")
            self.save_state(RebasePhase.INIT, "failed")
            self._generate_recovery_instructions()
            return RebaseResult.FAILED

    def _resume_rebase(self) -> RebaseResult:
        """Resume rebase from saved state"""
        self.log_info(f"Resuming rebase from phase: {self.state.phase.value}")

        if self.state.phase == RebasePhase.CONFLICT_RESOLUTION:
            if self.resolve_conflicts():
                return self._continue_rebase()
            else:
                return RebaseResult.CONFLICTS

        elif self.state.phase == RebasePhase.REBASE_IN_PROGRESS:
            if self._is_rebase_in_progress():
                if self.resolve_conflicts():
                    return self._continue_rebase()
                else:
                    return RebaseResult.CONFLICTS
            else:
                return self._complete_rebase()

        elif self.state.phase == RebasePhase.REBASE_CONTINUE:
            return self._continue_rebase()

        elif self.state.phase == RebasePhase.PUSH:
            return self._complete_rebase()

        else:
            self.log_warning("Unknown resume state, starting fresh")
            return self.perform_rebase(self.state.target_branch, self.state.force_push)

    def _continue_rebase(self) -> RebaseResult:
        """Continue rebase after conflict resolution"""
        try:
            self.run_command(["git", "rebase", "--continue"], capture_output=True)
            self.log_success("Rebase continued successfully")
            return self._complete_rebase()

        except GitRebaseError:
            # More conflicts or other issues
            if self.get_conflicted_files():
                self.save_state(RebasePhase.CONFLICT_RESOLUTION, "additional_conflicts")
                if self.resolve_conflicts():
                    return self._continue_rebase()
                else:
                    return RebaseResult.CONFLICTS
            else:
                self.log_error("Rebase continue failed without conflicts")
                return RebaseResult.FAILED

    def _complete_rebase(self) -> RebaseResult:
        """Complete the rebase operation"""
        self.save_state(RebasePhase.PUSH, "preparing_push")

        # Push changes if requested
        if self.state.force_push and not self.dry_run:
            try:
                self.run_command(
                    [
                        "git",
                        "push",
                        "--force-with-lease",
                        "origin",
                        self.state.current_branch,
                    ]
                )
                self.log_success("Changes pushed successfully")
            except GitRebaseError as e:
                self.log_warning(f"Push failed: {e}")

        # Generate summary
        self.save_state(RebasePhase.CLEANUP, "generating_summary")
        self._generate_summary()

        # Mark as complete
        self.save_state(RebasePhase.COMPLETE, "success")
        self.log_success("Rebase completed successfully!")

        return RebaseResult.SUCCESS

    def _check_prerequisites(self):
        """Check prerequisites for rebase"""
        # Check if working directory is clean
        result = self.run_command(["git", "status", "--porcelain"], capture_output=True)
        if result.stdout.strip():
            raise GitRebaseError(
                "Working directory is not clean. Please commit or stash changes."
            )

        # Check if target branch exists
        if not self.branch_exists(self.state.target_branch):
            raise GitRebaseError(
                f"Target branch '{self.state.target_branch}' does not exist"
            )

        # Fetch latest changes
        try:
            self.run_command(["git", "fetch", "origin"], capture_output=True)
        except GitRebaseError:
            self.log_warning("Failed to fetch latest changes")

    def _generate_recovery_instructions(self):
        """Generate recovery instructions for manual intervention"""
        instructions = [
            "# Rebase Recovery Instructions",
            "",
            "The rebase operation encountered issues and requires manual intervention.",
            f"Session ID: {self.state.session_id}",
            f"Timestamp: {self.state.timestamp}",
            "",
            "## Current State",
            f"- Phase: {self.state.phase.value}",
            f"- Step: {self.state.step}",
            f"- Current branch: {self.state.current_branch}",
            f"- Target branch: {self.state.target_branch}",
            f"- Backup branch: {self.state.backup_branch}",
            "",
            "## To Resume",
            f"Run: python3 {__file__} --resume",
            "",
            "## To Abort",
            "Run: git rebase --abort",
            f"Then: git checkout {self.state.backup_branch}",
            "",
            "## Manual Conflict Resolution",
        ]

        if self.state.conflicts:
            instructions.append("The following files have conflicts:")
            for conflict in self.state.conflicts:
                if not conflict.resolved:
                    instructions.append(
                        f"- {conflict.path} (strategy: {conflict.strategy.value})"
                    )
                    if conflict.error_message:
                        instructions.append(f"  Error: {conflict.error_message}")

        instructions.extend(
            [
                "",
                "After resolving conflicts manually:",
                "1. git add <resolved-files>",
                "2. git rebase --continue",
                f"3. python3 {__file__} --resume",
            ]
        )

        self.state.recovery_instructions = instructions

        # Write to file
        with open(self.state_dir / "recovery_instructions.md", "w") as f:
            f.write("\n".join(instructions))

    def _generate_summary(self):
        """Generate operation summary"""
        duration = datetime.now() - self.start_time

        summary = [
            "# Rebase Summary",
            "",
            f"**Session ID:** {self.state.session_id}",
            f"**Started:** {self.start_time.isoformat()}",
            f"**Duration:** {duration}",
            f"**Status:** {self.state.phase.value}",
            "",
            "## Branches",
            f"- **Source:** {self.state.source_branch}",
            f"- **Target:** {self.state.target_branch}",
            f"- **Backup:** {self.state.backup_branch}",
            "",
            "## Progress",
            f"- **Total commits:** {self.state.total_commits}",
            f"- **Processed commits:** {self.state.processed_commits}",
            f"- **Progress:** {self.state.progress_percent}%",
            "",
            "## Conflicts",
            f"- **Total conflicts:** {len(self.state.conflicts)}",
            f"- **Resolved conflicts:** {self.state.resolved_conflicts}",
            "",
        ]

        if self.state.conflicts:
            summary.append("### Conflict Details")
            for conflict in self.state.conflicts:
                status = "✓" if conflict.resolved else "✗"
                summary.append(
                    f"- {status} {conflict.path} ({conflict.strategy.value})"
                )
                if conflict.resolution_method:
                    summary.append(f"  - Resolution: {conflict.resolution_method}")
                if conflict.error_message:
                    summary.append(f"  - Error: {conflict.error_message}")

        if self.state.error_messages:
            summary.append("\n### Errors")
            for error in self.state.error_messages:
                summary.append(f"- {error}")

        summary.append("\n## Files")
        summary.append(f"- **Log:** {self.log_file}")
        summary.append(f"- **State:** {self.state_file}")
        summary.append(f"- **Progress:** {self.progress_file}")
        summary.append(f"- **Conflicts:** {self.conflict_log}")

        # Write summary
        with open(self.summary_file, "w") as f:
            f.write("\n".join(summary))

        self.log_info(f"Summary written to: {self.summary_file}")

    def show_status(self):
        """Show current rebase status"""
        if self.load_state():
            self.show_progress()

            if self.state.conflicts:
                self.log_info("\nConflict Status:")
                for conflict in self.state.conflicts:
                    status = "✓" if conflict.resolved else "✗"
                    self.log_info(f"  {status} {conflict.path}")
        else:
            self.log_info("No active rebase session found")

    def abort_rebase(self):
        """Abort current rebase and restore backup"""
        if self._is_rebase_in_progress():
            self.run_command(["git", "rebase", "--abort"])
            self.log_info("Rebase aborted")

        if self.state.backup_branch and self.branch_exists(self.state.backup_branch):
            self.run_command(["git", "checkout", self.state.backup_branch])
            self.log_info(f"Restored backup branch: {self.state.backup_branch}")

        self.cleanup_state()
        self.log_info("Rebase session cleaned up")


def create_argument_parser() -> argparse.ArgumentParser:
    """Create argument parser for the rebase tool"""
    parser = argparse.ArgumentParser(
        description="Smart Git Rebase Tool with intelligent conflict resolution",
        formatter_class=argparse.RawDescriptionHelpFormatter,
        epilog="""
Examples:
  %(prog)s main                    # Rebase current branch onto main
  %(prog)s --force main            # Rebase and force push
  %(prog)s --dry-run main          # Show what would be done
  %(prog)s --resume                # Resume interrupted rebase
  %(prog)s --status                # Show current rebase status
  %(prog)s --abort                 # Abort current rebase
  %(prog)s --cleanup               # Clean up state files
        """,
    )

    # Primary argument
    parser.add_argument("target_branch", nargs="?", help="Target branch to rebase onto")

    # Operation modes
    parser.add_argument(
        "--resume", action="store_true", help="Resume interrupted rebase session"
    )

    parser.add_argument(
        "--status", action="store_true", help="Show current rebase status"
    )

    parser.add_argument(
        "--abort", action="store_true", help="Abort current rebase and restore backup"
    )

    parser.add_argument("--cleanup", action="store_true", help="Clean up state files")

    # Options
    parser.add_argument(
        "--force", action="store_true", help="Force push after successful rebase"
    )

    parser.add_argument(
        "--dry-run",
        action="store_true",
        help="Show what would be done without executing",
    )

    parser.add_argument("--verbose", action="store_true", help="Enable verbose logging")

    parser.add_argument(
        "--mode",
        choices=["interactive", "automated", "smart"],
        default="smart",
        help="Rebase mode (default: smart)",
    )

    return parser


def main() -> int:
    """Main entry point"""
    parser = create_argument_parser()
    args = parser.parse_args()

    # Create rebase instance
    rebase = SmartRebase(verbose=args.verbose, dry_run=args.dry_run)

    try:
        # Handle different operations
        if args.cleanup:
            rebase.cleanup_state()
            return 0

        if args.status:
            rebase.show_status()
            return 0

        if args.abort:
            rebase.abort_rebase()
            return 0

        if args.resume:
            if rebase.load_state():
                result = rebase._resume_rebase()
                return 0 if result == RebaseResult.SUCCESS else 1
            else:
                print("No rebase session to resume")
                return 1

        # Regular rebase operation
        if not args.target_branch:
            parser.error("Target branch is required for rebase operation")

        result = rebase.perform_rebase(args.target_branch, force_push=args.force)

        if result == RebaseResult.SUCCESS:
            return 0
        elif result == RebaseResult.CONFLICTS:
            print(
                "\nConflicts require manual resolution. Run with --resume after fixing conflicts."
            )
            return 1
        else:
            print(f"\nRebase failed: {result.value}")
            return 1

    except KeyboardInterrupt:
        print("\nOperation interrupted by user")
        return 130
    except Exception as e:
        print(f"Error: {e}")
        return 1


if __name__ == "__main__":
    sys.exit(main())
