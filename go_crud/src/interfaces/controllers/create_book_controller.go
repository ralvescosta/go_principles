package controllers

import (
	"net/http"

	core "crud/src/__core__"
	applications "crud/src/applications/books"
	entities "crud/src/business/entities"
)

// Controller ...
type createBookController struct {
	usecase applications.IBooks
}

// Handle ...
func (c *createBookController) Handle(httpRequest *core.HTTPRequest) *core.HTTPResponse {
	castBody := httpRequest.Body.(*entities.InputCreateBook)

	_, err := c.usecase.RegisterABook(castBody)

	if err != nil {
		return &core.HTTPResponse{
			StatusCode: http.StatusInternalServerError,
		}
	}

	return &core.HTTPResponse{
		StatusCode: http.StatusCreated,
	}
}

// CreateABookController ...
func CreateABookController(usecase applications.IBooks) core.IController {
	return &createBookController{usecase: usecase}
}
