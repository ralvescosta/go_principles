package application

import (
	"clean/src/bussiness"
	"clean/src/infrastructure"
)

type ISearchGithubUserUsecase interface {
	Search(userName string) bussiness.GithubUserEntity
}

type usecase struct {
	repository infrastructure.IHttpGithubApiRepository
}

func (u *usecase) Search(userName string) bussiness.GithubUserEntity {

	user := u.repository.GithubApi(userName)

	return user
}

func Constructor(repository *infrastructure.IHttpGithubApiRepository) ISearchGithubUserUsecase {

	return &usecase{repository: *repository}
}
