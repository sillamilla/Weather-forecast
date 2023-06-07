package service

import (
	"Weather/internal/config"
	"Weather/internal/service/location"
	"Weather/internal/service/user"
	"Weather/internal/service/weather"
	"Weather/pkg/weather_client"
)

type Service struct {
	LocationSrv location.Service
	UserSrv     user.Service
	WeatherSrv  weather.Service
}

func New(cfg config.Config) *Service {
	client := weather_client.NewClient(cfg.WeatherApiUrl, cfg.WeatherApiToken)

	return &Service{
		UserSrv:     user.NewUserService(),
		LocationSrv: location.NewLocationService(),
		WeatherSrv:  weather.NewWeatherService(client),
	}
}
