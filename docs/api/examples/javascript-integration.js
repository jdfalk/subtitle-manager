// file: docs/api/examples/javascript-integration.js
// version: 1.0.0
// guid: 550e8400-e29b-41d4-a716-446655440023

/**
 * Comprehensive JavaScript integration examples for Subtitle Manager API.
 *
 * This script demonstrates various integration patterns including:
 * - Basic operations with error handling
 * - File processing workflows
 * - Real-time monitoring
 * - React.js integration patterns
 * - Node.js server integration
 */

const {
  SubtitleManagerClient,
  OperationType,
  TranslationProvider,
} = require('subtitle-manager-sdk');

/**
 * Example integration class showing best practices for using the
 * Subtitle Manager SDK in production Node.js applications.
 */
class SubtitleManagerIntegration {
  constructor(baseURL = 'http://localhost:8080', apiKey = null) {
    this.client = new SubtitleManagerClient({
      baseURL,
      apiKey: apiKey || process.env.SUBTITLE_MANAGER_API_KEY,
      timeout: 60000, // 60 seconds
      maxRetries: 3,
      retryDelay: 2000,
      debug: process.env.NODE_ENV === 'development',
    });

    console.log(
      `Initialized Subtitle Manager integration with base URL: ${baseURL}`
    );
  }

  /**
   * Check if the Subtitle Manager API is healthy
   */
  async healthCheck() {
    try {
      const systemInfo = await this.client.getSystemInfo();
      console.log(
        `Health check passed - System: ${systemInfo.os} ${systemInfo.arch}`
      );
      console.log(
        `Disk usage: ${systemInfo.diskFreeFormatted}/${systemInfo.diskTotalFormatted}`
      );
      return true;
    } catch (error) {
      console.error('Health check failed:', error.message);
      return false;
    }
  }

  /**
   * Authenticate with username/password instead of API key
   */
  async authenticateSession(username, password) {
    try {
      const loginResponse = await this.client.login(username, password);
      console.log(
        `Authenticated as ${loginResponse.username} with role: ${loginResponse.role}`
      );
      return true;
    } catch (error) {
      console.error('Authentication failed:', error.message);
      return false;
    }
  }

  /**
   * Process subtitle files with conversion and optional translation
   */
  async processSubtitleFile(file, targetLanguage = null) {
    try {
      console.log(`Processing subtitle file: ${file.name || 'unknown'}`);

      // Step 1: Convert to SRT format
      const srtBlob = await this.client.convertSubtitle(file);
      console.log('Converted to SRT format');

      // Step 2: Translate if requested
      if (targetLanguage) {
        console.log(`Translating to ${targetLanguage}`);
        const translatedBlob = await this.client.translateSubtitle(
          srtBlob,
          targetLanguage,
          TranslationProvider.GOOGLE
        );
        console.log(`Translation to ${targetLanguage} completed`);
        return translatedBlob;
      }

      return srtBlob;
    } catch (error) {
      console.error(`Failed to process subtitle file:`, error.message);
      return null;
    }
  }

  /**
   * Download subtitles for multiple media files with progress tracking
   */
  async batchDownloadSubtitles(mediaFiles) {
    const results = {};
    const totalFiles = mediaFiles.length;

    console.log(`Starting batch download for ${totalFiles} files`);

    let completed = 0;
    for await (const {
      index,
      result,
      error,
    } of this.client.downloadMultipleSubtitles(mediaFiles)) {
      const mediaFile = mediaFiles[index];
      completed++;

      console.log(`Processing ${completed}/${totalFiles}: ${mediaFile.path}`);

      if (error) {
        console.error(
          `Failed to download subtitles for ${mediaFile.path}:`,
          error.message
        );
        results[mediaFile.path] = { status: 'error', error: error.message };
      } else {
        if (result.success) {
          console.log(`Downloaded subtitle: ${result.subtitle_path}`);
          results[mediaFile.path] = {
            status: 'success',
            subtitle_path: result.subtitle_path,
            provider: result.provider,
          };
        } else {
          console.warn(`No subtitles found for ${mediaFile.path}`);
          results[mediaFile.path] = { status: 'not_found' };
        }
      }
    }

    const successCount = Object.values(results).filter(
      r => r.status === 'success'
    ).length;
    console.log(
      `Batch download completed: ${successCount}/${totalFiles} successful`
    );

    return results;
  }

