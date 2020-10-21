package adapters

import (
	core "crud/src/__core__"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// RouteAdapt ...
func RouteAdapt(controller core.IController, body interface{}) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {

		if body != nil {
			err := json.NewDecoder(req.Body).Decode(body)
			if err != nil {
				return
			}
		}

		vars := mux.Vars(req)

		httpRequest := &core.HTTPRequest{
			Body:   body,
			Params: vars,
		}

		result := controller.Handle(httpRequest)

		res.WriteHeader(result.StatusCode)
		if result.Body != nil {
			json.NewEncoder(res).Encode(result.Body)
		}
	}
}
