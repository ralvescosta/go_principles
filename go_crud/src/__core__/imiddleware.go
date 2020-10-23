package core

import "net/http"

// HTTPMiddleware ...
type HTTPMiddleware struct {
	Req  http.ResponseWriter
	Res  *http.Request
	Next http.Handler
}

// IMiddleware ...
type IMiddleware interface {
	Handle(input *HTTPMiddleware)
}
