package httpctx

import (
	"net/http"

	"golang.org/x/net/context"
)

// Handler is custom Handler interface for HTTP requests and includes context argument.
type Handler interface {
	ServeHTTP(context.Context, http.ResponseWriter, *http.Request) error
}

// HandlerFunc type is adaptor function so that regular functions can be used as handlers.
type HandlerFunc func(context.Context, http.ResponseWriter, *http.Request) error

// ServeHTTP returns f(ctx, w, req).
func (f HandlerFunc) ServeHTTP(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	return f(ctx, w, req)
}
