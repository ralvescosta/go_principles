package interfaces

import (
	"clean/src/application"
	"fmt"
)

type controller struct{}

type IController interface {
	Constructor(usecase application.ISearchGithubUserUsecase) func(userName string) string
}

func (*controller) Constructor(usecase application.ISearchGithubUserUsecase) func(userName string) string {
	return func(userName string) string {
		fmt.Println("hello from controller")

		usecase.Search(userName)

		return "hello from controller"
	}
}

func NewController() IController {

	return &controller{}
}
