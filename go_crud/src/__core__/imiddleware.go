package core

import "net/http"

// HTTPMiddleware ...
type HTTPMiddleware struct {
	Req  *http.Request
	Res  http.ResponseWriter
	Next http.Handler
}

// IMiddleware ...
type IMiddleware interface {
	Handle(input *HTTPMiddleware)
}
