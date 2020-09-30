package main

import (
	"clean/src/application"
	"clean/src/interfaces"
)

var (
	implUsecase    application.ISearchGithubUserUsecase = application.NewUsecase()
	implController interfaces.IController               = interfaces.NewController()
)

func main() {

	controller := implController.Constructor(implUsecase)
	controller("jaoao")
}
