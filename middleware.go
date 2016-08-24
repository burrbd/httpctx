package httpctx

import (
	"net/http"
	"time"

	"golang.org/x/net/context"
)

// TimeoutMiddlewareDecorator returns a handler decorator function with timeout set in context.
func TimeoutMiddlewareDecorator(d time.Duration) func(Handler) Handler {
	return func(h Handler) Handler {
		return HandlerFunc(func(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
			ctx, cancel := context.WithTimeout(ctx, d)
			defer cancel()
			return h.ServeHTTP(ctx, w, req)
		})
	}
}

// errorResolver is a function for resolving errors.
type errorResolver func(http.ResponseWriter, context.Context, error)

// ErrorMiddlewareDecorator catches and writes any errors to the response writer.
func ErrorMiddlewareDecorator(fn errorResolver) func(Handler) Handler {
	return func(h Handler) Handler {
		return HandlerFunc(func(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
			if err := h.ServeHTTP(ctx, w, req); err != nil {
				fn(w, ctx, err)
			}
			return nil
		})
	}
}
