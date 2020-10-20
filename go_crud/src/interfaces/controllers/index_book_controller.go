package controllers

import (
	"fmt"
	"net/http"

	core "crud/src/__core__"
	applications "crud/src/applications/books"
	entities "crud/src/business/entities"
)

// Controller ...
type indexBookController struct {
	usecase applications.IBooks
}

// Handle ...
func (c *indexBookController) Handle(body interface{}) *core.HTTPResponse {
	castBody := body.(*entities.InputCreateBook)

	fmt.Println(castBody)

	return &core.HTTPResponse{
		StatusCode: http.StatusCreated,
	}
}

// IndexBookController ...
func IndexBookController(usecase applications.IBooks) core.IController {
	return &indexBookController{usecase: usecase}
}
