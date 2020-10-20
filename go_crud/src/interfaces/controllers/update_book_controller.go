package controllers

import (
	"fmt"
	"net/http"

	core "crud/src/__core__"
	applications "crud/src/applications/books"
	entities "crud/src/business/entities"
)

// Controller ...
type updateBookController struct {
	usecase applications.IBooks
}

// Handle ...
func (c *updateBookController) Handle(body interface{}) *core.HTTPResponse {
	castBody := body.(*entities.InputCreateBook)

	fmt.Println(castBody)

	return &core.HTTPResponse{
		StatusCode: http.StatusCreated,
	}
}

// UpdateBookController ...
func UpdateBookController(usecase applications.IBooks) core.IController {
	return &showBookController{usecase: usecase}
}
