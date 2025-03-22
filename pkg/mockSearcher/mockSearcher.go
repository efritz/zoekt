// Package mockSearcher provides access to github.com/sourcegraph/zoekt/internal/mockSearcher
//
// This package is a mirror of the internal package with the same name, providing
// access to its exported identifiers.
package mockSearcher

import original "github.com/sourcegraph/zoekt/internal/mockSearcher"

// MockSearcher is re-exported from the internal package.
type MockSearcher = original.MockSearcher
