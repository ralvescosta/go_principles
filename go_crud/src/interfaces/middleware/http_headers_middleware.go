package interfaces

import (
	core "crud/src/__core__"
)

type middleware struct{}

// Handle ...
func (*middleware) Handle(input *core.HTTPMiddleware) {

	input.Req.Header().Add("Content-Type", "application/json; charset=utf-8")
	input.Req.Header().Add("X-DNS-Prefetch-Control", "off")
	input.Req.Header().Add("X-Frame-Options", "SAMEORIGIN")
	input.Req.Header().Add("Strict-Transport-Security", "max-age=15552000; includeSubDomains")
	input.Req.Header().Add("X-Download-Options", "noopen")
	input.Req.Header().Add("X-Content-Type-Options", "nosniff")
	input.Req.Header().Add("X-XSS-Protection", "1; mode=block")
	input.Req.Header().Add("Content-Security-Policy", "default-src 'none'")
	input.Req.Header().Add("X-Content-Security-Policy", "default-src 'none'")
	input.Req.Header().Add("X-WebKit-CSP", "default-src 'none'")
	input.Req.Header().Add("X-Permitted-Cross-Domain-Policies", "none")
	input.Req.Header().Add("Referrer-Policy", "origin-when-cross-origin,strict-origin-when-cross-origin")
	input.Req.Header().Add("Access-Control-Allow-Origin	", "*")
	input.Req.Header().Add("Vary", "Accept-Encoding")

	input.Next.ServeHTTP(input.Req, input.Res)
}

// Cors ...
func Cors() core.IMiddleware {
	return &middleware{}
}
