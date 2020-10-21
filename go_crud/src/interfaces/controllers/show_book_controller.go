package controllers

import (
	"net/http"

	core "crud/src/__core__"
	applications "crud/src/applications/books"
)

// Controller ...
type showBookController struct {
	usecase applications.IBooks
}

// Handle ...
func (c *showBookController) Handle(httpRequest *core.HTTPRequest) *core.HTTPResponse {

	result, err := c.usecase.GetAllBooks()
	if err != nil {
		return &core.HTTPResponse{
			StatusCode: http.StatusInternalServerError,
		}
	}

	if result == nil {
		return &core.HTTPResponse{
			StatusCode: http.StatusNotFound,
		}
	}

	return &core.HTTPResponse{
		StatusCode: http.StatusOK,
		Body:       result,
	}
}

// ShowBookController ...
func ShowBookController(usecase applications.IBooks) core.IController {
	return &showBookController{usecase: usecase}
}
