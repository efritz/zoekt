// Package tenant provides access to github.com/sourcegraph/zoekt/internal/tenant
//
// This package is a mirror of the internal package with the same name, providing
// access to its exported identifiers.
package tenant

import original "github.com/sourcegraph/zoekt/internal/tenant"

// ErrMissingTenant is re-exported from the internal package.
var ErrMissingTenant = original.ErrMissingTenant

// Propagator is re-exported from the internal package.
type Propagator = original.Propagator

// EnforceTenant is re-exported from the internal package.
var EnforceTenant = original.EnforceTenant

// FromContext is re-exported from the internal package.
var FromContext = original.FromContext

// HasAccess is re-exported from the internal package.
var HasAccess = original.HasAccess

// Log is re-exported from the internal package.
var Log = original.Log

// SrcPrefix is re-exported from the internal package.
var SrcPrefix = original.SrcPrefix

// StreamServerInterceptor is re-exported from the internal package.
var StreamServerInterceptor = original.StreamServerInterceptor

// UnaryServerInterceptor is re-exported from the internal package.
var UnaryServerInterceptor = original.UnaryServerInterceptor
