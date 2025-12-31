// file: pkg/webserver/system_disk_windows.go
// version: 1.0.0
// guid: 0f9e7a8b-1d2c-4b5a-9c0d-2e3f4a5b6c7d
//go:build windows

package webserver

func diskUsage() (free uint64, total uint64) {
	return 0, 0
}
