package controllers

import (
	"net/http"
	"strconv"

	core "crud/src/__core__"
	applications "crud/src/applications/books"
)

// Controller ...
type indexBookController struct {
	usecase applications.IBooks
}

// Handle ...
func (c *indexBookController) Handle(httpRequest *core.HTTPRequest) *core.HTTPResponse {
	id, err := strconv.ParseUint(httpRequest.Params["id"], 10, 64)
	if err != nil {
		return &core.HTTPResponse{
			StatusCode: http.StatusInternalServerError,
		}
	}

	result, err := c.usecase.FindABook(id)
	if err != nil {
		return &core.HTTPResponse{
			StatusCode: http.StatusInternalServerError,
		}
	}

	return &core.HTTPResponse{
		StatusCode: http.StatusOK,
		Body:       result,
	}
}

// IndexBookController ...
func IndexBookController(usecase applications.IBooks) core.IController {
	return &indexBookController{usecase: usecase}
}
