package service

import (
	"Weather/internal/config"
	"Weather/internal/service/location"
	"Weather/internal/service/user"
	"Weather/internal/service/weather"
	"Weather/pkg/forecast_client"
)

type Service struct {
	LocationSrv location.Service
	UserSrv     user.Service
	WeatherSrv  weather.Service
}

func New(cfg config.Config) *Service {
	client := forecast_client.NewClient(cfg.WeatherApiUrl, cfg.WeatherApiToken)

	return &Service{
		UserSrv:     user.NewUserService(*client),
		LocationSrv: location.NewLocationService(*client),
		WeatherSrv:  weather.NewWeatherService(*client),
	}
}