  /**
   * Monitor library scan with real-time progress updates
   */
  async monitorLibraryScan(path = null, progressCallback = null) {
    try {
      // Start the scan
      const scanResult = await this.client.startLibraryScan(path, true);
      console.log(`Started library scan with ID: ${scanResult.scan_id}`);

      // Monitor progress
      while (true) {
        const status = await this.client.getScanStatus();

        if (!status.scanning) {
          console.log('Library scan completed');
          if (progressCallback)
            progressCallback({ ...status, completed: true });
          return true;
        }

        console.log(`Scan progress: ${status.progressPercent}%`);
        if (status.current_path) {
          console.log(`Current path: ${status.current_path}`);
        }
        if (status.files_processed && status.files_total) {
          console.log(`Files: ${status.files_processed}/${status.files_total}`);
        }

        // Call progress callback if provided
        if (progressCallback) {
          progressCallback({ ...status, completed: false });
        }

        // Wait 5 seconds before checking again
        await new Promise(resolve => setTimeout(resolve, 5000));
      }
    } catch (error) {
      console.error('Library scan failed:', error.message);
      return false;
    }
  }

  /**
   * Analyze recent download history and provide statistics
   */
  async analyzeDownloadHistory(days = 7) {
    try {
      const endDate = new Date();
      const startDate = new Date(
        endDate.getTime() - days * 24 * 60 * 60 * 1000
      );

      // Get all history items using pagination
      const allItems = [];
      for await (const historyPage of this.client.getHistoryPages({
        start_date: startDate.toISOString(),
        end_date: endDate.toISOString(),
        limit: 100,
      })) {
        allItems.push(...historyPage);
      }

      // Analyze the data
      const analysis = {
        total_operations: allItems.length,
        by_type: {},
        by_status: {},
        by_provider: {},
        success_rate: 0,
        most_active_days: {},
        failed_operations: [],
      };

      allItems.forEach(item => {
        // Count by type
        analysis.by_type[item.type] = (analysis.by_type[item.type] || 0) + 1;

        // Count by status
        analysis.by_status[item.status] =
          (analysis.by_status[item.status] || 0) + 1;

        // Count by provider (for successful downloads)
        if (item.provider && item.isSuccess) {
          analysis.by_provider[item.provider] =
            (analysis.by_provider[item.provider] || 0) + 1;
        }

        // Track failed operations
        if (item.isFailed) {
          analysis.failed_operations.push({
            file_path: item.file_path,
            type: item.type,
            error: item.error_message,
            date: item.created_at,
          });
        }

        // Count by day
        const dayKey = item.createdAtDate.toISOString().split('T')[0];
        analysis.most_active_days[dayKey] =
          (analysis.most_active_days[dayKey] || 0) + 1;
      });

      // Calculate success rate
      const successful = analysis.by_status.success || 0;
      if (analysis.total_operations > 0) {
        analysis.success_rate = (successful / analysis.total_operations) * 100;
      }

      console.log(
        `Analyzed ${analysis.total_operations} operations from last ${days} days`
      );
      console.log(`Success rate: ${analysis.success_rate.toFixed(1)}%`);

      return analysis;
    } catch (error) {
      console.error('Failed to analyze history:', error.message);
      return {};
    }
  }

  /**
   * Extract embedded subtitles and translate to multiple languages
   */
  async extractAndTranslatePipeline(videoFile, targetLanguages) {
    const results = {};

    try {
      console.log(
        `Starting extraction and translation pipeline for: ${videoFile.name || 'video file'}`
      );

      // Step 1: Extract embedded subtitles
      const extractedSubs = await this.client.extractSubtitles(videoFile);
      console.log('Successfully extracted embedded subtitles');

      // Step 2: Translate to each target language
      for (const language of targetLanguages) {
        try {
          console.log(`Translating to ${language}`);

          const translated = await this.client.translateSubtitle(
            extractedSubs,
            language,
            TranslationProvider.GOOGLE
          );

          results[language] = translated;
          console.log(`Translation to ${language} completed`);
        } catch (error) {
          console.error(`Translation to ${language} failed:`, error.message);
          results[language] = null;
        }
      }

      return results;
    } catch (error) {
      console.error('Extraction failed:', error.message);
      return targetLanguages.reduce(
        (acc, lang) => ({ ...acc, [lang]: null }),
        {}
      );
    }
  }

