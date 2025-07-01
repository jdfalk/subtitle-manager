// file: sdks/javascript/src/types.ts
// version: 1.0.0
// guid: 550e8400-e29b-41d4-a716-446655440013

/**
 * Type definitions for the Subtitle Manager SDK
 */

// Enums
export enum LogLevel {
  DEBUG = 'debug',
  INFO = 'info',
  WARN = 'warn',
  ERROR = 'error',
}

export enum OperationType {
  DOWNLOAD = 'download',
  CONVERT = 'convert',
  TRANSLATE = 'translate',
  EXTRACT = 'extract',
}

export enum OperationStatus {
  SUCCESS = 'success',
  FAILED = 'failed',
  PENDING = 'pending',
}

export enum UserRole {
  READ = 'read',
  BASIC = 'basic',
  ADMIN = 'admin',
}

export enum TranslationProvider {
  GOOGLE = 'google',
  OPENAI = 'openai',
}

// Base interfaces
export interface ApiResponse<T = any> {
  data: T;
  status: number;
  statusText: string;
  headers: Record<string, string>;
}

export interface ApiError {
  error: string;
  message: string;
}

export interface PaginatedResponse<T> {
  items: T[];
  total: number;
  page: number;
  limit: number;
}

// Authentication types
export interface LoginRequest {
  username: string;
  password: string;
}

export interface LoginResponse {
  user_id: number;
  username: string;
  role: UserRole;
}

// System types
export interface SystemInfo {
  go_version: string;
  os: string;
  arch: string;
  goroutines: number;
  disk_free: number;
  disk_total: number;
  memory_usage?: number;
  uptime?: string;
  version?: string;
}

export interface LogEntry {
  timestamp: string;
  level: LogLevel;
  component: string;
  message: string;
  fields: Record<string, any>;
}

// History types
export interface HistoryItem {
  id: number;
  type: OperationType;
  file_path: string;
  subtitle_path?: string;
  language?: string;
  provider?: string;
  status: OperationStatus;
  created_at: string;
  user_id: number;
  error_message?: string;
}

// Scan types
export interface ScanStatus {
  scanning: boolean;
  progress: number;
  current_path?: string;
  files_processed?: number;
  files_total?: number;
  start_time?: string;
  estimated_completion?: string;
}

export interface ScanResult {
  scan_id: string;
}

// Download types
export interface DownloadRequest {
  path: string;
  language: string;
  providers?: string[];
}

export interface DownloadResult {
  success: boolean;
  subtitle_path?: string;
  provider?: string;
}

// OAuth types
export interface OAuthCredentials {
  client_id: string;
  client_secret: string;
  redirect_url?: string;
}

// File operation types
export interface ConvertRequest {
  file: File | Blob;
}

export interface TranslateRequest {
  file: File | Blob;
  language: string;
  provider?: TranslationProvider;
}

export interface ExtractRequest {
  file: File | Blob;
  language?: string;
  track?: number;
}

// Query parameters
export interface HistoryQueryParams {
  page?: number;
  limit?: number;
  type?: OperationType;
  start_date?: string;
  end_date?: string;
}

export interface LogsQueryParams {
  level?: LogLevel;
  limit?: number;
}

export interface ScanRequestBody {
  path?: string;
  force?: boolean;
}

// Client configuration
export interface ClientConfig {
  baseURL: string;
  apiKey?: string;
  sessionCookie?: string;
  timeout?: number;
  maxRetries?: number;
  retryDelay?: number;
  userAgent?: string;
  debug?: boolean;
}

// HTTP method types
export type HttpMethod = 'GET' | 'POST' | 'PUT' | 'PATCH' | 'DELETE';

// Request configuration
export interface RequestConfig {
  method: HttpMethod;
  url: string;
  params?: Record<string, any>;
  data?: any;
  headers?: Record<string, string>;
  timeout?: number;
  responseType?: 'json' | 'blob' | 'text' | 'arraybuffer';
}

// Retry configuration
export interface RetryConfig {
  retries: number;
  retryDelay: number;
  retryCondition?: (error: any) => boolean;
}
