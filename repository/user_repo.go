package repository

import (
	"database/sql"

	"github.com/devstackq/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) CreateUser(user models.User) error {
	//query in db, create
	_, err := ur.db.Exec("INSERT INTO users(uuid, firstname, lastname, email, age, createdtime) values (?,?,?,?,?,?)",
		user.ID, user.FirstName, user.LastName, user.Email, user.Age, user.CreatedTime)
	if err != nil {
		return err
	}
	return nil
}
