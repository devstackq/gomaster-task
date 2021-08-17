package service

import (
	"github.com/devstackq/models"
	"github.com/devstackq/repository"
)

//bussynnes logic
type User interface {
	CreateUser(models.User) error
}

type Service struct {
	User
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		User: NewUserService(r.User),
	}
}
