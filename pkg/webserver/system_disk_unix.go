// file: pkg/webserver/system_disk_unix.go
// version: 1.0.0
// guid: 3c0e5d14-6fd2-4f6b-9b51-3f7fce71c9c2
//go:build !windows

package webserver

import "syscall"

func diskUsage() (free uint64, total uint64) {
	var statfs syscall.Statfs_t
	if err := syscall.Statfs("/", &statfs); err != nil {
		return 0, 0
	}
	return statfs.Bfree * uint64(statfs.Bsize), statfs.Blocks * uint64(statfs.Bsize)
}
