# file: sdks/python/subtitle_manager_sdk/__init__.py
# version: 1.0.0
# guid: 550e8400-e29b-41d4-a716-446655440004

"""
Subtitle Manager SDK

A comprehensive Python SDK for the Subtitle Manager API.
Provides type-safe access to all API endpoints with automatic retry,
error handling, and pagination support.
"""

from .client import SubtitleManagerClient
from .exceptions import (
    SubtitleManagerError,
    AuthenticationError,
    AuthorizationError,
    NotFoundError,
    RateLimitError,
    APIError,
)
from .models import (
    SystemInfo,
    HistoryItem,
    ScanStatus,
    OAuthCredentials,
    LogEntry,
)

__version__ = "1.0.0"
__author__ = "Subtitle Manager Team"
__email__ = "support@subtitlemanager.com"
__description__ = "Python SDK for Subtitle Manager API"

__all__ = [
    "SubtitleManagerClient",
    "SubtitleManagerError",
    "AuthenticationError",
    "AuthorizationError",
    "NotFoundError",
    "RateLimitError",
    "APIError",
    "SystemInfo",
    "HistoryItem",
    "ScanStatus",
    "OAuthCredentials",
    "LogEntry",
]
