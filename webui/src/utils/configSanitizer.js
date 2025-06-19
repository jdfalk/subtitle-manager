// file: webui/src/utils/configSanitizer.js

import { useState } from 'react';

/**
 * Utility functions for sanitizing configuration data to hide sensitive information
 */

/**
 * Sanitizes configuration object by masking sensitive values
 * @param {Object} obj - Configuration object to sanitize
 * @param {boolean} showSensitive - Whether to show sensitive values or mask them
 * @returns {Object} Sanitized configuration object
 */
export const sanitizeConfig = (obj, showSensitive = false) => {
  const sensitive = ['password', 'apikey', 'api_key', 'token', 'secret', 'key'];

  const sanitizeValue = (key, val) => {
    if (showSensitive) return val;
    const lower = key.toLowerCase();
    if (sensitive.some(s => lower.includes(s))) {
      if (typeof val === 'string') {
        if (!val) return '';
        const last = val.slice(-4);
        return `****${last}`;
      }
      if (val === null || typeof val === 'undefined') return '';
      return '****';
    }
    return val;
  };

  const walk = input => {
    if (typeof input !== 'object' || input === null) return input;
    if (Array.isArray(input)) return input.map(walk);
    const result = {};
    for (const [k, v] of Object.entries(input)) {
      if (typeof v === 'object' && v !== null) {
        if (Array.isArray(v)) {
          // Check if the key itself is sensitive, if so sanitize the array values
          const lower = k.toLowerCase();
          if (sensitive.some(s => lower.includes(s))) {
            result[k] = v.map(item => sanitizeValue(k, item));
          } else {
            result[k] = v.map(walk);
          }
        } else {
          result[k] = walk(v);
        }
      } else {
        result[k] = sanitizeValue(k, v);
      }
    }
    return result;
  };

  return walk(obj);
};

/**
 * Custom hook for managing sensitive data display state
 * @returns {Object} Hook state and methods
 */
export const useSensitiveDataToggle = (initialState = false) => {
  const [showSensitive, setShowSensitive] = useState(initialState);

  const toggleSensitive = () => setShowSensitive(prev => !prev);

  return {
    showSensitive,
    setShowSensitive,
    toggleSensitive,
  };
};
