package service

import (
	"github.com/devstackq/models"
	"github.com/devstackq/repository"
)

type UserService struct {
	repository repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo}
}

func (us *UserService) CreateUser(u models.User) error {
	return nil
}
