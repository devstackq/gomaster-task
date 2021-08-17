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
	_, err := ur.db.Exec("INSERT INTO users(uuid, firstname, lastname, email, age, createdtime) VALUES ($1,$2,$3,$4,$5,$6)",
		user.UUID, user.FirstName, user.LastName, user.Email, user.Age, user.CreatedTime)
	if err != nil {
		return err
	}
	return nil
}
func (ur *UserRepository) GetUserById(id int) (*models.User, error) {
	var user models.User
	sqlStatement := `SELECT uuid, firstname, lastname, email, age, createdtime FROM users WHERE id=$1;`
	row := ur.db.QueryRow(sqlStatement, id)
	err := row.Scan(&user.UUID, &user.FirstName, &user.LastName, &user.Email, &user.Age, &user.CreatedTime)
	if err != nil {
		return nil, sql.ErrNoRows
	}
	return &user, nil
}
func (ur *UserRepository) UpdateUserByUUID(user *models.User) error {
	sqlStatement := `UPDATE users SET firstname = $1, lastname= $2,  email=$3,  age=$4 WHERE uuid = $5`
	_, err := ur.db.Exec(sqlStatement, user.FirstName, user.LastName, user.Email, user.Age, user.UUID)
	if err != nil {
		panic(err)
	}
	if err != nil {
		return sql.ErrNoRows
	}
	return nil
}
