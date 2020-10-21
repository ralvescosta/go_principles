package controllers

import (
	"fmt"
	"net/http"

	core "crud/src/__core__"
	applications "crud/src/applications/books"
	entities "crud/src/business/entities"
)

// Controller ...
type deleteBookController struct {
	usecase applications.IBooks
}

// Handle ...
func (c *deleteBookController) Handle(httpRequest *core.HTTPRequest) *core.HTTPResponse {
	castBody := httpRequest.Body.(*entities.InputCreateBook)

	fmt.Println(castBody)

	return &core.HTTPResponse{
		StatusCode: http.StatusCreated,
	}
}

// DeleteABookController ...
func DeleteABookController(usecase applications.IBooks) core.IController {
	return &deleteBookController{usecase: usecase}
}
