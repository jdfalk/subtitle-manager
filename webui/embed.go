package webui

import "embed"

// FS holds the built web UI assets.
//
//go:embed dist/* dist/assets/*
var FS embed.FS
