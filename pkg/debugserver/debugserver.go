// Package debugserver provides access to github.com/sourcegraph/zoekt/internal/debugserver
//
// This package is a mirror of the internal package with the same name, providing
// access to its exported identifiers.
package debugserver

import original "github.com/sourcegraph/zoekt/internal/debugserver"

// DebugPage is re-exported from the internal package.
type DebugPage = original.DebugPage

// AddHandlers is re-exported from the internal package.
var AddHandlers = original.AddHandlers
