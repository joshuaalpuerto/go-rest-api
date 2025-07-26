package middlewares

import (
	"net/http"
)

// Middleware represents a middleware function
type MiddlewareFunc func(handler http.HandlerFunc) http.HandlerFunc

// Chain applies multiple middlewares to a handler
func Chain(h http.HandlerFunc, mw ...MiddlewareFunc) http.HandlerFunc {
	for i := len(mw) - 1; i >= 0; i-- {
		mwFunc := mw[i]
		if mwFunc != nil {
			h = mwFunc(h)
		}
	}
	return h
}
