package bussiness

type ISearchGithubUserUsecase interface {
	search(userName string) string
}
