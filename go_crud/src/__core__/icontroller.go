package core

// HTTPResponse ....
type HTTPResponse struct {
	StatusCode int
	Body       interface{}
	Headers    interface{}
}

// IController ...
type IController interface {
	Handle(body interface{}) *HTTPResponse
}
