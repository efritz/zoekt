// Package tenanttest provides access to github.com/sourcegraph/zoekt/internal/tenant/tenanttest
//
// This package is a mirror of the internal package with the same name, providing
// access to its exported identifiers.
package tenanttest

import original "github.com/sourcegraph/zoekt/internal/tenant/tenanttest"

// TestTenantCounter is re-exported from the internal package.
var TestTenantCounter = original.TestTenantCounter

// MockEnforce is re-exported from the internal package.
var MockEnforce = original.MockEnforce

// NewTestContext is re-exported from the internal package.
var NewTestContext = original.NewTestContext

// ResetTestTenants is re-exported from the internal package.
var ResetTestTenants = original.ResetTestTenants
