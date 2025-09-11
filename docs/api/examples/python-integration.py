# file: docs/api/examples/python-integration.py
# version: 1.0.0
# guid: 550e8400-e29b-41d4-a716-446655440022

#!/usr/bin/env python3
"""
Comprehensive Python integration examples for Subtitle Manager API.

This script demonstrates various integration patterns including:
- Basic operations
- File processing workflows
- Library management
- Error handling
- Async operations
"""

import os
import logging
from pathlib import Path
from typing import List, Dict, Optional
from datetime import datetime, timedelta

# Import the SDK (install with: pip install subtitle-manager-sdk)
from subtitle_manager_sdk import SubtitleManagerClient
from subtitle_manager_sdk.models import OperationStatus, TranslationProvider
from subtitle_manager_sdk.exceptions import (
    SubtitleManagerError,
    AuthenticationError,
    RateLimitError,
)

# Configure logging
logging.basicConfig(
    level=logging.INFO, format="%(asctime)s - %(name)s - %(levelname)s - %(message)s"
)
logger = logging.getLogger(__name__)


class SubtitleManagerIntegration:
    """
    Example integration class showing best practices for using the
    Subtitle Manager SDK in production applications.
    """

    def __init__(
        self, base_url: str = "http://localhost:8080", api_key: Optional[str] = None
    ):
        """Initialize the integration with proper configuration."""
        self.client = SubtitleManagerClient(
            base_url=base_url,
            api_key=api_key or os.getenv("SUBTITLE_MANAGER_API_KEY"),
            timeout=60,  # 60 second timeout
            max_retries=3,
            retry_backoff=2.0,
        )
        logger.info(
            f"Initialized Subtitle Manager integration with base URL: {base_url}"
        )

    async def health_check(self) -> bool:
        """Check if the Subtitle Manager API is healthy."""
        try:
            system_info = self.client.get_system_info()
            logger.info(
                f"Health check passed - System: {system_info.os} {system_info.arch}"
            )
            logger.info(
                f"Disk usage: {system_info.disk_free_formatted}/{system_info.disk_total_formatted}"
            )
            return True
        except Exception as e:
            logger.error(f"Health check failed: {e}")
            return False

    def authenticate_session(self, username: str, password: str) -> bool:
        """Authenticate with username/password instead of API key."""
        try:
            login_response = self.client.login(username, password)
            logger.info(
                f"Authenticated as {login_response.username} with role: {login_response.role}"
            )
            return True
        except AuthenticationError as e:
            logger.error(f"Authentication failed: {e}")
            return False

    def process_subtitle_file(
        self, file_path: Path, target_language: Optional[str] = None
    ) -> Optional[bytes]:
        """
        Process a subtitle file: convert to SRT and optionally translate.

        Args:
            file_path: Path to the subtitle file
            target_language: Optional target language for translation

        Returns:
            Processed subtitle content as bytes
        """
        try:
            logger.info(f"Processing subtitle file: {file_path}")

            # Step 1: Convert to SRT format
            with open(file_path, "rb") as f:
                srt_content = self.client.convert_subtitle(f, filename=file_path.name)

            logger.info(f"Converted {file_path} to SRT format")

            # Step 2: Translate if requested
            if target_language:
                logger.info(f"Translating to {target_language}")
                from io import BytesIO

                srt_file = BytesIO(srt_content)

                translated_content = self.client.translate_subtitle(
                    file=srt_file,
                    language=target_language,
                    provider=TranslationProvider.GOOGLE,
                    filename=f"translated_{file_path.stem}.srt",
                )

                logger.info(f"Translation to {target_language} completed")
                return translated_content

            return srt_content

        except SubtitleManagerError as e:
            logger.error(f"Failed to process subtitle file {file_path}: {e}")
            return None
        except Exception as e:
            logger.error(f"Unexpected error processing {file_path}: {e}")
            return None

    def batch_download_subtitles(
        self, media_files: List[Dict[str, str]]
    ) -> Dict[str, Dict]:
        """
        Download subtitles for multiple media files.

        Args:
            media_files: List of dicts with 'path' and 'language' keys

        Returns:
            Dictionary with results for each file
        """
        results = {}
        total_files = len(media_files)

        logger.info(f"Starting batch download for {total_files} files")

        for i, media_file in enumerate(media_files, 1):
            file_path = media_file["path"]
            language = media_file["language"]
            providers = media_file.get("providers")

            logger.info(f"Processing {i}/{total_files}: {file_path}")

            try:
                result = self.client.download_subtitles(
                    path=file_path, language=language, providers=providers
                )

                if result.success:
                    logger.info(f"Downloaded subtitle: {result.subtitle_path}")
                    results[file_path] = {
                        "status": "success",
                        "subtitle_path": result.subtitle_path,
                        "provider": result.provider,
                    }
                else:
                    logger.warning(f"No subtitles found for {file_path}")
                    results[file_path] = {"status": "not_found"}

            except RateLimitError as e:
                logger.warning(f"Rate limited, waiting {e.retry_after} seconds")
                import time

                time.sleep(e.retry_after)
                # Retry the same file
                i -= 1
                continue

            except SubtitleManagerError as e:
                logger.error(f"Failed to download subtitles for {file_path}: {e}")
                results[file_path] = {"status": "error", "error": str(e)}

        success_count = sum(1 for r in results.values() if r["status"] == "success")
        logger.info(
            f"Batch download completed: {success_count}/{total_files} successful"
        )

        return results

    def monitor_library_scan(self, path: Optional[str] = None) -> bool:
        """
        Start a library scan and monitor its progress.

        Args:
            path: Optional specific path to scan

        Returns:
            True if scan completed successfully
        """
        try:
            # Start the scan
            scan_result = self.client.start_library_scan(path=path, force=True)
            logger.info(f"Started library scan with ID: {scan_result.scan_id}")

            # Monitor progress
            while True:
                status = self.client.get_scan_status()

                if not status.scanning:
                    logger.info("Library scan completed")
                    return True

                logger.info(f"Scan progress: {status.progress:.1%}")
                if status.current_path:
                    logger.info(f"Current path: {status.current_path}")
                if status.files_processed and status.files_total:
                    logger.info(f"Files: {status.files_processed}/{status.files_total}")

                # Wait 5 seconds before checking again
                import time

                time.sleep(5)

        except Exception as e:
            logger.error(f"Library scan failed: {e}")
            return False

    def analyze_download_history(self, days: int = 7) -> Dict:
        """
        Analyze recent download history and provide statistics.

        Args:
            days: Number of days to analyze

        Returns:
            Dictionary with analysis results
        """
        try:
            # Calculate date range
            end_date = datetime.now()
            start_date = end_date - timedelta(days=days)

            # Get history with pagination
            all_items = []
            page = 1

            while True:
                history_response = self.client.get_history(
                    page=page, limit=100, start_date=start_date, end_date=end_date
                )

                all_items.extend(history_response.items)

                if not history_response.has_next_page:
                    break

                page += 1

            # Analyze the data
            analysis = {
                "total_operations": len(all_items),
                "by_type": {},
                "by_status": {},
                "by_provider": {},
                "success_rate": 0,
                "most_active_days": {},
                "failed_operations": [],
            }

            for item in all_items:
                # Count by type
                analysis["by_type"][item.type] = (
                    analysis["by_type"].get(item.type, 0) + 1
                )

                # Count by status
                analysis["by_status"][item.status] = (
                    analysis["by_status"].get(item.status, 0) + 1
                )

                # Count by provider (for successful downloads)
                if item.provider and item.status == OperationStatus.SUCCESS:
                    analysis["by_provider"][item.provider] = (
                        analysis["by_provider"].get(item.provider, 0) + 1
                    )

                # Track failed operations
                if item.status == OperationStatus.FAILED:
                    analysis["failed_operations"].append(
                        {
                            "file_path": item.file_path,
                            "type": item.type,
                            "error": item.error_message,
                            "date": item.created_at,
                        }
                    )

                # Count by day
                day_key = item.created_at_date.strftime("%Y-%m-%d")
                analysis["most_active_days"][day_key] = (
                    analysis["most_active_days"].get(day_key, 0) + 1
                )

            # Calculate success rate
            successful = analysis["by_status"].get(OperationStatus.SUCCESS, 0)
            if analysis["total_operations"] > 0:
                analysis["success_rate"] = (
                    successful / analysis["total_operations"]
                ) * 100

            logger.info(
                f"Analyzed {analysis['total_operations']} operations from last {days} days"
            )
            logger.info(f"Success rate: {analysis['success_rate']:.1f}%")

            return analysis

        except Exception as e:
            logger.error(f"Failed to analyze history: {e}")
            return {}

    def extract_and_translate_pipeline(
        self, video_file: Path, target_languages: List[str]
    ) -> Dict[str, Optional[bytes]]:
        """
        Extract embedded subtitles and translate to multiple languages.

        Args:
            video_file: Path to video file
            target_languages: List of target language codes

        Returns:
            Dictionary mapping language codes to translated subtitle content
        """
        results = {}

        try:
            logger.info(
                f"Starting extraction and translation pipeline for: {video_file}"
            )

            # Step 1: Extract embedded subtitles
            with open(video_file, "rb") as f:
                extracted_subs = self.client.extract_subtitles(
                    file=f, filename=video_file.name
                )

            logger.info("Successfully extracted embedded subtitles")

            # Step 2: Translate to each target language
            from io import BytesIO

            for language in target_languages:
                try:
                    logger.info(f"Translating to {language}")

                    subs_file = BytesIO(extracted_subs)
                    translated = self.client.translate_subtitle(
                        file=subs_file,
                        language=language,
                        provider=TranslationProvider.GOOGLE,
                        filename=f"{video_file.stem}.{language}.srt",
                    )

                    results[language] = translated
                    logger.info(f"Translation to {language} completed")

                except SubtitleManagerError as e:
                    logger.error(f"Translation to {language} failed: {e}")
                    results[language] = None

            return results

        except SubtitleManagerError as e:
            logger.error(f"Extraction failed: {e}")
            return {lang: None for lang in target_languages}

    def create_processing_report(self, output_file: Path) -> None:
        """
        Create a comprehensive processing report.

        Args:
            output_file: Path to save the report
        """
        try:
            logger.info("Generating processing report")

            # Gather system information
            system_info = self.client.get_system_info()

            # Get recent history
            history_analysis = self.analyze_download_history(days=30)

            # Get current scan status
            scan_status = self.client.get_scan_status()

            # Get recent logs
            recent_logs = self.client.get_logs(limit=50)
            error_logs = [log for log in recent_logs if log.is_error]

            # Create report
            report = f"""
# Subtitle Manager Processing Report
Generated: {datetime.now().isoformat()}

## System Information
- OS: {system_info.os} {system_info.arch}
- Go Version: {system_info.go_version}
- Version: {system_info.version or "Unknown"}
- Uptime: {system_info.uptime or "Unknown"}
- Memory Usage: {system_info.memory_usage or "Unknown"} bytes
- Disk Usage: {system_info.disk_free_formatted}/{system_info.disk_total_formatted}

## Library Status
- Currently Scanning: {scan_status.scanning}
- Scan Progress: {scan_status.progress:.1%}
- Files Processed: {scan_status.files_processed or "N/A"}
- Files Total: {scan_status.files_total or "N/A"}

## Activity Summary (Last 30 Days)
- Total Operations: {history_analysis.get("total_operations", 0)}
- Success Rate: {history_analysis.get("success_rate", 0):.1f}%

### Operations by Type:
"""
            for op_type, count in history_analysis.get("by_type", {}).items():
                report += f"- {op_type}: {count}\n"

            report += "\n### Operations by Status:\n"
            for status, count in history_analysis.get("by_status", {}).items():
                report += f"- {status}: {count}\n"

            report += "\n### Provider Performance:\n"
            for provider, count in history_analysis.get("by_provider", {}).items():
                report += f"- {provider}: {count} successful downloads\n"

            if error_logs:
                report += f"\n## Recent Errors ({len(error_logs)} found)\n"
                for log in error_logs[:10]:  # Show last 10 errors
                    report += (
                        f"- {log.timestamp_date}: {log.component} - {log.message}\n"
                    )

            # Save report
            with open(output_file, "w") as f:
                f.write(report)

            logger.info(f"Processing report saved to: {output_file}")

        except Exception as e:
            logger.error(f"Failed to generate report: {e}")


