package usecase

import (
	"gin-db/model"
	"gin-db/repository"
)

type CredentialUsecase interface {
	Register(credential *model.Credential) (*model.Credential,error)
	IsEmailExist(email string) (bool,error)
	GetCredentialByEmail(email string,credential model.Credential) (model.Credential,error)
}

type credentialUsecase struct {
	repo repository.CredentialRepository
}

func NewCredentialUsecase(repo repository.CredentialRepository) CredentialUsecase {
	return &credentialUsecase{repo: repo}
}

func (cu *credentialUsecase) Register(credential *model.Credential) (*model.Credential,error) {
	return cu.repo.Register(credential)
}

func (cu *credentialUsecase) IsEmailExist(email string) (bool,error) {
	return cu.repo.IsEmailExist(email)
}

func (cu *credentialUsecase) GetCredentialByEmail(email string,credential model.Credential) (model.Credential,error) {
	return cu.repo.GetCredentialByEmail(email,credential)
}

