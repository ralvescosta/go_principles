package adapters

import (
	core "crud/src/__core__"

	"net/http"
)

// MiddlewareAdapt ...
func MiddlewareAdapt(middleware core.IMiddleware) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(req http.ResponseWriter, res *http.Request) {

			middleware.Handle(&core.HTTPMiddleware{
				Req:  req,
				Res:  res,
				Next: next,
			})

		})
	}
}