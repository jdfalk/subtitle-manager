// file: sdks/javascript/tests/client.test.ts
// version: 1.0.0
// guid: 550e8400-e29b-41d4-a716-446655440019

/**
 * Tests for the Subtitle Manager JavaScript SDK
 */

import nock from 'nock';
import { SubtitleManagerClient } from '../src/client';
import {
  AuthenticationError,
  AuthorizationError,
  NotFoundError,
  RateLimitError,
  ValidationError,
} from '../src/errors';
import {
  SystemInfo,
  HistoryItem,
  ScanStatus,
  LoginResponse,
  DownloadResult,
} from '../src/models';
import { OperationType, OperationStatus, UserRole } from '../src/types';

describe('SubtitleManagerClient', () => {
  let client: SubtitleManagerClient;
  const baseURL = 'http://test.example.com';

  beforeEach(() => {
    client = new SubtitleManagerClient({
      baseURL,
      apiKey: 'test-api-key',
    });
  });

  afterEach(() => {
    nock.cleanAll();
  });

  describe('Authentication', () => {
    test('should authenticate with API key', async () => {
      const systemInfo = {
        go_version: 'go1.24.0',
        os: 'linux',
        arch: 'amd64',
        goroutines: 10,
        disk_free: 1000000,
        disk_total: 2000000,
      };

      nock(baseURL)
        .get('/api/system')
        .matchHeader('X-API-Key', 'test-api-key')
        .reply(200, systemInfo);

      const result = await client.getSystemInfo();
      expect(result).toBeInstanceOf(SystemInfo);
      expect(result.go_version).toBe('go1.24.0');
    });

    test('should login with username and password', async () => {
      const loginResponse = {
        user_id: 1,
        username: 'testuser',
        role: 'admin',
      };

      nock(baseURL)
        .post('/api/login', { username: 'testuser', password: 'password' })
        .reply(200, loginResponse);

      const result = await client.login('testuser', 'password');
      expect(result).toBeInstanceOf(LoginResponse);
      expect(result.user_id).toBe(1);
      expect(result.username).toBe('testuser');
      expect(result.role).toBe(UserRole.ADMIN);
      expect(result.isAdmin).toBe(true);
    });

    test('should handle authentication error', async () => {
      nock(baseURL).get('/api/system').reply(401, {
        error: 'unauthorized',
        message: 'Authentication required',
      });

      await expect(client.getSystemInfo()).rejects.toThrow(AuthenticationError);
    });

    test('should handle authorization error', async () => {
      nock(baseURL)
        .post('/api/oauth/github/generate')
        .reply(403, { error: 'forbidden', message: 'Admin access required' });

      await expect(client.generateGitHubOAuth()).rejects.toThrow(
        AuthorizationError
      );
    });

    test('should handle rate limit error', async () => {
      nock(baseURL).get('/api/system').reply(
        429,
        { error: 'rate_limited', message: 'Rate limit exceeded' },
        {
          'Retry-After': '60',
        }
      );

      const error = await client.getSystemInfo().catch(e => e);
      expect(error).toBeInstanceOf(RateLimitError);
      expect(error.retryAfter).toBe(60);
    });
  });

  describe('System Information', () => {
    test('should get system information', async () => {
      const systemData = {
        go_version: 'go1.24.0',
        os: 'linux',
        arch: 'amd64',
        goroutines: 15,
        disk_free: 5000000000,
        disk_total: 10000000000,
        memory_usage: 512000000,
        uptime: '2 days',
        version: '1.0.0',
      };

      nock(baseURL).get('/api/system').reply(200, systemData);

      const result = await client.getSystemInfo();
      expect(result).toBeInstanceOf(SystemInfo);
      expect(result.go_version).toBe('go1.24.0');
      expect(result.goroutines).toBe(15);
      expect(result.disk_free).toBe(5000000000);
      expect(result.version).toBe('1.0.0');
      expect(result.diskUsagePercent).toBe(50); // 50% usage
    });

    test('should get application logs', async () => {
      const logData = [
        {
          timestamp: '2024-01-01T12:00:00Z',
          level: 'info',
          component: 'webserver',
          message: 'Server started',
          fields: { port: 8080 },
        },
        {
          timestamp: '2024-01-01T12:01:00Z',
          level: 'error',
          component: 'auth',
          message: 'Login failed',
          fields: { username: 'baduser' },
        },
      ];

      nock(baseURL)
        .get('/api/logs')
        .query({ level: 'info', limit: 50 })
        .reply(200, logData);

      const result = await client.getLogs({ level: 'info', limit: 50 });
      expect(result).toHaveLength(2);
      expect(result[0].level).toBe('info');
      expect(result[0].component).toBe('webserver');
      expect(result[1].level).toBe('error');
      expect(result[1].isError).toBe(true);
    });
  });

  describe('File Operations', () => {
    test('should convert subtitle file', async () => {
      const srtContent = '1\n00:00:01,000 --> 00:00:03,000\nHello World\n\n';
      const blob = new Blob([srtContent], { type: 'application/x-subrip' });

      nock(baseURL)
        .post('/api/convert')
        .reply(200, srtContent, { 'Content-Type': 'application/x-subrip' });

      const inputFile = new File(
        ['WEBVTT\n\n00:01.000 --> 00:03.000\nHello World'],
        'test.vtt'
      );
      const result = await client.convertSubtitle(inputFile);
      expect(result).toBeInstanceOf(Blob);
    });

    test('should download subtitles', async () => {
      const downloadData = {
        success: true,
        subtitle_path: '/path/to/movie.en.srt',
        provider: 'opensubtitles',
      };

      nock(baseURL)
        .post('/api/download', {
          path: '/movies/example.mkv',
          language: 'en',
          providers: ['opensubtitles', 'subscene'],
        })
        .reply(200, downloadData);

      const result = await client.downloadSubtitles(
        '/movies/example.mkv',
        'en',
        ['opensubtitles', 'subscene']
      );

      expect(result).toBeInstanceOf(DownloadResult);
      expect(result.success).toBe(true);
      expect(result.subtitle_path).toBe('/path/to/movie.en.srt');
      expect(result.provider).toBe('opensubtitles');
    });
  });

  describe('Library Management', () => {
    test('should get scan status', async () => {
      const scanData = {
        scanning: true,
        progress: 0.75,
        current_path: '/movies/subfolder',
        files_processed: 150,
        files_total: 200,
        start_time: '2024-01-01T12:00:00Z',
        estimated_completion: '2024-01-01T12:30:00Z',
      };

      nock(baseURL).get('/api/scan/status').reply(200, scanData);

      const result = await client.getScanStatus();
      expect(result).toBeInstanceOf(ScanStatus);
      expect(result.scanning).toBe(true);
      expect(result.progress).toBe(0.75);
      expect(result.progressPercent).toBe(75);
      expect(result.files_processed).toBe(150);
      expect(result.remainingFiles).toBe(50);
    });

    test('should start library scan', async () => {
      const scanResult = { scan_id: 'scan-123' };

      nock(baseURL)
        .post('/api/scan', { path: '/movies/new', force: true })
        .reply(200, scanResult);

      const result = await client.startLibraryScan('/movies/new', true);
      expect(result.scan_id).toBe('scan-123');
    });
  });

  describe('History', () => {
    test('should get operation history', async () => {
      const historyData = {
        items: [
          {
            id: 1,
            type: 'download',
            file_path: '/movies/example.mkv',
            subtitle_path: '/movies/example.en.srt',
            language: 'en',
            provider: 'opensubtitles',
            status: 'success',
            created_at: '2024-01-01T12:00:00Z',
            user_id: 1,
          },
          {
            id: 2,
            type: 'convert',
            file_path: '/subtitles/test.vtt',
            subtitle_path: '/subtitles/test.srt',
            language: null,
            provider: null,
            status: 'success',
            created_at: '2024-01-01T12:05:00Z',
            user_id: 1,
          },
        ],
        total: 2,
        page: 1,
        limit: 20,
      };

      nock(baseURL)
        .get('/api/history')
        .query({ page: 1, limit: 20, type: 'download' })
        .reply(200, historyData);

      const result = await client.getHistory({
        page: 1,
        limit: 20,
        type: OperationType.DOWNLOAD,
      });

      expect(result.total).toBe(2);
      expect(result.page).toBe(1);
      expect(result.items).toHaveLength(2);
      expect(result.hasNextPage).toBe(false);
      expect(result.totalPages).toBe(1);

      const firstItem = result.items[0];
      expect(firstItem).toBeInstanceOf(HistoryItem);
      expect(firstItem.type).toBe(OperationType.DOWNLOAD);
      expect(firstItem.status).toBe(OperationStatus.SUCCESS);
      expect(firstItem.isSuccess).toBe(true);
      expect(firstItem.provider).toBe('opensubtitles');
    });
  });

  describe('Utility Methods', () => {
    test('should perform health check - success', async () => {
      const systemInfo = {
        go_version: 'go1.24.0',
        os: 'linux',
        arch: 'amd64',
        goroutines: 10,
        disk_free: 1000000,
        disk_total: 2000000,
      };

      nock(baseURL).get('/api/system').reply(200, systemInfo);

      const result = await client.healthCheck();
      expect(result).toBe(true);
    });

    test('should perform health check - failure', async () => {
      nock(baseURL).get('/api/system').reply(500, { error: 'internal_error' });

      const result = await client.healthCheck();
      expect(result).toBe(false);
    });

    test('should set API key', () => {
      client.setApiKey('new-api-key');
      // API key should be set in axios defaults
    });

    test('should clear authentication', () => {
      client.clearAuth();
      // Authentication headers should be cleared
    });
  });

  describe('Error Handling', () => {
    test('should handle not found error', async () => {
      nock(baseURL)
        .get('/api/nonexistent')
        .reply(404, { error: 'not_found', message: 'Resource not found' });

      // Create a direct request to test error handling
      await expect(
        client['request']({ method: 'GET', url: '/api/nonexistent' })
      ).rejects.toThrow(NotFoundError);
    });

    test('should handle validation error', async () => {
      nock(baseURL).post('/api/download').reply(400, {
        error: 'validation_error',
        message: 'Invalid request data',
      });

      await expect(client.downloadSubtitles('', '')).rejects.toThrow(
        ValidationError
      );
    });
  });

  describe('Pagination Generator', () => {
    test('should iterate through history pages', async () => {
      // Mock first page
      nock(baseURL)
        .get('/api/history')
        .query({ page: 1, limit: 2 })
        .reply(200, {
          items: [
            {
              id: 1,
              type: 'download',
              file_path: '/movies/1.mkv',
              status: 'success',
              created_at: '2024-01-01T12:00:00Z',
              user_id: 1,
            },
            {
              id: 2,
              type: 'download',
              file_path: '/movies/2.mkv',
              status: 'success',
              created_at: '2024-01-01T12:01:00Z',
              user_id: 1,
            },
          ],
          total: 3,
          page: 1,
          limit: 2,
        });

      // Mock second page
      nock(baseURL)
        .get('/api/history')
        .query({ page: 2, limit: 2 })
        .reply(200, {
          items: [
            {
              id: 3,
              type: 'download',
              file_path: '/movies/3.mkv',
              status: 'success',
              created_at: '2024-01-01T12:02:00Z',
              user_id: 1,
            },
          ],
          total: 3,
          page: 2,
          limit: 2,
        });

      const pages: HistoryItem[][] = [];
      for await (const page of client.getHistoryPages({ limit: 2 })) {
        pages.push(page);
      }

      expect(pages).toHaveLength(2);
      expect(pages[0]).toHaveLength(2);
      expect(pages[1]).toHaveLength(1);
      expect(pages[0][0].id).toBe(1);
      expect(pages[1][0].id).toBe(3);
    });
  });
});
