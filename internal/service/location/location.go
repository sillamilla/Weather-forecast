package location

import (
	"Weather/internal/helper"
	"Weather/internal/models"
	"Weather/pkg/forecast_client"
	models2 "Weather/pkg/forecast_client/models"
)

type Service interface {
	FiendLocation(userRequest models.UserRequest) ([]models2.GeoLocation, error)
}

type locationService struct {
	client forecast_client.Client
}

func NewLocationService(client forecast_client.Client) Service {
	return locationService{client: client}
}

func (ls locationService) FiendLocation(userRequest models.UserRequest) ([]models2.GeoLocation, error) {
	locations, err := ls.client.FindLocation(userRequest.City)
	location := helper.SameCheck(locations)
	if err != nil {
		return nil, err
	}

	return location, nil
}
