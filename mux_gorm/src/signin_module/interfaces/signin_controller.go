package signin_interfaces

import (
	"encoding/json"
	usecases "gomux_gorm/src/signin_module/application/usecases"
	bussiness "gomux_gorm/src/signin_module/bussiness/entities"

	"net/http"
)

type ISigninController interface {
	Handle(res http.ResponseWriter, req *http.Request)
}

type controller struct {
	usecase *usecases.ISigninUsecase
}

func (c *controller) Handle(res http.ResponseWriter, req *http.Request) {

	var body bussiness.RegisterUsersEntity
	err := json.NewDecoder(req.Body).Decode(&body)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error unmarshaling array"}`))
		return
	}

	result := (*c.usecase).SigninUsecase(&body)

	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(result)
}

func SigninController(usecase *usecases.ISigninUsecase) ISigninController {
	return &controller{usecase}
}
