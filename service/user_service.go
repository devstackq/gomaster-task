package service

import (
	"net/http"
	"time"

	"github.com/devstackq/models"
	"github.com/devstackq/repository"
	uuid "github.com/satori/go.uuid"
)

type UserService struct {
	repository repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo}
}

func (us *UserService) CreateUser(user models.User) (int, error) {
	//valid params todo:
	var err error

	uuid := uuid.Must(uuid.NewV4(), err).String()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	user.CreatedTime = time.Now()
	user.ID = uuid
	//if ok -> query db
	us.repository.CreateUser(user)

	return 200, nil
}