  /**
   * Create a comprehensive processing report
   */
  async createProcessingReport() {
    try {
      console.log('Generating processing report');

      // Gather system information
      const systemInfo = await this.client.getSystemInfo();

      // Get recent history
      const historyAnalysis = await this.analyzeDownloadHistory(30);

      // Get current scan status
      const scanStatus = await this.client.getScanStatus();

      // Get recent logs
      const recentLogs = await this.client.getLogs({ limit: 50 });
      const errorLogs = recentLogs.filter(log => log.isError);

      // Create report
      const report = {
        generated: new Date().toISOString(),
        system: {
          os: `${systemInfo.os} ${systemInfo.arch}`,
          go_version: systemInfo.go_version,
          version: systemInfo.version || 'Unknown',
          uptime: systemInfo.uptime || 'Unknown',
          memory_usage: systemInfo.memory_usage || 'Unknown',
          disk_usage: `${systemInfo.diskFreeFormatted}/${systemInfo.diskTotalFormatted}`,
        },
        library: {
          currently_scanning: scanStatus.scanning,
          scan_progress: `${scanStatus.progressPercent}%`,
          files_processed: scanStatus.files_processed || 'N/A',
          files_total: scanStatus.files_total || 'N/A',
        },
        activity: {
          total_operations: historyAnalysis.total_operations || 0,
          success_rate: `${(historyAnalysis.success_rate || 0).toFixed(1)}%`,
          by_type: historyAnalysis.by_type || {},
          by_status: historyAnalysis.by_status || {},
          by_provider: historyAnalysis.by_provider || {},
        },
        recent_errors: errorLogs.slice(0, 10).map(log => ({
          timestamp: log.timestampDate.toISOString(),
          component: log.component,
          message: log.message,
        })),
      };

      console.log('Processing report generated:', report);
      return report;
    } catch (error) {
      console.error('Failed to generate report:', error.message);
      return null;
    }
  }
}

/**
 * React.js integration example
 */
class ReactSubtitleManager {
  constructor() {
    this.client = new SubtitleManagerClient({
      baseURL: process.env.REACT_APP_API_URL || 'http://localhost:8080',
      apiKey: process.env.REACT_APP_API_KEY,
      debug: process.env.NODE_ENV === 'development',
    });
  }

  // React hook for file upload and conversion
  useSubtitleConverter() {
    const [converting, setConverting] = useState(false);
    const [progress, setProgress] = useState(0);
    const [error, setError] = useState(null);

    const convertFile = async (file, targetLanguage = null) => {
      setConverting(true);
      setError(null);
      setProgress(0);

      try {
        // Step 1: Convert to SRT (50% progress)
        const srtBlob = await this.client.convertSubtitle(file);
        setProgress(50);

        let finalBlob = srtBlob;

        // Step 2: Translate if needed (remaining 50%)
        if (targetLanguage) {
          finalBlob = await this.client.translateSubtitle(
            srtBlob,
            targetLanguage,
            TranslationProvider.GOOGLE
          );
        }
        setProgress(100);

        // Download the result
        const url = URL.createObjectURL(finalBlob);
        const a = document.createElement('a');
        a.href = url;
        a.download = `converted${targetLanguage ? `.${targetLanguage}` : ''}.srt`;
        a.click();
        URL.revokeObjectURL(url);

        return true;
      } catch (err) {
        setError(err.message);
        return false;
      } finally {
        setConverting(false);
        setProgress(0);
      }
    };

    return { converting, progress, error, convertFile };
  }

  // React hook for library monitoring
  useLibraryMonitor() {
    const [scanStatus, setScanStatus] = useState(null);
    const [isLoading, setIsLoading] = useState(false);

    const startScan = async (path = null) => {
      setIsLoading(true);
      try {
        await this.client.startLibraryScan(path, true);

        // Start polling for status
        const interval = setInterval(async () => {
          const status = await this.client.getScanStatus();
          setScanStatus(status);

          if (!status.scanning) {
            clearInterval(interval);
            setIsLoading(false);
          }
        }, 5000);

        return interval;
      } catch (error) {
        setIsLoading(false);
        throw error;
      }
    };

    return { scanStatus, isLoading, startScan };
  }
}

/**
 * Express.js server integration example
 */
class ExpressSubtitleAPI {
  constructor(app) {
    this.app = app;
    this.client = new SubtitleManagerClient({
      baseURL: process.env.SUBTITLE_MANAGER_URL || 'http://localhost:8080',
      apiKey: process.env.SUBTITLE_MANAGER_API_KEY,
    });

    this.setupRoutes();
  }

