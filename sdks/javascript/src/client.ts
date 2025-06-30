// file: sdks/javascript/src/client.ts
// version: 1.0.0
// guid: 550e8400-e29b-41d4-a716-446655440016

/**
 * Main client class for the Subtitle Manager SDK
 */

import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse } from 'axios';
import {
  ClientConfig,
  RequestConfig,
  LoginRequest,
  DownloadRequest,
  ConvertRequest,
  TranslateRequest,
  ExtractRequest,
  ScanRequestBody,
  HistoryQueryParams,
  LogsQueryParams,
  OperationType,
  LogLevel,
  TranslationProvider,
  ApiError,
} from './types';
import {
  SubtitleManagerError,
  AuthenticationError,
  AuthorizationError,
  NotFoundError,
  RateLimitError,
  ValidationError,
  NetworkError,
  TimeoutError,
  APIError,
} from './errors';
import {
  SystemInfo,
  HistoryItem,
  ScanStatus,
  LogEntry,
  LoginResponse,
  DownloadResult,
  ScanResult,
  OAuthCredentials,
  PaginatedResponse,
} from './models';

/**
 * Subtitle Manager API Client
 * 
 * Provides type-safe access to all API endpoints with automatic retry,
 * error handling, and modern async/await patterns.
 */
export class SubtitleManagerClient {
  private readonly axios: AxiosInstance;
  private readonly config: Required<ClientConfig>;

  constructor(config: ClientConfig) {
    // Set default configuration
    this.config = {
      baseURL: config.baseURL,
      apiKey: config.apiKey || process.env.SUBTITLE_MANAGER_API_KEY || '',
      sessionCookie: config.sessionCookie || '',
      timeout: config.timeout || 30000,
      maxRetries: config.maxRetries || 3,
      retryDelay: config.retryDelay || 1000,
      userAgent: config.userAgent || 'subtitle-manager-js-sdk/1.0.0',
      debug: config.debug || false,
    };

    // Create axios instance
    this.axios = axios.create({
      baseURL: this.config.baseURL,
      timeout: this.config.timeout,
      headers: {
        'User-Agent': this.config.userAgent,
        'Accept': 'application/json',
      },
    });

    // Set authentication headers
    if (this.config.apiKey) {
      this.axios.defaults.headers.common['X-API-Key'] = this.config.apiKey;
    }

    if (this.config.sessionCookie) {
      this.axios.defaults.headers.common['Cookie'] = `session=${this.config.sessionCookie}`;
    }

    // Setup request interceptor for logging
    if (this.config.debug) {
      this.axios.interceptors.request.use((config) => {
        console.log(`[SDK] ${config.method?.toUpperCase()} ${config.url}`);
        return config;
      });
    }

    // Setup response interceptor for error handling
    this.axios.interceptors.response.use(
      (response) => response,
      (error) => this.handleError(error)
    );
  }

  /**
   * Handle API errors and convert to appropriate exception types
   */
  private handleError(error: any): never {
    if (error.response) {
      const status = error.response.status;
      const data = error.response.data as ApiError;
      const message = data?.message || `HTTP ${status}`;
      const errorCode = data?.error || 'unknown_error';

      switch (status) {
        case 401:
          throw new AuthenticationError(message, status, errorCode);
        case 403:
          throw new AuthorizationError(message, status, errorCode);
        case 404:
          throw new NotFoundError(message, status, errorCode);
        case 429: {
          const retryAfter = error.response.headers['retry-after'];
          const retryAfterSeconds = retryAfter ? parseInt(retryAfter, 10) : undefined;
          throw new RateLimitError(message, retryAfterSeconds, status, errorCode);
        }
        case 400:
          throw new ValidationError(message, status, errorCode);
        default:
          throw new APIError(message, status, errorCode);
      }
    } else if (error.request) {
      throw new NetworkError(`Network error: ${error.message}`, error);
    } else if (error.code === 'ECONNABORTED') {
      throw new TimeoutError(`Request timeout after ${this.config.timeout}ms`);
    } else {
      throw new SubtitleManagerError(`Unexpected error: ${error.message}`);
    }
  }

