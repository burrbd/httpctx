package httpctx

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/context"
)

const httpRouterParamsKey = "params"

// NewHTTPRouterAdaptor constructs a new Adaptor struct.
func NewHTTPRouterAdaptor(ctx context.Context) *HTTPRouterAdaptor {
	return &HTTPRouterAdaptor{ctx}
}

// HTTPRouterAdaptor is an adaptor to convert a httpctx.Handler to httprouter.Handle.
type HTTPRouterAdaptor struct {
	ctx context.Context
}

// Handle converts httpctx.Handler to httprouter.Handle.
func (a *HTTPRouterAdaptor) Handle(h Handler) httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		ctx := context.WithValue(a.ctx, httpRouterParamsKey, ps)
		h.ServeHTTP(ctx, w, req)
	}
}
