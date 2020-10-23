package adapters

import (
	core "crud/src/__core__"

	"github.com/stretchr/testify/mock"
)

// MiddlewareSpy ...
type MiddlewareSpy struct {
	mock.Mock
}

// Handle ...
func (m *MiddlewareSpy) Handle(input *core.HTTPMiddleware) {
	m.Called(input)
}
