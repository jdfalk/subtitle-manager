// file: pkg/security/security.go
package security

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/viper"
)

// ValidateURL checks if raw is a safe HTTP or HTTPS URL.
// It blocks known metadata services and dangerous ports.
// The sanitized URL string is returned when valid.
func ValidateURL(raw string) (string, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return "", fmt.Errorf("invalid URL format: %w", err)
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return "", fmt.Errorf("only HTTP and HTTPS schemes are allowed, got: %s", u.Scheme)
	}
	if u.Hostname() == "" {
		return "", fmt.Errorf("hostname cannot be empty")
	}

	hostname := strings.ToLower(u.Hostname())
	blockedHosts := []string{
		"169.254.169.254",
		"metadata.google.internal",
		"metadata",
	}
	for _, blocked := range blockedHosts {
		if hostname == blocked {
			return "", fmt.Errorf("hostname %s is not allowed", hostname)
		}
	}

	if u.Port() != "" {
		blockedPorts := []string{"22", "23", "3389", "5900", "5901"}
		port := u.Port()
		for _, blocked := range blockedPorts {
			if port == blocked {
				return "", fmt.Errorf("port %s is not allowed", port)
			}
		}
	}

	if !strings.HasPrefix(u.Path, "/") {
		u.Path = "/" + u.Path
	}

	return u.String(), nil
}

// GetAllowedBaseDirs returns directories considered safe for file browsing.
func GetAllowedBaseDirs() []string {
	var dirs []string
	if mediaDir := viper.GetString("media_directory"); mediaDir != "" {
		dirs = append(dirs, mediaDir)
	}
	if subDir := viper.GetString("subtitle_directory"); subDir != "" {
		dirs = append(dirs, subDir)
	}
	commonDirs := []string{
		"/media", "/mnt/media", "/home/media", "/var/media",
		"/Movies", "/TV", "/Videos",
	}
	if runtime.GOOS == "windows" {
		commonDirs = []string{
			"C:\\Media", "C:\\Movies", "C:\\TV", "D:\\Media", "D:\\Movies", "D:\\TV",
		}
	}
	for _, dir := range commonDirs {
		if info, err := os.Stat(dir); err == nil && info.IsDir() {
			dirs = append(dirs, dir)
		}
	}
	if home, err := os.UserHomeDir(); err == nil {
		dirs = append(dirs, home)
	}
	return dirs
}

// ValidateAndSanitizePath cleans userPath and ensures it resides in an allowed directory.
// An absolute, sanitized path is returned on success.
func ValidateAndSanitizePath(userPath string) (string, error) {
	cleanPath := filepath.Clean(userPath)
	if !filepath.IsAbs(cleanPath) {
		return "", fmt.Errorf("path must be absolute")
	}
	absPath := filepath.Clean(cleanPath)
	allowedBaseDirs := GetAllowedBaseDirs()

	if cleanPath == "/" || cleanPath == "." ||
		(runtime.GOOS == "windows" && len(filepath.VolumeName(cleanPath)) == 2 &&
			strings.Trim(filepath.Clean(cleanPath), "\\") == filepath.VolumeName(cleanPath)) {
		for _, baseDir := range allowedBaseDirs {
			if absPath == baseDir || strings.HasPrefix(baseDir, absPath) {
				return absPath, nil
			}
		}
	}

	for _, baseDir := range allowedBaseDirs {
		relPath, err := filepath.Rel(baseDir, absPath)
		if err == nil && !strings.HasPrefix(relPath, "..") {
			if strings.Contains(relPath, "..") {
				return "", fmt.Errorf("path traversal detected: %s", cleanPath)
			}
			return absPath, nil
		}
	}
	return "", fmt.Errorf("path not in allowed directories: %s", cleanPath)
}

// ValidateWebSocketOrigin validates the Origin header for WebSocket connections
// to prevent cross-site WebSocket hijacking attacks. It allows connections from:
// - Same origin (localhost with various ports for development)
// - Configured allowed origins from environment/config
// - Local network origins for development (127.0.0.1, localhost)
func ValidateWebSocketOrigin(origin, host string) bool {
	if origin == "" {
		// Allow empty origin for same-origin requests (some browsers)
		return true
	}

	originURL, err := url.Parse(origin)
	if err != nil {
		return false
	}

	hostURL, err := url.Parse("http://" + host)
	if err != nil {
		return false
	}

	// Allow same origin
	if originURL.Host == hostURL.Host {
		return true
	}

	// Allow localhost variations for development
	allowedHosts := []string{
		"localhost",
		"127.0.0.1",
		"::1",
	}

	originHost := strings.Split(originURL.Host, ":")[0]
	for _, allowed := range allowedHosts {
		if originHost == allowed {
			return true
		}
	}

	// Check configured allowed origins from environment
	if allowedOrigins := viper.GetString("allowed_websocket_origins"); allowedOrigins != "" {
		for _, allowed := range strings.Split(allowedOrigins, ",") {
			allowed = strings.TrimSpace(allowed)
			if origin == allowed || originURL.Host == allowed {
				return true
			}
		}
	}

	return false
}
