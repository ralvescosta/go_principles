package main

import (
	"clean/src/application"
	"clean/src/infrastructure"
	"clean/src/interfaces"
)

var (
	implRepository infrastructure.IHttpGithubApiRepository = infrastructure.Constructor()
	implUsecase    application.ISearchGithubUserUsecase    = application.Constructor(implRepository)
	implController interfaces.IController                  = interfaces.Constructor(implUsecase)
)

func main() {

	implController.Handle("jaoao")
}
