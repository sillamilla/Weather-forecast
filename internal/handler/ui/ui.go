package ui

import (
	"Weather/internal/models"
	"Weather/internal/repository/users"
	"Weather/internal/service"
	models2 "Weather/pkg/forecast_client/models"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Handler struct {
	rp   users.UserRepo
	srv  service.Service
	home *template.Template
}

func NewHandler(service service.Service, rp users.UserRepo) Handler {
	home, err := template.ParseFiles("./internal/templates/homePage.html")
	if err != nil {
		panic(err)
	}
	return Handler{
		srv:  service,
		home: home,
		rp:   rp,
	}
}

func (h Handler) MainPageProcess(w http.ResponseWriter, r *http.Request) {
	cityName := r.URL.Query().Get("name")
	if cityName == "" {
		cityName = "London"
	}
	userRequest := models.UserRequest{
		City: cityName,
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
