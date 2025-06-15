// file: cmd/version.go
package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

// Version information (set from main package)
var (
	appVersion = "dev"
	buildTime  = "unknown"
	gitCommit  = "unknown"
)

// SetVersionInfo sets version information from the main package
//
// Parameters:
//   - version: Application version string
//   - buildTimeParam: Build timestamp string
//   - gitCommitParam: Git commit hash string
//
// This function is called from main.go to pass build-time version information
// set via ldflags to the cmd package for use in version commands.
func SetVersionInfo(version, buildTimeParam, gitCommitParam string) {
	appVersion = version
	buildTime = buildTimeParam
	gitCommit = gitCommitParam
}

// GetVersionInfo returns the current version information
//
// Returns:
//   - version: Application version string
//   - buildTime: Build timestamp string
//   - gitCommit: Git commit hash string
//
// This function provides access to version information for other packages
// that may need to display or log version details.
func GetVersionInfo() (version, buildTime, gitCommit string) {
	return appVersion, buildTime, gitCommit
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Long: `Display detailed version information including application version,
build time, Git commit, and Go runtime version.

This command shows comprehensive version details useful for debugging
and support purposes.`,
	Run: func(cmd *cobra.Command, args []string) {
		printVersion()
	},
}

// printVersion outputs formatted version information
//
// Displays:
//   - Application name and version
//   - Build timestamp
//   - Git commit hash
//   - Go runtime version
//   - Operating system and architecture
//
// This function provides a consistent format for version output
// used by both the version command and --version flag.
func printVersion() {
	fmt.Printf("Subtitle Manager %s\n", appVersion)
	fmt.Printf("Build Time: %s (UTC)\n", buildTime)
	fmt.Printf("Git Commit: %s\n", gitCommit)
	fmt.Printf("Go Version: %s\n", runtime.Version())
	fmt.Printf("OS/Arch: %s/%s\n", runtime.GOOS, runtime.GOARCH)
}

func init() {
	// Add version subcommand
	rootCmd.AddCommand(versionCmd)
}
