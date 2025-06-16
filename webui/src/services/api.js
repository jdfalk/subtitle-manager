/**
 * API service configuration and utilities
 * Handles communication with the Subtitle Manager backend
 */

/**
 * API service configuration
 */
const API_BASE_URL = import.meta?.env?.VITE_API_URL || window.location.origin;

/**
 * Enhanced fetch wrapper with error handling and logging
 * @param {string} url - The endpoint URL
 * @param {RequestInit} options - Fetch options
 * @returns {Promise<Response>} - Fetch response
 */
async function apiClient(url, options = {}) {
  const fullUrl = url.startsWith('http') ? url : `${API_BASE_URL}${url}`;

  console.log(`API Request: ${options.method || 'GET'} ${fullUrl}`);

  const config = {
    headers: {
      'Content-Type': 'application/json',
      ...options.headers,
    },
    ...options,
  };

  try {
    const response = await fetch(fullUrl, config);
    console.log(`API Response: ${response.status} ${fullUrl}`);
    return response;
  } catch (error) {
    console.error('API Request Error:', error);

    if (error.name === 'TypeError' && error.message.includes('fetch')) {
      console.warn('Backend service appears to be offline');
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
      console.warn('Backend health check failed:', error);
      return false;
    }
  },
};

export default apiService;
