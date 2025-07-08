# file: sdks/python/subtitle_manager_sdk/models.py
# version: 1.0.0
# guid: 550e8400-e29b-41d4-a716-446655440006

"""
Data models for the Subtitle Manager SDK.
"""

from datetime import datetime
from typing import Optional, Dict, Any, List
from dataclasses import dataclass
from enum import Enum


class LogLevel(str, Enum):
    """Log levels."""
    DEBUG = "debug"
    INFO = "info"
    WARN = "warn"
    ERROR = "error"


class OperationType(str, Enum):
    """Operation types for history."""
    DOWNLOAD = "download"
    CONVERT = "convert"
    TRANSLATE = "translate"
    EXTRACT = "extract"


class OperationStatus(str, Enum):
    """Operation status for history."""
    SUCCESS = "success"
    FAILED = "failed"
    PENDING = "pending"


class UserRole(str, Enum):
    """User permission levels."""
    READ = "read"
    BASIC = "basic"
    ADMIN = "admin"


class TranslationProvider(str, Enum):
    """Translation providers."""
    GOOGLE = "google"
    OPENAI = "openai"


@dataclass
class SystemInfo:
    """System information model."""
    go_version: str
    os: str
    arch: str
    goroutines: int
    disk_free: int
    disk_total: int
    memory_usage: Optional[int] = None
    uptime: Optional[str] = None
    version: Optional[str] = None
    
    @classmethod
    def from_dict(cls, data: Dict[str, Any]) -> 'SystemInfo':
        """Create SystemInfo from dictionary."""
        return cls(
            go_version=data["go_version"],
            os=data["os"],
            arch=data["arch"],
            goroutines=data["goroutines"],
            disk_free=data["disk_free"],
            disk_total=data["disk_total"],
            memory_usage=data.get("memory_usage"),
            uptime=data.get("uptime"),
            version=data.get("version"),
        )


@dataclass
class HistoryItem:
    """History item model."""
    id: int
    type: OperationType
    file_path: str
    subtitle_path: Optional[str]
    language: Optional[str]
    provider: Optional[str]
    status: OperationStatus
    created_at: datetime
    user_id: int
    error_message: Optional[str] = None
    
    @classmethod
    def from_dict(cls, data: Dict[str, Any]) -> 'HistoryItem':
        """Create HistoryItem from dictionary."""
        return cls(
            id=data["id"],
            type=OperationType(data["type"]),
            file_path=data["file_path"],
            subtitle_path=data.get("subtitle_path"),
            language=data.get("language"),
            provider=data.get("provider"),
            status=OperationStatus(data["status"]),
            created_at=datetime.fromisoformat(data["created_at"].replace('Z', '+00:00')),
            user_id=data["user_id"],
            error_message=data.get("error_message"),
        )


@dataclass
class ScanStatus:
    """Library scan status model."""
    scanning: bool
    progress: float
    current_path: Optional[str] = None
    files_processed: Optional[int] = None
    files_total: Optional[int] = None
    start_time: Optional[datetime] = None
    estimated_completion: Optional[datetime] = None
    
    @classmethod
    def from_dict(cls, data: Dict[str, Any]) -> 'ScanStatus':
        """Create ScanStatus from dictionary."""
        start_time = None
        if data.get("start_time"):
            start_time = datetime.fromisoformat(data["start_time"].replace('Z', '+00:00'))
            
        estimated_completion = None
        if data.get("estimated_completion"):
            estimated_completion = datetime.fromisoformat(data["estimated_completion"].replace('Z', '+00:00'))
        
        return cls(
            scanning=data["scanning"],
            progress=data["progress"],
            current_path=data.get("current_path"),
            files_processed=data.get("files_processed"),
            files_total=data.get("files_total"),
            start_time=start_time,
            estimated_completion=estimated_completion,
        )


@dataclass
class OAuthCredentials:
    """OAuth2 credentials model."""
    client_id: str
    client_secret: str
    redirect_url: Optional[str] = None
    
    @classmethod
    def from_dict(cls, data: Dict[str, Any]) -> 'OAuthCredentials':
        """Create OAuthCredentials from dictionary."""
        return cls(
            client_id=data["client_id"],
            client_secret=data["client_secret"],
            redirect_url=data.get("redirect_url"),
        )


@dataclass
class LogEntry:
    """Log entry model."""
    timestamp: datetime
    level: LogLevel
    component: str
    message: str
    fields: Dict[str, Any]
    
    @classmethod
    def from_dict(cls, data: Dict[str, Any]) -> 'LogEntry':
        """Create LogEntry from dictionary."""
        return cls(
            timestamp=datetime.fromisoformat(data["timestamp"].replace('Z', '+00:00')),
            level=LogLevel(data["level"]),
            component=data["component"],
            message=data["message"],
            fields=data.get("fields", {}),
        )


@dataclass
class LoginResponse:
    """Login response model."""
    user_id: int
    username: str
    role: UserRole
    
    @classmethod
    def from_dict(cls, data: Dict[str, Any]) -> 'LoginResponse':
        """Create LoginResponse from dictionary."""
        return cls(
            user_id=data["user_id"],
            username=data["username"],
            role=UserRole(data["role"]),
        )


@dataclass
class DownloadResult:
    """Download operation result."""
    success: bool
    subtitle_path: Optional[str] = None
    provider: Optional[str] = None
    
    @classmethod
    def from_dict(cls, data: Dict[str, Any]) -> 'DownloadResult':
        """Create DownloadResult from dictionary."""
        return cls(
            success=data["success"],
            subtitle_path=data.get("subtitle_path"),
            provider=data.get("provider"),
        )


@dataclass
class ScanResult:
    """Scan operation result."""
    scan_id: str
    
    @classmethod
    def from_dict(cls, data: Dict[str, Any]) -> 'ScanResult':
        """Create ScanResult from dictionary."""
        return cls(scan_id=data["scan_id"])


@dataclass
class PaginatedResponse:
    """Paginated response wrapper."""
    items: List[Any]
    total: int
    page: int
    limit: int
    
    @classmethod
    def from_dict(cls, data: Dict[str, Any], item_class=None) -> 'PaginatedResponse':
        """Create PaginatedResponse from dictionary."""
        items = data["items"]
        if item_class and hasattr(item_class, 'from_dict'):
            items = [item_class.from_dict(item) for item in items]
        
        return cls(
            items=items,
            total=data["total"],
            page=data["page"],
            limit=data["limit"],
        )