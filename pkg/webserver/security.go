// file: pkg/webserver/security.go
// version: 1.0.1
// guid: f6e0dba3-dddf-4a3b-9ee6-f3ed345ac9ab

package webserver

import "net/http"

// securityHeadersMiddleware sets common security headers on all responses.
// It helps mitigate XSS, clickjacking and other browser based attacks.
// The CSP is intentionally permissive during development and should be tightened
// before production deployment.
func securityHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Referrer-Policy", "no-referrer")

		// Temporarily allow all sources to simplify development. This
		// policy should be restricted once development is complete.
		csp := "default-src * data: blob:; " +
			"script-src * 'unsafe-inline' 'unsafe-eval' data: blob:; " +
			"style-src * 'unsafe-inline' data:; " +
			"img-src * data: blob:; " +
			"connect-src *; " +
			"font-src * data:; " +
			"object-src *; " +
			"base-uri *"

		w.Header().Set("Content-Security-Policy", csp)
		w.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
		next.ServeHTTP(w, r)
	})
}
