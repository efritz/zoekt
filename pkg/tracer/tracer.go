// Package tracer provides access to github.com/sourcegraph/zoekt/internal/tracer
//
// This package is a mirror of the internal package with the same name, providing
// access to its exported identifiers.
package tracer

import original "github.com/sourcegraph/zoekt/internal/tracer"

// Init is re-exported from the internal package.
var Init = original.Init
