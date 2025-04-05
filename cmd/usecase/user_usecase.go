package usecase

import (
	"go-api/cmd/model"
	"go-api/cmd/repository"
)

type UserUsecase struct {
	userRepo repository.UserRepositoryInterface
}

func NewUserUsecase(repo repository.UserRepositoryInterface) *UserUsecase {
	return &UserUsecase{
		userRepo: repo,
	}
}

func (uu *UserUsecase) GetUsers() ([]model.User, error) {
	return uu.userRepo.GetUsers()
}
