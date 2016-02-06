// Package middleware provides a simple middleware framework based on http.Handler and
// collects a few useful small middleware. It is heavily based on the concepts of
// justinas/alice but with a minimal API
package middleware

import (
	"net/http"
)

// MiddleWare is a function that takes a http.Handler and wraps your middleware around it.
// Call next.ServeHTTP() in your middleware if you want to continue the middleware chain
// or return without calling it to stop propagation.
type MiddleWare func(next http.Handler) http.Handler

// Chain is a list of middleware that are called in order
type Chain []MiddleWare

// New creates a new middleware chain
func New(middleware ...MiddleWare) Chain {
	c := Chain{}
	return append(c, middleware...)
}

// Serve wraps the provided handler in the middleware chain and returns a handler
func (chain Chain) Serve(h http.Handler) http.Handler {
	ret := h
	for _, each := range chain {
		ret = each(ret)
	}
	return ret
}
