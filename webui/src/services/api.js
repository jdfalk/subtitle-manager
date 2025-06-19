/**
 * API service configuration and utilities
 * Handles communication with the Subtitle Manager backend
 */

/**
 * API service configuration
 */
/**
 * Determine the base path for API requests by inspecting the current URL.
 * If the first path segment isn't one of the known application routes, it is
 * treated as an installation prefix specified via `base_url` on the server.
 *
 * @returns {string} Base path beginning with a slash or an empty string.
 */
export function getBasePath() {
  const known = new Set([
    '',
    'dashboard',
    'library',
    'wanted',
    'history',
    'settings',
    'system',
    'tools',
    'offline-info',
    'setup',
  ]);
  const parts = window.location.pathname.split('/').filter(Boolean);
  if (parts.length > 0 && !known.has(parts[0])) {
    return `/${parts[0]}`;
  }
  return '';
}

const API_BASE_URL =
  import.meta?.env?.VITE_API_URL || `${window.location.origin}${getBasePath()}`;

/**
 * Enhanced fetch wrapper with error handling and logging
 * @param {string} url - The endpoint URL
 * @param {RequestInit} options - Fetch options
 * @returns {Promise<Response>} - Fetch response
 */
async function apiClient(url, options = {}) {
  const fullUrl = url.startsWith('http') ? url : `${API_BASE_URL}${url}`;

  // Only log in development mode
  if (import.meta.env.DEV) {
    console.log(`API Request: ${options.method || 'GET'} ${fullUrl}`);
  }

  const config = {
    headers: {
      'Content-Type': 'application/json',
      ...options.headers,
    },
    ...options,
  };

  try {
    const response = await fetch(fullUrl, config);
    if (import.meta.env.DEV) {
      console.log(`API Response: ${response.status} ${fullUrl}`);
    }
    return response;
  } catch (error) {
    // Only log errors in development mode to avoid console errors in production/testing
    if (import.meta.env.DEV) {
      console.error('API Request Error:', error);

      if (error.name === 'TypeError' && error.message.includes('fetch')) {
        console.warn('Backend service appears to be offline');
      }
    }

    throw error;
  }
}

/**
 * API service object with common HTTP methods
 */
export const apiService = {
  /**
   * Perform GET request
   * @param {string} url - The endpoint URL
   * @param {RequestInit} config - Optional fetch config
   * @returns {Promise<Response>} - Promise with response
   */
  async get(url, config = {}) {
    return apiClient(url, { method: 'GET', ...config });
  },

  /**
   * Perform POST request
   * @param {string} url - The endpoint URL
   * @param {any} data - Request payload
   * @param {RequestInit} config - Optional fetch config
   * @returns {Promise<Response>} - Promise with response
   */
  async post(url, data = {}, config = {}) {
    return apiClient(url, {
      method: 'POST',
      body: JSON.stringify(data),
      ...config,
    });
  },

  /**
   * Perform PUT request
   * @param {string} url - The endpoint URL
   * @param {any} data - Request payload
   * @param {RequestInit} config - Optional fetch config
   * @returns {Promise<Response>} - Promise with response
   */
  async put(url, data = {}, config = {}) {
    return apiClient(url, {
      method: 'PUT',
      body: JSON.stringify(data),
      ...config,
    });
  },

  /**
   * Perform DELETE request
   * @param {string} url - The endpoint URL
   * @param {RequestInit} config - Optional fetch config
   * @returns {Promise<Response>} - Promise with response
   */
  async delete(url, config = {}) {
    return apiClient(url, { method: 'DELETE', ...config });
  },

  /**
   * Check if the backend is available
   * @returns {Promise<boolean>} - Whether backend is available
   */
  async checkBackendHealth() {
    try {
      // Use setup/status endpoint for health check as it's unauthenticated
      const response = await this.get('/api/setup/status');
      return response.ok;
    } catch (error) {
      if (import.meta.env.DEV) {
        console.warn('Backend health check failed:', error);
      }
      return false;
    }
  },
};

export default apiService;
