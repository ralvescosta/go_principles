package core

// HTTPRequest ....
type HTTPRequest struct {
	Body   interface{}
	Params map[string]string
}

// HTTPResponse ....
type HTTPResponse struct {
	StatusCode int
	Body       interface{}
	Headers    interface{}
}

// IController ...
type IController interface {
	Handle(httpRequest *HTTPRequest) *HTTPResponse
}
