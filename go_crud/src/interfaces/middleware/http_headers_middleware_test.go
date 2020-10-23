package interfaces

import (
	core "crud/src/__core__"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mocksMiddlewareStruct struct {
	httpMiddleware *core.HTTPMiddleware
	res            *httptest.ResponseRecorder
	headers        map[string][]string
}

func makeMiddlewareMocks() *mocksMiddlewareStruct {

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	res := httptest.NewRecorder()
	handler := http.NewServeMux()

	headers := make(map[string][]string)

	headers["Access-Control-Allow-Origin"] = []string{"*"}
	headers["Content-Security-Policy"] = []string{"default-src 'none'"}
	headers["Content-Type"] = []string{"text/plain; charset=utf-8"}
	headers["Referrer-Policy"] = []string{"origin-when-cross-origin,strict-origin-when-cross-origin"}
	headers["Strict-Transport-Security"] = []string{"max-age=15552000; includeSubDomains"}
	headers["Vary"] = []string{"Accept-Encoding"}
	headers["X-Content-Security-Policy"] = []string{"default-src 'none'"}
	headers["X-Content-Type-Options"] = []string{"nosniff"}
	headers["X-Dns-Prefetch-Control"] = []string{"off"}
	headers["X-Download-Options"] = []string{"noopen"}
	headers["X-Frame-Options"] = []string{"SAMEORIGIN"}
	headers["X-Permitted-Cross-Domain-Policies"] = []string{"none"}
	headers["X-Webkit-Csp"] = []string{"default-src 'none'"}
	headers["X-Xss-Protection"] = []string{"1; mode=block"}

	httpMiddleware := &core.HTTPMiddleware{
		Req:  req,
		Res:  res,
		Next: handler,
	}

	return &mocksMiddlewareStruct{
		httpMiddleware: httpMiddleware,
		res:            res,
		headers:        headers,
	}
}

func TestCorsMiddleware(t *testing.T) {

	mocks := makeMiddlewareMocks()

	sut := Cors()

	sut.Handle(mocks.httpMiddleware)

	result := mocks.res.Result()

	var headers map[string][]string
	headers = result.Header

	assert.Equal(t, mocks.headers, headers)
}
