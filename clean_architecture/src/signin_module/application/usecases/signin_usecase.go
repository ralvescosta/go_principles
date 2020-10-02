package signin_application

import (
	"errors"
	"fmt"
	bussiness "gomux_gorm/src/signin_module/bussiness/entities"
	repositories "gomux_gorm/src/signin_module/frameworks/repositories"
)

type usecase struct {
	repository *repositories.IUserRepository
}

type ISigninUsecase interface {
	SigninUsecase(user *bussiness.RegisterUsersEntity) (*bussiness.RegisterUsersEntity, error)
}

func (u *usecase) SigninUsecase(user *bussiness.RegisterUsersEntity) (*bussiness.RegisterUsersEntity, error) {
	fmt.Println("SigninUsecase")

	userAlreadyRegistered := (*u.repository).FindByEmail(user.Email)

	if userAlreadyRegistered.ID != 0 {
		return nil, errors.New("user already exist")
	}

	(*u.repository).Create(user)

	return user, nil
}

func SigninUsecaseConstructor(repository *repositories.IUserRepository) ISigninUsecase {
	return &usecase{repository}
}
