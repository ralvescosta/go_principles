package controllers

import "net/http"

// IController ...
type IController interface {
	Handle(res http.ResponseWriter, req *http.Request)
}