  /**
   * Make HTTP request with automatic retry
   */
  private async request<T>(config: RequestConfig): Promise<T> {
    const axiosConfig: AxiosRequestConfig = {
      method: config.method,
      url: config.url,
      params: config.params,
      data: config.data,
      headers: config.headers,
      timeout: config.timeout || this.config.timeout,
      responseType: config.responseType || 'json',
    };

    let lastError: any;
    for (let attempt = 1; attempt <= this.config.maxRetries; attempt++) {
      try {
        const response: AxiosResponse<T> = await this.axios.request(axiosConfig);
        return response.data;
      } catch (error: any) {
        lastError = error;

        // Don't retry on certain errors
        if (
          error instanceof AuthenticationError ||
          error instanceof AuthorizationError ||
          error instanceof ValidationError ||
          attempt === this.config.maxRetries
        ) {
          throw error;
        }

        // Wait before retrying
        await this.delay(this.config.retryDelay * attempt);
      }
    }

    throw lastError;
  }

  /**
   * Utility method to add delay
   */
  private delay(ms: number): Promise<void> {
    return new Promise((resolve) => setTimeout(resolve, ms));
  }

  // Authentication methods

  /**
   * Authenticate with username and password
   */
  async login(username: string, password: string): Promise<LoginResponse> {
    const data = await this.request<LoginResponse>({
      method: 'POST',
      url: '/api/login',
      data: { username, password } as LoginRequest,
    });

    return new LoginResponse(data);
  }

  /**
   * Logout and invalidate current session
   */
  async logout(): Promise<void> {
    await this.request({
      method: 'POST',
      url: '/api/logout',
    });

    // Clear session cookie
    delete this.axios.defaults.headers.common['Cookie'];
  }

  // System information methods

  /**
   * Get system information
   */
  async getSystemInfo(): Promise<SystemInfo> {
    const data = await this.request<SystemInfo>({
      method: 'GET',
      url: '/api/system',
    });

    return new SystemInfo(data);
  }

  /**
   * Get application logs
   */
  async getLogs(params: LogsQueryParams = {}): Promise<LogEntry[]> {
    const data = await this.request<LogEntry[]>({
      method: 'GET',
      url: '/api/logs',
      params,
    });

    return data.map((entry) => new LogEntry(entry));
  }

  // Configuration methods

  /**
   * Get application configuration
   */
  async getConfig(): Promise<Record<string, any>> {
    return this.request({
      method: 'GET',
      url: '/api/config',
    });
  }

  // Subtitle operations

  /**
   * Convert subtitle file to SRT format
   */
  async convertSubtitle(file: File | Blob, filename?: string): Promise<Blob> {
    const formData = new FormData();
    formData.append('file', file, filename);

    return this.request<Blob>({
      method: 'POST',
      url: '/api/convert',
      data: formData,
      headers: { 'Content-Type': 'multipart/form-data' },
      responseType: 'blob',
    });
  }

  /**
   * Translate subtitle file to target language
   */
  async translateSubtitle(
    file: File | Blob,
    language: string,
    provider: TranslationProvider = TranslationProvider.GOOGLE,
    filename?: string
  ): Promise<Blob> {
    const formData = new FormData();
    formData.append('file', file, filename);
    formData.append('language', language);
    formData.append('provider', provider);

    return this.request<Blob>({
      method: 'POST',
      url: '/api/translate',
      data: formData,
      headers: { 'Content-Type': 'multipart/form-data' },
      responseType: 'blob',
    });
  }

  /**
   * Extract embedded subtitles from video file
   */
  async extractSubtitles(
    file: File | Blob,
    language?: string,
    track: number = 0,
    filename?: string
  ): Promise<Blob> {
    const formData = new FormData();
    formData.append('file', file, filename);
    formData.append('track', track.toString());
    if (language) {
      formData.append('language', language);
    }

    return this.request<Blob>({
      method: 'POST',
      url: '/api/extract',
      data: formData,
      headers: { 'Content-Type': 'multipart/form-data' },
      responseType: 'blob',
    });
  }

  // Download and scanning

  /**
   * Download subtitles for media file
   */
  async downloadSubtitles(
    path: string,
    language: string,
    providers?: string[]
  ): Promise<DownloadResult> {
    const data = await this.request<DownloadResult>({
      method: 'POST',
      url: '/api/download',
      data: { path, language, providers } as DownloadRequest,
    });

    return new DownloadResult(data);
  }