def main():
    """
    Main function demonstrating various integration scenarios.
    """
    # Initialize integration
    integration = SubtitleManagerIntegration()

    # Example 1: Health check
    if not integration.health_check():
        logger.error("Subtitle Manager is not healthy, exiting")
        return

    # Example 2: Process subtitle files
    subtitle_files = [Path("example.vtt"), Path("example.ass"), Path("example.smi")]

    for subtitle_file in subtitle_files:
        if subtitle_file.exists():
            # Convert and translate to Spanish
            processed = integration.process_subtitle_file(subtitle_file, "es")
            if processed:
                output_file = subtitle_file.with_suffix(".es.srt")
                with open(output_file, "wb") as f:
                    f.write(processed)
                logger.info(f"Saved translated subtitle: {output_file}")

    # Example 3: Batch download subtitles
    media_files = [
        {"path": "/movies/example1.mkv", "language": "en"},
        {
            "path": "/movies/example2.mkv",
            "language": "en",
            "providers": ["opensubtitles"],
        },
        {"path": "/tv/series/s01e01.mkv", "language": "es"},
    ]

    download_results = integration.batch_download_subtitles(media_files)
    logger.info(f"Download results: {download_results}")

    # Example 4: Monitor library scan
    # integration.monitor_library_scan('/movies/new')

    # Example 5: Analyze history
    analysis = integration.analyze_download_history(days=7)
    logger.info(f"History analysis: {analysis}")

    # Example 6: Extract and translate pipeline
    video_file = Path("example_video.mkv")
    if video_file.exists():
        translations = integration.extract_and_translate_pipeline(
            video_file, ["es", "fr", "de"]
        )
        for lang, content in translations.items():
            if content:
                output_file = video_file.with_suffix(f".{lang}.srt")
                with open(output_file, "wb") as f:
                    f.write(content)
                logger.info(f"Saved translation: {output_file}")

    # Example 7: Generate report
    report_file = Path("subtitle_manager_report.md")
    integration.create_processing_report(report_file)


if __name__ == "__main__":
    main()
