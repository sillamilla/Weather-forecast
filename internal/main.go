package main

import (
	"Weather/internal/config"
	"Weather/internal/handler"
	"Weather/internal/service"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	srv := service.New()

	config.GetConfig().

	r := mux.NewRouter()
	h := handler.New(srv)

	r.HandleFunc("/", h.ServiceStuff).Methods(http.MethodPost)

	http.ListenAndServe(":8080", r)
}
