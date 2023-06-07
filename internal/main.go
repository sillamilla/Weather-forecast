package main

import (
	"Weather/internal/config"
	"Weather/internal/handler"
	"Weather/internal/service"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	srv := service.New(*config.GetConfig())
	h := handler.New(srv)
	r := mux.NewRouter()

	r.HandleFunc("/", h.ServiceStuff).Methods(http.MethodGet)

	http.ListenAndServe(":8080", r)
}
