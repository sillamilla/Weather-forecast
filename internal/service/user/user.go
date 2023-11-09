package user

import (
	"Weather/pkg/forecast_client"
)

type Service interface {
}

type userService struct {
	client forecast_client.Client
}

func NewUserService(client forecast_client.Client) Service {
	return userService{client: client}
}
