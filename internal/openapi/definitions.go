package openapi

import "embed"

// Definitions embeds all files and directories in openapi directory.
//
//go:embed *
var Definitions embed.FS
