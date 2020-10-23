package controllers

import (
	"net/http"
	"strconv"

	core "crud/src/__core__"
	applications "crud/src/applications/books"
)

// Controller ...
type deleteBookController struct {
	usecase applications.IBooks
}

// Handle ...
func (c *deleteBookController) Handle(httpRequest *core.HTTPRequest) *core.HTTPResponse {
	id, err := strconv.ParseUint(httpRequest.Params["id"], 10, 64)
	if err != nil {
		return &core.HTTPResponse{
			StatusCode: http.StatusInternalServerError,
		}
	}

	_, err = c.usecase.DeleteABook(id)
	if err != nil {
		return &core.HTTPResponse{
			StatusCode: http.StatusInternalServerError,
		}
	}

	return &core.HTTPResponse{
		StatusCode: http.StatusOK,
	}
}

// DeleteABookController ...
func DeleteABookController(usecase applications.IBooks) core.IController {
	return &deleteBookController{usecase: usecase}
}
