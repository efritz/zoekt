// Package shards provides access to github.com/sourcegraph/zoekt/internal/shards
//
// This package is a mirror of the internal package with the same name, providing
// access to its exported identifiers.
package shards

import original "github.com/sourcegraph/zoekt/internal/shards"

// DirectoryWatcher is re-exported from the internal package.
type DirectoryWatcher = original.DirectoryWatcher

// NewDirectorySearcher is re-exported from the internal package.
var NewDirectorySearcher = original.NewDirectorySearcher

// NewDirectorySearcherFast is re-exported from the internal package.
var NewDirectorySearcherFast = original.NewDirectorySearcherFast
