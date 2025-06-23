// file: webui/src/utils/security.js
/**
 * Validate that a URL uses HTTP(S) and matches allowed hostnames.
 * Mirrors the logic from the Go security package.
 * @param {string} urlStr - URL to validate
 * @param {string[]} allowedHosts - list of allowed hostnames
 * @returns {boolean} true if valid and host is allowed
 */
export function validateAllowedHostname(urlStr, allowedHosts) {
  try {
    const u = new URL(urlStr, 'http://example.com');
    if (u.protocol !== 'http:' && u.protocol !== 'https:') {
      return false;
    }
    const host = u.hostname.toLowerCase();
    return allowedHosts.some(h => host === h.toLowerCase());
  } catch {
    return false;
  }
}
