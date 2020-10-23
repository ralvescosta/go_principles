package interfaces

import (
	core "crud/src/__core__"
)

type middleware struct{}

// Handle ...
func (*middleware) Handle(input *core.HTTPMiddleware) {

	input.Res.Header().Add("Content-Type", "application/json; charset=utf-8")
	input.Res.Header().Add("X-DNS-Prefetch-Control", "off")
	input.Res.Header().Add("X-Frame-Options", "SAMEORIGIN")
	input.Res.Header().Add("Strict-Transport-Security", "max-age=15552000; includeSubDomains")
	input.Res.Header().Add("X-Download-Options", "noopen")
	input.Res.Header().Add("X-Content-Type-Options", "nosniff")
	input.Res.Header().Add("X-XSS-Protection", "1; mode=block")
	input.Res.Header().Add("Content-Security-Policy", "default-src 'none'")
	input.Res.Header().Add("X-Content-Security-Policy", "default-src 'none'")
	input.Res.Header().Add("X-WebKit-CSP", "default-src 'none'")
	input.Res.Header().Add("X-Permitted-Cross-Domain-Policies", "none")
	input.Res.Header().Add("Referrer-Policy", "origin-when-cross-origin,strict-origin-when-cross-origin")
	input.Res.Header().Add("Access-Control-Allow-Origin", "*")
	input.Res.Header().Add("Vary", "Accept-Encoding")

	input.Next.ServeHTTP(input.Res, input.Req)
}

// Cors ...
func Cors() core.IMiddleware {
	return &middleware{}
}
