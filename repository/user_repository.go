package repository

import (
	"database/sql"
	"gin-db/model"
	_ "github.com/lib/pq"
)

type UserRepository interface {
	GetUser() (*sql.Rows,error)
	GetDetailUser(id int,user *model.User) (*model.User,error)
	CreateUser(user *model.User) (*model.User,error)
	UpdateUser(id int,user *model.User) (*model.User,error)
	DeleteUser(id int) (bool,error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) GetUser() (*sql.Rows, error) {
	// Get all data from customer table
	select_all := "SELECT user_id,u.credential_id,name,age,gender,u.created_at,u.updated_at,c.credential_id,c.email,c.password,c.created_at,c.updated_at FROM users AS u INNER JOIN credential AS c ON u.credential_id = c.credential_id;"

	rows,err := ur.db.Query(select_all)
	if err != nil {
		return nil,err
	}
	return rows,nil
}

func (ur *userRepository) GetDetailUser(id int,user *model.User) (*model.User,error) {
	select_by_id := "SELECT user_id,name,age,gender,u.created_at,u.updated_at,c.credential_id,c.email,c.password,c.created_at,c.updated_at FROM users AS u INNER JOIN credential AS c ON u.credential_id = c.credential_id; WHERE user_id = $1;"
	
	err := ur.db.QueryRow(select_by_id,id).Scan(&user.User_id,&user.Name,&user.Age,&user.Gender,&user.Created_at,&user.Updated_at,&user.Credential.Credential_id,&user.Credential.Email,&user.Credential.Password,&user.Created_at,&user.Credential.Updated_at)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil , nil
		}

		return nil , err
	}

	return user , nil
}

func (ur *userRepository) CreateUser(user *model.User) (*model.User, error) {
	// insert product data into db
	insert_query := "INSERT INTO users (credential_id,name,age,gender) VALUES ($1, $2, $3,$4) RETURNING user_id,created_at,updated_at;"

	err := ur.db.QueryRow(insert_query, user.Credential.Credential_id,user.Name,user.Age,user.Gender).Scan(&user.User_id,&user.Created_at,&user.Updated_at)
	if err != nil {
		return nil, err // Handle error if the query fails
	}
	return user, nil
}

func (ur *userRepository) UpdateUser(id int, user *model.User) (*model.User, error) {
	update := "UPDATE users SET name = $2,age = $3,gender = $4,updated_at = $5 WHERE user_id = $1 RETURNING user_id,created_at,updated_at;"

	err := ur.db.QueryRow(update,id,user.Name,user.Age,user.Gender,user.Updated_at).Scan(&user.User_id,&user.Created_at,&user.Updated_at)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) DeleteUser(id int) (bool, error) {
	query := "DELETE FROM users WHERE user_id = $1"
	// Execute the query and scan the result
	_, err := ur.db.Exec(query, id)
	if err != nil {
		// Return any other errors encountered
		return false, err
	}

	return true, nil
}