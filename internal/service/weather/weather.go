package weather

import (
	"Weather/pkg/forecast_client"
	"Weather/pkg/forecast_client/models"
)

type Service interface {
	WeatherStatus(geoLocation models.GeoLocation) (models.WeatherInfo, error)
}

type weatherService struct {
	client forecast_client.Client
}

func NewWeatherService(client forecast_client.Client) Service {
	return weatherService{client: client}
}

func (ws weatherService) WeatherStatus(geoLocation models.GeoLocation) (models.WeatherInfo, error) {
	weatherInfo, err := ws.client.WeatherStatus(geoLocation)
	if err != nil {
		return weatherInfo, err
	}

	return weatherInfo, nil
}
