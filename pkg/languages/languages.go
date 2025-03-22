// Package languages provides access to github.com/sourcegraph/zoekt/internal/languages
//
// This package is a mirror of the internal package with the same name, providing
// access to its exported identifiers.
package languages

import original "github.com/sourcegraph/zoekt/internal/languages"

// GetLanguage is re-exported from the internal package.
var GetLanguage = original.GetLanguage

// GetLanguageByAlias is re-exported from the internal package.
var GetLanguageByAlias = original.GetLanguageByAlias
