package weather

import (
	"Weather/internal/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Service interface {
	WeatherStatus(geoLocation models.GeoLocation) (models.WeatherInfo, error)
}

type weatherService struct {
	ws Service
}

func NewWeatherService(ws Service) Service {
	return weatherService{ws: ws}
}

var key = os.Getenv("key")

//todo key

func (ws weatherService) WeatherStatus(geoLocation models.GeoLocation) (models.WeatherInfo, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s", geoLocation.Lat, geoLocation.Lon, key)

	getResponse, err := http.Get(url)
	if err != nil {
		return models.WeatherInfo{}, err
	}
	defer getResponse.Body.Close()

	readResponse, err := io.ReadAll(getResponse.Body)
	if err != nil {
		return models.WeatherInfo{}, err
	}

	var weatherInfo models.WeatherInfo
	err = json.Unmarshal(readResponse, &weatherInfo)
	if err != nil {
		return models.WeatherInfo{}, err
	}

	return weatherInfo, err
}
