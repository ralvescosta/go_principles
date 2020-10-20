package adapters

import (
	core "crud/src/__core__"

	"encoding/json"
	"net/http"
)

// RouteAdapt ...
func RouteAdapt(controller core.IController, body interface{}) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		err := json.NewDecoder(req.Body).Decode(body)
		if err != nil {
			return
		}

		result := controller.Handle(body)

		res.WriteHeader(result.StatusCode)
		if result.Body != nil {
			json.NewEncoder(res).Encode(result.Body)
		}
	}
}
