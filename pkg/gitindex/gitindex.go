// Package gitindex provides access to github.com/sourcegraph/zoekt/internal/gitindex
//
// This package is a mirror of the internal package with the same name, providing
// access to its exported identifiers.
package gitindex

import original "github.com/sourcegraph/zoekt/internal/gitindex"

// BlobLocation is re-exported from the internal package.
type BlobLocation = original.BlobLocation

// Filter is re-exported from the internal package.
type Filter = original.Filter

// Options is re-exported from the internal package.
type Options = original.Options

// RepoCache is re-exported from the internal package.
type RepoCache = original.RepoCache

// RepoWalker is re-exported from the internal package.
type RepoWalker = original.RepoWalker

// SubmoduleEntry is re-exported from the internal package.
type SubmoduleEntry = original.SubmoduleEntry

// CloneRepo is re-exported from the internal package.
var CloneRepo = original.CloneRepo

// DeleteRepos is re-exported from the internal package.
var DeleteRepos = original.DeleteRepos

// FindGitRepos is re-exported from the internal package.
var FindGitRepos = original.FindGitRepos

// IndexGitRepo is re-exported from the internal package.
var IndexGitRepo = original.IndexGitRepo

// ListRepos is re-exported from the internal package.
var ListRepos = original.ListRepos

// NewFilter is re-exported from the internal package.
var NewFilter = original.NewFilter

// NewRepoCache is re-exported from the internal package.
var NewRepoCache = original.NewRepoCache

// NewRepoWalker is re-exported from the internal package.
var NewRepoWalker = original.NewRepoWalker

// ParseGitModules is re-exported from the internal package.
var ParseGitModules = original.ParseGitModules

// Path is re-exported from the internal package.
var Path = original.Path

// SetTemplatesFromOrigin is re-exported from the internal package.
var SetTemplatesFromOrigin = original.SetTemplatesFromOrigin
