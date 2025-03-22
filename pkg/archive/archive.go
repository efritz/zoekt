// Package archive provides access to github.com/sourcegraph/zoekt/internal/archive
//
// This package is a mirror of the internal package with the same name, providing
// access to its exported identifiers.
package archive

import original "github.com/sourcegraph/zoekt/internal/archive"

// Archive is re-exported from the internal package.
type Archive = original.Archive

// File is re-exported from the internal package.
type File = original.File

// Options is re-exported from the internal package.
type Options = original.Options

// Index is re-exported from the internal package.
var Index = original.Index

// OpenReader is re-exported from the internal package.
var OpenReader = original.OpenReader
