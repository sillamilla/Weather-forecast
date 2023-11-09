package main

import (
	"Weather/internal/config"
	"Weather/internal/handler/ui"
	"Weather/internal/service"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	srv := service.New(*config.GetConfig())
	h := ui.NewHandler(*srv)
	r := mux.NewRouter()

	r.HandleFunc("/", (h.MainPageProcess)).Methods(http.MethodGet)
	http.ListenAndServe(":8080", r)
}