  /**
   * Start library scan
   */
  async startLibraryScan(path?: string, force: boolean = false): Promise<ScanResult> {
    const data = await this.request<ScanResult>({
      method: 'POST',
      url: '/api/scan',
      data: { path, force } as ScanRequestBody,
    });

    return new ScanResult(data);
  }

  /**
   * Get current scan status
   */
  async getScanStatus(): Promise<ScanStatus> {
    const data = await this.request<ScanStatus>({
      method: 'GET',
      url: '/api/scan/status',
    });

    return new ScanStatus(data);
  }

  // History methods

  /**
   * Get operation history
   */
  async getHistory(params: HistoryQueryParams = {}): Promise<PaginatedResponse<HistoryItem>> {
    const data = await this.request<PaginatedResponse<HistoryItem>>({
      method: 'GET',
      url: '/api/history',
      params,
    });

    const items = data.items.map((item) => new HistoryItem(item));
    return new PaginatedResponse({ ...data, items });
  }

  // OAuth2 management (admin only)

  /**
   * Generate GitHub OAuth2 credentials (admin only)
   */
  async generateGitHubOAuth(): Promise<OAuthCredentials> {
    const data = await this.request<OAuthCredentials>({
      method: 'POST',
      url: '/api/oauth/github/generate',
    });

    return new OAuthCredentials(data);
  }

  /**
   * Regenerate GitHub OAuth2 client secret (admin only)
   */
  async regenerateGitHubOAuth(): Promise<OAuthCredentials> {
    const data = await this.request<OAuthCredentials>({
      method: 'POST',
      url: '/api/oauth/github/regenerate',
    });

    return new OAuthCredentials(data);
  }

  /**
   * Reset GitHub OAuth2 configuration (admin only)
   */
  async resetGitHubOAuth(): Promise<void> {
    await this.request({
      method: 'POST',
      url: '/api/oauth/github/reset',
    });
  }

  // Utility methods

  /**
   * Check if the API is accessible
   */
  async healthCheck(): Promise<boolean> {
    try {
      await this.getSystemInfo();
      return true;
    } catch {
      return false;
    }
  }

  /**
   * Wait for library scan to complete
   */
  async waitForScanCompletion(
    pollInterval: number = 5000,
    timeout: number = 3600000 // 1 hour
  ): Promise<ScanStatus> {
    const startTime = Date.now();

    while (Date.now() - startTime < timeout) {
      const status = await this.getScanStatus();
      if (!status.scanning) {
        return status;
      }
      await this.delay(pollInterval);
    }

    throw new TimeoutError(`Scan did not complete within ${timeout}ms`);
  }

  /**
   * Download multiple files with progress tracking
   */
  async *downloadMultipleSubtitles(
    requests: Array<{ path: string; language: string; providers?: string[] }>
  ): AsyncGenerator<{ index: number; result?: DownloadResult; error?: Error }, void, unknown> {
    for (let i = 0; i < requests.length; i++) {
      try {
        const result = await this.downloadSubtitles(
          requests[i].path,
          requests[i].language,
          requests[i].providers
        );
        yield { index: i, result };
      } catch (error) {
        yield { index: i, error: error as Error };
      }
    }
  }

  /**
   * Get paginated history with automatic pagination
   */
  async *getHistoryPages(
    params: HistoryQueryParams = {}
  ): AsyncGenerator<HistoryItem[], void, unknown> {
    let page = params.page || 1;
    const limit = params.limit || 20;

    while (true) {
      const response = await this.getHistory({ ...params, page, limit });
      yield response.items;

      if (!response.hasNextPage) {
        break;
      }
      page++;
    }
  }

  /**
   * Set API key for authentication
   */
  setApiKey(apiKey: string): void {
    this.axios.defaults.headers.common['X-API-Key'] = apiKey;
  }

  /**
   * Set session cookie for authentication
   */
  setSessionCookie(sessionCookie: string): void {
    this.axios.defaults.headers.common['Cookie'] = `session=${sessionCookie}`;
  }

  /**
   * Clear all authentication
   */
  clearAuth(): void {
    delete this.axios.defaults.headers.common['X-API-Key'];
    delete this.axios.defaults.headers.common['Cookie'];
  }
}