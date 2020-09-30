package application

import "fmt"

type usecase struct{}

type ISearchGithubUserUsecase interface {
	Search(userName string) string
}

func (*usecase) Search(userName string) string {
	fmt.Println("Hello from Usecase")
	return "Hello from Usecase"

}

func NewUsecase() ISearchGithubUserUsecase {
	return &usecase{}
}
