package service

import (
	"Weather/internal/service/location"
	"Weather/internal/service/user"
	"Weather/internal/service/weather"
)

type Service struct {
	LocationSrv location.Service
	UserSrv     user.Service
	WeatherSrv  weather.Service
}

func New(sr *Service) *Service {
	return &Service{
		UserSrv:     user.NewUserService(sr.UserSrv),
		LocationSrv: location.NewLocationService(sr.LocationSrv),
		WeatherSrv:  weather.NewWeatherService(sr.WeatherSrv),
	}
}
