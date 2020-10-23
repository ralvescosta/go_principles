package controllers

import (
	"net/http"
	"strconv"

	core "crud/src/__core__"
	applications "crud/src/applications/books"
	entities "crud/src/business/entities"
)

// Controller ...
type updateBookController struct {
	usecase applications.IBooks
}

// Handle ...
func (c *updateBookController) Handle(httpRequest *core.HTTPRequest) *core.HTTPResponse {
	id, err := strconv.ParseUint(httpRequest.Params["id"], 10, 64)
	if err != nil {
		return &core.HTTPResponse{
			StatusCode: http.StatusInternalServerError,
		}
	}
	castBody := httpRequest.Body.(*entities.InputCreateBook)

	result, err := c.usecase.UpdateABook(id, castBody)
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

// UpdateBookController ...
func UpdateBookController(usecase applications.IBooks) core.IController {
	return &updateBookController{usecase: usecase}
}
