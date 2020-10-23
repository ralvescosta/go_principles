package adapters

import (
	"net/http/httptest"
	"testing"

	core "crud/src/__core__"
	stuntman "crud/src/frameworks/__test__"

	"github.com/stretchr/testify/assert"
)

type mocksMiddlewareStruct struct {
	httpMiddleware *core.HTTPMiddleware
	res            *httptest.ResponseRecorder
	middlewareSpy  *stuntman.MiddlewareSpy
}

func makeMiddlewareMocks() *mocksMiddlewareStruct {

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	res := httptest.NewRecorder()

	httpMiddleware := &core.HTTPMiddleware{
		Req: req,
		Res: res,
	}

	middlewareSpy := &stuntman.MiddlewareSpy{}

	return &mocksMiddlewareStruct{
		httpMiddleware: httpMiddleware,
		res:            res,
		middlewareSpy:  middlewareSpy,
	}
}

func TestMiddlewareAdapt(t *testing.T) {
	mocks := makeMiddlewareMocks()

	mocks.middlewareSpy.On("Handle", mocks.httpMiddleware)

	sut := MiddlewareAdapt(mocks.middlewareSpy)

	sut(mocks.httpMiddleware.Next)

	result := mocks.res.Result()

	assert.Equal(t, result.StatusCode, 200)
}
