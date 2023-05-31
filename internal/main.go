package main

import (
	"Weather/internal/handler"
	"Weather/internal/service"
	"Weather/internal/service/location"
	"Weather/internal/service/user"
	"Weather/internal/service/weather"
	"github.com/gorilla/mux"
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
	h := handler.New(srv)

	r.HandleFunc("/", h.ServiceStuff).Methods(http.MethodPost)

	http.ListenAndServe(":8080", r)
}
