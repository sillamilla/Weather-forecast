package helper

import (
	"Weather/internal/models"
	"encoding/json"
	"io"
	"net/http"
)

func ValueFromUrl(url string, information interface{}) error {
	getResponse, err := http.Get(url)
	if err != nil {
		return err
	}
	defer getResponse.Body.Close()

	readResponse, err := io.ReadAll(getResponse.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(readResponse, information)
	if err != nil {
		return err
	}

	return nil
}

func SameCheck(weathersInfo []models.WeatherInfo) []models.WeatherInfo {
	sortMap := make(map[string]models.WeatherInfo)
	for i, value := range weathersInfo {
		sortMap[value.Name] = weathersInfo[i]
	}

	var weather []models.WeatherInfo
	for _, value := range sortMap {
		weather = append(weather, sortMap[value.Name])
	}

	return weather
}
