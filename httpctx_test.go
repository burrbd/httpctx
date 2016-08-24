package httpctx_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"golang.org/x/net/context"

	"github.com/burrbd/httpctx"
)

func TestAdaptor(t *testing.T) {
	var isCalled bool
	a := httpctx.NewHTTPRouterAdaptor(context.Background())
	h := a.Handle(httpctx.HandlerFunc(func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		isCalled = true
		return nil
	}))
	req, _ := http.NewRequest("GET", "http://dummyurl", nil)
	h(httptest.NewRecorder(), req, nil)
	if !isCalled {
		t.Error("expected handler func to be called")
	}
}

func TestTimeoutMiddleware(t *testing.T) {
	timeout := 1 * time.Second
	sleep := 2 * time.Second
	toMiddleware := httpctx.TimeoutMiddlewareDecorator(timeout)
	h := toMiddleware(httpctx.HandlerFunc(func(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
		select {
		case <-time.Tick(sleep):
			t.Error("expected to receive off ctx.Done() channel")
		case <-ctx.Done():
		}
		return nil
	}))
	req, _ := http.NewRequest("GET", "http://dummyurl", nil)
	h.ServeHTTP(context.Background(), httptest.NewRecorder(), req)
}

func TestErrorMiddlewareHandlesError(t *testing.T) {
	fn := func(w http.ResponseWriter, ctx context.Context, err error) {
		if err == nil {
			t.Fatalf("expected an error")
		}
		if err.Error() != "foo" {
			t.Errorf("expected a foo error")
		}
	}
	errHandler := httpctx.ErrorMiddlewareDecorator(fn)
	h := errHandler(httpctx.HandlerFunc(func(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
		return errors.New("foo")
	}))
	req, _ := http.NewRequest("GET", "http://dummyurl", nil)
	h.ServeHTTP(context.Background(), httptest.NewRecorder(), req)
}

func TestErrorMiddlewareIsNotCalledWhenNoError(t *testing.T) {
	var isCalled bool
	fn := func(w http.ResponseWriter, ctx context.Context, err error) {
		isCalled = true
	}
	errHandler := httpctx.ErrorMiddlewareDecorator(fn)
	h := errHandler(httpctx.HandlerFunc(func(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
		return nil
	}))
	req, _ := http.NewRequest("GET", "http://dummyurl", nil)
	h.ServeHTTP(context.Background(), httptest.NewRecorder(), req)
	if isCalled {
		t.Error("error resolver should not be called")
	}
}
