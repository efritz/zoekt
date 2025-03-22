// Package profiler provides access to github.com/sourcegraph/zoekt/internal/profiler
//
// This package is a mirror of the internal package with the same name, providing
// access to its exported identifiers.
package profiler

import original "github.com/sourcegraph/zoekt/internal/profiler"

// Init is re-exported from the internal package.
var Init = original.Init
