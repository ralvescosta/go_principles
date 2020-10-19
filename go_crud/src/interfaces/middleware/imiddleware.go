package interfaces

import "net/http"

// IMiddleware ...
type IMiddleware interface {
	Handle(next http.Handler) http.Handler
}
