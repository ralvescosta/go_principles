package adapters

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	core "crud/src/__core__"
	stuntman "crud/src/frameworks/__test__"
)

type mocksControllerStruct struct {
	httpRequest   *core.HTTPRequest
	httpResponse  *core.HTTPResponse
	controllerSpy *stuntman.CreateControllerSpy
}

func makeControllerMocks() *mocksControllerStruct {

	httpRequest := &core.HTTPRequest{}
	httpResponse := &core.HTTPResponse{
		StatusCode: 201,
	}

	controllerSpy := &stuntman.CreateControllerSpy{}

	return &mocksControllerStruct{
		httpRequest:   httpRequest,
		httpResponse:  httpResponse,
		controllerSpy: controllerSpy,
	}
}

func TestRouteAdapt(t *testing.T) {
	mocks := makeControllerMocks()

	mocks.controllerSpy.On("Handle", mocks.httpRequest).Return(mocks.httpResponse)

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	res := httptest.NewRecorder()

	sut := RouteAdapt(mocks.controllerSpy, nil)

	sut(res, req)

	result := res.Result()

	assert.Equal(t, result.StatusCode, mocks.httpResponse.StatusCode)
}
