package usecase

import (
	"gin-db/model"
	"gin-db/repository"
)

type CredentialUsecase interface {
	Register(credential *model.Credential) (*model.Credential,error)
	IsEmailExist(email string) (bool,error)
}

type credentialUsecase struct {
	repo repository.CredentialRepository
}

func NewCredentialUsecase(repo repository.CredentialRepository) CredentialUsecase {
	return &credentialUsecase{repo: repo}
}

func (cu *credentialUsecase) Register(credential *model.Credential) (*model.Credential,error) {
	credential,err := cu.repo.Register(credential)
	if err != nil {
		return nil,err
	}

	return credential,nil
}

func (cu *credentialUsecase) IsEmailExist(email string) (bool,error) {
	isEmailExist,err := cu.repo.IsEmailExist(email)
	if err != nil {
		return false,err
	}
	if !isEmailExist {
		return false,nil
	}

	return true, nil
}

