package ui

import (
	"Weather/internal/models"
	"Weather/internal/service"
	models2 "Weather/pkg/weather_client/models"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Handler struct {
	srv  service.Service
	home *template.Template
}

func NewHandler(service service.Service) Handler {
	home, err := template.ParseFiles("./internal/templates/homePage.html")
	if err != nil {
		panic(err)
	}
	return Handler{
		srv:  service,
		home: home,
	}
}

func (h Handler) MainPageProcess(w http.ResponseWriter, r *http.Request) {
	userRequest := models.UserRequest{
		City: "London",
	}

	locations, err := h.srv.LocationSrv.FiendLocation(userRequest)
	if err != nil {
		log.Fatal(err)
	}

	var weathersInfo []models2.WeatherInfo
	var errors []error

	for _, location := range locations {
		weatherInfo, err := h.srv.WeatherSrv.WeatherStatus(location)
		if err != nil {
			errors = append(errors, err)
		}

		weathersInfo = append(weathersInfo, weatherInfo)
	}

	if len(errors) != 0 {
		fmt.Println(errors)
	}

	err = h.home.Execute(w, weathersInfo)
	if err != nil {
		log.Println(err)
		return
	}
}
