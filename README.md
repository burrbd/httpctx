# Context HTTP

This repository contains some sample code used for my personal
projects. If you wish to use this package, I suggest copying
into your code directly, rather than importing.

Handlers are like `http.Handler` and `http.HandlerFunc` from the
standard library but include an additional `context.Context`
parameter and return an ```error```.

```go
ServeHTTP(context.Context, http.ResponseWriter, *http.Request) error
```

Middleware includes an `errorResolver` func type that can be used
with an error middleware to handle errors centrally.

There's also an `HTTPRouterAdaptor` which accepts a `Handler` and
returns an `httprouter.Handle`, so that you can do:

```go
a := NewHTTPRouterAdaptor(context.Background())
handler := a.Handle(HandlerFunc(func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	// some code
}))
router := httprouter.New()
router.GET("/endpoint", handler)
```
