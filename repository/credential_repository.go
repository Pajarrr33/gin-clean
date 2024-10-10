package repository

import (
	"database/sql"
	"gin-db/model"
	_ "github.com/lib/pq"
)

type CredentialRepository interface {
	Register(credential *model.Credential) (*model.Credential,error)
	IsEmailExist(email string) (bool, error)
	GetCredentialByEmail(email string,credential model.Credential) (model.Credential,error)
}

type credentialRepository struct {
	db *sql.DB
}

func NewCredentialRepo(db *sql.DB) CredentialRepository {
	return &credentialRepository{db: db}
}

func (cr *credentialRepository) Register(credential *model.Credential) (*model.Credential,error) {
	query := "INSERT INTO credential (email,password) VALUES ($1,$2) RETURNING credential_id,created_at,updated_at"

	err := cr.db.QueryRow(query,credential.Email,credential.Password).Scan(&credential.Credential_id,&credential.Created_at,&credential.Updated_at)
	if err != nil {
		return nil,err
	}
	return credential,nil
}

func (cr *credentialRepository) IsEmailExist(email string) (bool, error) {
	query := "SELECT email FROM credential WHERE email = $1"

	err := cr.db.QueryRow(query,email).Scan(&email)
	if err != nil {
		if err == sql.ErrNoRows {
			// No customer found
			return false, nil // No error, just return false
		}
		return false,err
	}
	
	return true,nil
}

func (cr *credentialRepository) GetCredentialByEmail(email string,credential model.Credential) (model.Credential,error) {
	query := "SELECT credential_id,email,password,created_at,updated_at FROM credential WHERE email = $1"

	err := cr.db.QueryRow(query,email).Scan(&credential.Credential_id,&credential.Email,&credential.Password,&credential.Created_at,&credential.Updated_at)
	if err != nil {
		return credential,err
	}
	
	return credential,nil
}