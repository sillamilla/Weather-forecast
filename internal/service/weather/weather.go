package weather

import (
	"Weather/internal/config"
	"Weather/pkg/weather_client/models"
	"Weather/pkg/weather_client"
	"os"
)

type Service interface {
	WeatherStatus(geoLocation models.GeoLocation) (models.WeatherInfo, error)
}

type weatherService struct {
	openWeatherClient weatherThirdPartyService
}

type weatherThirdPartyService interface {
	FindLocation(city string) ([]models.GeoLocation, error)
	WeatherStatus(geoLocation models.GeoLocation) (models.WeatherInfo, error)
}

func NewWeatherService(service weatherThirdPartyService) Service {
	return weatherService{
		openWeatherClient: service,
	}
}

var key = os.Getenv("key")

//todo key

func (ws weatherService) WeatherStatus(geoLocation models.GeoLocation) (models.WeatherInfo, error) {
	weatherInfo, err :=
	return weatherInfo, err
}
