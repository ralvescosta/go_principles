package signin_application

import (
	"fmt"
	bussiness "gomux_gorm/src/signin_module/bussiness/entities"
	repositories "gomux_gorm/src/signin_module/frameworks/repositories"
)

type usecase struct {
	repository *repositories.IUserRepository
}

type ISigninUsecase interface {
	SigninUsecase(user *bussiness.RegisterUsersEntity) *bussiness.RegisterUsersEntity
}

func (*usecase) SigninUsecase(user *bussiness.RegisterUsersEntity) *bussiness.RegisterUsersEntity {
	fmt.Println("SigninUsecase")

	return user
}

func SigninUsecaseConstructor(repository *repositories.IUserRepository) ISigninUsecase {
	return &usecase{repository}
}
