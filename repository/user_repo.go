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
	return nil
}
