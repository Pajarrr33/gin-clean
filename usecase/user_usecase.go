package usecase

import (
	"gin-db/model"
	"gin-db/repository"
)

type UserUsecase interface {
	CreateUser(user *model.User) (*model.User,error)
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

func (uu *userUsecase) CreateUser(user *model.User) (*model.User,error) {
	return uu.repo.CreateUser(user)
}