package main

import (
	"Weather/internal/helper"
	"Weather/internal/models"
	"Weather/internal/service"
	"Weather/internal/service/location"
	"Weather/internal/service/user"
	"Weather/internal/service/weather"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

func main() {

	sr := &service.Service{
		UserSrv:     user.NewUserService(nil),
		LocationSrv: location.NewLocationService(nil),
		WeatherSrv:  weather.NewWeatherService(nil),
	}
	srv := service.New(sr)
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cityName := r.URL.Query().Get("name")
		if cityName == "" {
			cityName = "London"
		}
		userRequest := models.UserRequest{
			City: cityName,
		}

		locations, err := srv.LocationSrv.FiendLocation(userRequest)
		if err != nil {
			log.Fatal(err)
		}

		var weathersInfo []models.WeatherInfo
		var errors []error

		for _, location := range locations {
			weatherInfo, err := srv.WeatherSrv.WeatherStatus(location)
			if err != nil {
				errors = append(errors, err)
			}

			weathersInfo = append(weathersInfo, weatherInfo)
		}
		if len(errors) != 0 {
			fmt.Println(errors)
		}

		weather := helper.SameCheck(weathersInfo)

		homePage, err := template.ParseFiles("./internal/templates/homePage.html")
		if err != nil {
			panic(err)
		}
		err = homePage.Execute(w, weather)
		if err != nil {
			log.Println(err)
			return
		}
	}).Methods(http.MethodGet)

	http.ListenAndServe(":8080", r)
}
