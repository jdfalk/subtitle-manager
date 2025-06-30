# file: sdks/python/subtitle_manager_sdk/exceptions.py
# version: 1.0.0
# guid: 550e8400-e29b-41d4-a716-446655440005

"""
Exception classes for the Subtitle Manager SDK.
"""

from typing import Optional, Dict, Any


class SubtitleManagerError(Exception):
    """Base exception for all Subtitle Manager SDK errors."""
    
    def __init__(self, message: str, status_code: Optional[int] = None, 
                 error_code: Optional[str] = None, details: Optional[Dict[str, Any]] = None):
        self.message = message
        self.status_code = status_code
        self.error_code = error_code
        self.details = details or {}
        super().__init__(message)


class AuthenticationError(SubtitleManagerError):
    """Raised when authentication fails (401)."""
    pass


class AuthorizationError(SubtitleManagerError):
    """Raised when user lacks required permissions (403)."""
    pass


class NotFoundError(SubtitleManagerError):
    """Raised when requested resource is not found (404)."""
    pass


class RateLimitError(SubtitleManagerError):
    """Raised when rate limit is exceeded (429)."""
    
    def __init__(self, message: str, retry_after: Optional[int] = None, **kwargs):
        self.retry_after = retry_after
        super().__init__(message, **kwargs)


class APIError(SubtitleManagerError):
    """Raised for general API errors (4xx, 5xx)."""
    pass


class ValidationError(SubtitleManagerError):
    """Raised when request validation fails."""
    pass


class NetworkError(SubtitleManagerError):
    """Raised when network-related errors occur."""
    pass


class TimeoutError(SubtitleManagerError):
    """Raised when request times out."""
    pass