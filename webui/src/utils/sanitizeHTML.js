// file: webui/src/utils/sanitizeHTML.js

import DOMPurify from 'dompurify';

/**
 * Sanitize HTML to prevent XSS.
 * @param {string} input - Untrusted HTML string.
 * @returns {string} Sanitized HTML safe for rendering.
 */
export const sanitizeHTML = input => DOMPurify.sanitize(input);
