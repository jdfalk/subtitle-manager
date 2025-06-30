# file: sdks/python/subtitle_manager_sdk/client.py
# version: 1.0.0
# guid: 550e8400-e29b-41d4-a716-446655440007

"""
Main client class for the Subtitle Manager SDK.
"""

import os
import time
import json
import logging
from typing import Optional, Dict, Any, List, Union, BinaryIO, Tuple
from urllib.parse import urljoin, urlparse
from datetime import datetime, timedelta

try:
    import requests
    from requests.adapters import HTTPAdapter
    from requests.packages.urllib3.util.retry import Retry
except ImportError:
    raise ImportError("requests library is required. Install with: pip install requests")

from .exceptions import (
    SubtitleManagerError,
    AuthenticationError,
    AuthorizationError,
    NotFoundError,
    RateLimitError,
    APIError,
    NetworkError,
    TimeoutError,
    ValidationError,
)
from .models import (
    SystemInfo,
    HistoryItem,
    ScanStatus,
    OAuthCredentials,
    LogEntry,
    LoginResponse,
    DownloadResult,
    ScanResult,
    PaginatedResponse,
    OperationType,
    LogLevel,
    TranslationProvider,
)


class SubtitleManagerClient:
    """
    Client for the Subtitle Manager API.
    
    Provides type-safe access to all API endpoints with automatic retry,
    error handling, and pagination support.
    """
    
    def __init__(
        self,
        base_url: str = "http://localhost:8080",
        api_key: Optional[str] = None,
        session_cookie: Optional[str] = None,
        timeout: int = 30,
        max_retries: int = 3,
        retry_backoff: float = 1.0,
        user_agent: Optional[str] = None,
    ):
        """
        Initialize the Subtitle Manager client.
        
        Args:
            base_url: Base URL of the Subtitle Manager API
            api_key: API key for authentication
            session_cookie: Session cookie for authentication
            timeout: Request timeout in seconds
            max_retries: Maximum number of retry attempts
            retry_backoff: Backoff factor for retries
            user_agent: Custom user agent string
        """
        self.base_url = base_url.rstrip('/')
        self.api_key = api_key or os.getenv('SUBTITLE_MANAGER_API_KEY')
        self.session_cookie = session_cookie
        self.timeout = timeout
        self.max_retries = max_retries
        self.retry_backoff = retry_backoff
        
        # Configure session
        self.session = requests.Session()
        
        # Setup retry strategy
        retry_strategy = Retry(
            total=max_retries,
            backoff_factor=retry_backoff,
            status_forcelist=[429, 500, 502, 503, 504],
            method_whitelist=["HEAD", "GET", "OPTIONS", "POST", "PUT", "PATCH", "DELETE"],
        )
        adapter = HTTPAdapter(max_retries=retry_strategy)
        self.session.mount("http://", adapter)
        self.session.mount("https://", adapter)
        
        # Set default headers
        headers = {
            "User-Agent": user_agent or f"subtitle-manager-python-sdk/1.0.0",
            "Accept": "application/json",
        }
        
        if self.api_key:
            headers["X-API-Key"] = self.api_key
            
        if self.session_cookie:
            self.session.cookies.set("session", self.session_cookie)
            
        self.session.headers.update(headers)
        
        # Configure logging
        self.logger = logging.getLogger(__name__)
    
    def _make_request(
        self,
        method: str,
        endpoint: str,
        params: Optional[Dict[str, Any]] = None,
        data: Optional[Dict[str, Any]] = None,
        files: Optional[Dict[str, Any]] = None,
        json_data: Optional[Dict[str, Any]] = None,
        headers: Optional[Dict[str, str]] = None,
        timeout: Optional[int] = None,
    ) -> requests.Response:
        """
        Make an HTTP request to the API.
        
        Args:
            method: HTTP method
            endpoint: API endpoint (without base URL)
            params: Query parameters
            data: Form data
            files: Files to upload
            json_data: JSON data
            headers: Additional headers
            timeout: Request timeout
            
        Returns:
            Response object
            
        Raises:
            Various SubtitleManagerError subclasses based on response
        """
        url = urljoin(self.base_url, endpoint)
        
        request_kwargs = {
            "params": params,
            "timeout": timeout or self.timeout,
        }
        
        if headers:
            request_kwargs["headers"] = headers
            
        if files:
            request_kwargs["files"] = files
        elif json_data:
            request_kwargs["json"] = json_data
        elif data:
            request_kwargs["data"] = data
        
        try:
            self.logger.debug(f"Making {method} request to {url}")
            response = self.session.request(method, url, **request_kwargs)
            self._handle_response(response)
            return response
            
        except requests.exceptions.Timeout:
            raise TimeoutError(f"Request to {url} timed out after {self.timeout} seconds")
        except requests.exceptions.ConnectionError as e:
            raise NetworkError(f"Connection error: {str(e)}")
        except requests.exceptions.RequestException as e:
            raise NetworkError(f"Request failed: {str(e)}")
    
    def _handle_response(self, response: requests.Response) -> None:
        """
        Handle API response and raise appropriate exceptions.
        
        Args:
            response: Response object
            
        Raises:
            Various SubtitleManagerError subclasses based on status code
        """
        if response.status_code < 400:
            return
            
        try:
            error_data = response.json()
            error_code = error_data.get("error", "unknown_error")
            message = error_data.get("message", f"HTTP {response.status_code}")
        except (ValueError, json.JSONDecodeError):
            error_code = "unknown_error"
            message = f"HTTP {response.status_code}: {response.text}"
        
        if response.status_code == 401:
            raise AuthenticationError(message, response.status_code, error_code)
        elif response.status_code == 403:
            raise AuthorizationError(message, response.status_code, error_code)
        elif response.status_code == 404:
            raise NotFoundError(message, response.status_code, error_code)
        elif response.status_code == 429:
            retry_after = response.headers.get("Retry-After")
            retry_after_int = int(retry_after) if retry_after else None
            raise RateLimitError(message, retry_after_int, response.status_code, error_code)
        else:
            raise APIError(message, response.status_code, error_code)
    
    def _get_json(self, response: requests.Response) -> Dict[str, Any]:
        """
        Extract JSON from response.
        
        Args:
            response: Response object
            
        Returns:
            JSON data as dictionary
        """
        try:
            return response.json()
        except (ValueError, json.JSONDecodeError):
            raise APIError("Invalid JSON in response", response.status_code)
    
    # Authentication methods
    
    def login(self, username: str, password: str) -> LoginResponse:
        """
        Authenticate with username and password.
        
        Args:
            username: Username or email
            password: Password
            
        Returns:
            Login response with user information
        """
        response = self._make_request(
            "POST",
            "/api/login",
            json_data={"username": username, "password": password}
        )
        
        # Extract session cookie
        session_cookie = response.cookies.get("session")
        if session_cookie:
            self.session_cookie = session_cookie
            self.session.cookies.set("session", session_cookie)
        
        return LoginResponse.from_dict(self._get_json(response))
    
    def logout(self) -> None:
        """Logout and invalidate current session."""
        self._make_request("POST", "/api/logout")
        self.session_cookie = None
        self.session.cookies.clear()
    
    # System information methods
    
    def get_system_info(self) -> SystemInfo:
        """
        Get system information.
        
        Returns:
            System information
        """
        response = self._make_request("GET", "/api/system")
        return SystemInfo.from_dict(self._get_json(response))
    
    def get_logs(
        self,
        level: Optional[LogLevel] = None,
        limit: int = 100
    ) -> List[LogEntry]:
        """
        Get application logs.
        
        Args:
            level: Minimum log level
            limit: Maximum number of entries
            
        Returns:
            List of log entries
        """
        params = {"limit": limit}
        if level:
            params["level"] = level.value
            
        response = self._make_request("GET", "/api/logs", params=params)
        log_data = self._get_json(response)
        return [LogEntry.from_dict(entry) for entry in log_data]
    
    # Configuration methods
    
    def get_config(self) -> Dict[str, Any]:
        """
        Get application configuration.
        
        Returns:
            Configuration dictionary
        """
        response = self._make_request("GET", "/api/config")
        return self._get_json(response)
    
    # Subtitle operations
    
    def convert_subtitle(
        self,
        file: Union[str, BinaryIO],
        filename: Optional[str] = None
    ) -> bytes:
        """
        Convert subtitle file to SRT format.
        
        Args:
            file: File path or file-like object
            filename: Original filename (required for file-like objects)
            
        Returns:
            Converted SRT file content
        """
        if isinstance(file, str):
            with open(file, 'rb') as f:
                files = {"file": (os.path.basename(file), f, "application/octet-stream")}
                response = self._make_request("POST", "/api/convert", files=files)
        else:
            if not filename:
                raise ValidationError("filename required for file-like objects")
            files = {"file": (filename, file, "application/octet-stream")}
            response = self._make_request("POST", "/api/convert", files=files)
        
        return response.content
    
    def translate_subtitle(
        self,
        file: Union[str, BinaryIO],
        language: str,
        provider: TranslationProvider = TranslationProvider.GOOGLE,
        filename: Optional[str] = None
    ) -> bytes:
        """
        Translate subtitle file to target language.
        
        Args:
            file: File path or file-like object
            language: Target language code (ISO 639-1)
            provider: Translation provider
            filename: Original filename (required for file-like objects)
            
        Returns:
            Translated SRT file content
        """
        if isinstance(file, str):
            with open(file, 'rb') as f:
                files = {"file": (os.path.basename(file), f, "application/octet-stream")}
                data = {"language": language, "provider": provider.value}
                response = self._make_request("POST", "/api/translate", files=files, data=data)
        else:
            if not filename:
                raise ValidationError("filename required for file-like objects")
            files = {"file": (filename, file, "application/octet-stream")}
            data = {"language": language, "provider": provider.value}
            response = self._make_request("POST", "/api/translate", files=files, data=data)
        
        return response.content
    
    def extract_subtitles(
        self,
        file: Union[str, BinaryIO],
        language: Optional[str] = None,
        track: int = 0,
        filename: Optional[str] = None
    ) -> bytes:
        """
        Extract embedded subtitles from video file.
        
        Args:
            file: File path or file-like object
            language: Preferred subtitle language
            track: Subtitle track number (0-based)
            filename: Original filename (required for file-like objects)
            
        Returns:
            Extracted SRT file content
        """
        data = {"track": track}
        if language:
            data["language"] = language
            
        if isinstance(file, str):
            with open(file, 'rb') as f:
                files = {"file": (os.path.basename(file), f, "application/octet-stream")}
                response = self._make_request("POST", "/api/extract", files=files, data=data)
        else:
            if not filename:
                raise ValidationError("filename required for file-like objects")
            files = {"file": (filename, file, "application/octet-stream")}
            response = self._make_request("POST", "/api/extract", files=files, data=data)
        
        return response.content
    
    # Download and scanning
    
    def download_subtitles(
        self,
        path: str,
        language: str,
        providers: Optional[List[str]] = None
    ) -> DownloadResult:
        """
        Download subtitles for media file.
        
        Args:
            path: Media file path
            language: Subtitle language (ISO 639-1)
            providers: Preferred providers (optional)
            
        Returns:
            Download result
        """
        data = {"path": path, "language": language}
        if providers:
            data["providers"] = providers
            
        response = self._make_request("POST", "/api/download", json_data=data)
        return DownloadResult.from_dict(self._get_json(response))
    
    def start_library_scan(
        self,
        path: Optional[str] = None,
        force: bool = False
    ) -> ScanResult:
        """
        Start library scan.
        
        Args:
            path: Specific path to scan (optional)
            force: Force rescan of existing files
            
        Returns:
            Scan result with job ID
        """
        data = {"force": force}
        if path:
            data["path"] = path
            
        response = self._make_request("POST", "/api/scan", json_data=data)
        return ScanResult.from_dict(self._get_json(response))
    
    def get_scan_status(self) -> ScanStatus:
        """
        Get current scan status.
        
        Returns:
            Scan status information
        """
        response = self._make_request("GET", "/api/scan/status")
        return ScanStatus.from_dict(self._get_json(response))
    
    # History methods
    
    def get_history(
        self,
        page: int = 1,
        limit: int = 20,
        operation_type: Optional[OperationType] = None,
        start_date: Optional[datetime] = None,
        end_date: Optional[datetime] = None
    ) -> PaginatedResponse:
        """
        Get operation history.
        
        Args:
            page: Page number
            limit: Items per page
            operation_type: Filter by operation type
            start_date: Filter operations after this date
            end_date: Filter operations before this date
            
        Returns:
            Paginated history response
        """
        params = {"page": page, "limit": limit}
        
        if operation_type:
            params["type"] = operation_type.value
        if start_date:
            params["start_date"] = start_date.isoformat()
        if end_date:
            params["end_date"] = end_date.isoformat()
            
        response = self._make_request("GET", "/api/history", params=params)
        data = self._get_json(response)
        return PaginatedResponse.from_dict(data, HistoryItem)
    
    # OAuth2 management (admin only)
    
    def generate_github_oauth(self) -> OAuthCredentials:
        """
        Generate GitHub OAuth2 credentials (admin only).
        
        Returns:
            OAuth2 credentials
        """
        response = self._make_request("POST", "/api/oauth/github/generate")
        return OAuthCredentials.from_dict(self._get_json(response))
    
    def regenerate_github_oauth(self) -> OAuthCredentials:
        """
        Regenerate GitHub OAuth2 client secret (admin only).
        
        Returns:
            Updated OAuth2 credentials
        """
        response = self._make_request("POST", "/api/oauth/github/regenerate")
        return OAuthCredentials.from_dict(self._get_json(response))
    
    def reset_github_oauth(self) -> None:
        """Reset GitHub OAuth2 configuration (admin only)."""
        self._make_request("POST", "/api/oauth/github/reset")
    
    # Utility methods
    
    def health_check(self) -> bool:
        """
        Check if the API is accessible.
        
        Returns:
            True if API is accessible, False otherwise
        """
        try:
            self.get_system_info()
            return True
        except Exception:
            return False
    
    def wait_for_scan_completion(
        self,
        poll_interval: int = 5,
        timeout: int = 3600
    ) -> ScanStatus:
        """
        Wait for library scan to complete.
        
        Args:
            poll_interval: Polling interval in seconds
            timeout: Maximum wait time in seconds
            
        Returns:
            Final scan status
            
        Raises:
            TimeoutError: If scan doesn't complete within timeout
        """
        start_time = time.time()
        
        while time.time() - start_time < timeout:
            status = self.get_scan_status()
            if not status.scanning:
                return status
            time.sleep(poll_interval)
        
        raise TimeoutError(f"Scan did not complete within {timeout} seconds")
    
    def __enter__(self):
        """Context manager entry."""
        return self
    
    def __exit__(self, exc_type, exc_val, exc_tb):
        """Context manager exit."""
        self.session.close()
    
    def __repr__(self):
        """String representation."""
        return f"SubtitleManagerClient(base_url='{self.base_url}')"