  setupRoutes() {
    // Health check endpoint
    this.app.get('/api/subtitle-manager/health', async (req, res) => {
      try {
        const healthy = await this.client.healthCheck();
        res.json({ healthy, timestamp: new Date().toISOString() });
      } catch (error) {
        res.status(500).json({ error: error.message });
      }
    });

    // Convert subtitle endpoint
    this.app.post('/api/subtitle-manager/convert', async (req, res) => {
      try {
        if (!req.files || !req.files.subtitle) {
          return res.status(400).json({ error: 'No subtitle file provided' });
        }

        const file = req.files.subtitle;
        const targetLanguage = req.body.language;

        // Convert the file
        const blob = new Blob([file.data], { type: file.mimetype });
        let result = await this.client.convertSubtitle(blob, file.name);

        // Translate if requested
        if (targetLanguage) {
          result = await this.client.translateSubtitle(
            result,
            targetLanguage,
            TranslationProvider.GOOGLE
          );
        }

        // Return the converted file
        const buffer = await result.arrayBuffer();
        res.setHeader('Content-Type', 'application/x-subrip');
        res.setHeader(
          'Content-Disposition',
          `attachment; filename="converted${targetLanguage ? `.${targetLanguage}` : ''}.srt"`
        );
        res.send(Buffer.from(buffer));
      } catch (error) {
        res.status(500).json({ error: error.message });
      }
    });

    // Download subtitles endpoint
    this.app.post('/api/subtitle-manager/download', async (req, res) => {
      try {
        const { path, language, providers } = req.body;

        if (!path || !language) {
          return res
            .status(400)
            .json({ error: 'Path and language are required' });
        }

        const result = await this.client.downloadSubtitles(
          path,
          language,
          providers
        );
        res.json(result);
      } catch (error) {
        res.status(500).json({ error: error.message });
      }
    });

    // Get history endpoint with pagination
    this.app.get('/api/subtitle-manager/history', async (req, res) => {
      try {
        const { page = 1, limit = 20, type, start_date, end_date } = req.query;

        const history = await this.client.getHistory({
          page: parseInt(page),
          limit: parseInt(limit),
          type: type || undefined,
          start_date: start_date || undefined,
          end_date: end_date || undefined,
        });

        res.json(history);
      } catch (error) {
        res.status(500).json({ error: error.message });
      }
    });

    // WebSocket endpoint for real-time updates
    this.app.ws('/api/subtitle-manager/events', (ws, req) => {
      let scanInterval = null;

      ws.on('message', async message => {
        try {
          const data = JSON.parse(message);

          if (data.action === 'start_scan') {
            // Start library scan and send updates
            await this.client.startLibraryScan(data.path, true);

            scanInterval = setInterval(async () => {
              try {
                const status = await this.client.getScanStatus();
                ws.send(JSON.stringify({ type: 'scan_status', data: status }));

                if (!status.scanning) {
                  clearInterval(scanInterval);
                  scanInterval = null;
                }
              } catch (error) {
                ws.send(
                  JSON.stringify({
                    type: 'error',
                    data: { message: error.message },
                  })
                );
              }
            }, 2000);
          }
        } catch (error) {
          ws.send(
            JSON.stringify({ type: 'error', data: { message: error.message } })
          );
        }
      });

      ws.on('close', () => {
        if (scanInterval) {
          clearInterval(scanInterval);
        }
      });
    });
  }
}

/**
 * Main function demonstrating various integration scenarios
 */
async function main() {
  // Initialize integration
  const integration = new SubtitleManagerIntegration();

  // Example 1: Health check
  if (!(await integration.healthCheck())) {
    console.error('Subtitle Manager is not healthy, exiting');
    return;
  }

  // Example 2: Batch download subtitles
  const mediaFiles = [
    { path: '/movies/example1.mkv', language: 'en' },
    {
      path: '/movies/example2.mkv',
      language: 'en',
      providers: ['opensubtitles'],
    },
    { path: '/tv/series/s01e01.mkv', language: 'es' },
  ];

  const downloadResults = await integration.batchDownloadSubtitles(mediaFiles);
  console.log('Download results:', downloadResults);

  // Example 3: Monitor library scan with progress callback
  await integration.monitorLibraryScan('/movies/new', progress => {
    if (progress.completed) {
      console.log('Scan completed successfully!');
    } else {
      console.log(`Progress update: ${progress.progressPercent}%`);
    }
  });

  // Example 4: Analyze history
  const analysis = await integration.analyzeDownloadHistory(7);
  console.log('History analysis:', analysis);

  // Example 5: Generate report
  const report = await integration.createProcessingReport();
  console.log('Processing report:', report);
}

// Export for use in other modules
module.exports = {
  SubtitleManagerIntegration,
  ReactSubtitleManager,
  ExpressSubtitleAPI,
};

// Run main function if this file is executed directly
if (require.main === module) {
  main().catch(console.error);
}
