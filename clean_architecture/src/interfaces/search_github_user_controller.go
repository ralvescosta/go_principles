package interfaces

import (
	"clean/src/application"
	"fmt"
)

type controller struct {
	usecase application.ISearchGithubUserUsecase
}

type IController interface {
	Handle(userName string) string
}

func (c *controller) Handle(userName string) string {

	user := c.usecase.Search(userName)

	fmt.Println("controller", user)

	return "hello from controller"
}

func Constructor(usecase *application.ISearchGithubUserUsecase) IController {

	return &controller{usecase: *usecase}
}
