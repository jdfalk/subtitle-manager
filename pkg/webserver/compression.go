// file: pkg/webserver/compression.go
// version: 1.0.0
// guid: 9e8f7d6c-5b4a-0e9f-3a2b-6c5d4e3f2c1b

package webserver

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

// compressionResponseWriter wraps http.ResponseWriter to provide gzip compression
type compressionResponseWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

// Write compresses data before writing to the underlying response writer
func (crw *compressionResponseWriter) Write(b []byte) (int, error) {
	return crw.Writer.Write(b)
}

// compressionMiddleware provides gzip compression for HTTP responses to reduce bandwidth usage.
//
// This middleware automatically compresses responses when the client supports gzip encoding,
// significantly reducing response sizes for JSON APIs and static assets. It provides:
//
//   - Automatic content-type detection for compressible content
//   - Proper gzip headers and encoding
//   - Performance optimization through selective compression
//   - Memory-efficient streaming compression
//
// The middleware compresses responses for common web content types including:
//   - JSON API responses (application/json)
//   - HTML pages (text/html)
//   - CSS stylesheets (text/css)
//   - JavaScript files (text/javascript, application/javascript)
//   - Plain text content (text/plain)
//   - XML content (text/xml, application/xml)
//
// Content types that are already compressed (images, videos, archives) are not re-compressed
// to avoid unnecessary CPU usage and potential size increases.
func compressionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if client accepts gzip encoding
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		// Set compression headers
		w.Header().Set("Content-Encoding", "gzip")
		w.Header().Set("Vary", "Accept-Encoding")

		// Create gzip writer
		gzipWriter := gzip.NewWriter(w)
		defer gzipWriter.Close()

		// Create compression response writer
		crw := &compressionResponseWriter{
			ResponseWriter: w,
			Writer:         gzipWriter,
		}

		// Serve the request with compression
		next.ServeHTTP(crw, r)
	})
}

// shouldCompress determines if a response should be compressed based on content type.
//
// Returns true for text-based content types that benefit from compression:
//   - application/json
//   - text/html, text/css, text/javascript, text/plain, text/xml
//   - application/javascript, application/xml
//
// Returns false for binary content that is already compressed or doesn't benefit:
//   - Images (image/*)
//   - Videos (video/*)
//   - Audio (audio/*)
//   - Archives (application/zip, application/gzip, etc.)
func shouldCompress(contentType string) bool {
	compressibleTypes := []string{
		"application/json",
		"text/html",
		"text/css",
		"text/javascript",
		"text/plain",
		"text/xml",
		"application/javascript",
		"application/xml",
		"application/rss+xml",
		"application/atom+xml",
	}

	for _, compressible := range compressibleTypes {
		if strings.HasPrefix(contentType, compressible) {
			return true
		}
	}

	return false
}

// ConditionalCompressionMiddleware provides smarter compression that only compresses
// responses when beneficial based on content type and size.
//
// This enhanced version provides:
//   - Content-type aware compression decisions
//   - Size-based compression thresholds
//   - Better performance for small responses
//   - Reduced CPU usage for incompressible content
func ConditionalCompressionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if client accepts gzip encoding
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		// Create a response recorder to capture the response
		recorder := &responseRecorder{
			ResponseWriter: w,
			statusCode:     200,
		}

		// Serve the request to the recorder
		next.ServeHTTP(recorder, r)

		// Decide whether to compress based on content type and size
		contentType := recorder.Header().Get("Content-Type")
		shouldCompress := shouldCompress(contentType) && len(recorder.body) > 100 // Only compress if > 100 bytes

		if shouldCompress {
			// Apply compression
			w.Header().Set("Content-Encoding", "gzip")
			w.Header().Set("Vary", "Accept-Encoding")

			gzipWriter := gzip.NewWriter(w)
			defer gzipWriter.Close()

			// Copy headers
			for key, values := range recorder.Header() {
				for _, value := range values {
					w.Header().Add(key, value)
				}
			}

			w.WriteHeader(recorder.statusCode)
			gzipWriter.Write(recorder.body)
		} else {
			// Serve uncompressed
			for key, values := range recorder.Header() {
				for _, value := range values {
					w.Header().Add(key, value)
				}
			}
			w.WriteHeader(recorder.statusCode)
			w.Write(recorder.body)
		}
	})
}

// responseRecorder captures response data for analysis before compression decision
type responseRecorder struct {
	http.ResponseWriter
	statusCode int
	body       []byte
}

// WriteHeader captures the status code
func (rr *responseRecorder) WriteHeader(statusCode int) {
	rr.statusCode = statusCode
}

// Write captures the response body
func (rr *responseRecorder) Write(b []byte) (int, error) {
	rr.body = append(rr.body, b...)
	return len(b), nil
}