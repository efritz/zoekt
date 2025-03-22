// Package trace provides access to github.com/sourcegraph/zoekt/internal/trace
//
// This package is a mirror of the internal package with the same name, providing
// access to its exported identifiers.
package trace

import original "github.com/sourcegraph/zoekt/internal/trace"

// Trace is re-exported from the internal package.
type Trace = original.Trace

// Tracer is re-exported from the internal package.
type Tracer = original.Tracer

// ContextWithSpanContext is re-exported from the internal package.
var ContextWithSpanContext = original.ContextWithSpanContext

// ContextWithTrace is re-exported from the internal package.
var ContextWithTrace = original.ContextWithTrace

// GetOpenTracer is re-exported from the internal package.
var GetOpenTracer = original.GetOpenTracer

// Middleware is re-exported from the internal package.
var Middleware = original.Middleware

// New is re-exported from the internal package.
var New = original.New

// Printf is re-exported from the internal package.
var Printf = original.Printf

// SpanContextFromContext is re-exported from the internal package.
var SpanContextFromContext = original.SpanContextFromContext

// StartSpanFromContext is re-exported from the internal package.
var StartSpanFromContext = original.StartSpanFromContext

// StartSpanFromContextWithTracer is re-exported from the internal package.
var StartSpanFromContextWithTracer = original.StartSpanFromContextWithTracer

// TraceFromContext is re-exported from the internal package.
var TraceFromContext = original.TraceFromContext

// WithOpenTracingEnabled is re-exported from the internal package.
var WithOpenTracingEnabled = original.WithOpenTracingEnabled
