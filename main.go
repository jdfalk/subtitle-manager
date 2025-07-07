// Package main is the entry point for the subtitle-manager application.
// It initializes version information and starts the CLI command handler.

package main

import "github.com/jdfalk/subtitle-manager/cmd"

// Version information (set via ldflags during build)
var Version = "dev"       // Application version
var BuildTime = "unknown" // Build timestamp
var GitCommit = "unknown" // Git commit hash

func main() {
	// Pass version information to cmd package
	cmd.SetVersionInfo(Version, BuildTime, GitCommit)
	cmd.Execute()
}
