package infrastructure

import (
	"clean/src/bussiness"
)

type IHttpGithubApiRepository interface {
	GithubApi(userName string) bussiness.GithubUserEntity
}

type repository struct{}

func (*repository) GithubApi(userName string) bussiness.GithubUserEntity {
	user := bussiness.GithubUserEntity{Id: 1, Name: userName, Bio: "blablabla"}

	return user
}

func Constructor() IHttpGithubApiRepository {
	return &repository{}
}
