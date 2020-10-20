package controllers

import (
	core "crud/src/__core__"
	applications "crud/src/applications/books"
)

// Controller ...
type controller struct {
	usecase applications.IBooks
}

// Handle ...
func (c *controller) Handle(body interface{}) *core.HTTPResponse {

}
