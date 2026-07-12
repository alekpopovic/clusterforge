package templates

import "embed"

// FS contains the built-in environment templates used by installed CLI binaries.
//
//go:embed env
var FS embed.FS
