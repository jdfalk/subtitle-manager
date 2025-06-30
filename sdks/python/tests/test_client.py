# file: sdks/python/tests/test_client.py
# version: 1.0.0
# guid: 550e8400-e29b-41d4-a716-446655440010

"""
Tests for the Subtitle Manager SDK client.
"""

import pytest
import responses
from datetime import datetime
from io import BytesIO

from subtitle_manager_sdk import SubtitleManagerClient
from subtitle_manager_sdk.exceptions import (
    AuthenticationError,
    AuthorizationError,
    NotFoundError,
    RateLimitError,
    APIError,
)
from subtitle_manager_sdk.models import (
    SystemInfo,
    HistoryItem,
    ScanStatus,
    LogEntry,
    LoginResponse,
    DownloadResult,
    OperationType,
    OperationStatus,
    LogLevel,
    UserRole,
)


class TestSubtitleManagerClient:
    """Test cases for SubtitleManagerClient."""
    
    def setup_method(self):
        """Set up test fixtures."""
        self.client = SubtitleManagerClient(
            base_url="http://test.example.com",
            api_key="test-api-key"
        )
    
    @responses.activate
    def test_authentication_with_api_key(self):
        """Test API key authentication."""
        responses.add(
            responses.GET,
            "http://test.example.com/api/system",
            json={"go_version": "go1.24.0", "os": "linux", "arch": "amd64", 
                  "goroutines": 10, "disk_free": 1000000, "disk_total": 2000000},
            status=200
        )
        
        system_info = self.client.get_system_info()
        assert system_info.go_version == "go1.24.0"
        assert system_info.os == "linux"
        
        # Check that API key header was sent
        assert len(responses.calls) == 1
        assert responses.calls[0].request.headers["X-API-Key"] == "test-api-key"
    
    @responses.activate
    def test_login_success(self):
        """Test successful login."""
        responses.add(
            responses.POST,
            "http://test.example.com/api/login",
            json={"user_id": 1, "username": "testuser", "role": "admin"},
            status=200,
            headers={"Set-Cookie": "session=test-session-cookie; HttpOnly"}
        )
        
        login_response = self.client.login("testuser", "password")
        
        assert login_response.user_id == 1
        assert login_response.username == "testuser"
        assert login_response.role == UserRole.ADMIN
        assert self.client.session_cookie == "test-session-cookie"
    
    @responses.activate
    def test_authentication_error(self):
        """Test authentication error handling."""
        responses.add(
            responses.GET,
            "http://test.example.com/api/system",
            json={"error": "unauthorized", "message": "Authentication required"},
            status=401
        )
        
        with pytest.raises(AuthenticationError) as exc_info:
            self.client.get_system_info()
        
        assert exc_info.value.status_code == 401
        assert exc_info.value.error_code == "unauthorized"
    
    @responses.activate
    def test_authorization_error(self):
        """Test authorization error handling."""
        responses.add(
            responses.POST,
            "http://test.example.com/api/oauth/github/generate",
            json={"error": "forbidden", "message": "Admin access required"},
            status=403
        )
        
        with pytest.raises(AuthorizationError) as exc_info:
            self.client.generate_github_oauth()
        
        assert exc_info.value.status_code == 403
        assert exc_info.value.error_code == "forbidden"
    
    @responses.activate
    def test_rate_limit_error(self):
        """Test rate limit error handling."""
        responses.add(
            responses.GET,
            "http://test.example.com/api/system",
            json={"error": "rate_limited", "message": "Rate limit exceeded"},
            status=429,
            headers={"Retry-After": "60"}
        )
        
        with pytest.raises(RateLimitError) as exc_info:
            self.client.get_system_info()
        
        assert exc_info.value.status_code == 429
        assert exc_info.value.retry_after == 60
    
    @responses.activate
    def test_not_found_error(self):
        """Test not found error handling."""
        responses.add(
            responses.GET,
            "http://test.example.com/api/nonexistent",
            json={"error": "not_found", "message": "Resource not found"},
            status=404
        )
        
        with pytest.raises(NotFoundError) as exc_info:
            self.client._make_request("GET", "/api/nonexistent")
        
        assert exc_info.value.status_code == 404
    
    @responses.activate
    def test_get_system_info(self):
        """Test getting system information."""
        system_data = {
            "go_version": "go1.24.0",
            "os": "linux",
            "arch": "amd64",
            "goroutines": 15,
            "disk_free": 5000000000,
            "disk_total": 10000000000,
            "memory_usage": 512000000,
            "uptime": "2 days",
            "version": "1.0.0"
        }
        
        responses.add(
            responses.GET,
            "http://test.example.com/api/system",
            json=system_data,
            status=200
        )
        
        system_info = self.client.get_system_info()
        
        assert isinstance(system_info, SystemInfo)
        assert system_info.go_version == "go1.24.0"
        assert system_info.goroutines == 15
        assert system_info.disk_free == 5000000000
        assert system_info.version == "1.0.0"
    
    @responses.activate
    def test_get_logs(self):
        """Test getting application logs."""
        log_data = [
            {
                "timestamp": "2024-01-01T12:00:00Z",
                "level": "info",
                "component": "webserver",
                "message": "Server started",
                "fields": {"port": 8080}
            },
            {
                "timestamp": "2024-01-01T12:01:00Z",
                "level": "error",
                "component": "auth",
                "message": "Login failed",
                "fields": {"username": "baduser"}
            }
        ]
        
        responses.add(
            responses.GET,
            "http://test.example.com/api/logs",
            json=log_data,
            status=200
        )
        
        logs = self.client.get_logs(level=LogLevel.INFO, limit=50)
        
        assert len(logs) == 2
        assert all(isinstance(log, LogEntry) for log in logs)
        assert logs[0].level == LogLevel.INFO
        assert logs[0].component == "webserver"
        assert logs[1].level == LogLevel.ERROR
        
        # Check query parameters
        request = responses.calls[0].request
        assert "level=info" in request.url
        assert "limit=50" in request.url
    
    @responses.activate
    def test_convert_subtitle_with_file_path(self):
        """Test converting subtitle file from file path."""
        srt_content = b"1\n00:00:01,000 --> 00:00:03,000\nHello World\n\n"
        
        responses.add(
            responses.POST,
            "http://test.example.com/api/convert",
            body=srt_content,
            status=200,
            headers={"Content-Type": "application/x-subrip"}
        )
        
        # Mock file reading
        import tempfile
        import os
        
        with tempfile.NamedTemporaryFile(suffix=".vtt", delete=False) as temp_file:
            temp_file.write(b"WEBVTT\n\n00:01.000 --> 00:03.000\nHello World\n")
            temp_file_path = temp_file.name
        
        try:
            result = self.client.convert_subtitle(temp_file_path)
            assert result == srt_content
        finally:
            os.unlink(temp_file_path)
    
    @responses.activate
    def test_convert_subtitle_with_file_object(self):
        """Test converting subtitle file from file object."""
        srt_content = b"1\n00:00:01,000 --> 00:00:03,000\nHello World\n\n"
        
        responses.add(
            responses.POST,
            "http://test.example.com/api/convert",
            body=srt_content,
            status=200,
            headers={"Content-Type": "application/x-subrip"}
        )
        
        file_obj = BytesIO(b"WEBVTT\n\n00:01.000 --> 00:03.000\nHello World\n")
        result = self.client.convert_subtitle(file_obj, filename="test.vtt")
        
        assert result == srt_content
    
    @responses.activate
    def test_download_subtitles(self):
        """Test downloading subtitles."""
        download_data = {
            "success": True,
            "subtitle_path": "/path/to/movie.en.srt",
            "provider": "opensubtitles"
        }
        
        responses.add(
            responses.POST,
            "http://test.example.com/api/download",
            json=download_data,
            status=200
        )
        
        result = self.client.download_subtitles(
            path="/movies/example.mkv",
            language="en",
            providers=["opensubtitles", "subscene"]
        )
        
        assert isinstance(result, DownloadResult)
        assert result.success is True
        assert result.subtitle_path == "/path/to/movie.en.srt"
        assert result.provider == "opensubtitles"
        
        # Check request data
        import json
        request_body = json.loads(responses.calls[0].request.body)
        assert request_body["path"] == "/movies/example.mkv"
        assert request_body["language"] == "en"
        assert request_body["providers"] == ["opensubtitles", "subscene"]
    
    @responses.activate
    def test_get_history(self):
        """Test getting operation history."""
        history_data = {
            "items": [
                {
                    "id": 1,
                    "type": "download",
                    "file_path": "/movies/example.mkv",
                    "subtitle_path": "/movies/example.en.srt",
                    "language": "en",
                    "provider": "opensubtitles",
                    "status": "success",
                    "created_at": "2024-01-01T12:00:00Z",
                    "user_id": 1
                },
                {
                    "id": 2,
                    "type": "convert",
                    "file_path": "/subtitles/test.vtt",
                    "subtitle_path": "/subtitles/test.srt",
                    "language": None,
                    "provider": None,
                    "status": "success",
                    "created_at": "2024-01-01T12:05:00Z",
                    "user_id": 1
                }
            ],
            "total": 2,
            "page": 1,
            "limit": 20
        }
        
        responses.add(
            responses.GET,
            "http://test.example.com/api/history",
            json=history_data,
            status=200
        )
        
        result = self.client.get_history(
            page=1,
            limit=20,
            operation_type=OperationType.DOWNLOAD
        )
        
        assert result.total == 2
        assert result.page == 1
        assert len(result.items) == 2
        
        first_item = result.items[0]
        assert isinstance(first_item, HistoryItem)
        assert first_item.type == OperationType.DOWNLOAD
        assert first_item.status == OperationStatus.SUCCESS
        assert first_item.provider == "opensubtitles"
        
        # Check query parameters
        request = responses.calls[0].request
        assert "page=1" in request.url
        assert "limit=20" in request.url
        assert "type=download" in request.url
    
    @responses.activate
    def test_scan_status(self):
        """Test getting scan status."""
        scan_data = {
            "scanning": True,
            "progress": 0.75,
            "current_path": "/movies/subfolder",
            "files_processed": 150,
            "files_total": 200,
            "start_time": "2024-01-01T12:00:00Z",
            "estimated_completion": "2024-01-01T12:30:00Z"
        }
        
        responses.add(
            responses.GET,
            "http://test.example.com/api/scan/status",
            json=scan_data,
            status=200
        )
        
        status = self.client.get_scan_status()
        
        assert isinstance(status, ScanStatus)
        assert status.scanning is True
        assert status.progress == 0.75
        assert status.files_processed == 150
        assert status.files_total == 200
        assert isinstance(status.start_time, datetime)
    
    @responses.activate
    def test_health_check_success(self):
        """Test successful health check."""
        responses.add(
            responses.GET,
            "http://test.example.com/api/system",
            json={"go_version": "go1.24.0", "os": "linux", "arch": "amd64", 
                  "goroutines": 10, "disk_free": 1000000, "disk_total": 2000000},
            status=200
        )
        
        assert self.client.health_check() is True
    
    @responses.activate
    def test_health_check_failure(self):
        """Test failed health check."""
        responses.add(
            responses.GET,
            "http://test.example.com/api/system",
            json={"error": "internal_error", "message": "Server error"},
            status=500
        )
        
        assert self.client.health_check() is False
    
    def test_context_manager(self):
        """Test client as context manager."""
        with SubtitleManagerClient("http://test.example.com") as client:
            assert client.base_url == "http://test.example.com"
        # Session should be closed after exiting context
    
    def test_repr(self):
        """Test string representation."""
        client = SubtitleManagerClient("http://test.example.com")
        assert repr(client) == "SubtitleManagerClient(base_url='http://test.example.com')"


if __name__ == "__main__":
    pytest.main([__file__])