package main

import (
	"Weather/internal/config"
	"Weather/internal/handler/ui"
	"Weather/internal/repository"
	"Weather/internal/service"
	"Weather/pkg/sql_lite"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	database, err := sql_lite.New()
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.New(database)
	srv := service.New(*config.GetConfig(), repo)
	h := ui.NewHandler(*srv, repo.UserSrv)
	r := mux.NewRouter()

	r.HandleFunc("/", h.Authorization(h.MainPageProcess)).Methods(http.MethodGet)
	http.ListenAndServe(":8080", r)
}
