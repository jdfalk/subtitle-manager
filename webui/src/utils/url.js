// file: webui/src/utils/url.js
// version: 1.0.0
// guid: 148d07ae-f97c-4342-bb06-dfb662b10111

export function normalizeUrl(rawUrl) {
  const url = rawUrl.trim();
  if (!url) {
    return '';
  }
  if (!/^https?:\/\//i.test(url)) {
    return `http://${url}`;
  }
  return url;
}
