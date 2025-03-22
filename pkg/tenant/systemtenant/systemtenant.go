// Package systemtenant provides access to github.com/sourcegraph/zoekt/internal/tenant/systemtenant
//
// This package is a mirror of the internal package with the same name, providing
// access to its exported identifiers.
package systemtenant

import original "github.com/sourcegraph/zoekt/internal/tenant/systemtenant"

// Is is re-exported from the internal package.
var Is = original.Is

// WithUnsafeContext is re-exported from the internal package.
var WithUnsafeContext = original.WithUnsafeContext
