// Package otlpenv provides access to github.com/sourcegraph/zoekt/internal/otlpenv
//
// This package is a mirror of the internal package with the same name, providing
// access to its exported identifiers.
package otlpenv

import original "github.com/sourcegraph/zoekt/internal/otlpenv"

// ProtocolGRPC is re-exported from the internal package.
const ProtocolGRPC = original.ProtocolGRPC

// ProtocolHTTPJSON is re-exported from the internal package.
const ProtocolHTTPJSON = original.ProtocolHTTPJSON

// ProtocolHTTPProto is re-exported from the internal package.
const ProtocolHTTPProto = original.ProtocolHTTPProto

// Protocol is re-exported from the internal package.
type Protocol = original.Protocol

// GetEndpoint is re-exported from the internal package.
var GetEndpoint = original.GetEndpoint

// GetProtocol is re-exported from the internal package.
var GetProtocol = original.GetProtocol

// IsInsecure is re-exported from the internal package.
var IsInsecure = original.IsInsecure
