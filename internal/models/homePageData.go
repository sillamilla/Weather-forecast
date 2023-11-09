package models

import "Weather/pkg/forecast_client/models"

type PageData struct {
	//PinedLocations []string
	WeatherInfo []models.WeatherInfo
}
