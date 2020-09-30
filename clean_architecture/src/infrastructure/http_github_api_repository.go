package infrastructure

import (
	"clean/src/bussiness"
)

type HttpGithubApiRepository struct{}

func (*HttpGithubApiRepository) githubApi(userName string) bussiness.GithubUserEntity {
	user := bussiness.GithubUserEntity{Id: 1, Name: userName, Bio: "blablabla"}

	return user
}

func NewHttpRepository() *HttpGithubApiRepository {
	return &HttpGithubApiRepository{}
}
