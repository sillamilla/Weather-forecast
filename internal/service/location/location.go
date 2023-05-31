package location

import (
	"Weather/internal/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Service interface {
	FiendLocation(userRequest models.UserRequest) ([]models.GeoLocation, error)
}

type locationService struct {
	ls Service
}

func NewLocationService(ls Service) Service {
	return locationService{ls: ls}
}

var key = os.Getenv("key")

func (s locationService) FiendLocation(userRequest models.UserRequest) ([]models.GeoLocation, error) {
	url := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%s&limit=5&appid=%s", userRequest.City, key)

	getResponse, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer getResponse.Body.Close()

	readResponse, err := io.ReadAll(getResponse.Body)
	if err != nil {
		return nil, err
	}

	var locations []models.GeoLocation
	err = json.Unmarshal(readResponse, &locations)
	if err != nil {
		return nil, err
	}

	return locations, nil
}
