package service

import (
	"github.com/devstackq/models"
	"github.com/devstackq/repository"
)

//business logic
type User interface {
	CreateUser(models.User) (int, error)
}

type Service struct {
	User
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		User: NewUserService(r.User),
	}
}
