package controllers

import (
	"fmt"
	"net/http"

	core "crud/src/__core__"
	applications "crud/src/applications/books"
	entities "crud/src/business/entities"
)

// Controller ...
type showBookController struct {
	usecase applications.IBooks
}

// Handle ...
func (c *showBookController) Handle(body interface{}) *core.HTTPResponse {
	castBody := body.(*entities.InputCreateBook)

	fmt.Println(castBody)

	return &core.HTTPResponse{
		StatusCode: http.StatusCreated,
	}
}

// ShowBookController ...
func ShowBookController(usecase applications.IBooks) core.IController {
	return &showBookController{usecase: usecase}
}
