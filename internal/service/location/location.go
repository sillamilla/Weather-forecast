package location

import (
	"Weather/internal/config"
	"Weather/internal/models"
	models2 "Weather/pkg/weather_client/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Service interface {
	FiendLocation(userRequest models.UserRequest) ([]models2.GeoLocation, error)
}

type locationService struct {
}

func NewLocationService() Service {
	return locationService{}
}

func (s locationService) FiendLocation(userRequest models.UserRequest) ([]models2.GeoLocation, error) {
	weatherApiUrl := config.GetConfig().WeatherApiUrl
	weatherApiToken := config.GetConfig().WeatherApiToken

	url := fmt.Sprintf("%s/geo/1.0/direct?q=%s&limit=5&appid=%s", weatherApiUrl, userRequest.City, weatherApiToken)

	getResponse, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer getResponse.Body.Close()

	readResponse, err := io.ReadAll(getResponse.Body)
	if err != nil {
		return nil, err
	}

	var locations []models2.GeoLocation
	err = json.Unmarshal(readResponse, &locations)
	if err != nil {
		return nil, err
	}

	return locations, nil
}
