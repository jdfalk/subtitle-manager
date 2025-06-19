// file: pkg/webserver/security.go

package webserver

import "net/http"

// securityHeadersMiddleware sets common security headers on all responses.
// It helps mitigate XSS, clickjacking and other browser based attacks.
// The CSP allows necessary resources for the React/Material-UI frontend while maintaining security.
func securityHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-Content-Type-Options", "nosniff")

		// Content Security Policy that allows:
		// - Self-hosted content (scripts, styles, images, etc.)
		// - Google Fonts (Material-UI requirement)
		// - Inline styles (Material-UI dynamic styles)
		// - Data URLs for images (common in modern web apps)
		// - Blob URLs (file operations and dynamic content)
		// - HTTPS images (for external content like movie posters)
		csp := "default-src 'self'; " +
			"script-src 'self' 'unsafe-inline' blob:; " +
			"style-src 'self' 'unsafe-inline' https://fonts.googleapis.com; " +
			"font-src 'self' https://fonts.gstatic.com; " +
			"img-src 'self' data: https: blob:; " +
			"connect-src 'self'; " +
			"object-src 'none'; " +
			"base-uri 'self'"

		w.Header().Set("Content-Security-Policy", csp)
		w.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
		next.ServeHTTP(w, r)
	})
}
