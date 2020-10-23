package adapters

import (
	"github.com/stretchr/testify/mock"

	core "crud/src/__core__"
)

// CreateControllerSpy ...
type CreateControllerSpy struct {
	mock.Mock
}

// Handle ...
func (c *CreateControllerSpy) Handle(httpRequest *core.HTTPRequest) *core.HTTPResponse {
	args := c.Called(httpRequest)

	return args.Get(0).(*core.HTTPResponse)
}
