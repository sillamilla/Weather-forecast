package handler

import (
	"Weather/internal/handler/ui"
	"Weather/internal/service"
	"html/template"
)

type Handler struct {
	ServiceStuff service.Service
	Home         *template.Template
	srv          *service.Service
}

func New(srv *service.Service) Handler {
	home, err := template.ParseFiles("./internal/templates/homePage.html")
	if err != nil {
		panic(err)
	}

	return Handler{
		ServiceStuff: ui.NewHandler(*srv), //todo Я ХОЧУ 1 ХЕНДЛЕР
		Home:         home,
		srv:          srv,
	}
}
