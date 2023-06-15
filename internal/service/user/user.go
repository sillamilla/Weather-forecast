package user

import (
	"Weather/internal/models"
	"Weather/internal/repository/users"
)

type Service interface {
	AddUser(newCookie string) error
	GetUser(cookieValue string) (models.User, error)
}

type userService struct {
	rp users.UserRepo
}

func NewUserService(rp users.UserRepo) Service {
	return userService{rp: rp}
}

func (us userService) AddUser(newCookie string) error {
	err := us.rp.AddUser(newCookie)
	if err != nil {
		return err //todo error wrap
	}

	return nil
}

func (us userService) GetUser(cookieValue string) (models.User, error) {
	user, err := us.rp.GetUser(cookieValue)
	if err != nil {
		return user, err //todo error wrap
	}

	return user, nil
}
