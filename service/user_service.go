package service

import (
	"log"
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
	user.UUID = uuid
	//if ok -> query db
	us.repository.CreateUser(user)
	log.Print(user, 2)
	return 200, nil
}

func (us *UserService) GetUserById(id int) (*models.User, int, error) {
	//valid params todo:
	user, err := us.repository.GetUserById(id)
	if err != nil {
		return nil, 500, err
	}
	return user, 200, nil
}

func (us *UserService) UpdateUserByUUID(user *models.User) (int, error) {
	err := us.repository.UpdateUserByUUID(user)
	if err != nil {
		return 500, err
	}
	return 200, nil
}
