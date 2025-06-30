// file: sdks/javascript/src/models.ts
// version: 1.0.0
// guid: 550e8400-e29b-41d4-a716-446655440015

/**
 * Data model classes for the Subtitle Manager SDK
 */

import {
  SystemInfo as ISystemInfo,
  HistoryItem as IHistoryItem,
  ScanStatus as IScanStatus,
  LogEntry as ILogEntry,
  LoginResponse as ILoginResponse,
  DownloadResult as IDownloadResult,
  ScanResult as IScanResult,
  OAuthCredentials as IOAuthCredentials,
  PaginatedResponse as IPaginatedResponse,
  OperationType,
  OperationStatus,
  LogLevel,
  UserRole,
} from './types';

export class SystemInfo implements ISystemInfo {
  public readonly go_version: string;
  public readonly os: string;
  public readonly arch: string;
  public readonly goroutines: number;
  public readonly disk_free: number;
  public readonly disk_total: number;
  public readonly memory_usage?: number;
  public readonly uptime?: string;
  public readonly version?: string;

  constructor(data: ISystemInfo) {
    this.go_version = data.go_version;
    this.os = data.os;
    this.arch = data.arch;
    this.goroutines = data.goroutines;
    this.disk_free = data.disk_free;
    this.disk_total = data.disk_total;
    this.memory_usage = data.memory_usage;
    this.uptime = data.uptime;
    this.version = data.version;
  }

  public get diskUsagePercent(): number {
    return ((this.disk_total - this.disk_free) / this.disk_total) * 100;
  }

  public formatDiskSize(bytes: number): string {
    const sizes = ['B', 'KB', 'MB', 'GB', 'TB'];
    if (bytes === 0) return '0 B';
    const i = Math.floor(Math.log(bytes) / Math.log(1024));
    return Math.round(bytes / Math.pow(1024, i) * 100) / 100 + ' ' + sizes[i];
  }

  public get diskFreeFormatted(): string {
    return this.formatDiskSize(this.disk_free);
  }

  public get diskTotalFormatted(): string {
    return this.formatDiskSize(this.disk_total);
  }
}

export class HistoryItem implements IHistoryItem {
  public readonly id: number;
  public readonly type: OperationType;
  public readonly file_path: string;
  public readonly subtitle_path?: string;
  public readonly language?: string;
  public readonly provider?: string;
  public readonly status: OperationStatus;
  public readonly created_at: string;
  public readonly user_id: number;
  public readonly error_message?: string;

  constructor(data: IHistoryItem) {
    this.id = data.id;
    this.type = data.type;
    this.file_path = data.file_path;
    this.subtitle_path = data.subtitle_path;
    this.language = data.language;
    this.provider = data.provider;
    this.status = data.status;
    this.created_at = data.created_at;
    this.user_id = data.user_id;
    this.error_message = data.error_message;
  }

  public get createdAtDate(): Date {
    return new Date(this.created_at);
  }

  public get isSuccess(): boolean {
    return this.status === OperationStatus.SUCCESS;
  }

  public get isFailed(): boolean {
    return this.status === OperationStatus.FAILED;
  }

  public get isPending(): boolean {
    return this.status === OperationStatus.PENDING;
  }
}

export class ScanStatus implements IScanStatus {
  public readonly scanning: boolean;
  public readonly progress: number;
  public readonly current_path?: string;
  public readonly files_processed?: number;
  public readonly files_total?: number;
  public readonly start_time?: string;
  public readonly estimated_completion?: string;

  constructor(data: IScanStatus) {
    this.scanning = data.scanning;
    this.progress = data.progress;
    this.current_path = data.current_path;
    this.files_processed = data.files_processed;
    this.files_total = data.files_total;
    this.start_time = data.start_time;
    this.estimated_completion = data.estimated_completion;
  }

  public get progressPercent(): number {
    return Math.round(this.progress * 100);
  }

  public get startTimeDate(): Date | undefined {
    return this.start_time ? new Date(this.start_time) : undefined;
  }

  public get estimatedCompletionDate(): Date | undefined {
    return this.estimated_completion ? new Date(this.estimated_completion) : undefined;
  }

  public get remainingFiles(): number | undefined {
    if (this.files_total && this.files_processed) {
      return this.files_total - this.files_processed;
    }
    return undefined;
  }
}

export class LogEntry implements ILogEntry {
  public readonly timestamp: string;
  public readonly level: LogLevel;
  public readonly component: string;
  public readonly message: string;
  public readonly fields: Record<string, any>;

  constructor(data: ILogEntry) {
    this.timestamp = data.timestamp;
    this.level = data.level;
    this.component = data.component;
    this.message = data.message;
    this.fields = data.fields;
  }

  public get timestampDate(): Date {
    return new Date(this.timestamp);
  }

  public get isError(): boolean {
    return this.level === LogLevel.ERROR;
  }

  public get isWarning(): boolean {
    return this.level === LogLevel.WARN;
  }
}

export class LoginResponse implements ILoginResponse {
  public readonly user_id: number;
  public readonly username: string;
  public readonly role: UserRole;

  constructor(data: ILoginResponse) {
    this.user_id = data.user_id;
    this.username = data.username;
    this.role = data.role;
  }

  public get isAdmin(): boolean {
    return this.role === UserRole.ADMIN;
  }

  public get hasBasicAccess(): boolean {
    return this.role === UserRole.BASIC || this.role === UserRole.ADMIN;
  }

  public get hasReadAccess(): boolean {
    return this.role === UserRole.READ || this.role === UserRole.BASIC || this.role === UserRole.ADMIN;
  }
}

export class DownloadResult implements IDownloadResult {
  public readonly success: boolean;
  public readonly subtitle_path?: string;
  public readonly provider?: string;

  constructor(data: IDownloadResult) {
    this.success = data.success;
    this.subtitle_path = data.subtitle_path;
    this.provider = data.provider;
  }
}

export class ScanResult implements IScanResult {
  public readonly scan_id: string;

  constructor(data: IScanResult) {
    this.scan_id = data.scan_id;
  }
}

export class OAuthCredentials implements IOAuthCredentials {
  public readonly client_id: string;
  public readonly client_secret: string;
  public readonly redirect_url?: string;

  constructor(data: IOAuthCredentials) {
    this.client_id = data.client_id;
    this.client_secret = data.client_secret;
    this.redirect_url = data.redirect_url;
  }
}

export class PaginatedResponse<T> implements IPaginatedResponse<T> {
  public readonly items: T[];
  public readonly total: number;
  public readonly page: number;
  public readonly limit: number;

  constructor(data: IPaginatedResponse<T>) {
    this.items = data.items;
    this.total = data.total;
    this.page = data.page;
    this.limit = data.limit;
  }

  public get hasNextPage(): boolean {
    return this.page * this.limit < this.total;
  }

  public get hasPreviousPage(): boolean {
    return this.page > 1;
  }

  public get totalPages(): number {
    return Math.ceil(this.total / this.limit);
  }

  public get startIndex(): number {
    return (this.page - 1) * this.limit + 1;
  }

  public get endIndex(): number {
    return Math.min(this.page * this.limit, this.total);
  }
}