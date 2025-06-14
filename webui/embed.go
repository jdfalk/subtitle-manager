package webui

import "embed"

// FS holds the built web UI assets.
//
//go:embed dist
var FS embed.FS
