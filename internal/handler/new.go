package handler

import (
	"Weather/internal/service"
	"html/template"
	"net/http"
)

type Handler struct {
	Home *template.Template
	srv  *service.Service
}

func (h Handler) ServiceStuff(writer http.ResponseWriter, request *http.Request) {

}

func New(srv *service.Service) Handler {
	home, err := template.ParseFiles("./internal/templates/homePage.html")
	if err != nil {
		panic(err)
	}

	return Handler{
		Home: home,
		srv:  srv,
	}
}
